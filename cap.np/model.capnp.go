package np

// AUTO GENERATED - DO NOT EDIT

import (
	strconv "strconv"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type FileBlock struct{ capnp.Struct }

// FileBlock_TypeID is the unique identifier for the type FileBlock.
const FileBlock_TypeID = 0xd5a2538369c2f82a

func NewFileBlock(s *capnp.Segment) (FileBlock, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return FileBlock{st}, err
}

func NewRootFileBlock(s *capnp.Segment) (FileBlock, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return FileBlock{st}, err
}

func ReadRootFileBlock(msg *capnp.Message) (FileBlock, error) {
	root, err := msg.RootPtr()
	return FileBlock{root.Struct()}, err
}

func (s FileBlock) String() string {
	str, _ := text.Marshal(0xd5a2538369c2f82a, s.Struct)
	return str
}

func (s FileBlock) Hash() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s FileBlock) HasHash() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s FileBlock) SetHash(v []byte) error {
	return s.Struct.SetData(0, v)
}

func (s FileBlock) Key() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.Data()), err
}

func (s FileBlock) HasKey() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s FileBlock) SetKey(v []byte) error {
	return s.Struct.SetData(1, v)
}

// FileBlock_List is a list of FileBlock.
type FileBlock_List struct{ capnp.List }

// NewFileBlock creates a new list of FileBlock.
func NewFileBlock_List(s *capnp.Segment, sz int32) (FileBlock_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return FileBlock_List{l}, err
}

func (s FileBlock_List) At(i int) FileBlock { return FileBlock{s.List.Struct(i)} }

func (s FileBlock_List) Set(i int, v FileBlock) error { return s.List.SetStruct(i, v.Struct) }

// FileBlock_Promise is a wrapper for a FileBlock promised by a client call.
type FileBlock_Promise struct{ *capnp.Pipeline }

func (p FileBlock_Promise) Struct() (FileBlock, error) {
	s, err := p.Pipeline.Struct()
	return FileBlock{s}, err
}

type File struct{ capnp.Struct }

// File_TypeID is the unique identifier for the type File.
const File_TypeID = 0xecfda38634f4591a

func NewFile(s *capnp.Segment) (File, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return File{st}, err
}

func NewRootFile(s *capnp.Segment) (File, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return File{st}, err
}

func ReadRootFile(msg *capnp.Message) (File, error) {
	root, err := msg.RootPtr()
	return File{root.Struct()}, err
}

func (s File) String() string {
	str, _ := text.Marshal(0xecfda38634f4591a, s.Struct)
	return str
}

func (s File) BlockSize() uint8 {
	return s.Struct.Uint8(0)
}

func (s File) SetBlockSize(v uint8) {
	s.Struct.SetUint8(0, v)
}

func (s File) Blocks() (FileBlock_List, error) {
	p, err := s.Struct.Ptr(0)
	return FileBlock_List{List: p.List()}, err
}

func (s File) HasBlocks() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s File) SetBlocks(v FileBlock_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewBlocks sets the blocks field to a newly
// allocated FileBlock_List, preferring placement in s's segment.
func (s File) NewBlocks(n int32) (FileBlock_List, error) {
	l, err := NewFileBlock_List(s.Struct.Segment(), n)
	if err != nil {
		return FileBlock_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// File_List is a list of File.
type File_List struct{ capnp.List }

// NewFile creates a new list of File.
func NewFile_List(s *capnp.Segment, sz int32) (File_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return File_List{l}, err
}

func (s File_List) At(i int) File { return File{s.List.Struct(i)} }

func (s File_List) Set(i int, v File) error { return s.List.SetStruct(i, v.Struct) }

// File_Promise is a wrapper for a File promised by a client call.
type File_Promise struct{ *capnp.Pipeline }

func (p File_Promise) Struct() (File, error) {
	s, err := p.Pipeline.Struct()
	return File{s}, err
}

type Link struct{ capnp.Struct }

// Link_TypeID is the unique identifier for the type Link.
const Link_TypeID = 0xe419a0e5a661965c

func NewLink(s *capnp.Segment) (Link, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Link{st}, err
}

func NewRootLink(s *capnp.Segment) (Link, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Link{st}, err
}

func ReadRootLink(msg *capnp.Message) (Link, error) {
	root, err := msg.RootPtr()
	return Link{root.Struct()}, err
}

func (s Link) String() string {
	str, _ := text.Marshal(0xe419a0e5a661965c, s.Struct)
	return str
}

func (s Link) Target() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Link) HasTarget() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Link) TargetBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Link) SetTarget(v string) error {
	return s.Struct.SetText(0, v)
}

