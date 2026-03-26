package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Menu() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Select an option [start/status/exit]: ")

	choice, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(choice), nil
}
