package CatDB

import (
	"encoding/binary"
	"os"
)

const PAGESIZE = 4096
const EIGHTBYTES = 8

type ConvertServer struct {
	file *os.File
}

func uint64ToBytes(num uint64) []byte {
	bytes := make([]byte, EIGHTBYTES)
	binary.BigEndian.PutUint64(bytes, num)
	return bytes
}

func uint64FromBytes(bytes []byte) uint64 {
	return binary.BigEndian.Uint64(bytes)
}

func (self *ConvertServer) convertBlockToBuffer(block *diskBlock) []byte {
	buf := make([]byte, PAGESIZE)
	offset := uint64(0)
	copy(buf[offset:], uint64ToBytes(block.id))
	offset += 8
	copy(buf[offset:], uint16ToBytes(block.leafSize))
	offset += 2
	copy(buf[offset:], uint16ToBytes(block.childrenSize))
	offset += 2

	for i := 0; i < len(block.data); i++ {
		copy(buf[offset:], converFromPairToBytes(block.data[i]))
		offset += MAXPAIRSIZE
	}

	for i := 0; i < len(block.childrenIDs); i++ {
		copy(buf[offset:], uint64ToBytes(block.childrenIDs[i]))
		offset += 8
	}

	return buf
}

func (self *ConvertServer) convertBufferToBlock(buf []byte) *diskBlock {
	return nil
}

//
func (self *ConvertServer) writeBlockToDisk(block *diskBlock) error {
	offset := block.id * PAGESIZE
	_, err := self.file.Seek(int64(offset), 0)
	if err != nil {
		return err
	}
	buf := self.convertBlockToBuffer(block)
	_, err = self.file.Write(buf)
	if err != nil {
		return err
	}

	return nil
}

func (self *ConvertServer) readBlockFromDisk(index uint64) (*diskBlock, error) {
	offset := index * PAGESIZE
	_, err := self.file.Seek(int64(offset), 0)
	if err != nil {
		return nil, err
	}
	buf := make([]byte, PAGESIZE)
	_, err = self.file.Read(buf)
	if err != nil {
		return nil, err
	}

	block := self.convertBufferToBlock(buf)

	return block, nil
}
