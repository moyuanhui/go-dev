package balance

import "strconv"

type Instance struct {
	hostname string
	port     int
}

func NewInstance(hostname string, port int) *Instance {
	return &Instance{
		hostname: hostname,
		port:     port,
	}
}

func (p *Instance) GetHost() string {
	return p.hostname
}

func (p *Instance) GetPort() int {
	return p.port
}

func (p *Instance) String() string {
	return p.hostname + ":" + strconv.Itoa(p.port)
}
