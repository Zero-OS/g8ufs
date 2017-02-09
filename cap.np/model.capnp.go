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

const schema_ae9223e76351538a = "x\xda\xa4U_l\x14\xd5\x17>\xdf\xbd3;\xddd" +
	"\xcb\xee\xfcfI\x81\xd0\xec\x0f~\xfd%\xda\xd0\x8a\x05" +
	"\x13\xd2\x90,\"*%\x10\xb9,<`x`\xba;" +
	"m\x87nw\xd7\xddY\xb1\xc4\xa6@\x80Hcb$" +
	"\xf8'\x06#Qb|0\x92\xa0>\xa0!\xea\xab\x0f" +
	"<\xf0\xa4\x06\x1e\x88\xd6\xaa\x11\xa2\x11\xa3&\xe0\x983" +
	"\xbb;\xb3]\xfaf\x9fv\xbe{\xee9\xdf9\xe7\xbb" +
	"_\xd7\xf7\xc8-\xe2a\xfd!\x9dHm\xd6c\xfe\x91" +
	"k\xbb\x9e\xfa\xea\xda\xd0i2\xd3\xc2\x9f\xdc<q\xf5" +
	"\xb5\xf7\xbc\xebD\xb0f\xc4\x97\xd6ia\x10Y'\xc5" +
	",\xc1\xdf~\xafp9wj\xed\x1c\xa9\x04\xa4?\x97" +
	"S\xf9\x85\xff\x9d\xf9\x80t\x8dC\xae\x88\xe3\xd6\x17\x1c" +
	"\xbc\xe1\x8a\xf0A\xf8\xe3\x06~\xbb\xba\xe6\xc2\x053\x81" +
	"\xb6Pp\xe8M\xed\x0d\xeb\xc7\xe0\xd2\xbc\x96%\xf8\xd6" +
	"\xf5\xb5\xef\xf7%>\xb9H\xea?\x10\xfe\xd9\xf9\xdb\x1f" +
	"\x1e{K|N\xfb\x84\x01\x8d\xc8\xd2\xf5\x9f\x08V\\" +
	"_ D\x87*\x01\xd1\x967`yC?c\xcd\xeb" +
	"L\xe1\xa6\x9e\x01!jF%p\x1f\x8b\xf9\xd8\xdb\xd6" +
	"\xadX\x0f\x91\xf5{\xec0\xc1?\xf0\xaa\xfd\xee\xfc\xf9" +
	"\x95\xdf\xd1\x12\x8c\x951g\xed7\xf8\xd7>\x83\x19\xbf" +
	"\xf4\xed\xe4\xc2\xae\x97\x97\x7fO*\x05\xf8\xe7>\xfbf" +
	"\xf9\xd9W>Zh\x06\xd7\x8dK\xd6L\x10<mp" +
	"\xe2\xf0\xb8\x83\x85\xe4\x90\xaf\x8d\xe3\xd6\x0d\xa3\x87h\xc3" +
	"\xbc\x11p^\xb5\xff\xce\xc6S\xef\xdc\xfb\xb9\xb3\xc3 " +
	"\xb5\x1e\x9f\xb3\xba\xe3\xfc+\x1e\xbfH\xf0\x87?=\xd7" +
	";\xda\xb3\xe7vg\xea`\x1c\x1f\xc7/YW\x82\xe0" +
	"\xcb\xf1\x8b4\xe0O\x95\x0bNq0o\x8bJ\xa92" +
	"\x9c\xab8y\xd7.\x0e\xee\x9d\xae8D\xbb\x01\x95\x86" +
	" 2\x1f\x19&\x02\xcc\x81!\"\x08\xf3\xff[\x89 " +
	"\xcd\xde\x1dD\xd0\xcc\x95[\x89\xb2\xb5r~\xd2\xf12" +
	"\xa3\xc5r~r6?aW\x0b\xce\xb3\xfe\x98;V" +
	"\xae\xb8\x9c\x89f\xeb\xa5\xc9R\xf9p),\x07.\xb7" +
	"\xcd\xad\x06E\xfa\xa4F\xa4\x81\xc8\xbc\xd5O\xa4~\x90" +
	"Pw\x04L \x0d\x06\x7f\xddA\xa4~\x91Pw\x05" +
	"L!\x1a\x94\xfeb\xf0O\x89\\\x1a\x02\xa6\x94iH" +
	"\"\xcb\xc40Q.\x01\x89\xdc\x0a\x08@K\x07bY" +
	"\x8e~\xa2\\\x8a\xe1\xd5\x1c\xae\x8b4t\"ke\x10" +
	"\x9ef\xfc\xbf\x8c\xc7\x0e\xa6\x11#\xb2z\x03|\x05\xe3" +
	"}\x8c\x1bZ:\x98\xf4\x1a\xcc\x11\xe5\xfa\x18_\xcfx" +
	"\x97\x9eF\x17\x915\x80CD\xb9u\x8co\x82@\xb2" +
	"dO9H\x90@\x82\xe0\x17\xcby\xdbs\xcb%\"" +
	"\x0a\xb1|\xb9\xe49%\xaf\xc6\xd82\xc2n\x09\xa4\"" +
	"\x19\x13\x18\xccV\xec\xaaS\xf2Zw\x925\xf7\x88\x83" +
	"8\x09\xc4\x09Y;_\x9ct\xa6\xd1E\x02]\x84\xac" +
	"[\xdb\xe9\x96&\x01\x12`\xb9L\x95\x0b\xee\x98\x9b\xb7" +
	"\xc1u\xf7\xbaS\x0eQ+\xd4\xcfW\x9d\x06\x9d$\x1f" +
	"\x84\xf0\xa2\xc5\xe4\xea\xa3\xdb\xa4[\xe5\xddh\xe1n\xba" +
	"\xd7\x12\xa9.\x09\x95\x160\xb8x\xab\x99E\x12\x1a)" +
	"\x95\x0b\xce\xa0\xedyUw4Y\xf7\x9c\x9aJIm" +
	"\xb5\xef7wis\x92\x03\x12jB\xa0\x17\x7f3\xcc" +
	"\xdbtx\xef\x07%TQ\xa0W\xdccX\x12\x99." +
	"\xc3\x05\x09U\x11\xe8\x95w\x19\xd6\x88\xcc\xa9\xadDj" +
	"BBy\x02F\xc1\xad\"\xd5\xf2\x15\x02R\x84\xe4\x98" +
	"[t\x90\x8a^M\x13.\xf2\x8cR\xd1\x9bn\xc0\xb3" +
	"\xb5\x86\xea\x91j\xf79>Y<\x14\xee\x0cN\xf0(" +
	"\xc2\x99\xcc0\xc1\xe7$\xd4\x09\x01\xfe\x8b\\\xcb<\xf6" +
	"4\x09S \x10\x9a95\x1c16\xa5\x08Tf>" +
	"3G\xa4<\x09uT\xc0\xd4d 1s\xe6\x10\x91" +
	"z^B\xbd\xd0\xa9\xa3\xc6L\xeb\x1eI\xa7\xd6\xb1\xff" +
	"\x7f\xbf\xf0\x8a\x937\\\xbb\xa84\xa0\xcd\xfa\xd1\x9fd" +
	"+P]a\xc3\x0fr\xc3}\x12j=7\xdc\xd8\xe9" +
	"\x00c\x0fH\xa8\x8d\x02Io\xba\xe2 \x19\xe5  " +
	"IH\x16l\xcfF7\x09tw\x96f\xe16\xbc\xa6" +
	"Mj\xc3\x91\xd4\xb2\x9e]\x1dw\xbc\xfb\xd4\x16\\~" +
	"\xf4\xb1\x91\xc1\xcc\x1ew|\xc2\xe3\x04m4\x87\x96\xa0" +
	"9J\xa4\xd6I\xa8M\x02\x99*\xdf\x09s\xd6kN" +
	"u\xbcZ\xae\x93Qq\x0b0H\xc0X\xa2\x12\x05\xd3" +
	"\x09]\xde\xc4P\xa3t\xbb\x1e\x86\"=\x84\xfeul" +
	"(Z)\x9a\xf6u\x92gvTB\xbd\xd8\x14\x04\xcb" +
	"\xfd4\xb7}BB\x9dgA4\xc5\xfe\xe6*\"\xf5" +
	"\xba\x84\xba \x90\xa9\xb7+\"3\xde\xfe\x95d\xba-" +
	"\xee\xd9\xa0\xbfZ\xe4-!\xeb\x86\xb7H\xb7\xb0\xb4\x10" +
	"\x9ep\x8bM\xe7O\x84M=\xbe\x87Hm\x93P\xbb" +
	"\xa3a\xeeb\xaa\xdb%T\x81=\xb9\xf9\x8a\xed\xfe\xe8" +
	"q\xfb\xc1\xff\x82\x9c{\x84\xe0 F\x021B6\xc0" +
	"BV,\x87e\x1d\xc6\xb6\x98\xcd\xbe\x9aS\xcd<Y" +
	"-\xd7+\x1d\x8c\xb8\xd0\x16\x09\xb5\xb3m\xcc#Cm" +
	"4\x9bc\x0ei\xee\xedxM\x19w\xba<Rh}" +
	"e\xf9\xab\xe4\xb5H\xfc\x13\x00\x00\xff\xff4O\x0c\xe1"

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
