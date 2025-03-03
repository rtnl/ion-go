package ion

// #include "ion.h"
import "C"

type BufferState struct {
	inner *C.t_ion_buffer_state
}

func NewBufferState(inner *C.t_ion_buffer_state) (s *BufferState) {
	s = new(BufferState)

	s.inner = inner

	return
}

func (s *BufferState) GetEntryLevel() int32 {
	return int32(s.inner.entry_level)
}
