package classreader

import (
	"encoding/binary"
)

type ClassReader struct {
	data []byte
}

func New(data []byte) *ClassReader {
	return &ClassReader{
		data: data,
	}
}

func (reader *ClassReader) ReadUint8() uint8 {
	bytes := reader.data[0]
	reader.data = reader.data[1:]

	return bytes
}

func (reader *ClassReader) ReadUint16() uint16 {
	data := binary.BigEndian.Uint16(reader.data)
	reader.data = reader.data[2:]

	return data
}

func (reader *ClassReader) ReadUint32() uint32 {
	bytes := binary.BigEndian.Uint32(reader.data)
	reader.data = reader.data[4:]

	return bytes
}

func (reader *ClassReader) ReadUint64() uint64 {
	bytes := binary.BigEndian.Uint64(reader.data)
	reader.data = reader.data[8:]

	return bytes
}

func (reader *ClassReader) ReadUint16s() []uint16 {
	length := reader.ReadUint16()
	bytes := make([]uint16, length)
	for i := range bytes {
		bytes[i] = reader.ReadUint16()
	}

	return bytes
}

func (reader *ClassReader) readBytes(length uint32) []byte {
	bytes := reader.data[:length]
	reader.data = reader.data[length:]

	return bytes
}
