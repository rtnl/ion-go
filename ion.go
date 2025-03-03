package ion

// #cgo CFLAGS: -I/usr/local/include
// #cgo LDFLAGS: -L/usr/local/lib -lion -Wl,-rpath,/usr/local/lib
// #include "ion.h"
import (
	"C"
)
