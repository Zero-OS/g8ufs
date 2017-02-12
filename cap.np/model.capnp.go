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

const schema_ae9223e76351538a = "x\xda\xacVo\x88T\xd5\x1b~\x9fs\xee\xfcY\x98" +
	"u\xe6\xfe\xee\x88\x7f\xf8-SVP\xd2n\xba*\xd4" +
	"`\x8c\x9a\x95+J\x9e\x1d%\x0a?tw\xe6\xee\xee" +
	"ufg\x86;w\xb2\x95dSTt)B\xd9\xfe" +
	"\xa2\xa4\xe5\x86~\xa8\xc0\xfa\x12!\xd5\xa7(JP\x0a" +
	"$4\x82\xb2\x85\x08\x092Rp\xbb\xf1\xce\xdd\xbdw" +
	"v\xda\x8f}\xbb\xf7\xb9\xefy\x9f\xe7<\xef\xfb\x9es" +
	"W\xc4\xe5:\xb12\xf2@\x84H\xad\x8dD\xbd=\x97" +
	"\xb6>q\xf9R\xef\x11\xd2\xd3\xc2+\xad\x1d\xbe\xf0\xda" +
	"\x19\xf7\x0a\x11\x8c\xbd\xe2+\xe3\x88\x88\x11\x19\x07\xc5\x18" +
	"\xc1\xdb4]\xfc8\x7fh\xd98\xa9\x04\x847\x9eW" +
	"\x85\xa9\xbb\x8e\xbdO\x91\x08\x87\x9c\x17\xfb\x8d\xcf9x" +
	"\xd5y\xf1\x05\x08\x7f]\xc5\x1f\x17\xee<}ZO\xa0" +
	"%\x14\x1c:\xa9\xbdi\xbc\xa7\xf1\xd3Y-G\xf0&" +
	"\xae]\xffp\xdf[\xe23\xce+[\x82%\x87|\xa9" +
	"\x1d3.r\xf0\xaao\xb4'A\x08\x15\xaa\x04\xfe\x95" +
	"\xba3\xfa\xb6\xb10\xba\x88\xc8\xe8\x8a\xee&x;_" +
	"5\xdf\xbdvr\xc9\xcf4\x8f\x8c\xd1\xe8\xb8\xb1/\xca" +
	"O{\xa3,\xe3\xe5\x9fJS[\x8f.\xfc\x85T\x0a" +
	"\xf0\x8e\x7f\xfa\xfd\xc2\x89W>\x9a\x9a\x09>\x11=g" +
	"L6\x83O5\x13\x07\x9f\xdbT45#\xb6\xdf\x88" +
	"\xc4\x16\x11\xad\xea\x8ceX\xf3\xd2\xa7n\xac>\xf4\xce" +
	"\xf4o\xf3j\xee\x8e\x8f\x1bk\xe2\xfc\xb42\xce\xa9\xb3" +
	"\x9f\x1c\xef\x1aX\xd4\x7f\xbd=\xb8Y\x89S\xf1s\xc6" +
	"\xd9f\xf0d\xfc\x03\x82\x97\xf8![{\xa8Q\xbfE" +
	"\xea\x7f\x90\xa1\x93;d\x0c\x1a4\xe3\xe1\x8e_\x09\xc6" +
	"\xfa\x8e)\xea\xf6F\xaaE\xab\xdcS0E\xadR\xcb" +
	"\xe6kV\xc16\xcb=\xdbGk\x16\xd16@\xa5!" +
	"\x88\xf45Y\"@\xef\xee%\x82\xd0\xef\xd9@\x04\xa9" +
	"wm&\x82\xa6/\xd9@\x94\xabW\x0b%\xcb\xcd\x0c" +
	"\x94\xab\x85\xd2Xa\xd8t\x8a\xd6\xb3\xde\xa0=X\xad" +
	"\xd9\x9c\x89\xc6\x1a\x95R\xa5\xba\xbb\x12\xd0\x81\xe96\xda" +
	"N\x93\xe4\x0e\xa9\x11i \xd2/.'R_K\xa8" +
	"\xcb\x02:\x90\x06\x83\xdfm&R\xdfJ\xa8\x1f\x05t" +
	"!|IW\x19\xbc\"\xa1n\x0a\xe8R\xa6!\x89\xf4" +
	"?\xb3D\xeaw\x09u[\x00Z\x1a\x1a\x91~\x8bS" +
	"\xde\x90\xe8\x87\x80\x1e\xd1\xd2\x88\x10\xe9\xd3\x1cxS\"" +
	"\xaf1\x1a\x15iD\xb9B\x18'\xcak\x90\xc8\xa7\x18" +
	"\x8f\xc9\xb4\xdf?\xd8E\x94O0\xbe\x18\x02\xc9\x8a9" +
	"b!A\x02\x09\x82W\xae\x16L\xd7\xaeV\x88(\xc0" +
	"\x0a\xd5\x8akU\xdc:c\x0b\x08\xdb$\x90\x0a\xcb@" +
	"`0W3\x1d\xab\xe2\xce\xaeI\xd6\xed=\x16:H" +
	"\xa0\x83\x903\x0b\xe5\x925\x1a\xe4\x1b\xa9\x16\xedA\xbb" +
	"`\x82\x89\xb6\xdb#\x16\x11\xe2$\x10g.\xc7\xf2\xf9" +
	"\x93\xfc!\x80\xe7\x18\x9do\x0cl\x94\xb6\xc3^k\x81" +
	"\xd7\x9d\xcb\x88T\\B\xa5\x05bml\xe1\xd2\xbeJ" +
	"\xb5\x08\x8bW.\x0eV\xbe\xc1\x96NH\xa8\x93\x02\xb3" +
	"E:\xc1\xd8\xeb\x12\xea\xb4\x00\x04ZzP?\xf54" +
	"\x09]\xfa\x1e\xeb/\xb1\xf3\x87%\xd4\x84\x80\xae\xf9\x06" +
	"\xebG\xc7\xc3\x84\xcd\x1a\xc59\xe3.\"u\\B\x9d" +
	"i\xb3|\x8eU\x9e\xe9\xba\x8e=\xd0pIZ\xf5\xff" +
	"\xdc\xb7\x9aU\x88\xd9fYi@\xcb\x89\x88\xe5I\x9e" +
	"\x10\x15\x0f\x1c\xb9\x8fw\x7f\xb7\x84Z\x11:\xd2\xcd\xd8" +
	"\xbd\x12j\xb5@\xd2\x1d\xadYH\x869\x08H\x12\x92" +
	"E\xd35\xd1I\x02\x9d\xed\xd4[\xecJ\xc9\x1f\xc1\x96" +
	"\x8ae\xc3\x8a\xe5\\\xd3\x19\xb2\xdc\xf9\x8b\xb6\xfe\x91\xbe" +
	"\x9eL\xbf=4\xecr\x82\x16\x99\xbd\xf3\xc8\x1c R" +
	"\xf7K\xa8\x07\x052\x0e\xaf\x09r6\xea\x963\xe4T" +
	"\x1b\x14\xab\xd9E\xc4H 6\x0f\x135\xdd\x09\xceI" +
	"\x1d\xbd>\xb5J\x07\xbc{\x99\xf79\x09u\xa0e\xac" +
	"\xf71\xf8\xbc\x84:\xcc\x1d\xe3O\xf5A\xf6\xec\x05\x09" +
	"\xf5\"O\xb5\xf0\xa7\xfa\x08o\xfb\xc0Lwh\xf0\xc7" +
	"\xfa\xc4\xd2\xb0\xdf2\x8d\xd6\xf6\xc8\x0c\xcdi\x16\x96;" +
	"\xab=\xd7\xdc_=\x9c\xc9@\xb5?\x93\xd2.\xce\xdf" +
	"\x08\x8f\xd9\xe5\x99\x03\xb1\xc5\xcc\xfe\xb0\xbe\xb3{Z\x99" +
	"\x9d1s\x93\x80\xd7<\x0d\xf3\xf6\x1e\x82\x85(\x09D" +
	"\x09\xb9&\x16\x08\xe0\xca/h\xe7\xdaQ\xb7\x9c\xcc\xe3" +
	"N\xb5Qc\xbeD\xc0\xf7({\xb3NBmi1" +
	"\xb1\x8fM\xdc(\xa1\xb6\x85&ne\x11\x9b$\xd4\xf6" +
	"\xb6\xc1\xc9\xd8\xa3\xd5\xbe\xe2\xec[\x8e\xdf*n0H" +
	"sn\x02\x1e{\xab\xc7\x9f\xadd\xc3\xb5\xea*%\xb5" +
	"\xff{\x1e|\x0a\x93\xcf\x8e\x9d\x12jX\xa0\x0b\x7f3" +
	"\xcc\x95\xb2X\xe23\x12\xaa,\xd0%\xa6\xbd\x99Z\xd9" +
	"\x0c\x17%TM\xa0K\xdef\x98\x0f\xe1\x91\x0dDj" +
	"XB\xb9\x02\xb1\xa2\xed 5\xfb{@@\x8a\x90\x1c" +
	"\xb4\xcb\x16R\xe1=9\x03\x97\xedJ\x09\xa9\xf0\x16\xf7" +
	"\xe1\xb1\xba\x7fy!\xd5\xfa\xbb\xc2_\xfe\x09\x00\x00\xff" +
	"\xff\x08~\x0b\x1f"

func init() {
	schemas.Register(schema_ae9223e76351538a,
		0x8932d2d84f4dd27a,
		0x8a228653b964fd48,
		0xa4a421ce00f301dd,
		0xc0029f81b3eee594,
		0xdc74a897ce683c6b,
		0xe419a0e5a661965c,
		0xe615914de76be38f,
		0xe7b4959415dabf9c,
		0xecfda38634f4591a,
		0xee5217621d9cbb3a,
		0xf9737539703ade0c)
}
