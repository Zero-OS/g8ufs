package np

// AUTO GENERATED - DO NOT EDIT

import (
	strconv "strconv"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type File struct{ capnp.Struct }

// File_TypeID is the unique identifier for the type File.
const File_TypeID = 0xecfda38634f4591a

func NewFile(s *capnp.Segment) (File, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return File{st}, err
}

func NewRootFile(s *capnp.Segment) (File, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
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

func (s File) Blocks() (capnp.DataList, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.DataList{List: p.List()}, err
}

func (s File) HasBlocks() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s File) SetBlocks(v capnp.DataList) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewBlocks sets the blocks field to a newly
// allocated capnp.DataList, preferring placement in s's segment.
func (s File) NewBlocks(n int32) (capnp.DataList, error) {
	l, err := capnp.NewDataList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.DataList{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s File) Size() uint64 {
	return s.Struct.Uint64(8)
}

func (s File) SetSize(v uint64) {
	s.Struct.SetUint64(8, v)
}

// File_List is a list of File.
type File_List struct{ capnp.List }

// NewFile creates a new list of File.
func NewFile_List(s *capnp.Segment, sz int32) (File_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1}, sz)
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
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Link{st}, err
}

func NewRootLink(s *capnp.Segment) (Link, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
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

func (s Link) DestDirKey() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Link) HasDestDirKey() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Link) DestDirKeyBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Link) SetDestDirKey(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Link) DestName() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s Link) HasDestName() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Link) DestNameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s Link) SetDestName(v string) error {
	return s.Struct.SetText(1, v)
}

// Link_List is a list of Link.
type Link_List struct{ capnp.List }

// NewLink creates a new list of Link.
func NewLink_List(s *capnp.Segment, sz int32) (Link_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
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
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 2})
	return Inode{st}, err
}

func NewRootInode(s *capnp.Segment) (Inode, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 2})
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

func (s Inode) Attributes() Inode_attributes { return Inode_attributes(s) }

func (s Inode_attributes) Which() Inode_attributes_Which {
	return Inode_attributes_Which(s.Struct.Uint16(0))
}
func (s Inode_attributes) Dir() (SubDir, error) {
	p, err := s.Struct.Ptr(1)
	return SubDir{Struct: p.Struct()}, err
}

