package simplehttp

import (
	"bufio"
	"log"
	"net"
)

const CONN_TYPE = "tcp"

func TcpServer() {
	listener, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	log.Println("Listening on", CONN_HOST+":"+CONN_PORT)
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		log.Println(conn)
		handleRequest(conn)

	}

}

func handleRequest(conn net.Conn) {
	reader := bufio.NewReader(conn)
	message, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(message)

	conn.Close()
}