// Link_List is a list of Link.
type Link_List struct{ capnp.List }

// NewLink creates a new list of Link.
func NewLink_List(s *capnp.Segment, sz int32) (Link_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Link_List{l}, err
}

func (s Link_List) At(i int) Link { return Link{s.List.Struct(i)} }

func (s Link_List) Set(i int, v Link) error { return s.List.SetStruct(i, v.Struct) }

// Link_Promise is a wrapper for a Link promised by a client call.
type Link_Promise struct{ *capnp.Pipeline }

func (p Link_Promise) Struct() (Link, error) {
	s, err := p.Pipeline.Struct()
	return Link{s}, err
}

type Special struct{ capnp.Struct }

// Special_TypeID is the unique identifier for the type Special.
const Special_TypeID = 0xdc74a897ce683c6b

func NewSpecial(s *capnp.Segment) (Special, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Special{st}, err
}

func NewRootSpecial(s *capnp.Segment) (Special, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Special{st}, err
}

func ReadRootSpecial(msg *capnp.Message) (Special, error) {
	root, err := msg.RootPtr()
	return Special{root.Struct()}, err
}

func (s Special) String() string {
	str, _ := text.Marshal(0xdc74a897ce683c6b, s.Struct)
	return str
}

func (s Special) Type() Special_Type {
	return Special_Type(s.Struct.Uint16(0))
}

func (s Special) SetType(v Special_Type) {
	s.Struct.SetUint16(0, uint16(v))
}

func (s Special) Data() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s Special) HasData() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Special) SetData(v []byte) error {
	return s.Struct.SetData(0, v)
}

// Special_List is a list of Special.
type Special_List struct{ capnp.List }

// NewSpecial creates a new list of Special.
func NewSpecial_List(s *capnp.Segment, sz int32) (Special_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return Special_List{l}, err
}

func (s Special_List) At(i int) Special { return Special{s.List.Struct(i)} }

func (s Special_List) Set(i int, v Special) error { return s.List.SetStruct(i, v.Struct) }

// Special_Promise is a wrapper for a Special promised by a client call.
type Special_Promise struct{ *capnp.Pipeline }

func (p Special_Promise) Struct() (Special, error) {
	s, err := p.Pipeline.Struct()
	return Special{s}, err
}

type Special_Type uint16

// Values of Special_Type.
const (
	Special_Type_socket   Special_Type = 0
	Special_Type_block    Special_Type = 1
	Special_Type_chardev  Special_Type = 2
	Special_Type_fifopipe Special_Type = 3
	Special_Type_unknown  Special_Type = 4
)

// String returns the enum's constant name.
func (c Special_Type) String() string {
	switch c {
	case Special_Type_socket:
		return "socket"
	case Special_Type_block:
		return "block"
	case Special_Type_chardev:
		return "chardev"
	case Special_Type_fifopipe:
		return "fifopipe"
	case Special_Type_unknown:
		return "unknown"

	default:
		return ""
	}
}

// Special_TypeFromString returns the enum value with a name,
// or the zero value if there's no such value.
func Special_TypeFromString(c string) Special_Type {
	switch c {
	case "socket":
		return Special_Type_socket
	case "block":
		return Special_Type_block
	case "chardev":
		return Special_Type_chardev
	case "fifopipe":
		return Special_Type_fifopipe
	case "unknown":
		return Special_Type_unknown

	default:
		return 0
	}
}

type Special_Type_List struct{ capnp.List }

func NewSpecial_Type_List(s *capnp.Segment, sz int32) (Special_Type_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return Special_Type_List{l.List}, err
}

func (l Special_Type_List) At(i int) Special_Type {
	ul := capnp.UInt16List{List: l.List}
	return Special_Type(ul.At(i))
}

func (l Special_Type_List) Set(i int, v Special_Type) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

type SubDir struct{ capnp.Struct }

// SubDir_TypeID is the unique identifier for the type SubDir.
const SubDir_TypeID = 0xa4a421ce00f301dd

func NewSubDir(s *capnp.Segment) (SubDir, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SubDir{st}, err
}

func NewRootSubDir(s *capnp.Segment) (SubDir, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return SubDir{st}, err
}

func ReadRootSubDir(msg *capnp.Message) (SubDir, error) {
	root, err := msg.RootPtr()
	return SubDir{root.Struct()}, err
}

func (s SubDir) String() string {
	str, _ := text.Marshal(0xa4a421ce00f301dd, s.Struct)
	return str
}

