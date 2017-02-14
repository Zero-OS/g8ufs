package meta

import (
	"bytes"
	"fmt"
	"github.com/codahale/blake2"
	"github.com/g8os/g8ufs/cap.np"
	"github.com/patrickmn/go-cache"
	rocksdb "github.com/tecbot/gorocksdb"
	"io"
	"path"
	"sync"
	"time"
	"zombiezen.com/go/capnproto2"
)

func NewRocksMeta(ns string, dbpath string) (MetaStore, error) {
	opt := rocksdb.NewDefaultOptions()
	db, err := rocksdb.OpenDbForReadOnly(opt, dbpath, true)
	if err != nil {
		return nil, err
	}

	return &rocksMetaStore{
		ns:    ns,
		db:    db,
		ro:    rocksdb.NewDefaultReadOptions(),
		cache: cache.New(5*time.Second, 1*time.Second),
	}, nil
}

type rocksMeta struct {
	name  string
	dir   np.Dir
	inode np.Inode

	store *rocksMetaStore
	o     sync.Once
}

func (rm *rocksMeta) load() {
	//load runs once and only take effect if
	//the meta wasn't already filed in by a previous Meta node.

	rm.o.Do(func() {
		if rm.dir.HasData() {
			return
		}

		if loaded, ok := rm.store.Get(rm.name); ok {
			self := loaded.(*rocksMeta)
			rm.dir = self.dir
			rm.inode = self.inode
		}
	})
}

func (rm *rocksMeta) IsDir() bool {
	rm.load()
	return !rm.inode.HasData()
}

func (rm *rocksMeta) String() string {
	return rm.name
}

//base name
func (rm *rocksMeta) Name() string {
	return path.Base(rm.name)
}

func (rm *rocksMeta) Blocks() []string {
	rm.load()
	var blocks []string
	if !rm.inode.HasData() {
		return blocks
	}

	attrs := rm.inode.Attributes()
	if !attrs.HasFile() {
		return blocks
	}
	file, _ := attrs.File()
	if !file.HasBlocks() {
		return blocks
	}

	cblocks, _ := file.Blocks()
	for i := 0; i < cblocks.Len(); i++ {
		bytes, _ := cblocks.At(i)
		blocks = append(blocks, string(bytes))
	}
	return blocks
}

func (rm *rocksMeta) Children() []Meta {
	rm.load()
	//if that is a file, we must have no children
	var children []Meta
	if rm.inode.HasData() {
		return children
	}

	if !rm.dir.HasContents() {
		return children
	}

	contents, _ := rm.dir.Contents()

	for i := 0; i < contents.Len(); i++ {
		inode := contents.At(i)
		name, _ := inode.Name()
		var child *rocksMeta
		if inode.Attributes().HasDir() {
			//we don't set the dir, to force it to load it's content when
			//it has too.
			child = &rocksMeta{
				store: rm.store,
				name:  path.Join(rm.name, name),
			}
		} else {
			//here we already have everything we know about this node (file)
			//so we just create a complete struct.
			child = &rocksMeta{
				store: rm.store,
				name:  path.Join(rm.name, name),
				dir:   rm.dir,
				inode: inode,
			}
		}

		children = append(children, child)
	}

	return children
}

func (rm *rocksMeta) Info() MetaInfo {
	rm.load()
	if !rm.inode.HasData() {
		//that must be a dir.
		return MetaInfo{
			Type:             DirType,
			Size:             rm.dir.Size(),
			CreationTime:     rm.dir.CreationTime(),
			ModificationTime: rm.dir.ModificationTime(),
		}
	}

	info := MetaInfo{
		CreationTime:     rm.inode.CreationTime(),
		ModificationTime: rm.inode.ModificationTime(),
	}

	attrs := rm.inode.Attributes()
	if !attrs.HasData() {
		log.Fatalf("'%s' attributes is empty", attrs)
	}

	if attrs.HasFile() {
		file, _ := attrs.File()
		info.Type = RegularType
		info.Size = rm.inode.Size()
		info.FileBlockSize = file.BlockSize()
	} else if attrs.HasLink() {
		link, _ := attrs.Link()
		info.Type = LinkType
		target, _ := link.Target()
		info.LinkTarget = target
	}

	return info
}

