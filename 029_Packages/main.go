package main

import (
	"fmt"

	"github.com/alok/podcast/auth"
	"github.com/alok/podcast/user"
	"github.com/fatih/color"
)

// Command for custom packages
// go mod init github.com/alok/podcast

// command for third party packages
// go get "package url"
// GOLANG only compile those file which have change not other files.


func main() {
	auth.LoginWithCredentials("alok", "kola")
	session := auth.GetSession()
	fmt.Println("session = ", session)
