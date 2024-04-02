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

func Prompt(message string, options []string) string {
	var response string
	for {
		if len(options) == 0 {
			fmt.Println(">>", message)
		} else {
			fmt.Printf(">> %s %s\n", message, options)
		}
		if _, e := fmt.Scan(&response); e != nil {
			panic(e)
		}
		if len(options) > 0 && IsSelection(response, options) {
			return response
		}
		if len(options) == 0 {
			return response
		}
	}
}
