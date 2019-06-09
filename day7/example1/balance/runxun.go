package balance

import "errors"

type RunxunBalance struct {
	curIndex int
}

func init() {
	RegisterBalancer("runxun", &RunxunBalance{})
}

func (p *RunxunBalance) DoBalance(insts []*Instance) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("Instance is empty")
		return
	}
	instLen := len(insts)

	if p.curIndex >= instLen {
		p.curIndex = 0
	}
	inst = insts[p.curIndex]
	p.curIndex = (p.curIndex + 1) % instLen
	return
}