func (s SubDir) Key() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s SubDir) HasKey() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s SubDir) KeyBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s SubDir) SetKey(v string) error {
	return s.Struct.SetText(0, v)
}

// SubDir_List is a list of SubDir.
type SubDir_List struct{ capnp.List }

// NewSubDir creates a new list of SubDir.
func NewSubDir_List(s *capnp.Segment, sz int32) (SubDir_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return SubDir_List{l}, err
}

func (s SubDir_List) At(i int) SubDir { return SubDir{s.List.Struct(i)} }

func (s SubDir_List) Set(i int, v SubDir) error { return s.List.SetStruct(i, v.Struct) }

// SubDir_Promise is a wrapper for a SubDir promised by a client call.
type SubDir_Promise struct{ *capnp.Pipeline }

func (p SubDir_Promise) Struct() (SubDir, error) {
	s, err := p.Pipeline.Struct()
	return SubDir{s}, err
}

type Inode struct{ capnp.Struct }
type Inode_attributes Inode
type Inode_attributes_Which uint16

const (
	Inode_attributes_Which_dir     Inode_attributes_Which = 0
	Inode_attributes_Which_file    Inode_attributes_Which = 1
	Inode_attributes_Which_link    Inode_attributes_Which = 2
	Inode_attributes_Which_special Inode_attributes_Which = 3
)

func (w Inode_attributes_Which) String() string {
	const s = "dirfilelinkspecial"
	switch w {
	case Inode_attributes_Which_dir:
		return s[0:3]
	case Inode_attributes_Which_file:
		return s[3:7]
	case Inode_attributes_Which_link:
		return s[7:11]
	case Inode_attributes_Which_special:
		return s[11:18]

	}
	return "Inode_attributes_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Inode_TypeID is the unique identifier for the type Inode.
const Inode_TypeID = 0xc0029f81b3eee594

func NewInode(s *capnp.Segment) (Inode, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 3})
	return Inode{st}, err
}

func NewRootInode(s *capnp.Segment) (Inode, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 3})
	return Inode{st}, err
}

func ReadRootInode(msg *capnp.Message) (Inode, error) {
	root, err := msg.RootPtr()
	return Inode{root.Struct()}, err
}

func (s Inode) String() string {
	str, _ := text.Marshal(0xc0029f81b3eee594, s.Struct)
	return str
}

func (s Inode) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Inode) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Inode) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Inode) SetName(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Inode) Size() uint64 {
	return s.Struct.Uint64(0)
}

func (s Inode) SetSize(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s Inode) Attributes() Inode_attributes { return Inode_attributes(s) }

func (s Inode_attributes) Which() Inode_attributes_Which {
	return Inode_attributes_Which(s.Struct.Uint16(8))
}
func (s Inode_attributes) Dir() (SubDir, error) {
	p, err := s.Struct.Ptr(1)
	return SubDir{Struct: p.Struct()}, err
}