type rocksMetaStore struct {
	ns    string
	db    *rocksdb.DB
	ro    *rocksdb.ReadOptions
	cache *cache.Cache
}

func (rs *rocksMetaStore) hash(namespace string, path string) (string, error) {
	bl2b := blake2.New(&blake2.Config{
		Size: 32,
	})

	_, err := io.WriteString(bl2b, fmt.Sprintf("%s%s", namespace, path))
	if err != nil {
		return "", err
	}

	hash := fmt.Sprintf("%x", bl2b.Sum(nil))
	if namespace != "" {
		return fmt.Sprintf("%s:%s", namespace, hash), nil
	} else {
		return hash, nil
	}
}

func (rs *rocksMetaStore) dirFromSlice(slice *rocksdb.Slice) np.Dir {
	msg, err := capnp.NewDecoder(bytes.NewBuffer(slice.Data())).Decode()
	if err != nil {
		log.Fatal("invalid capnp Dir message")
	}

	dir, err := np.ReadRootDir(msg)
	if err != nil {
		log.Fatal("failed to read Dir message")
	}

	return dir
}

func (rs *rocksMetaStore) aciFromSlice(slice *rocksdb.Slice) np.ACI {
	msg, err := capnp.NewDecoder(bytes.NewBuffer(slice.Data())).Decode()
	if err != nil {
		log.Fatal("invalid capnp ACI message")
	}

	aci, err := np.ReadRootACI(msg)
	if err != nil {
		log.Fatal("failed to read ACI messeage")
	}
	return aci
}

func (rs *rocksMetaStore) getACI(key string) np.ACI {
	slice, err := rs.db.Get(rs.ro, []byte(key))
	if err != nil {
		log.Fatalf("failed to get ACI (%s) from db: %s", key, err)
	}

	return rs.aciFromSlice(slice)
}

func (rs *rocksMetaStore) get(name string, level int) (*rocksMeta, bool) {
	if level == 0 {
		return nil, false
	}
	if obj, ok := rs.cache.Get(name); ok {
		log.Debugf("cache hit for name: '%s' (%d)", name, level)
		return obj.(*rocksMeta), true
	}

	hash, _ := rs.hash(rs.ns, name)
	slice, err := rs.db.Get(rs.ro, []byte(hash))
	if err != nil {
		log.Fatal(err)
	}
	defer slice.Free()

	if slice.Size() != 0 {
		dir := rs.dirFromSlice(slice)

		return &rocksMeta{
			store: rs,
			name:  name,
			dir:   dir,
		}, true
	}

	return rs.get(path.Dir(name), level-1)
}

func (rs *rocksMetaStore) Get(name string) (Meta, bool) {
	/*
		When we try to hit an object, this object can be either a directory or any other file type.
		The rocks-db has all values as directories, which means if we try to retrieve an object and we
		didn't find it, we should try to retrieve it's parent. Then find the file name in that directory entry.
	*/

	//we search only 2 levels
	meta, ok := rs.get(name, 2)
	if !ok {
		return nil, false
	}

	//direct hit. we return
	if meta.name == name {
		rs.cache.Set(name, meta, cache.DefaultExpiration)
		return meta, true
	}

	log.Debugf("searching children of '%s'", meta.name)
	//searching the directory contents.
	if !meta.dir.HasContents() {
		return nil, false
	}

	contents, _ := meta.dir.Contents()
	base := path.Base(name)

	for i := 0; i < contents.Len(); i++ {
		inode := contents.At(i)
		nodeName, _ := inode.Name()
		if nodeName == base {
			log.Debugf("found %s in children of %s", nodeName, meta.name)
			nodeMeta := &rocksMeta{
				name:  name,
				dir:   meta.dir,
				inode: inode,
				store: rs,
			}
			//cache it
			rs.cache.Set(name, nodeMeta, cache.DefaultExpiration)
			return nodeMeta, true
		}
	}

	return nil, false
}
