package binaryio

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ReadByte(t *testing.T) {
	var expected = byte(128)
	var reader = BigEndianBufferReader([]byte{expected})
	var actual, err = reader.ReadByte()
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func Test_ReadBytes(t *testing.T) {
	var expected = []byte{1, 2, 3, 4, 5}
	var reader = BigEndianBufferReader(expected)
	var actual, err = reader.ReadBytes(5)
	assert.NoError(t, err)
	assert.EqualValues(t, expected, actual)
}

func Test_ReadUint16(t *testing.T) {
	var expected = uint16(0x1234)
	var reader = BigEndianBufferReader([]byte{0x12, 0x34})
	var actual, err = reader.ReadUint16()
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func Test_ReadInt16(t *testing.T) {
	var expected = int16(-0x1234)
	var reader = BigEndianBufferReader([]byte{0xed, 0xcc})
	var actual, err = reader.ReadInt16()
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func Test_ReadUint32(t *testing.T) {
	var expected = uint32(0x12345678)
	var reader = BigEndianBufferReader([]byte{0x12, 0x34, 0x56, 0x78})
	var actual, err = reader.ReadUint32()
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func Test_ReadInt32(t *testing.T) {
	var expected = int32(-0x12345678)
	var reader = BigEndianBufferReader([]byte{0xed, 0xcb, 0xa9, 0x88})
	var actual, err = reader.ReadInt32()
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func Test_ReadUint64(t *testing.T) {
	var expected = uint64(0x1023456789abcdef)
	var reader = BigEndianBufferReader([]byte{0x10, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef})
	var actual, err = reader.ReadUint64()
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func Test_ReadInt64(t *testing.T) {
	var expected = int64(-0x1023456789abcdef)
	var reader = BigEndianBufferReader([]byte{0xef, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x11})
	var actual, err = reader.ReadInt64()
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Read(t *testing.T) {
	var expected = int64(-0x1023456789abcdef)
	var reader = BigEndianBufferReader([]byte{0xef, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x11})
	var actual int64
	var err = reader.Read(&actual)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Read_Array(t *testing.T) {
	var expected = []byte{0xef, 0xdc, 0xba, 0x98, 0x76, 0x54, 0x32, 0x11}
	var reader = BigEndianBufferReader(expected)
	var actual = make([]byte, 8)
	var err = reader.Read(&actual)
	assert.NoError(t, err)
	assert.EqualValues(t, expected, actual)
}
