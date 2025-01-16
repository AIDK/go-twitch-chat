package main

import (
	"bufio"
	"fmt"
	"net"
)

const SEVER string = "irc.chat.twitch.tv:6667"

func main() {

	// setup connection socket
	conn := connect_socket()
	defer conn.Close()

	// pass credentials
	set_credentials(conn)

	// read chat
	read_chat(conn)
}

func connect_socket() net.Conn {

	conn, err := net.Dial("tcp", SEVER)
	if err != nil {
		fmt.Println("error connecting to Twitch IRC", err)
		return nil
	}

	return conn

}

func set_credentials(conn net.Conn) {

	// for this exercise we connect using
	// an anonymous account (twitch allows this)
	username := "justinfan12345"

	// pass authentication details (auth, username, channel)
	fmt.Fprintf(conn, "PASS oauth:fake\r\n")
	fmt.Fprintf(conn, "NICK %s\r\n", username)

	// replace the channel name placeholder with the actual
	// channel you want to connect to
	fmt.Fprintf(conn, "JOIN #ThePrimeagen\r\n")
}

func read_chat(conn net.Conn) {

	reader := bufio.NewReader(conn)

	// loop forever 'listening' for any new messages
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading from Twitch IRC", err)
			break
		}
		fmt.Print(msg)
	}
}
