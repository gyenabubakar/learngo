package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type permissions map[string]bool

type user struct {
	Name        string      `json:"username"`
	Permissions permissions `json:"perms,omitempty"`
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	var __jsonBuf []byte

	for input.Scan() {
		__jsonBuf = append(__jsonBuf, input.Bytes()...)
	}

	if err := input.Err(); err != nil {
		fmt.Println("An error ocurred!", err)
		return
	}

	var users []user
	err := json.Unmarshal(__jsonBuf, &users)

	if err != nil {
		fmt.Println("An occurred during unmarshalling!", err)
		return
	}

	fmt.Printf("%#v\n", users)
}
