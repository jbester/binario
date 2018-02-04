package binaryio

import (
	"bytes"
	"encoding/binary"
)

type BinaryBufferWriter struct {
	BinaryWriter
	buffer *bytes.Buffer
}

// Create a new big endian writer using a bytes.Buffer as the underlying stream
func BigEndianBufferWriter() IBinaryBufferWriter {
	var buf = &bytes.Buffer{}
	return BinaryBufferWriter{buffer: buf, BinaryWriter: BinaryWriter{endian: binary.BigEndian, raw: buf}}
}

// Create a new little endian writer using a bytes.Buffer as the underlying stream
func LittleEndianBufferWriter() IBinaryBufferWriter {
	var buf = &bytes.Buffer{}
	return BinaryBufferWriter{buffer: buf, BinaryWriter: BinaryWriter{endian: binary.LittleEndian, raw: buf}}
}

// Create a new big endian writer using a bytes.Buffer as the underlying stream.  The buffer will have
// capacity specified.
func NewBigEndianBufferWriter(capacity int) IBinaryBufferWriter {
	var buf = bytes.NewBuffer(make([]byte, 0, capacity))
	return BinaryBufferWriter{buffer: buf, BinaryWriter: BinaryWriter{endian: binary.BigEndian, raw: buf}}
}

// Create a new little endian writer using a bytes.Buffer as the underlying stream.  The buffer will have
// capacity specified.
func NewLittleEndianBufferWriter(capacity int) IBinaryBufferWriter {
	var buf = bytes.NewBuffer(make([]byte, 0, capacity))
	return BinaryBufferWriter{buffer: buf, BinaryWriter: BinaryWriter{endian: binary.LittleEndian, raw: buf}}
}

// Return the byte representation of the underlying buffer
func (writer BinaryBufferWriter) Bytes() []byte {
	return writer.buffer.Bytes()
}
