package main

import (
	"fmt"
	"log"

	"github.com/mattkasun/hypr"
)

func main() {
	window, err := hypr.ActiveWindow()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(window)
	workspace, err := hypr.ActiveWorkspace()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(workspace)
	clients, err := hypr.Clients()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(clients)
	spaces, err := hypr.Workspaces()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(spaces)
}
