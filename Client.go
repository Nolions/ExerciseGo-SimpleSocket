package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

var msg string

func main() {
	// creatre tcp connection
	conn, err := net.Dial("tcp", "localhost:4321")

	if err != nil {
		fmt.Println(err)
		return
	}

	log.Println("cliner run")

	run(conn)
}

func run(c net.Conn) {
	// 關閉後後執行
	defer c.Close()

	//
	// get user's input data by bufio
	//
	// Bulid a bufio's reader
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Please input a data:")

		//
		// get user's input data by bufio
		//
		data, _ := reader.ReadString('\n')
		data = strings.TrimSpace(data)

		//
		// get user's input data by scanf
		//
		// fmt.Scan(&data)

		buf := make([]byte, 1024)

		c.Write([]byte(data))
		// 讀取資料
		n, err := c.Read(buf)

		if err != nil {
			fmt.Println("Fail to read data, %s\n", err)
			continue
		}

		fmt.Println("status:", string(buf[0:n]))
	}
}
