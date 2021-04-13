package CatDB

import (
	"fmt"
	"strconv"
	"testing"
)

func TestConvertBlockToBuffer(t *testing.T) {
	cs := ConvertServer{}
	pairs := make([]*KVPair, 0)

	for i := 50; i < 60; i++ {
		pair, _ := NewPair(strconv.Itoa(i), strconv.Itoa(i))
		pairs = append(pairs, pair)
	}

	block := diskBlock{
		id:           100,
		leafSize:     20,
		childrenSize: 30,
		data:         pairs,
		childrenIDs:  []uint64{1, 2, 3, 4, 5, 6, 7},
	}

	buf := cs.convertBlockToBuffer(&block)
	fmt.Print(buf)
}
