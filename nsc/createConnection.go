package nsc

import (
	"errors"
	"fmt"
	"net"
)

func CreateConnection (url string, port int) (conn net.Conn, err error) {
	conn, err = net.Dial("tcp", fmt.Sprintf("%s:%v", url, port))

	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from net.Dial: %s", err.Error()))
	}

	return
}
