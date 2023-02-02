package main

import (
	"bencher/c_lib"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func echo_banner() {
	println("+-----------------------------------------+")
	println("| Server Benchmark Toolkit                |")
	println("| For learning Go!                        |")
	println("+-----------------------------------------+")
}
func main() {
	echo_banner()
	//todo input validation
	if len(os.Args) < 4 {
		println("Usage:")
		println(os.Args[0] + " <target> <port> <protocol[tcp|udp]> thread")
		println("Example:")
		println(os.Args[0] + " 127.0.0.1 80 tcp 100")
		os.Exit(0)
	}

	var target = os.Args[1]
	var port = os.Args[2]
	var protocol = os.Args[3]

	var start_t = time.Now().UnixNano() / int64(time.Millisecond)
	var wait_group sync.WaitGroup
	var max_thread, _ = strconv.Atoi(os.Args[4])
	var c c_lib.Container
	var total_req = 0
	wait_group.Add(max_thread)
	for i := 0; i < max_thread; i++ {
		go c_lib.NewConnection(target, port, protocol, &wait_group, &c)
		total_req++

	}
	wait_group.Wait()
	var end_t = time.Now().UnixNano() / int64(time.Millisecond)
	println("")
	println("Test Result:")
	println("+---------------------------------------------+")
	fmt.Printf("Test execution time -> %d milliseconds \n", end_t-start_t)
	fmt.Printf("Total request %d, Passed %d , %d Failed.", total_req, c.Exec_counter,
		total_req-c.Exec_counter)

}
