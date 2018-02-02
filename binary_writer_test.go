package binaryio

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_WriteByte(t *testing.T) {
	var writer = BigEndianBufferWriter()
	writer.WriteByte(128)
	assert.Equal(t, writer.Bytes(), []byte{128})
}

func Test_WriteBytes(t *testing.T) {
	var writer = BigEndianBufferWriter()
	var expected = []byte{1, 2, 3, 4, 5}
	writer.WriteBytes(expected)
	assert.Equal(t, expected, writer.Bytes())
}

func Test_WriteUint16(t *testing.T) {
	var writer = BigEndianBufferWriter()
	var expected = []byte{0x12, 0x34}
	writer.WriteUint16(0x1234)
	assert.Equal(t, expected, writer.Bytes())
}

func Test_WriteUint16s(t *testing.T) {
	var writer = BigEndianBufferWriter()
	var expected = []byte{0x12, 0x34, 0xed, 0xcc}
	writer.WriteUint16s([]uint16{0x1234, 0xedcc})
	assert.Equal(t, expected, writer.Bytes())
}

func Test_WriteInt16(t *testing.T) {
	var writer = BigEndianBufferWriter()
	var expected = []byte{0xed, 0xcc}
	writer.WriteInt16(-0x1234)
	assert.Equal(t, expected, writer.Bytes())
}

func Test_WriteInt16s(t *testing.T) {
	var writer = BigEndianBufferWriter()
	var expected = []byte{0x12, 0x34, 0xed, 0xcc}
	writer.WriteInt16s([]int16{0x1234, -0x1234})
	assert.Equal(t, expected, writer.Bytes())
}

func Test_WriteUint32(t *testing.T) {
	var writer = BigEndianBufferWriter()
	var expected = []byte{0x12, 0x34, 0x56, 0x78}
	writer.WriteUint32(0x12345678)
	assert.Equal(t, expected, writer.Bytes())
}

func Test_WriteUint32s(t *testing.T) {
	var writer = BigEndianBufferWriter()
	var expected = []byte{0x12, 0x34, 0x56, 0x78, 0xed, 0xcb, 0xa9, 0x88}
	writer.WriteUint32s([]uint32{0x12345678, 0xedcba988})
	assert.Equal(t, expected, writer.Bytes())
}

func Test_WriteInt32(t *testing.T) {
	var writer = BigEndianBufferWriter()
	var expected = []byte{0xed, 0xcb, 0xa9, 0x88}
	writer.WriteInt32(-0x12345678)
	assert.Equal(t, expected, writer.Bytes())
}

func Test_WriteInt32s(t *testing.T) {
	var writer = BigEndianBufferWriter()
	var expected = []byte{0x12, 0x34, 0x56, 0x78, 0xed, 0xcb, 0xa9, 0x88}
	writer.WriteInt32s([]int32{0x12345678, -0x12345678})
	assert.Equal(t, expected, writer.Bytes())
}

func Test_WriteUint64(t *testing.T) {
	var writer = BigEndianBufferWriter()
	var expected = []byte{0x10, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef}
	writer.WriteUint64(0x1023456789abcdef)
	assert.Equal(t, expected, writer.Bytes())
}

func Test_WriteUint64s(t *testing.T) {
	var writer = BigEndianBufferWriter()
	var expected = []byte{0x10, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xef, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x11}
	writer.WriteUint64s([]uint64{0x1023456789abcdef, 0xefdcba9876543211})
	assert.Equal(t, expected, writer.Bytes())
}

func Test_WriteInt64(t *testing.T) {
	var writer = BigEndianBufferWriter()
	var expected = []byte{0xef, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x11}
	writer.WriteInt64(-0x1023456789abcdef)
	assert.Equal(t, expected, writer.Bytes())
}

func Test_WriteInt64s(t *testing.T) {
	var writer = BigEndianBufferWriter()
	var expected = []byte{0x10, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xef, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x11}
	writer.WriteInt64s([]int64{0x1023456789abcdef, -0x1023456789abcdef})
	assert.Equal(t, expected, writer.Bytes())
}

func Test_Write(t *testing.T) {
	var writer = BigEndianBufferWriter()
	var expected = []byte{0x10, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xef, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x11}
	writer.Write([]int64{0x1023456789abcdef, -0x1023456789abcdef})
	assert.Equal(t, expected, writer.Bytes())
}