func (s Inode_attributes) HasDir() bool {
	if s.Struct.Uint16(0) != 0 {
		return false
	}
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Inode_attributes) SetDir(v SubDir) error {
	s.Struct.SetUint16(0, 0)
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewDir sets the dir field to a newly
// allocated SubDir struct, preferring placement in s's segment.
func (s Inode_attributes) NewDir() (SubDir, error) {
	s.Struct.SetUint16(0, 0)
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
	if s.Struct.Uint16(0) != 1 {
		return false
	}
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Inode_attributes) SetFile(v File) error {
	s.Struct.SetUint16(0, 1)
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewFile sets the file field to a newly
// allocated File struct, preferring placement in s's segment.
func (s Inode_attributes) NewFile() (File, error) {
	s.Struct.SetUint16(0, 1)
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
	if s.Struct.Uint16(0) != 2 {
		return false
	}
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Inode_attributes) SetLink(v Link) error {
	s.Struct.SetUint16(0, 2)
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewLink sets the link field to a newly
// allocated Link struct, preferring placement in s's segment.
func (s Inode_attributes) NewLink() (Link, error) {
	s.Struct.SetUint16(0, 2)
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
	if s.Struct.Uint16(0) != 3 {
		return false
	}
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Inode_attributes) SetSpecial(v Special) error {
	s.Struct.SetUint16(0, 3)
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewSpecial sets the special field to a newly
// allocated Special struct, preferring placement in s's segment.
func (s Inode_attributes) NewSpecial() (Special, error) {
	s.Struct.SetUint16(0, 3)
	ss, err := NewSpecial(s.Struct.Segment())
	if err != nil {
		return Special{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s Inode) Aclkey() uint32 {
	return s.Struct.Uint32(4)
}

func (s Inode) SetAclkey(v uint32) {
	s.Struct.SetUint32(4, v)
}

func (s Inode) ModificationTime() uint32 {
	return s.Struct.Uint32(8)
}

func (s Inode) SetModificationTime(v uint32) {
	s.Struct.SetUint32(8, v)
}

func (s Inode) CreationTime() uint32 {
	return s.Struct.Uint32(12)
}

func (s Inode) SetCreationTime(v uint32) {
	s.Struct.SetUint32(12, v)
}

// Inode_List is a list of Inode.
type Inode_List struct{ capnp.List }

// NewInode creates a new list of Inode.
func NewInode_List(s *capnp.Segment, sz int32) (Inode_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 2}, sz)
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
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 4})
	return Dir{st}, err
}

func NewRootDir(s *capnp.Segment) (Dir, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 4})
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

func (s Dir) Aclkey() uint32 {
	return s.Struct.Uint32(8)
}

func (s Dir) SetAclkey(v uint32) {
	s.Struct.SetUint32(8, v)
}

func (s Dir) IsLink() bool {
	return s.Struct.Bit(96)
}

func (s Dir) SetIsLink(v bool) {
	s.Struct.SetBit(96, v)
}

func (s Dir) ModificationTime() uint32 {
	return s.Struct.Uint32(16)
}

func (s Dir) SetModificationTime(v uint32) {
	s.Struct.SetUint32(16, v)
}

func (s Dir) CreationTime() uint32 {
	return s.Struct.Uint32(20)
}

func (s Dir) SetCreationTime(v uint32) {
	s.Struct.SetUint32(20, v)
}

// Dir_List is a list of Dir.
type Dir_List struct{ capnp.List }

// NewDir creates a new list of Dir.
func NewDir_List(s *capnp.Segment, sz int32) (Dir_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 24, PointerCount: 4}, sz)
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

