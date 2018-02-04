package binaryio

// IBinaryReader wraps the encoding/binary interface to allow decoding
// of binary streams.
type IBinaryReader interface {
	//  Read any fixed size value. The  parameter should be a pointer to the underlying type.
	// // var value uint32
	// // if err :=  reader.Read(&value); err != nil {
	// //     ...
	// // }
	Read(value interface{}) error
	ReadByte() (byte, error)
	ReadBytes(expected int) ([]byte, error)
	ReadUint16() (uint16, error)
	ReadInt16() (int16, error)
	ReadUint16s(expected int) ([]uint16, error)
	ReadInt16s(expected int) ([]int16, error)
	ReadUint32() (uint32, error)
	ReadInt32() (int32, error)
	ReadUint32s(expected int) ([]uint32, error)
	ReadInt32s(expected int) ([]int32, error)
	ReadUint64() (uint64, error)
	ReadInt64() (int64, error)
	ReadUint64s(expected int) ([]uint64, error)
	ReadInt64s(expected int) ([]int64, error)
	//  Skip n bytes of the output
	Skip(count int) error
}

// IBinaryBufferReader implements the IBinaryReader interface over a bytes.Buffer
type IBinaryBufferReader interface {
	IBinaryReader
	// Get the number of bytes remaining in the buffer.
	Len() int
}
