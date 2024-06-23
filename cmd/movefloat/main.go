package main

import (
	"log"

	"github.com/mattkasun/hypr"
)

func main() {
	active, err := hypr.ActiveWorkspace()
	if err != nil {
		log.Fatal(err)
	}
	windows, err := hypr.Clients()
	if err != nil {
		log.Fatal(err)
	}
	floating := []hypr.Client{}
	for _, window := range windows {
		if window.Workspace.ID == active.ID && window.Floating {
			floating = append(floating, window)
		}
	}
	if len(floating) > 0 {
		move(floating, "F"+active.Name)
	} else {
		restore(active, windows)
	}
}

func move(windows []hypr.Client, workspaceName string) {
	for _, window := range windows {
		hypr.MoveToWorkspaceSilent(window, workspaceName)
	}
}

func restore(workspace hypr.Workspace, windows []hypr.Client) {
	clientsToMove := []hypr.Client{}
	name := "F" + workspace.Name
	for _, client := range windows {
		if client.Workspace.Name == name {
			clientsToMove = append(clientsToMove, client)
		}
	}
	move(clientsToMove, workspace.Name)
}
