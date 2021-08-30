package main

import (
	"encoding/json"
	"fmt"
)

type permissions map[string]bool

type user struct {
	Name        string      `json:"username"`
	Password    string      `json:"-"`
	Permissions permissions `json:"perms,omitempty"`
}

func main() {
	me := user{
		Name:     "Gyen Abubakar",
		Password: "Age8924@",
		Permissions: permissions{
			"admin": true,
			"premium": false,
		},
	}

	users := []user{
		me,
		{
			Name:     "Sena Godson",
			Password: "kjfj#@jb09",
			Permissions: permissions{
				"admin": false,
				"premium": true,
			},
		},
		{
			Name:     "DeGraft Arthur",
			Password: "kjfj!@JNK",
			Permissions: permissions{
				"admin": true,
				"premium": true,
			},
		},
	}

	fmt.Printf("%#v\n", users)

	__json, _ := json.MarshalIndent(users, "", "   ")

	fmt.Println(string(__json))
}
