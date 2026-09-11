package main

import (
	"fmt"

	"github.com/alok/podcast/auth"
	"github.com/alok/podcast/user"
	"github.com/fatih/color"
)

// Command for custom packages
// go mod init github.com/alok/podcast

func main() {
	auth.LoginWithCredentials("alok", "kola")
	session := auth.GetSession()
	fmt.Println(session)
}
