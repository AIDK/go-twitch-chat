package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	server := "irc.chat.twitch.tv:6667"

	// for this exercise we connect using
	// an anonymous account (twitch allows this)
	username := "justinfan12345"

	// connect
	conn, err := net.Dial("tcp", server)
	if err != nil {
		fmt.Println("error connecting to Twitch IRC", err)
		return
	}

	defer conn.Close()

	// pass authentication details (auth, username, channel)
	fmt.Fprintf(conn, "PASS oauth:fake\r\n")
	fmt.Fprintf(conn, "NICK %s\r\n", username)

	// replace the channel name placeholder with the actual
	// channel you want to connect to
	fmt.Fprintf(conn, "JOIN #<channel_name>\r\n")

	// read chat (forever)
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
