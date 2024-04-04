package stdin

import (
	"fmt"
	"strconv"
)

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
		fmt.Println(">>", message)
		if len(options) > 0 {
			for index, option := range options {
				fmt.Printf("%d.) %s\n", index+1, option)
			}
		}
		if _, err := fmt.Scan(&response); err != nil {
			panic(err)
		}
		if len(options) > 0 {
			index_to_int, _ := strconv.Atoi(response)
			if index_to_int >= 1 && index_to_int <= len(options) {
				response = options[index_to_int-1]
				if IsSelection(response, options) {
					return response
				}
			}
		} else {
			return response
		}
	}
}
