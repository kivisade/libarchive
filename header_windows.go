package libarchive

/*
#include <stdlib.h>
*/
import "C"

import (
	"time"
)

// ModTime returns entry modification time
// Ref.: https://github.com/libarchive/libarchive/blob/master/libarchive/archive_windows.c#L88
func (e *entryInfo) ModTime() time.Time {
	return time.Unix(int64(e.stat.st_mtime), 0) // unfortunately, e.stat.st_mtime_nsec doesn't seem to be available on the struct
}
