package gofish

import "time"

const (
	UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.106 Safari/537.36"
	Qps = 50
)

var rateLimiter = time.Tick(time.Second/Qps)

type GoFish struct {
	Request *Request
}

func NewGoFish() *GoFish {
	return &GoFish{}
}

func (g *GoFish) Visit() error  {
	return g.Request.Do()
}
