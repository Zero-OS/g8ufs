@0xae9223e76351538a;

using Go = import "/go.capnp";
$Go.package("np");
$Go.import("github.com/g8os/g8ufs/cap.np");

struct File {
    # blocksize in bytes = blocksize * 4 KB, blocksize is same for all parts of file
    # max blocksize = 128 MB
    blockSize   @0: UInt8;
    blocks      @1: List(Data);    # list of the hashes of the blocks
    size        @2: UInt64;        # in bytes
}

struct Link {
    target @0: Text;          # path to target
}

struct Special {
    type @0: Type;

    # 0 => socket       (S_IFSOCK)
    # 1 => block device (S_IFBLK)
    # 2 => char. device (S_IFCHR)
    # 3 => fifo pipe    (S_IFIFO)
    enum Type {
        socket   @0;
        block    @1;
        chardev  @2;
        fifopipe @3;
        unknown  @4;
    }

    data @1: Data;     # data relevant for type of item
}

struct SubDir {
    key  @0: Text;
}

struct Inode {
    name    @0: Text;

    attributes: union {
        dir     @1: SubDir;
        file    @2: File;
        link    @3: Link;
        special @4: Special;
    }

    aclkey           @5: UInt32; # is pointer to ACL
    modificationTime @6: UInt32;
    creationTime     @7: UInt32;
}

struct Dir {
    name     @0: Text;            # name of dir

    location @1: Text;            # location in filesystem = namespace
    contents @2: List(Inode);     # list of the contents (files, links, specials)
    parent   @3: Text;            # directory key of parent

    # directory's metadata

    size             @4: UInt64;  # in bytes
    aclkey           @5: UInt32;  # is pointer to ACL
    isLink           @6: Bool;    # if is link and not physically on disk
    modificationTime @7: UInt32;
    creationTime     @8: UInt32;
}

struct UserGroup {
    name   @0: Text;
    iyoId  @1: Text;    # itsyou.online id
    iyoInt @2: UInt64;  # itsyouonline unique id per user or group, is globally unique
}

struct ACI {
    # for backwards compatibility with posix
    uname @0: Text;
    gname @1: Text;
    mode  @2: UInt16;

    rights @3 :List(Right);
    struct Right {
        # text e.g. rwdl- (admin read write delete list -), freely to be chosen
        # admin means all rights (e.g. on / = namespace or filesystem level all rights for everything)
        # - means remove all previous ones (is to stop recursion), if usergroupid=0 then is for all users & all groups
        right       @0: Text;
        usergroupid @1: UInt16;
    }

    id @4: UInt32;
}