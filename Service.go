package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// create a Listener's service
	listener, err := net.Listen("tcp", "127.0.0.1:4321")

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Service run")

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Println(err)
			return
		}

		go run(conn)
	}
}

func run(c net.Conn) {
	// 關閉後後執行
	defer c.Close()

	// get Clinet ip address
	addr := c.RemoteAddr()

	log.Println("Connect :", addr)

	if c == nil {
		log.Println("Connection error")
	}

	for {
		// creater buf' Channel
		data := make([]byte, 1024)

		// 讀取數據
		read, err := c.Read(data)

		if read == 0 {
			log.Printf("%s has disconnect\n", addr)
			break
		}

		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Printf("%s say %s \n", addr, string(data[:read]))
		_, err = c.Write([]byte("success"))

		if err != nil {
			c.Write([]byte("false"))
			log.Println("ERROR : ", err)
			continue
		}
	}

}

// func connHandler(c net.Conn) {
// 	if c == nil {
// 		return
// 	}

// 	fmt.Println("Connect :", c.RemoteAddr())

// 	// creater buf' Channel
// 	buf := make([]byte, 1024)

// 	for {
// 		read, err := c.Read(buf)
// 		// get request's address
// 		addr := c.RemoteAddr()

// 		if read == 0 {
// 			fmt.Println("%s has disconnect\n", addr)
// 			break
// 		}

// 		if err != nil {
// 			fmt.Println(err)
// 			continue
// 		}

// 		fmt.Printf("Receive '%s' from '%s' on '%s'\n", string(buf[:read]), addr, time.Now())

// 		_, err = c.Write([]byte("success"))

// 		if err != nil {
// 			c.Write([]byte("false"))
// 			log.Println("ERROR : ", err)
// 			continue
// 		}
// 	}
// }
