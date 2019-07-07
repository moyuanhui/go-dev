package main

import (
	"fmt"
	"go_dev/day7/example1/balance"
	"hash/crc32"
	"math/rand"
)

type HashBalance struct {
}

func init() {
	balance.RegisterBalancer("hash", &HashBalance{})
}

func (p *HashBalance) DoBalance(insts []*balance.Instance, key ...string) (inst *balance.Instance, err error) {

	// defKey := rand.Int31n()
	defKey := fmt.Sprintf("%d", rand.Int())
	if len(key) > 0 {
		// defKey = fmt.Sprintf("%d", rand.Int())
		defKey = key[0]
		// err = fmt.Errorf("hash balance must pass hash key")
		return
	}
	lens := len(insts)
	if lens == 0 {
		err = fmt.Errorf("No backend instance")
		return
	}
	crcTable := crc32.MakeTable(crc32.IEEE)
	hashVal := crc32.Checksum([]byte(defKey), crcTable)
	index := int(hashVal) % lens
	inst = insts[index]
	return
}
