package cli

import (
	"context"
	"fmt"

	"github.com/manifoldco/promptui"
)

func Menu(ctx context.Context) error {

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			shouldExit, err := mainMenu()
			if err != nil {
				return err
			}
			if shouldExit {
				return nil
			}
		}
	}
}

func mainMenu() (bool, error) {
	prompt := promptui.Select{
		Label: "Select an action",
		Items: []string{"Create note", "List notes", "Delete note", "Exit"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return false, err
	}

	switch result {
	// case "Create note":
	// 	return false, createMenu()
	// case "List notes":
	// 	return false, listMenu()
	// case "Delete note":
	// 	return false, deleteMenu()
	case "Exit":
		fmt.Println("Bye")
		return true, nil
	default:
		return false, nil
	}
}
