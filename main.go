package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	server := "irc.chat.twitch.tv:6667"
	username := "justinfan12345"

	conn, err := net.Dial("tcp", server)
	if err != nil {
		fmt.Println("error connecting to Twitch IRC", err)
		return
	}

	defer conn.Close()

	fmt.Fprintf(conn, "PASS oauth:fake\r\n")
	fmt.Fprintf(conn, "NICK %s\r\n", username)
	fmt.Fprintf(conn, "JOIN #Break\r\n")

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading from Twitch IRC", err)
			break
		}
		fmt.Print(message)
	}
}
