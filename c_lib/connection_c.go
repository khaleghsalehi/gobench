package c_lib

import (
	"fmt"
	"net"
	"strings"
	"sync"
)

const (
	HttpServerSuccessResponse = "html"
	DEBUG                     = false
)

func NewConnection(target string, port string, protocol string, wg *sync.WaitGroup, c *Container) {
	//establish connection
	connection, err := net.Dial(protocol, target+":"+port)
	if err != nil {
		panic(err)
	}
	///send some data
	_, err = connection.Write([]byte("HELLO"))

	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		if DEBUG {
			fmt.Println("error reading buffer:", err.Error())
		}
	}
	if DEBUG {
		fmt.Println("received: ", string(buffer[:mLen]))
	}
	defer connection.Close()

	var tmp = string(buffer[:mLen])
	if strings.Contains(tmp, HttpServerSuccessResponse) ||
		len(string(buffer[:mLen])) > 0 {
		c.mu.Lock()
		c.Exec_counter++
		c.mu.Unlock()
	}
	wg.Done()

}
