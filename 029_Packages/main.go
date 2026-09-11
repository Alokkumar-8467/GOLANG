package main

import (
	"fmt"

	"github.com/alok/podcast/auth"
	"github.com/alok/podcast/user"
	"github.com/fatih/color"
)

// go mod init github.com/alok/podcast

func main() {
	auth.LoginWithCredentials("alok", "kola")
	session := auth.GetSession()
	fmt.Println(session)
}
