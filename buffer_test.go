package ion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBuffer(t *testing.T) {
	buffer := NewBuffer()
	assert.NotNil(t, buffer)
	assert.NotNil(t, buffer.inner)
}

func TestBuffer_Write(t *testing.T) {
	var (
		size int
		err  error
	)

	buffer := NewBuffer()

	src := []byte("In the beginning God created the heaven and the earth.")

	size, err = buffer.Write(src)
	assert.Nil(t, err)
	assert.Equal(t, len(src), size)

	dst := buffer.Consume()
	assert.NotNil(t, dst)
	assert.Equal(t, src, dst)
}

func TestBuffer_Read(t *testing.T) {
	var (
		size int
		err  error
	)

	buffer := NewBuffer()

	src := []byte("In the beginning God created the heaven and the earth.")
	dst := make([]byte, len(src))

	size, err = buffer.Write(src)
	assert.Nil(t, err)
	assert.Equal(t, len(src), size)

	size, err = buffer.Read(dst)
	assert.Nil(t, err)
	assert.Equal(t, len(src), size)
	assert.Equal(t, src, dst[:size])

	dst = make([]byte, len(src))
	size, err = buffer.Read(dst)
	assert.Nil(t, err)
	assert.Equal(t, size, 0)
	assert.Equal(t, dst, make([]byte, len(src)))
}

func TestBuffer_Peek(t *testing.T) {
	var (
		size int
		err  error
	)

	buffer := NewBuffer()

	src := []byte("In the beginning God created the heaven and the earth.")
	dst := make([]byte, len(src))

	size, err = buffer.Write(src)
	assert.Nil(t, err)
	assert.Equal(t, len(src), size)

	size, err = buffer.Peek(dst)
	assert.Nil(t, err)
	assert.Equal(t, len(src), size)
	assert.Equal(t, src, dst[:size])

	dst = make([]byte, len(src))
	size, err = buffer.Peek(dst)
	assert.Nil(t, err)
	assert.Equal(t, size, 0)
	assert.Equal(t, dst, make([]byte, len(src)))

	dst = make([]byte, len(src))
	size, err = buffer.Read(dst)
	assert.Nil(t, err)
	assert.Equal(t, len(src), size)
	assert.Equal(t, src, dst[:size])

	dst = make([]byte, len(src))
	size, err = buffer.Read(dst)
	assert.Nil(t, err)
	assert.Equal(t, size, 0)
	assert.Equal(t, dst, make([]byte, len(src)))
}
