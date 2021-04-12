package CatDB

import "os"

type ConvertServer struct {
	file *os.File
}

//
func (self *ConvertServer) writeBlockToDisk(index uint64) error {
	return nil
}

func (self *ConvertServer) readBlockFromDisk(index uint64) error {
	return nil
}
