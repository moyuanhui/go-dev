package balance

import (
	"errors"
	"math/rand"
)

type RandomBalance struct {
}

func init() {
	RegisterBalancer("random", &RandomBalance{})
}

func (p *RandomBalance) DoBalance(insts []*Instance) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("No Instance")
		return
	}

	instsLen := len(insts)
	index := rand.Intn(instsLen)
	inst = insts[index]
	return

}
