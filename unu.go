package main

import "net"
import "fmt"
import "bufio"
import "strings"

func main() {

	fmt.Println("Launching server...")

	ln, _ := net.Listen("tcp", ":8082")

	conn, _ := ln.Accept()

	for {
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print("Mesaj primit :", string(message))
		// sample process for string received
		newmessage := strings.ToUpper(message)

		if strings.ContainsAny(message, "a & e & i & u & o") {
			fmt.Println("Mesajul contine vocale ")
		}  else {
		fmt.Println("Mesajul nu contine vocale ")
		}


		conn.Write([]byte(newmessage + "\n"))
	}
}