package CatDB

//4096bytes
type diskBlock struct {
	id           uint64
	leafSize     uint16
	childrenSize uint16
	data         []*KVPair //30 pairs
	childrenIDs  []uint64  //31 children
}
