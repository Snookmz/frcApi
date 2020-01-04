package nsc

import (
	"errors"
	"fmt"
	"github.com/fatih/pool"
	"net"
)

func CreateConnectionPool (url string, port int) (p pool.Pool, err error) {
	var factory func() (net.Conn, error)
	factory = func() (net.Conn, error) {return net.Dial("tcp", fmt.Sprintf("%s:%v", url, port))}

	p, err = pool.NewChannelPool(5, 30, factory)
	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from pool.NewChannelPool: %s", err.Error()))
		return
	}

	return
}
