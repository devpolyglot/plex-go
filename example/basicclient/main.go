package main

import (
	"bufio"
	"fmt"
	"os"
	"syscall"

	"github.com/Lorac/plex-go/plex"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Print("Plex Username: ")
	username, _ := r.ReadString('\n')

	fmt.Print("Plex Password: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	password := string(bytePassword)

	tp := plex.PlexAuthTransport{}

	token, err := plex.RequestPlexToken(username, password, tp.Client())
	if err != nil {
		fmt.Println("Unable to fetch a token")
	}

	tp.XPlexToken = token

	client := plex.NewClient(tp.Client())
	//client.NewRequest("POST", "", nil)
}
