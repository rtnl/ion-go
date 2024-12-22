package ion

// #cgo CFLAGS: -I/usr/local/include
// #cgo LDFLAGS: -L/usr/local/include -lion -Wl,-rpath,/usr/local/include
// #include "ion.h"
import (
	"C"
)

func init() {

}
