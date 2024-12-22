package ion

// #include <ion.h>
import "C"
import (
	"fmt"
	"reflect"
)

func (b *Buffer) WriteObject(value any) (err error) {
	switch val := value.(type) {
	case any:
		return b.WriteU0()
	case nil:
		return b.WriteU0()
	case uint8:
		return b.WriteU8(val)
	case uint16:
		return b.WriteU16(val)
	case uint32:
		return b.WriteU32(val)
	case uint64:
		return b.WriteU64(val)
	default:
		return fmt.Errorf("unimplemented object type: %v", reflect.TypeOf(value))
	}
}

func (b *Buffer) WriteU0() (err error) {
	if b.inner == nil {
		err = fmt.Errorf("buffer is nil")
		return
	}

	err = Check(C.ion_buffer_io_write_u0(b.inner))
	if err != nil {
		return
	}

	return
}

func (b *Buffer) WriteU8(value uint8) (err error) {
	if b.inner == nil {
		err = fmt.Errorf("buffer is nil")
		return
	}

	err = Check(C.ion_buffer_io_write_u8(b.inner, C.uint8_t(value)))
	if err != nil {
		return
	}

	return
}

func (b *Buffer) WriteU16(value uint16) (err error) {
	if b.inner == nil {
		err = fmt.Errorf("buffer is nil")
		return
	}

	err = Check(C.ion_buffer_io_write_u16(b.inner, C.uint16_t(value)))
	if err != nil {
		return
	}

	return
}

func (b *Buffer) WriteU32(value uint32) (err error) {
	if b.inner == nil {
		err = fmt.Errorf("buffer is nil")
		return
	}

	err = Check(C.ion_buffer_io_write_u32(b.inner, C.uint32_t(value)))
	if err != nil {
		return
	}

	return
}

func (b *Buffer) WriteU64(value uint64) (err error) {
	if b.inner == nil {
		err = fmt.Errorf("buffer is nil")
		return
	}

	err = Check(C.ion_buffer_io_write_u64(b.inner, C.uint64_t(value)))
	if err != nil {
		return
	}

	return
}

func (b *Buffer) WriteArrayOpen(kind ObjectKind) (err error) {
	if b.inner == nil {
		err = fmt.Errorf("buffer is nil")
		return
	}

	err = Check(C.ion_buffer_io_write_arr_open(b.inner, kind))
	if err != nil {
		return
	}

	return
}

func (b *Buffer) WriteArrayClose() (err error) {
	if b.inner == nil {
		err = fmt.Errorf("buffer is nil")
		return
	}

	err = Check(C.ion_buffer_io_write_arr_close(b.inner))
	if err != nil {
		return
	}

	return
}

func (b *Buffer) WriteListOpen() (err error) {
	if b.inner == nil {
		err = fmt.Errorf("buffer is nil")
		return
	}

	err = Check(C.ion_buffer_io_write_list_open(b.inner))
	if err != nil {
		return
	}

	return
}

func (b *Buffer) WriteListClose() (err error) {
	if b.inner == nil {
		err = fmt.Errorf("buffer is nil")
		return
	}

	err = Check(C.ion_buffer_io_write_list_close(b.inner))
	if err != nil {
		return
	}

	return
}
