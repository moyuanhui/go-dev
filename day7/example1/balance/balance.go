package balance

type IBalancer interface {
	DoBalance(insts []*Instance) (instance *Instance, err error)
}