func (s Inode_attributes) HasDir() bool {
	if s.Struct.Uint16(8) != 0 {
		return false
	}
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Inode_attributes) SetDir(v SubDir) error {
	s.Struct.SetUint16(8, 0)
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewDir sets the dir field to a newly
// allocated SubDir struct, preferring placement in s's segment.
func (s Inode_attributes) NewDir() (SubDir, error) {
	s.Struct.SetUint16(8, 0)
	ss, err := NewSubDir(s.Struct.Segment())
	if err != nil {
		return SubDir{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s Inode_attributes) File() (File, error) {
	p, err := s.Struct.Ptr(1)
	return File{Struct: p.Struct()}, err
}

func (s Inode_attributes) HasFile() bool {
	if s.Struct.Uint16(8) != 1 {
		return false
	}
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Inode_attributes) SetFile(v File) error {
	s.Struct.SetUint16(8, 1)
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewFile sets the file field to a newly
// allocated File struct, preferring placement in s's segment.
func (s Inode_attributes) NewFile() (File, error) {
	s.Struct.SetUint16(8, 1)
	ss, err := NewFile(s.Struct.Segment())
	if err != nil {
		return File{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s Inode_attributes) Link() (Link, error) {
	p, err := s.Struct.Ptr(1)
	return Link{Struct: p.Struct()}, err
}

func (s Inode_attributes) HasLink() bool {
	if s.Struct.Uint16(8) != 2 {
		return false
	}
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Inode_attributes) SetLink(v Link) error {
	s.Struct.SetUint16(8, 2)
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewLink sets the link field to a newly
// allocated Link struct, preferring placement in s's segment.
func (s Inode_attributes) NewLink() (Link, error) {
	s.Struct.SetUint16(8, 2)
	ss, err := NewLink(s.Struct.Segment())
	if err != nil {
		return Link{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s Inode_attributes) Special() (Special, error) {
	p, err := s.Struct.Ptr(1)
	return Special{Struct: p.Struct()}, err
}

func (s Inode_attributes) HasSpecial() bool {
	if s.Struct.Uint16(8) != 3 {
		return false
	}
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Inode_attributes) SetSpecial(v Special) error {
	s.Struct.SetUint16(8, 3)
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewSpecial sets the special field to a newly
// allocated Special struct, preferring placement in s's segment.
func (s Inode_attributes) NewSpecial() (Special, error) {
	s.Struct.SetUint16(8, 3)
	ss, err := NewSpecial(s.Struct.Segment())
	if err != nil {
		return Special{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s Inode) Aclkey() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s Inode) HasAclkey() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Inode) AclkeyBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s Inode) SetAclkey(v string) error {
	return s.Struct.SetText(2, v)
}

func (s Inode) ModificationTime() uint32 {
	return s.Struct.Uint32(12)
}

func (s Inode) SetModificationTime(v uint32) {
	s.Struct.SetUint32(12, v)
}

func (s Inode) CreationTime() uint32 {
	return s.Struct.Uint32(16)
}

func (s Inode) SetCreationTime(v uint32) {
	s.Struct.SetUint32(16, v)
}

// Inode_List is a list of Inode.
type Inode_List struct{ capnp.List }

// NewInode creates a new list of Inode.
func NewInode_List(s *capnp.Segment, sz int32) (Inode_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 24, PointerCount: 3}, sz)
	return Inode_List{l}, err
}

func (s Inode_List) At(i int) Inode { return Inode{s.List.Struct(i)} }

func (s Inode_List) Set(i int, v Inode) error { return s.List.SetStruct(i, v.Struct) }

// Inode_Promise is a wrapper for a Inode promised by a client call.
type Inode_Promise struct{ *capnp.Pipeline }

func (p Inode_Promise) Struct() (Inode, error) {
	s, err := p.Pipeline.Struct()
	return Inode{s}, err
}

func (p Inode_Promise) Attributes() Inode_attributes_Promise {
	return Inode_attributes_Promise{p.Pipeline}
}

// Inode_attributes_Promise is a wrapper for a Inode_attributes promised by a client call.
type Inode_attributes_Promise struct{ *capnp.Pipeline }

func (p Inode_attributes_Promise) Struct() (Inode_attributes, error) {
	s, err := p.Pipeline.Struct()
	return Inode_attributes{s}, err
}

func (p Inode_attributes_Promise) Dir() SubDir_Promise {
	return SubDir_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p Inode_attributes_Promise) File() File_Promise {
	return File_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p Inode_attributes_Promise) Link() Link_Promise {
	return Link_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p Inode_attributes_Promise) Special() Special_Promise {
	return Special_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type Dir struct{ capnp.Struct }

// Dir_TypeID is the unique identifier for the type Dir.
const Dir_TypeID = 0x8a228653b964fd48

func NewDir(s *capnp.Segment) (Dir, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 5})
	return Dir{st}, err
}

func NewRootDir(s *capnp.Segment) (Dir, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 5})
	return Dir{st}, err
}

func ReadRootDir(msg *capnp.Message) (Dir, error) {
	root, err := msg.RootPtr()
	return Dir{root.Struct()}, err
}

func (s Dir) String() string {
	str, _ := text.Marshal(0x8a228653b964fd48, s.Struct)
	return str
}

func (s Dir) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Dir) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Dir) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Dir) SetName(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Dir) Location() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s Dir) HasLocation() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Dir) LocationBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s Dir) SetLocation(v string) error {
	return s.Struct.SetText(1, v)
}

func (s Dir) Contents() (Inode_List, error) {
	p, err := s.Struct.Ptr(2)
	return Inode_List{List: p.List()}, err
}

func (s Dir) HasContents() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Dir) SetContents(v Inode_List) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewContents sets the contents field to a newly
// allocated Inode_List, preferring placement in s's segment.
func (s Dir) NewContents(n int32) (Inode_List, error) {
	l, err := NewInode_List(s.Struct.Segment(), n)
	if err != nil {
		return Inode_List{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

func (s Dir) Parent() (string, error) {
	p, err := s.Struct.Ptr(3)
	return p.Text(), err
}

func (s Dir) HasParent() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Dir) ParentBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	return p.TextBytes(), err
}

func (s Dir) SetParent(v string) error {
	return s.Struct.SetText(3, v)
}

func (s Dir) Size() uint64 {
	return s.Struct.Uint64(0)
}

func (s Dir) SetSize(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s Dir) Aclkey() (string, error) {
	p, err := s.Struct.Ptr(4)
	return p.Text(), err
}

func (s Dir) HasAclkey() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s Dir) AclkeyBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(4)
	return p.TextBytes(), err
}

