package main

type PingModel interface {
	getPing() string
	getPong() string
}

type PingPong struct {
	ping string
	pong string
}

func (p *PingPong) getPing() string {
	return p.ping
}

func (p *PingPong) getPong() string {
	return p.pong
}
