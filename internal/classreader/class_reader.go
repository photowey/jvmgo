/*
 * Copyright 2022 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