func (s Dir) SetAclkey(v string) error {
	return s.Struct.SetText(4, v)
}

func (s Dir) ModificationTime() uint32 {
	return s.Struct.Uint32(8)
}

func (s Dir) SetModificationTime(v uint32) {
	s.Struct.SetUint32(8, v)
}

func (s Dir) CreationTime() uint32 {
	return s.Struct.Uint32(12)
}

func (s Dir) SetCreationTime(v uint32) {
	s.Struct.SetUint32(12, v)
}

// Dir_List is a list of Dir.
type Dir_List struct{ capnp.List }

// NewDir creates a new list of Dir.
func NewDir_List(s *capnp.Segment, sz int32) (Dir_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 5}, sz)
	return Dir_List{l}, err
}

func (s Dir_List) At(i int) Dir { return Dir{s.List.Struct(i)} }

func (s Dir_List) Set(i int, v Dir) error { return s.List.SetStruct(i, v.Struct) }

// Dir_Promise is a wrapper for a Dir promised by a client call.
type Dir_Promise struct{ *capnp.Pipeline }

func (p Dir_Promise) Struct() (Dir, error) {
	s, err := p.Pipeline.Struct()
	return Dir{s}, err
}

type UserGroup struct{ capnp.Struct }

// UserGroup_TypeID is the unique identifier for the type UserGroup.
const UserGroup_TypeID = 0xee5217621d9cbb3a

func NewUserGroup(s *capnp.Segment) (UserGroup, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return UserGroup{st}, err
}

func NewRootUserGroup(s *capnp.Segment) (UserGroup, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return UserGroup{st}, err
}

func ReadRootUserGroup(msg *capnp.Message) (UserGroup, error) {
	root, err := msg.RootPtr()
	return UserGroup{root.Struct()}, err
}

func (s UserGroup) String() string {
	str, _ := text.Marshal(0xee5217621d9cbb3a, s.Struct)
	return str
}

func (s UserGroup) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s UserGroup) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s UserGroup) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s UserGroup) SetName(v string) error {
	return s.Struct.SetText(0, v)
}

func (s UserGroup) IyoId() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s UserGroup) HasIyoId() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s UserGroup) IyoIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s UserGroup) SetIyoId(v string) error {
	return s.Struct.SetText(1, v)
}

func (s UserGroup) IyoInt() uint64 {
	return s.Struct.Uint64(0)
}

func (s UserGroup) SetIyoInt(v uint64) {
	s.Struct.SetUint64(0, v)
}

// UserGroup_List is a list of UserGroup.
type UserGroup_List struct{ capnp.List }

// NewUserGroup creates a new list of UserGroup.
func NewUserGroup_List(s *capnp.Segment, sz int32) (UserGroup_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return UserGroup_List{l}, err
}

func (s UserGroup_List) At(i int) UserGroup { return UserGroup{s.List.Struct(i)} }

func (s UserGroup_List) Set(i int, v UserGroup) error { return s.List.SetStruct(i, v.Struct) }

// UserGroup_Promise is a wrapper for a UserGroup promised by a client call.
type UserGroup_Promise struct{ *capnp.Pipeline }

func (p UserGroup_Promise) Struct() (UserGroup, error) {
	s, err := p.Pipeline.Struct()
	return UserGroup{s}, err
}

type ACI struct{ capnp.Struct }

// ACI_TypeID is the unique identifier for the type ACI.
const ACI_TypeID = 0xe7b4959415dabf9c

func NewACI(s *capnp.Segment) (ACI, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return ACI{st}, err
}

func NewRootACI(s *capnp.Segment) (ACI, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return ACI{st}, err
}

func ReadRootACI(msg *capnp.Message) (ACI, error) {
	root, err := msg.RootPtr()
	return ACI{root.Struct()}, err
}

func (s ACI) String() string {
	str, _ := text.Marshal(0xe7b4959415dabf9c, s.Struct)
	return str
}

func (s ACI) Uname() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s ACI) HasUname() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s ACI) UnameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s ACI) SetUname(v string) error {
	return s.Struct.SetText(0, v)
}

func (s ACI) Gname() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s ACI) HasGname() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s ACI) GnameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s ACI) SetGname(v string) error {
	return s.Struct.SetText(1, v)
}

