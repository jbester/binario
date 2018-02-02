package binaryio

import (
	"encoding/binary"
	"io"
)

type BinaryWriter struct {
	raw    io.Writer
	endian binary.ByteOrder
}

func NewBigEndianFromWriter(writer io.Writer) IBinaryWriter {
	return BinaryWriter{raw: writer, endian: binary.BigEndian}
}

func NewLittleEndianFromWriter(writer io.Writer) IBinaryWriter {
	return BinaryWriter{raw: writer, endian: binary.LittleEndian}
}

func (writer BinaryWriter) WriteByte(b byte) error {
	var n, err = writer.raw.Write([]byte{b})
	if n != 1 {
		return io.ErrShortWrite
	}
	if err != nil {
		return err
	}
	return nil
}

func (writer BinaryWriter) WriteBytes(bytes []byte) error {
	var n, err = writer.raw.Write(bytes)
	if n != len(bytes) {
		return io.ErrShortWrite
	}
	if err != nil {
		return err
	}
	return nil
}

func (writer BinaryWriter) WriteUint16(value uint16) error {
	return binary.Write(writer.raw, writer.endian, value)
}

func (writer BinaryWriter) WriteInt16(value int16) error {
	return binary.Write(writer.raw, writer.endian, value)
}

func (writer BinaryWriter) WriteUint16s(values []uint16) error {
	return binary.Write(writer.raw, writer.endian, values)
}

func (writer BinaryWriter) WriteInt16s(values []int16) error {
	return binary.Write(writer.raw, writer.endian, values)
}

func (writer BinaryWriter) WriteUint32(value uint32) error {
	return binary.Write(writer.raw, writer.endian, value)
}

func (writer BinaryWriter) WriteInt32(value int32) error {
	return binary.Write(writer.raw, writer.endian, value)
}

func (writer BinaryWriter) WriteUint32s(values []uint32) error {
	return binary.Write(writer.raw, writer.endian, values)
}

func (writer BinaryWriter) WriteInt32s(values []int32) error {
	return binary.Write(writer.raw, writer.endian, values)
}

func (writer BinaryWriter) WriteUint64(value uint64) error {
	return binary.Write(writer.raw, writer.endian, value)

}

func (writer BinaryWriter) WriteInt64(value int64) error {
	return binary.Write(writer.raw, writer.endian, value)
}

func (writer BinaryWriter) WriteUint64s(values []uint64) error {
	return binary.Write(writer.raw, writer.endian, values)
}

func (writer BinaryWriter) WriteInt64s(values []int64) error {
	return binary.Write(writer.raw, writer.endian, values)
}

func (writer BinaryWriter) Write(value interface{}) error {
	return binary.Write(writer.raw, writer.endian, value)
}
