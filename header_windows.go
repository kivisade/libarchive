package libarchive

/*
#include <stdlib.h>
*/
import "C"

import (
	"time"
)

// ModTime returns entry modification time
func (e *entryInfo) ModTime() time.Time {
	return e.mtime
}
