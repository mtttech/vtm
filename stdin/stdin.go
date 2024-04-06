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
		reader := bufio.NewReader(os.Stdin)
		fmt.Println(">>", message)
		if len(options) > 0 {
			for index, option := range options {
				fmt.Printf("%d.) %s\n", index+1, option)
			}
		}

		response, _ := reader.ReadString('\n')
		if response == "\n" {
			continue
		}

		response = strings.TrimSpace(response)
		if len(options) == 0 {
			return response
		}
		idx, _ := strconv.Atoi(response)
		if idx < 1 || idx > len(options) {
			continue
		}
		response = options[idx-1]
		if IsSelection(response, options) {
			return response
		}
	}
}
