package binaryio

import (
	"bytes"
	"encoding/binary"
)

type BinaryBufferWriter struct {
	BinaryWriter
	buffer *bytes.Buffer
}

func BigEndianBufferWriter() IBinaryBufferWriter {
	var buf = &bytes.Buffer{}
	return BinaryBufferWriter{buffer: buf, BinaryWriter: BinaryWriter{endian: binary.BigEndian, raw: buf}}
}

func LittleEndianBufferWriter() IBinaryBufferWriter {
	var buf = &bytes.Buffer{}
	return BinaryBufferWriter{buffer: buf, BinaryWriter: BinaryWriter{endian: binary.LittleEndian, raw: buf}}
}

func NewBigEndianBufferWriter(capacity int) IBinaryBufferWriter {
	var buf = bytes.NewBuffer(make([]byte, 0, capacity))
	return BinaryBufferWriter{buffer: buf, BinaryWriter: BinaryWriter{endian: binary.BigEndian, raw: buf}}
}

func NewLittleEndianBufferWriter(capacity int) IBinaryBufferWriter {
	var buf = bytes.NewBuffer(make([]byte, 0, capacity))
	return BinaryBufferWriter{buffer: buf, BinaryWriter: BinaryWriter{endian: binary.LittleEndian, raw: buf}}
}

func (writer BinaryBufferWriter) Bytes() []byte {
	return writer.buffer.Bytes()
}
