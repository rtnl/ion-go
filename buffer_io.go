package ion

// #include <ion.h>
import "C"
import (
	"container/list"
	"fmt"
	"log/slog"
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

func (b *Buffer) ReadObject(kind ObjectKind) (value any, err error) {
	switch kind {
	case U0:
		return b.ReadU0()
	case U8:
		return b.ReadU8()
	case U16:
		return b.ReadU16()
	case U32:
		return b.ReadU32()
	case U64:
		return b.ReadU64()
	default:
		return nil, fmt.Errorf("unimplemented object type: %v", reflect.TypeOf(value))
	}
}

func (b *Buffer) PeekObject(kind ObjectKind) (value any, err error) {
	switch kind {
	case U0:
		return b.PeekU0()
	case U8:
		return b.PeekU8()
	case U16:
		return b.PeekU16()
	case U32:
		return b.PeekU32()
	case U64:
		return b.PeekU64()
	default:
		return nil, fmt.Errorf("unimplemented object type: %v", reflect.TypeOf(value))
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

func (b *Buffer) ReadU0() (value any, err error) {
	if b.inner == nil {
		err = fmt.Errorf("buffer is nil")
		return
	}

	err = Check(C.ion_buffer_io_read_u0(b.inner))
	if err != nil {
		return
	}

	return
}

func (b *Buffer) ReadU8() (value uint8, err error) {
	var (
		v C.uint8_t
	)

	if b.inner == nil {
		err = fmt.Errorf("buffer is nil")
		return
	}

	err = Check(C.ion_buffer_io_read_u8(b.inner, &v))
	if err != nil {
		return
	}

	value = uint8(v)

	return
}

func (b *Buffer) ReadU16() (value uint16, err error) {
	var (
		v C.uint16_t
	)

	if b.inner == nil {
		err = fmt.Errorf("buffer is nil")
		return
	}

	err = Check(C.ion_buffer_io_read_u16(b.inner, &v))
	if err != nil {
		return
	}

	value = uint16(v)

	return
}

func (b *Buffer) ReadU32() (value uint32, err error) {
	var (
		v C.uint32_t
	)

	if b.inner == nil {
		err = fmt.Errorf("buffer is nil")
		return
	}

	err = Check(C.ion_buffer_io_read_u32(b.inner, &v))
	if err != nil {
		return
	}

	value = uint32(v)

	return
}

func (b *Buffer) ReadU64() (value uint64, err error) {
	var (
		v C.uint64_t
	)

	if b.inner == nil {
		err = fmt.Errorf("buffer is nil")
		return
	}

	err = Check(C.ion_buffer_io_read_u64(b.inner, &v))
	if err != nil {
		return
	}

	value = uint64(v)

	return
}

func (b *Buffer) ReadArrayOpen() (kind ObjectKind, length int, err error) {
	var (
		k C.t_ion_object_kind
		l C.uint8_t
	)

	if b.inner == nil {
		err = fmt.Errorf("buffer is nil")
		return
	}

	err = Check(C.ion_buffer_io_read_arr_open(b.inner, &k, &l))
	if err != nil {
		return
	}

	kind = ObjectKind(k)
	length = int(l)

	return
}

func (b *Buffer) ReadArrayCheck(kind ObjectKind, length int) (err error) {
	err = Check(C.ion_buffer_io_read_arr_check(b.inner, kind, C.uint8_t(length)))
	if err != nil {
		return
	}

	return
}

func (b *Buffer) ReadArrayClose() (err error) {
	if b.inner == nil {
		err = fmt.Errorf("buffer is nil")
		return
	}

	err = Check(C.ion_buffer_io_read_arr_close(b.inner))
	if err != nil {
		return
	}

	return
}

func BufferReadArray[T any](b *Buffer, kind ObjectKind) (value []T, err error) {
	var (
		length  int
		valKind ObjectKind
		valRaw  any
	)

	valKind, length, err = b.ReadArrayOpen()
	if err != nil {
		return
	}

	if valKind != kind {
		err = fmt.Errorf("inner type does not match: %v %v", kind, valKind)
		return
	}

	err = b.ReadArrayCheck(kind, length)
	if err != nil {
		return
	}

	value = make([]T, length)
	for x := range length {
		valRaw, err = b.ReadObject(kind)
		if err != nil {
			return
		}

		value[x] = valRaw.(T)
	}

	err = b.ReadArrayClose()
	if err != nil {
		return
	}

	return
}

func (b *Buffer) ReadListOpen() (length int, err error) {
	var (
		l C.uint8_t
	)

	if b.inner == nil {
		err = fmt.Errorf("buffer is nil")
		return
	}

	err = Check(C.ion_buffer_io_read_list_open(b.inner, &l))
	if err != nil {
		return
	}

	length = int(l)

	return
}

func (b *Buffer) ReadListClose() (err error) {
	if b.inner == nil {
		err = fmt.Errorf("buffer is nil")
	}

	err = Check(C.ion_buffer_io_read_list_close(b.inner))
	if err != nil {
		return
	}

	return
}

func (b *Buffer) ReadList(kinds []ObjectKind) (value *list.List, err error) {
	var (
		length int
		item   any
	)

	length, err = b.ReadListOpen()
	if err != nil {
		return
	}

	value = list.New()

	for x := range length {
		item, err = b.ReadObject(kinds[x])
		if err != nil {
			return
		}

		value.PushBack(item)
	}

	err = b.ReadListClose()
	if err != nil {
		return
	}

	return
}

func (b *Buffer) PeekU0() (value any, err error) {
	if b.inner == nil {
		err = fmt.Errorf("buffer is nil")
		return
	}

	err = Check(C.ion_buffer_io_peek_u0(b.inner))
	if err != nil {
		return
	}

	return
}

func (b *Buffer) PeekU8() (value uint8, err error) {
	var (
		v C.uint8_t
	)

	if b.inner == nil {
		err = fmt.Errorf("buffer is nil")
		return
	}

	err = Check(C.ion_buffer_io_peek_u8(b.inner, &v))
	if err != nil {
		return
	}

	value = uint8(v)

	return
}

func (b *Buffer) PeekU16() (value uint16, err error) {
	var (
		v C.uint16_t
	)

	if b.inner == nil {
		err = fmt.Errorf("buffer is nil")
		return
	}

	err = Check(C.ion_buffer_io_peek_u16(b.inner, &v))
	if err != nil {
		return
	}

	value = uint16(v)

	return
}

func (b *Buffer) PeekU32() (value uint32, err error) {
	var (
		v C.uint32_t
	)

	if b.inner == nil {
		err = fmt.Errorf("buffer is nil")
		return
	}

	err = Check(C.ion_buffer_io_peek_u32(b.inner, &v))
	if err != nil {
		return
	}

	value = uint32(v)

	return
}

func (b *Buffer) PeekU64() (value uint64, err error) {
	var (
		v C.uint64_t
	)

	if b.inner == nil {
		err = fmt.Errorf("buffer is nil")
		return
	}

	err = Check(C.ion_buffer_io_peek_u64(b.inner, &v))
	if err != nil {
		return
	}

	value = uint64(v)

	return
}

func (b *Buffer) PeekArrayOpen() (kind ObjectKind, length int, err error) {
	var (
		k C.t_ion_object_kind
		l C.uint8_t
	)

	if b.inner == nil {
		err = fmt.Errorf("buffer is nil")
		return
	}

	err = Check(C.ion_buffer_io_peek_arr_open(b.inner, &k, &l))
	if err != nil {
		return
	}

	kind = ObjectKind(k)
	length = int(l)

	return
}

func (b *Buffer) PeekArrayCheck(kind ObjectKind, length int) (err error) {
	err = Check(C.ion_buffer_io_peek_arr_check(b.inner, kind, C.uint8_t(length)))
	if err != nil {
		return
	}

	return
}

func (b *Buffer) PeekArrayClose() (err error) {
	if b.inner == nil {
		err = fmt.Errorf("buffer is nil")
		return
	}

	err = Check(C.ion_buffer_io_peek_arr_close(b.inner))
	if err != nil {
		return
	}

	return
}

func BufferPeekArray[T any](b *Buffer, kind ObjectKind) (value []T, err error) {
	var (
		length  int
		valKind ObjectKind
		valRaw  any
	)

	valKind, length, err = b.PeekArrayOpen()
	slog.Debug("peek array open",
		slog.Any("valKind", valKind),
		slog.Any("length", length),
		slog.Any("err", err))
	if err != nil {
		return
	}

	if valKind != kind {
		err = fmt.Errorf("inner type does not match: %v %v", kind, valKind)
		return
	}

	err = b.PeekArrayCheck(kind, length)
	if err != nil {
		return
	}

	value = make([]T, length)
	for x := range length {
		valRaw, err = b.PeekObject(kind)
		slog.Debug("peek array open object",
			slog.Any("valRaw", valRaw),
			slog.Any("err", err))
		if err != nil {
			return
		}

		value[x] = valRaw.(T)
	}

	err = b.PeekArrayClose()
	if err != nil {
		return
	}

	return
}

func (b *Buffer) PeekListOpen() (length int, err error) {
	var (
		l C.uint8_t
	)

	if b.inner == nil {
		err = fmt.Errorf("buffer is nil")
		return
	}

	err = Check(C.ion_buffer_io_peek_list_open(b.inner, &l))
	if err != nil {
		return
	}

	length = int(l)

	return
}

func (b *Buffer) PeekListClose() (err error) {
	if b.inner == nil {
		err = fmt.Errorf("buffer is nil")
	}

	err = Check(C.ion_buffer_io_peek_list_close(b.inner))
	if err != nil {
		return
	}

	return
}

func (b *Buffer) PeekList(kinds []ObjectKind) (value *list.List, err error) {
	var (
		length int
		item   any
	)

	length, err = b.ReadListOpen()
	if err != nil {
		return
	}

	value = list.New()

	for x := range length {
		item, err = b.PeekObject(kinds[x])
		if err != nil {
			return
		}

		value.PushBack(item)
	}

	err = b.PeekListClose()
	if err != nil {
		return
	}

	return
}
