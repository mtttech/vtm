package stdin

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	for {
		stdin := bufio.NewReader(os.Stdin)
		fmt.Println(">>", message)
		if len(options) > 0 {
			for index, option := range options {
				fmt.Printf("%d.) %s\n", index+1, option)
			}
		}

		response, err := stdin.ReadString('\n')
		if response == "\n" {
			continue
		}
		if err != nil {
			panic(err)
		}

		response = strings.TrimSpace(response)

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
