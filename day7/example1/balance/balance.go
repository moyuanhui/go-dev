package balance

type IBalancer interface {
	DoBalance(insts []*Instance, key ...string) (instance *Instance, err error)
}
