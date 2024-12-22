package ion

// #include <ion.h>
import "C"

type ObjectKind = C.t_ion_object_kind

const (
	U0 ObjectKind = iota
	U8
	U16
	U32
	U64
	ARR
	LIST
)
