package nsc

import (
	"errors"
	"fmt"
	"io"
	"net"
)

func sendToConnection (conn net.Conn, data []byte) (buf []byte, err error) {

	var count int
	count, err = fmt.Fprint(conn, string(data))

	if err != nil {
		err = errors.New(fmt.Sprintf("error returned from fmt.Fprint: %s", err.Error()))
		return
	}

	if count == 0 {
		err = errors.New("error with fmt.Fprint - no bytes written to connection")
		return
	}

	buf = make([]byte, 0, 4096)
	tmp := make([]byte, 256)
	for {
		n, err := conn.Read(tmp)
		if err != nil {
			if err != io.EOF {
				err = errors.New(fmt.Sprintf("error returned from conn.Read: %s", err.Error()))
			}
			break
		}
		buf = append(buf, tmp[:n]...)
	}

	return
}
