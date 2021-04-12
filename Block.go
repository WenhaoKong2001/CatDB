package CatDB

//4096bytes
type diskBlock struct {
	id           uint64    //4096-8=4088
	leafSize     uint8     //4088-1=4087
	childrenSize uint8     //4087-1=4086
	data         []*KVPair //30 pairs
	childrenIDs  []uint64  //31 children  248
}
