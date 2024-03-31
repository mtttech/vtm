package stdin

import "fmt"

func IsSelection(option string, options []string) bool {
	for _, value := range options {
		if option == value {
			return true
		}
	}
	return false
}

func Input(message string) string {
	var stdin string
	fmt.Println(">>", message)
	fmt.Scan(&stdin)
	return stdin
}
