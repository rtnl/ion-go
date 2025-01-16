package ion

// #include "ion.h"
import "C"
import (
	"fmt"
	"runtime"
	"unsafe"
)

type Buffer struct {
	inner *C.t_ion_buffer
}

func NewBuffer() (b *Buffer) {
	b = new(Buffer)

	b.inner = C.ion_buffer_new()

	runtime.SetFinalizer(b, func(v *Buffer) {
		if v.inner != nil {
			C.ion_buffer_free(v.inner)
		}
	})

	return
}

func (b *Buffer) Clone() (other *Buffer) {
	other = NewBuffer()

	if b == nil {
		return
	}

	if b.inner == nil {
		return
	}

	other.inner = C.ion_buffer_clone(b.inner)

	return
}

func (b *Buffer) Reduce() (err error) {
	err = Check(C.ion_buffer_reduce(b.inner))
	if err != nil {
		return
	}

	return
}

func (b *Buffer) Consume() (result []byte) {
	var (
		size C.size_t
		raw  *C.uint8_t
		arr  []C.uchar
	)

	if b.inner == nil {
		return nil
	}

	raw = C.ion_buffer_consume(b.inner, &size)
	arr = unsafe.Slice(raw, int(size))

	result = make([]byte, size)
	for x, it := range arr {
		result[x] = byte(it)
	}

	b.inner = nil
	return
}

func (b *Buffer) SeekRead(index uint8) (err error) {
	if b.inner == nil {
		return
	}

	err = Check(C.ion_buffer_seek_read(b.inner, C.uint8_t(index)))
	if err != nil {
		return
	}

	return
}

func (b *Buffer) SeekWrite(index uint8) (err error) {
	if b.inner == nil {
		return
	}

	err = Check(C.ion_buffer_seek_write(b.inner, C.uint8_t(index)))
	if err != nil {
		return
	}

	return
}

func (b *Buffer) Write(data []byte) (n int, err error) {
	if b.inner == nil {
		err = fmt.Errorf("buffer is null")
		return
	}

	err = Check(C.ion_buffer_write(b.inner, unsafe.Pointer(unsafe.SliceData(data)), C.size_t(len(data))))
	if err != nil {
		return
	}

	n = len(data)
	return
}

func (b *Buffer) Read(data []byte) (n int, err error) {
	var (
		size int
		curr C.size_t
	)

	if b.inner == nil {
		err = fmt.Errorf("buffer is null")
		return
	}

	curr = max(b.inner.body.curr_w-b.inner.body.curr_r, 0)
	size = min(int(curr), len(data))

	err = Check(C.ion_buffer_read(b.inner, unsafe.Pointer(unsafe.SliceData(data)), C.size_t(size)))
	if err != nil {
		return
	}

	n = size
	return
}