func (s ACI) Mode() uint16 {
	return s.Struct.Uint16(0)
}

func (s ACI) SetMode(v uint16) {
	s.Struct.SetUint16(0, v)
}

func (s ACI) Rights() (ACI_Right_List, error) {
	p, err := s.Struct.Ptr(2)
	return ACI_Right_List{List: p.List()}, err
}

func (s ACI) HasRights() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s ACI) SetRights(v ACI_Right_List) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewRights sets the rights field to a newly
// allocated ACI_Right_List, preferring placement in s's segment.
func (s ACI) NewRights(n int32) (ACI_Right_List, error) {
	l, err := NewACI_Right_List(s.Struct.Segment(), n)
	if err != nil {
		return ACI_Right_List{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

func (s ACI) Id() uint32 {
	return s.Struct.Uint32(4)
}

func (s ACI) SetId(v uint32) {
	s.Struct.SetUint32(4, v)
}

// ACI_List is a list of ACI.
type ACI_List struct{ capnp.List }

// NewACI creates a new list of ACI.
func NewACI_List(s *capnp.Segment, sz int32) (ACI_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3}, sz)
	return ACI_List{l}, err
}

func (s ACI_List) At(i int) ACI { return ACI{s.List.Struct(i)} }

func (s ACI_List) Set(i int, v ACI) error { return s.List.SetStruct(i, v.Struct) }

// ACI_Promise is a wrapper for a ACI promised by a client call.
type ACI_Promise struct{ *capnp.Pipeline }

func (p ACI_Promise) Struct() (ACI, error) {
	s, err := p.Pipeline.Struct()
	return ACI{s}, err
}

type ACI_Right struct{ capnp.Struct }

// ACI_Right_TypeID is the unique identifier for the type ACI_Right.
const ACI_Right_TypeID = 0xe615914de76be38f

func NewACI_Right(s *capnp.Segment) (ACI_Right, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return ACI_Right{st}, err
}

func NewRootACI_Right(s *capnp.Segment) (ACI_Right, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return ACI_Right{st}, err
}

func ReadRootACI_Right(msg *capnp.Message) (ACI_Right, error) {
	root, err := msg.RootPtr()
	return ACI_Right{root.Struct()}, err
}

func (s ACI_Right) String() string {
	str, _ := text.Marshal(0xe615914de76be38f, s.Struct)
	return str
}

func (s ACI_Right) Right() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s ACI_Right) HasRight() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s ACI_Right) RightBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s ACI_Right) SetRight(v string) error {
	return s.Struct.SetText(0, v)
}

func (s ACI_Right) Usergroupid() uint16 {
	return s.Struct.Uint16(0)
}

func (s ACI_Right) SetUsergroupid(v uint16) {
	s.Struct.SetUint16(0, v)
}

// ACI_Right_List is a list of ACI_Right.
type ACI_Right_List struct{ capnp.List }

// NewACI_Right creates a new list of ACI_Right.
func NewACI_Right_List(s *capnp.Segment, sz int32) (ACI_Right_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return ACI_Right_List{l}, err
}

func (s ACI_Right_List) At(i int) ACI_Right { return ACI_Right{s.List.Struct(i)} }

func (s ACI_Right_List) Set(i int, v ACI_Right) error { return s.List.SetStruct(i, v.Struct) }

// ACI_Right_Promise is a wrapper for a ACI_Right promised by a client call.
type ACI_Right_Promise struct{ *capnp.Pipeline }

func (p ACI_Right_Promise) Struct() (ACI_Right, error) {
	s, err := p.Pipeline.Struct()
	return ACI_Right{s}, err
}

