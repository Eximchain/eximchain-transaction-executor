package main

import (
	"flag"
	"fmt"
	"log"
)

type UserCommand struct {
	email  string
	delete bool
	update bool
	list   bool
}

func RunUserCommand(args []string) {
	userCommand := flag.NewFlagSet("user", flag.ExitOnError)
	emailFlag := userCommand.String("email", "", "user email")
	deleteFlag := userCommand.Bool("delete", false, "delete user by email")
	updateFlag := userCommand.Bool("update", false, "update user token")
	listFlag := userCommand.Bool("list", false, "list all users")
	userCommand.Parse(args)

	command := UserCommand{email: *emailFlag, delete: *deleteFlag, update: *updateFlag, list: *listFlag}

	db := &BoltDB{}
	err := db.Open()
	if err != nil {
		log.Fatal(err)
	}

	if command.list {
		db.ListUsers()
		return
	}

	if len(command.email) == 0 {
		fmt.Println("user email is empty")
		return
	}

	if command.delete {
		token, err := db.GetTokenByEmail(command.email)
		if err != nil {
			log.Fatal(err)
		}

		if token != "" {
			err := db.DeleteUserByToken(token)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(command.email + " deleted")
		} else {
			fmt.Println("user not found")
		}
	} else if command.update {
		token, err := db.GetTokenByEmail(command.email)
		if token != "" {
			err = db.DeleteUserByToken(token)
			if err != nil {
				log.Fatal(err)
			}
		}

		token, err = db.CreateUser(command.email)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(command.email, token)
	} else {
		token, err := db.GetTokenByEmail(command.email)
		if err != nil {
			log.Fatal(err)
		}

		if token == "" {
			fmt.Println(command.email + " not found")
		} else {
			fmt.Println(command.email, token)
		}
	}
}
