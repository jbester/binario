package binaryio

type IBinaryWriter interface {
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

type IBinaryBufferWriter interface {
	IBinaryWriter
	Bytes() []byte
}
