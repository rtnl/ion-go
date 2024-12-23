package ion

// #include <ion.h>
import "C"
import (
	"fmt"
	"reflect"
)

func (b *Buffer) WriteObjectAny(value any) (err error) {
	switch val := value.(type) {
	case uint8:
		return b.WriteU8(val)
	case uint16:
		return b.WriteU16(val)
	case uint32:
		return b.WriteU32(val)
	case uint64:
		return b.WriteU64(val)
	case int:
		return b.WriteU64(uint64(val))
	case uint:
		return b.WriteU64(uint64(val))
	default:
		return fmt.Errorf("unimplemented object type: %v", reflect.TypeOf(value))
	}
}

func (b *Buffer) WriteObject(kind ObjectKind, value any) (err error) {
	switch kind {
	case U0:
		return b.WriteU0()
	case U8:
		return b.WriteU8(value.(uint8))
	case U16:
		return b.WriteU16(value.(uint16))
	case U32:
		return b.WriteU32(value.(uint32))
	case U64:
		return b.WriteU64(value.(uint64))
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

func BufferWriteArray[T any](b *Buffer, kind ObjectKind, value []T) (err error) {
	err = b.WriteArrayOpen(kind)
	if err != nil {
		return
	}

	for _, it := range value {
		err = b.WriteObject(kind, it)
		if err != nil {
			return
		}
	}

	err = b.WriteArrayClose()
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

func BufferWriteList[T any](b *Buffer, value []T) (err error) {
	err = b.WriteListOpen()
	if err != nil {
		return
	}

	for _, it := range value {
		err = b.WriteObjectAny(it)
		if err != nil {
			return
		}
	}

	err = b.WriteListClose()
	if err != nil {
		return
	}

	return
}
