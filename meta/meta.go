package meta

import (
	"fmt"
	"github.com/op/go-logging"
	"syscall"
)

var (
	log         = logging.MustGetLogger("meta")
	ErrNotFound = fmt.Errorf("not found")
)

type NodeType uint32

const (
	DirType         = NodeType(syscall.S_IFDIR)
	RegularType     = NodeType(syscall.S_IFREG)
	BlockDeviceType = NodeType(syscall.S_IFBLK)
	CharDeviceType  = NodeType(syscall.S_IFCHR)
	SocketType      = NodeType(syscall.S_IFSOCK)
	FIFOType        = NodeType(syscall.S_IFIFO)
	LinkType        = NodeType(syscall.S_IFLNK)
)

type MetaData struct {
	Hash        string // file hash
	Size        uint64 // file size in bytes
	Uname       string // username (used for permissions)
	Uid         uint32
	Gname       string // groupname (used for permissions)
	Gid         uint32
	IsDir       bool
	Permissions uint32 // perissions (octal style)
	Ctime       uint64 // creation time
	Mtime       uint64 // modification time
	UserKey     string
	StoreKey    string
	Inode       uint64
}

type MetaInfo struct {
	//Common
	CreationTime     uint32
	ModificationTime uint32
	ACL              string
	Type             NodeType

	//Specific Attr

	//Link
	LinkTarget string

	//File
	FileSize      uint64
	FileBlockSize uint8

	//Special
	SpecialData string
}

type Meta interface {
	fmt.Stringer
	//base name
	Name() string
	Hash() string
	IsDir() bool

	Stat() MetaData //deprecated
	Info() MetaInfo

	Children() []Meta
}

type MetaStore interface {
	// Populate(entry Entry) error
	Get(name string) (Meta, bool)
}
