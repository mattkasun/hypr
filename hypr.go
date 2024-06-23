package hypr

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

var hypr string

func init() {
	hypr, _ = exec.LookPath("hyprctl")
}

// ActiveWorkspace return the active workspace and its properties
func ActiveWorkspace() (Workspace, error) {
	workspace := Workspace{}
	bytes, err := exec.Command(hypr, "activeworkspace", "-j").Output()
	if err != nil {
		return workspace, err
	}
	err = json.Unmarshal(bytes, &workspace)
	return workspace, err
}

// ActiveWindow returns the active window and its properties
func ActiveWindow() (Client, error) {
	client := Client{}
	bytes, err := exec.Command(hypr, "activewindow", "-j").Output()
	if err != nil {
		return client, err
	}
	err = json.Unmarshal(bytes, &client)
	return client, err
}

// Clients returns all windows
func Clients() ([]Client, error) {
	clients := []Client{}
	bytes, err := exec.Command(hypr, "clients", "-j").Output()
	if err != nil {
		return clients, err
	}
	err = json.Unmarshal(bytes, &clients)
	return clients, err
}

// Workspaces returns all workspaces
func Workspaces() ([]Workspace, error) {
	workspaces := []Workspace{}
	bytes, err := exec.Command(hypr, "workspaces", "-j").Output()
	if err != nil {
		return workspaces, err
	}
	err = json.Unmarshal(bytes, &workspaces)
	return workspaces, err
}

// MoveToWorkspaceSilent moves a window to a workspace without switching to new workspace
func MoveToWorkspaceSilent(window Client, workspace string) {
	_, _ = exec.Command(hypr, "dispatch", "movetoworkspacesilent", fmt.Sprintf("name:%s,pid:%d", workspace, window.PID)).Output()
}
