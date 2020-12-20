package main

import (
	"log"
	"net"
	"os"
	"time"
)

func main() {
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go copyToStderr(conn)

	}
}

func copyToStderr(conn net.Conn) {
	defer conn.Close()
	for {
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			log.Printf("Copied %d bytes; finished with err = %v", n, err)
			return
		}
		os.Stderr.Write(buf[:n])
	}

}

// func main() {
// 	dataStream, err := net.Listen("tcp", ":8080")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	defer dataStream.Close()

// 	for {
// 		conn, err := dataStream.Accept()
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		go handle(conn)
// 	}
// }

// func handle(conn net.Conn) {
// 	defer conn.Close()
// 	for {
// 		data, err := bufio.NewReader(conn).ReadString('\n')
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		fmt.Println(data)
// 	}

// }