const schema_ae9223e76351538a = "x\xda\xacVml\x14\xd5\x1a~\x9fsfwv\x93" +
	"\xed\xdd\x9d;K\xf8\xc8m\xf6^\xee\xbd\x894\xb6\xa1" +
	"\x05\x8c4\x9a\x85\x8aJ\x09DN\x17b4\xfcp\xba" +
	";m\x87\xdd\xeenvg\xc5\x12M\x05\x11\xa1\x91\x18" +
	"H\xfd\xc0@\x04$\x06\xfehD\xfdA\x8c\xf1\xe3\x8f" +
	"\xd1(\x09D\x124`L\x14\x1b\x8c\x12\x13!BB" +
	"\x1d\xf3\xee\xb4\xb3\xdbZ\xfe\xf9\xef\xcc3\xefy\x9f\xe7" +
	"<\xef{\xde\x99\xa5w\xc9U\xa23\xf4\\\x88H\xad" +
	"\x0a\x85\xbd\xed\xe76<t\xe1\\\xd7^2\x92\xc2\xcb" +
	"\xdf3t\xe6\xe5\x13\xeeE\"\x98'\xc5\xe7\xe6{B" +
	"'2\xdf\x16\xa3\x04o\xedd\xeetf\xf7\xe21R" +
	"1\x08o,\xa3\xb2\x13\xff=\xf0&\x85B\x1crE" +
	"\xec4\x7f\xe1\xe0eW\xc4\xa7 \xfc~\x09\xbf\x9d\xf9" +
	"\xcf\xf1\xe3F\x0cM\xa1\xe0\xd0\xb3\xda\xab\xe6\xd7\x1a\xaf" +
	"\xceki\x827~\xf9\xea;;^\x13\x1fq^\xd9" +
	"\x14,9\xe4\xbav\xc0\x9c\xe4\xe0e7\xb5\x87A\xf0" +
	"\xdan|\xe2<\x939v\x9eff\xae\xeb\\\x19>" +
	"e\xae\x0e\xf3\xea\xde\xf06B\xe34*\x86\xbf\xc88" +
	"\x18>f\x1e\x0d\xcf'2O\xd6\x83\xb7\xbcd\xbdq" +
	"\xf9\xc8\xc2\x1fh\x0e\xc9!}\xccl\xd1y\x15\xd5Y" +
	"\xf2\x0b\xdf\xe7'6\xec\x9f\xf7#\xa9\x04\xe0\x1d\xfa\xf0" +
	"\x9by\xe3/\xbe;1\x15\xbcD?ev\xd6\x83\xdb" +
	"uN\x1c\xbc\x9e\xa5\xa2~\xbe}\xfaNs\xbf>\x9f" +
	"h\xd9A=\xc5\xe7[\xf4\xc8\xb5\xe5\xbb_\x9f\xfcy" +
	"N\xcd\xa7#c\xe6\xc7\x11^}\x10\xe1\xd4\xdd\xef\x1f" +
	"j\xed\x9f\xdfwuvp\xdd\x8d\xf6\xe8)sE\x94" +
	"W\x9d\xd1\xb7\x08^\xec\xdb\xee\xf2\xcaZ\xf5&\xa9\x7f" +
	"B6\\\xdf,uh\xd0\xcc\xcf\xa2?\x11\xcc/\xa3" +
	"\x13\xd4\xee\x0d\x97rv\xa1#k\x89r\xb1\xdc\x9d)" +
	"\xdbY\xc7*tl\x1a)\xdbD\x1b\x01\x95\x84 2" +
	"Vt\x13\x01F{\x17\x11\x84\xf1\xff\x1e\"H\xa3u" +
	"\x1d\x114ca\x0fQ\xbaZ\xca\xe6m7\xd5_(" +
	"e\xf3\xa3\xd9!\xab\x92\xb3\x1f\xf7\x06\x9c\x81R\xd9\xe1" +
	"L4Z+\xe6\x8b\xa5m\xc5\x80\x0eL\xb7\xc6\xa9\xd4" +
	"I\xfe-5\"\x0dD\xc6\xd96\"\xf5\x85\x84\xba " +
	"`\x00I0x~\x1d\x91\xfaJB}'`\x08\xe1" +
	"K\xba\xc4\xe0E\x09uC\xc0\x902\x09Id\\\xef" +
	"&R\xbfJ\xa8[\x02\xd0\x92\xd0\x88\x8c\x9b\x9c\xf2\x9a" +
	"D\x1f\x04\x8c\x90\x96D\x88\xc8\x98\xe4\xc0\x1b\x12\x19\x8d" +
	"\xd1\xb0H\"Ld\x02cD\x19\x0d\x12\x99\x04\xe3\xba" +
	"L\xd6k\xd1\x82\xadD\x99\x18\xe3\x0b \x10/Z\xc3" +
	"6b$\x10#x\x85R\xd6r\x9dR\x91\x88\x02," +
	"[*\xbav\xd1\xad2\xf6\x0f\xc2F\x09$\x1ae " +
	"0\x98.[\x15\xbb\xe8N\xef\x89W\x9d\xed6\xa2$" +
	"\x10%\xa4\xadl!o\x8f\x04\xf9\x86K9g\xc0\xc9" +
	"Z`\xa2M\xce\xb0M\x84\x08\x09D\x98\xabb\xfb\xfc" +
	"q~\x11\xc03\x8c\xce\xd4\xfa\xd7H\xa7\xc2^k\x81" +
	"\xd7-\x8b\x89TDB%\x05\xf4Yl\x8d\xad\xbd\xc5" +
	"R\x0e6\xef\\\x10\xec<\xc8\x96\x8eK\xa8#\x02\xd3" +
	"E:\xcc\xd8+\x12\xea\xb8\x00\x04\x9az\xd08\xfa(" +
	"\x09C\xfa\x1e\x1b\xfb\xd8\xf9=\x12j\\\xc0\xd0|\x83" +
	"\x8d\xfdc\x8d\x84\xf5\x1aE8\xe3V\"uHB\x9d" +
	"\x98e\xf9\x0c\xab<\xcbu+N\x7f\xcd%iW\xff" +
	"n\xdf\x1ep\x0av\xaa\x87\x9b\x9a\x0d\x88\x04\x06,\xe1" +
	"\xc3\xfeOB-mj\xd3v\xf6\xf3\x0e\x09\xb5\\ " +
	">dU\x87\xd0B\x02-\xe4\x9b;\xb5\x9eU\x97\xb2" +
	"\x9d\xd5\x1d\xab\xa04\xa0i:\xa3-\xce7\xf06\x84" +
	"\x01_[\x13\x9f;R\xb6\x11o\xe4  N\x88\xe7" +
	",\xd7\x9a\x9bz\xbdS\xcc\xfbW\xbc\xa9#\xba\x1b\x1d" +
	"\x91v\xad\xca\xa0\xed\xce\xdd\x14\xab\xef\xeb\xedH\xf59" +
	"\x83C\xee,_\xba\xe6\x90\xd9O\xa4\xee\x94Pw\x0b" +
	"\xa4*\xbc'\xc8Y\xab\xda\x95\xc1J\xa9Fz\xd9\xc9" +
	"A'\x01}\x0e&\xaa\xbb\x13\xcca\x03]>\xb5J" +
	"\x06\xbcO1\xef\x13\x12jWS=v0\xf8\xa4\x84" +
	"\xda\xc3\x1d\xe9O\x8dg\xd9\xb3\xa7%\xd4\xf3<5\x84" +
	"?5\xf6\xf2\xb1wMu\x9f\x06\x7fl\x1c^\xd4\xe8" +
	"\xe7T\xad\xb9\xfdR\x833\x9a\x91\xe5NkO\xd7\xcf" +
	"Wm\xdc\xf9@\xb5\x7f\xe7\xa5\x93\xbb}\xa3\xf9\xd5h" +
	"2\xb3\xafQ\xdf\xe93uvO\x99\xb9V\xc0\xabO" +
	"\xdb\x8c\xb3\x9d`#L\x02aB\xba\x8e5\x09\x08\xbe" +
	"\xa1\xbe\x80\x99\xac\x9b\xabv%\xf5`\xa5T+3s" +
	",`\xbe\x9f]Z%\xa1\xd67\xd9\xd9\xcbv\xae\x91" +
	"P\x1b\x1bvn`9k%\xd4\xa6YW4\xe5\x8c" +
	"\x94zs\xd3Oi~*\xba\xc1\x95\x9d\xf1\xcd\xe1\x01" +
	"cw\xf8\xb78^s\xed\xaaJH\xed_\x9e\x07\x9f" +
	"\xc2\xe2[\xb5EB\x0d\x09\xb4\xe2\x0f\x86\xb9f6K" +
	"|LB\x15\x04Z\xc5\xa47U5\x87\xe1\x9c\x84*" +
	"\x0b\xb4\xca[\x0c\xf3\xb8\x1f\xee!RC\x12\xca\x15\xd0" +
	"sN\x05\x89\xe9\x9f\x16\x02\x12\x84\xf8\x80S\xb0\x91h" +
	"|\x91\xa7\xe0\x82S\xcc#\xd1\xf8_\xf0\xe1\xd1\xaa\xff" +
	"\x99D\xa2\xf9'\x8a\xdf\xfc\x19\x00\x00\xff\xffO)2" +
	"\x1b"

func init() {
	schemas.Register(schema_ae9223e76351538a,
		0x8932d2d84f4dd27a,
		0x8a228653b964fd48,
		0xa4a421ce00f301dd,
		0xc0029f81b3eee594,
		0xd5a2538369c2f82a,
		0xdc74a897ce683c6b,
		0xe419a0e5a661965c,
		0xe615914de76be38f,
		0xe7b4959415dabf9c,
		0xecfda38634f4591a,
		0xee5217621d9cbb3a,
		0xf9737539703ade0c)
}
