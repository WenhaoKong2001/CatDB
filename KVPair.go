package CatDB

import (
	"encoding/binary"
	"errors"
)

//127 - 3 = 124
type KVPair struct {
	Keylen   uint16 //2
	Valuelen uint16 //2
	Key      string //30
	Value    string //90
}

const maxPairSize = 124
const maxKey = 30
const maxValue = 30

func (pair *KVPair) SetKey(key string) error {
	if len(key) > 30 {
		return errors.New("key is too big")
	}
	pair.Key = key
	pair.Keylen = uint16(len(key))
	return nil
}

func (pair *KVPair) SetValue(value string) error {
	if len(value) > 93 {
		return errors.New("value is too big")
	}
	pair.Value = value
	pair.Valuelen = uint16(len(value))
	return nil
}

func NewPair(key string, value string) (*KVPair, error) {
	pair := KVPair{}
	err := pair.SetKey(key)
	if err != nil {
		return nil, err
	}

	err = pair.SetValue(key)
	if err != nil {
		return nil, err
	}

	return &pair, nil
}

func convertFromBytesToPair(bytes []byte) *KVPair {
	offset := uint16(0)
	pair := KVPair{}
	pair.Keylen = uint16FromBytes(bytes[offset:])
	offset += 2
	pair.Valuelen = uint16FromBytes(bytes[offset:])
	offset += 2
	pair.Key = string(bytes[offset:pair.Keylen])
	offset += pair.Keylen
	pair.Value = string(bytes[offset:pair.Valuelen])
	offset += pair.Valuelen
	return &pair
}

func converFromPairToBytes(pair *KVPair) []byte {
	offset := uint16(0)
	bytes := make([]byte, maxPairSize)
	copy(bytes[offset:], uint16ToBytes(pair.Keylen))
	offset += 2
	copy(bytes[offset:], uint16ToBytes(pair.Valuelen))
	offset += 2
	copy(bytes[offset:], []byte(pair.Key))
	offset += pair.Keylen
	copy(bytes[offset:], []byte(pair.Value))
	offset += pair.Valuelen
	return nil
}

func uint16ToBytes(len uint16) []byte {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, len)
	return buf
}

func uint16FromBytes(bytes []byte) uint16 {
	len := binary.BigEndian.Uint16(bytes)
	return len
}
