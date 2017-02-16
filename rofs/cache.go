package rofs

import (
	"github.com/g8os/g8ufs/meta"
	"io"
	"os"
	"path"
	//"syscall"
	"github.com/golang/snappy"
	"github.com/xxtea/xxtea-go/xxtea"
	"io/ioutil"
	"syscall"
)

func (fs *filesystem) path(hash string) string {
	return path.Join(fs.cache, hash)
}

func (fs *filesystem) exists(hash string) bool {
	name := path.Join(fs.cache, hash)
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func (fs *filesystem) checkAndGet(m meta.Meta) (*os.File, error) {
	//atomic check and download a file
	name := fs.path(m.ID())
	info := m.Info()
	f, err := os.OpenFile(name, os.O_CREATE|os.O_RDWR, os.ModePerm&os.FileMode(0755))
	if err != nil {
		return nil, err
	}
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		return nil, err
	}

	defer syscall.Flock(int(f.Fd()), syscall.LOCK_UN)

	fstat, err := f.Stat()

	if err != nil {
		return nil, err
	}

	if fstat.Size() == int64(info.Size) {
		return f, nil
	}

	if err := fs.download(f, m); err != nil {
		f.Close()
		os.Remove(name)
		return nil, err
	}

	f.Seek(0, os.SEEK_SET)
	return f, nil
}

func (fs *filesystem) downloadBlock(block meta.BlockInfo, writer io.Writer) error {
	body, err := fs.storage.Get(string(block.Key))
	if err != nil {
		return err
	}

	defer body.Close()

	data, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	data = xxtea.Decrypt(data, block.Decipher)
	raw, err := snappy.Decode(nil, data)
	if err != nil {
		return err
	}

	if _, err := writer.Write(raw); err != nil {
		return err
	}

	return nil
}

// download file from stor
func (fs *filesystem) download(file *os.File, m meta.Meta) error {
	for _, block := range m.Blocks() {
		if err := fs.downloadBlock(block, file); err != nil {
			return err
		}
	}

	return nil
}
