package binaryio

// IBinaryWriter exposes methods to write fixed sized numbers to
// an underlying writer implementation
type IBinaryWriter interface {
	//  Write any fixed sized type or array of fixed size types
	Write(interface{}) error
	WriteByte(byte) error
	WriteBytes([]byte) error
	WriteInt16(int16) error
	WriteUint16(uint16) error
	WriteInt16s([]int16) error
	WriteUint16s([]uint16) error
	WriteInt32(int32) error
	WriteUint32(uint32) error
	WriteInt32s([]int32) error
	WriteUint32s([]uint32) error
	WriteInt64(int64) error
	WriteUint64(uint64) error
	WriteInt64s([]int64) error
	WriteUint64s([]uint64) error
}

// IBinaryBufferWriter implements the IBinaryWriter interfaces on top of bytes.Buffer
type IBinaryBufferWriter interface {
	IBinaryWriter
	// Return the binary representation of
	Bytes() []byte
}
