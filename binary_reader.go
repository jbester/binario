package binaryio

import (
	"encoding/binary"
	"io"

	"bytes"
	"github.com/pkg/errors"
)

type BinaryReader struct {
	raw    io.Reader
	endian binary.ByteOrder
}

type BinaryBufferReader struct {
	BinaryReader
	buffer *bytes.Buffer
}

func BigEndianBufferReader(buf []byte) IBinaryBufferReader {
	var buffer = bytes.NewBuffer(buf)
	return BinaryBufferReader{
		BinaryReader: BinaryReader{
			raw:    buffer,
			endian: binary.BigEndian,
		},
		buffer: buffer,
	}
}

func LittleEndianBufferReader(buf []byte) IBinaryBufferReader {
	var buffer = bytes.NewBuffer(buf)
	return BinaryBufferReader{
		BinaryReader: BinaryReader{
			raw:    buffer,
			endian: binary.LittleEndian,
		},
		buffer: buffer,
	}
}

func BigEndianReaderFrom(reader io.Reader) IBinaryReader {
	return BinaryReader{raw: reader, endian: binary.BigEndian}
}

func LittleEndianReaderFrom(reader io.Reader) IBinaryReader {
	return BinaryReader{raw: reader, endian: binary.LittleEndian}
}

var ErrShortRead = errors.New("short read")

func (reader BinaryBufferReader) Len() int {
	return reader.Len()
}

func (reader BinaryReader) ReadByte() (byte, error) {
	var ReadSize = Size(byte(0))
	var b = make([]byte, ReadSize)
	var n, err = reader.raw.Read(b)
	if err != nil {
		return 0, err
	}
	if n != ReadSize {
		return 0, ErrShortRead
	}

	return b[0], nil
}

func (reader BinaryReader) ReadBytes(expected int) ([]byte, error) {
	var b = make([]byte, expected)
	var n, err = reader.raw.Read(b)
	if err != nil {
		return nil, err
	}
	if n != expected {
		return nil, ErrShortRead
	}

	return b, nil
}

func (reader BinaryReader) ReadUint16() (uint16, error) {
	var value uint16
	var err = binary.Read(reader.raw, reader.endian, &value)
	return value, err
}

func (reader BinaryReader) ReadInt16() (int16, error) {
	var value int16
	var err = binary.Read(reader.raw, reader.endian, &value)
	return value, err
}

func (reader BinaryReader) ReadUint16s(expected int) ([]uint16, error) {
	var value = make([]uint16, expected)
	var err = binary.Read(reader.raw, reader.endian, &value)
	return value, err
}

func (reader BinaryReader) ReadInt16s(expected int) ([]int16, error) {
	var value = make([]int16, expected)
	var err = binary.Read(reader.raw, reader.endian, &value)
	return value, err
}

func (reader BinaryReader) ReadUint32() (uint32, error) {
	var value uint32
	var err = binary.Read(reader.raw, reader.endian, &value)
	return value, err
}

func (reader BinaryReader) ReadInt32() (int32, error) {
	var value int32
	var err = binary.Read(reader.raw, reader.endian, &value)
	return value, err
}

func (reader BinaryReader) ReadUint32s(expected int) ([]uint32, error) {
	var value = make([]uint32, expected)
	var err = binary.Read(reader.raw, reader.endian, &value)
	return value, err
}

func (reader BinaryReader) ReadInt32s(expected int) ([]int32, error) {
	var value = make([]int32, expected)
	var err = binary.Read(reader.raw, reader.endian, &value)
	return value, err
}

func (reader BinaryReader) ReadUint64() (uint64, error) {
	var value uint64
	var err = binary.Read(reader.raw, reader.endian, &value)
	return value, err
}

func (reader BinaryReader) ReadInt64() (int64, error) {
	var value int64
	var err = binary.Read(reader.raw, reader.endian, &value)
	return value, err
}

func (reader BinaryReader) ReadUint64s(expected int) ([]uint64, error) {
	var value = make([]uint64, expected)
	var err = binary.Read(reader.raw, reader.endian, &value)
	return value, err
}

func (reader BinaryReader) ReadInt64s(expected int) ([]int64, error) {
	var value = make([]int64, expected)
	var err = binary.Read(reader.raw, reader.endian, &value)
	return value, err
}

func (reader BinaryReader) Skip(count int) error {
	var b = make([]byte, count)
	var n, err = reader.raw.Read(b)
	if err != nil {
		return err
	}
	if n != count {
		return ErrShortRead
	}
	return nil
}

func (reader BinaryReader) Read(value interface{}) error {
	return binary.Read(reader.raw, reader.endian, value)
}