const schema_ae9223e76351538a = "x\xda\xa4U]\x8c\x13e\x17>\xcf\xfbN\xff\xbe\x14" +
	"\xda\xf9\xa6\x84\x9f\xb0\xe9\x07\xdf~\x09l\xd8\xfd\x96\x82" +
	"\x09\xd9\x90\x14qU\x16Y\xe5\xa5p\x01\xe1\x82\xd9v" +
	"\x96\x1d\xdamk;\x15\x97H\x16\x08\x10\xd9\x98\x18\x09" +
	"\xfe\xc4`$J\x8c\x17F\x12\xf4\x06\x0dQn\xbd\xe0" +
	"\x82+c\xe0\x82\xe8\xbaj\x84h\xc4(\x09\xeb\x983" +
	"mg\xba\x85;\xe7j\xe6y\xcf\x9c\xf3\x9cs\x9es" +
	"\xde\xc15r\x8bX\x1f\xfa\x7f\x88Hm\x0e\x85\xdd#" +
	"7F\x9f\xfb\xfaF\xe6\x0c\xe9)\xe1\x167O\\\x7f" +
	"\xf3C\xe7&\x11\x8c\xa3\xe2+\xe3\x8c\x88\x10\x19\xa7\xc4" +
	"4\xc1\xdd6_\xb8\x92;\xbdz\x86T\x1c\xd2\x9d\xc9" +
	"\xa9\xfc\xdc\x7f\xcf~L!\x8dM\xae\x8a\x13\xc656" +
	"\xdepU\xb8 \xfcq\x0b\xbf]_u\xf1\xa2\x1eG" +
	"\x87)\xd8\xf4\xb6\xf6\xb6\xf1\xa3\xf7\xd3\xac\x96%\xb8\xc6" +
	"\xcd\xd5\x1f\xf5\xc6?\xbbD\xea\xdf\x10\xee\xb9\xd9\xbb\x9f" +
	"\x1c\x7fW|I{D\x04\x1a\x91\x11\x0a\xfdD0b" +
	"\xa19Bp\xa8\xe2\x10\x1d~=\x96\xb7Bg\x8d\xd9" +
	"\x10S\xb8\x1dJ\x83\x10$\xa3\xe2x\x88\xc5l\xf8=" +
	"\xe3Nx)\x91\xf1{\xf80\xc1\xdd\xff\x86\xf9\xc1\xec" +
	"\x85\xe5\xdf\xd1B\xc6\x9eg\x15\x991\xf6F\xf8mO" +
	"\x84m_\xfd\xb687\xfa\xda\x92\xefI%\x01\xf7\xfc" +
	"\x17\xdf,9\xf7\xfa\xa7s-\xc7W\"\x97\x8dk\x9e" +
	"\xf1U\xcf\xd8?\xeeb!\xd9\xa4'z\xc2X\x15]" +
	"J\xb4am\xd4\xe3\xbcb\xef\xbd\x8d\xa7\xdf\x9f\xff\xb9" +
	";C\xcf\xf5Hl\xc6P1~\x1b\x8d]\"\xb8C" +
	"\x9f\x9f\xef\x19[\xba\xebn\xb7k\x8f\xf4\xfd\xd8e\x03" +
	"\xff\xe2\xb7\xf9\xd8%\xeaw'+\x05\xab4\x907E" +
	"\xb5\\\x1d\xcaU\xad\xbcm\x96\x06vOU-\xa2\x9d" +
	"\x80JA\x10\xe9\x8f\x0d\x11\x01z\x7f\x86\x08B\xff\xdf" +
	"V\"H\xbdg;\x114}\xf9V\xa2l\xbd\x92/" +
	"ZNz\xacT\xc9\x17\xa7\xf3\x13f\xad`\xbd\xe0\x8e" +
	"\xdb\xe3\x95\xaa\xcd\x9eh\xbaQ.\x96+\x87\xcb~8" +
	"p\xb8a\xbb\xe6\x05\xe9\x95\x1a\x91\x06\"\xfdN\x1f\x91" +
	"\xfaAB\xdd\x13\xd0\x81\x14\x18\xfcu;\x91\xfaEB" +
	"=\x10\xd0\x85hR\xba\xcf\xe0\x9f\x12\xb9\x14\x04t)" +
	"S\x90D\x86\x8e!\xa2\\\x1c\x12\xb9e\x10\x80\x96\xf2" +
	"\xc4\xb2\x04}D\xb9$\xc3+\xd9<$R\x08\x11\x19" +
	"\xcb=\xf3\x14\xe3\xffa<| \x850\xd7\xdf\xc3\x97" +
	"1\xde\xcbxDKy\x95^\x85\x19\xa2\\/\xe3\x83" +
	"\x8cGC)D\x89\x8c~\x1c\"\xca\xadc|\x13\x04" +
	"\x12es\xd2B\x9c\x04\xe2\x04\xb7T\xc9\x9b\x8e])" +
	"\x13\x91\x8f\xe5+e\xc7*;u\xc6\x16\x13vJ " +
	"\x19\xc8\x98\xc0`\xb6j\xd6\xac\xb2\xd3\xfe'Q\xb7\x8f" +
	"X\x88\x91@\x8c\x905\xf3\xa5\xa25\x85(\x09D\x09" +
	"Y\xbb\xbe\xc3.\x17\x01\x12`\xb9LV\x0a\xf6\xb8\x9d" +
	"7\xc1qw\xdb\x93\x16Q\xdb\xd4\xcd\xd7\xac&\x9d\x04" +
	"\x1f\xf8\xf0\x82\xc6\xe4\x1ac\xc3\xd2\xaeqo4\xbf7" +
	"\x8bV\x13\xa9\xa8\x84J\x09D8x;\x99\x05\x12\x1a" +
	")W\x0a\xd6\x80\xe985{,\xd1p\xac\xbaJJ" +
	"m\xa5\xeb\xb6zi\xb2\x93\xfd\x12jB\xa0\x07\x7f1" +
	"\xcc\xdd\xb4\xb8\xef\x07$TI\xa0G\xcc3,\x89t" +
	"\x9b\xe1\x82\x84\xaa\x0a\xf4\xc8\x07\x0ckD\xfa\xe4V\"" +
	"5!\xa1\x1c\x81H\xc1\xae!\xd9\xde+\x04$\x09\x89" +
	"q\xbbd!\x19LM\x0b.q\x8d\x92\xc1L7\xe1" +
	"\xe9zS\xf5Hv\xee9>YX\x14\xce\x0c\x967" +
	"\x14~M\x8e2\xc1\x17%\xd4I\x01~\x82\xad\xa5\x1f" +
	"\xdfGB\x17\xf0\x84\xa6O\x0e\x05\x8cu)<\x95\xe9" +
	"\xcf\xcf\x10)GB\x1d\x13\xd05\xe9IL?z\x88" +
	"H\xbd$\xa1^\xee\xd6Q\xb3\xa6\x0d\x87\xa4U\xef\xea" +
	"\xff?ox\xd5\xcaGl\xb3\xa44\xa0c\xf5\xa3/" +
	"\xc1\xab@E\xfd\x84\xd7r\xc2\xbd\x12j\x90\x13n\xf6" +
	"\xb4\x9f\xb15\x12j\xa3@\xc2\x99\xaaZH\x04>\x08" +
	"H\x10\x12\x05\xd31\xb1\x88\x04\x16u\x87f\xe16w" +
	"MG\x94}\x81G\x7f\x0d\xac\xe7\x89\x1f\x94P\x9b\x05" +
	"\xdc\x82Uw\x86\xed\xda3$;\x84\xc8\xe0\xb3&'" +
	"O\x0f\x89\xd3\x8b\xf5\xf8\x13#\x03\xe9]\xf6\xc1\x09\xa7" +
	"+^\xe6\x11Y\x8d\x11\xa9u\x12j\x93@\xba\xc6\xff" +
	"\xf8>\x1bu\xabv\xb0ViP\xa4j\x17\x10!\x81" +
	"\xc8#\"\x91WL\xffR\xd0\x91i\x86\xee\x94O&" +
	"\x90\x8f\x9f\xe7\xf1L\xa0\x00\xb4\xb6\xdd).\xf11\x09" +
	"\xf5JK?<\x1dgXT'%\xd4\x05\xd6Ok" +
	"6\xdeYA\xa4\xde\x92P\x17\x05\xd2\x8dN\x01\xa5\x0f" +
	"v~%\x98n\x9b{\xd6\xcb\xaf\x1e\xac\"\x9fus" +
	"\x15I\xbb\xf0h\xdd<e\x97Z\x17E\xdcO\xea\xc9" +
	"]DjXB\xed\x0c\x8a9\xcaT\xb7I\xa8\x02\xaf" +
	"\xf0\xd6\xd0\x9b}\xc1.p\xbd\xab#g\x1f!X\x08" +
	"\x93@\x98\x90\xf50\x9f\x15\xabgq\xd7\x1e\\\xc8f" +
	"O\xdd\xaa\xa5\x9f\xaeU\x1a\xd5.F\x1ch\x8b\x84\xda" +
	"\xd1Q\xe6\x91L\x07\xcdV\x99}\x9a\xbb\xbb\x86/m" +
	"OUF\x0a\xed\xaf,\x7f\x95\x9d6\x89\xbf\x03\x00\x00" +
	"\xff\xffG\xc6\x15*"

func init() {
	schemas.Register(schema_ae9223e76351538a,
		0x8932d2d84f4dd27a,
		0x8a228653b964fd48,
		0xa4a421ce00f301dd,
		0xafba0c24ac22dc13,
		0xc0029f81b3eee594,
		0xdc74a897ce683c6b,
		0xe419a0e5a661965c,
		0xe615914de76be38f,
		0xe7b4959415dabf9c,
		0xecfda38634f4591a,
		0xee5217621d9cbb3a)
}
