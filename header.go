package libarchive

/*
#cgo pkg-config: libarchive
#include <archive.h>
#include <archive_entry.h>
#include <stdlib.h>
*/
import "C"

import (
	"os"
	"path/filepath"
	"syscall"
	"time"
)

// ArchiveEntry represents an libarchive archive_entry
type ArchiveEntry interface {
	// FileInfo describing archive_entry
	Stat() os.FileInfo
	// The name of the entry
	PathName() string
	Symlink() string
}

type entryImpl struct {
	entry *C.struct_archive_entry
}

type entryInfo struct {
	stat  *C.struct_stat
	name  string
	size  int64
	mtime time.Time
}

func (h *entryImpl) Stat() os.FileInfo {
	return &entryInfo{
		stat:  C.archive_entry_stat(h.entry),
		name:  filepath.Base(h.PathName()),
		size:  int64(C.archive_entry_size(h.entry)),
		mtime: time.Unix(int64(C.archive_entry_mtime(h.entry)), int64(C.archive_entry_mtime_nsec(h.entry))),
	}
}

func (h *entryImpl) PathName() string {
	name := C.archive_entry_pathname_utf8(h.entry)

	return C.GoString(name)
}

func (h *entryImpl) Symlink() string {
	name := C.archive_entry_symlink(h.entry)

	return C.GoString(name)
}

func (e *entryInfo) Name() string {
	return e.name
}

func (e *entryInfo) Size() int64 {
	return e.size
}

func (e *entryInfo) Mode() os.FileMode {
	mode := os.FileMode(e.stat.st_mode & 0777)
	switch uint(e.stat.st_mode) & uint(syscall.S_IFMT) {
	case syscall.S_IFLNK:
		mode |= os.ModeSymlink
	case syscall.S_IFSOCK:
		mode |= os.ModeSocket
	case syscall.S_IFCHR:
		mode |= os.ModeDevice | os.ModeCharDevice
	case syscall.S_IFBLK:
		mode |= os.ModeDevice
	case syscall.S_IFDIR:
		mode |= os.ModeDir
	case syscall.S_IFIFO:
		mode |= os.ModeNamedPipe
	}
	return mode
}

func (e *entryInfo) IsDir() bool {
	return e.stat.st_mode&syscall.S_IFDIR != 0
}

func (e *entryInfo) Sys() interface{} {
	return e.stat
}
