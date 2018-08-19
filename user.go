package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"reflect"
	"strings"
)

const executorSocket = "/tmp/executor.sock"

type UserCommand struct {
	email  string
	delete bool
	update bool
	list   bool
}

func ipcServer(db *BoltDB, c net.Conn) {
	defer c.Close()
	dec := gob.NewDecoder(c)
	var args []string
	err := dec.Decode(&args)
	if err != nil {
		log.Println("decode error", err)
		return
	}
	log.Println(args)
	// Echo the reply to the server and the client
	writer := io.MultiWriter(c, os.Stdout)
	runUserCommand(db, writer, args)
}

func acceptLoop(db *BoltDB, l net.Listener) {
	// Unix sockets must be unlink()ed before being reused again.
	// Close() will do unlinking if listener is of type UnixListener.
	defer l.Close()

	for {
		fd, err := l.Accept()
		if err != nil {
			if !strings.HasSuffix(err.Error(), "use of closed network connection") {
				log.Println("accept error", err, reflect.TypeOf(err))
			}
			return
		}

		ipcServer(db, fd)
	}
}

func listenIPC(db *BoltDB) net.Listener {
	l, err := net.Listen("unix", executorSocket)
	if err != nil {
		log.Println("listen error", err)
		return nil
	}

	go acceptLoop(db, l)
	return l
}

func sendIPC(args []string) string {
	c, err := net.Dial("unix", executorSocket)
	if err != nil {
		log.Println("dial error", err)
		return ""
	}
	defer c.Close()
	enc := gob.NewEncoder(c)
	err = enc.Encode(args)
	if err != nil {
		log.Println("encode error", err)
		return ""
	}
	data, err := ioutil.ReadAll(c)
	if err != nil {
		log.Println("read error", err)
		return ""
	}
	return string(data)
}

func openUserDB() (*BoltDB, error) {
	db := &BoltDB{}
	err := db.Open("eximchain.db")
	return db, err
}

func RunUserCommand(args []string) {
	db, err := openUserDB()
	if err != nil {
		// If the open timed out, the server is likely running; send over IPC
		if err.Error() == "timeout" {
			fmt.Print(sendIPC(args))
			return
		}
		log.Println("open database error", err)
		return
	}

	runUserCommand(db, os.Stdout, args)
}

func runUserCommand(db *BoltDB, out io.Writer, args []string) {
	userCommand := flag.NewFlagSet("user", flag.ExitOnError)
	emailFlag := userCommand.String("email", "", "user email")
	deleteFlag := userCommand.Bool("delete", false, "delete user by email")
	updateFlag := userCommand.Bool("update", false, "update user token")
	listFlag := userCommand.Bool("list", false, "list all users")
	userCommand.Parse(args)

	command := UserCommand{email: *emailFlag, delete: *deleteFlag, update: *updateFlag, list: *listFlag}

	if command.list {
		err := db.ListUsers(out)
		if err != nil {
			log.Println("ListUsers", err)
		}
		return
	}

	if len(command.email) == 0 {
		fmt.Fprintln(out, "user email is empty")
		return
	}

	if command.delete {
		token, err := db.GetTokenByEmail(command.email)
		if err != nil {
			log.Println("GetTokenByEmail", err)
		}

		if token != "" {
			err := db.DeleteUserByToken(token)
			if err != nil {
				if err != nil {
					log.Println("DeleteUserByToken error", err)
				}
			}

			fmt.Fprintln(out, command.email+" deleted")
		} else {
			fmt.Fprintln(out, "user not found")
		}
	} else if command.update {
		token, err := db.GetTokenByEmail(command.email)
		if token != "" {
			err = db.DeleteUserByToken(token)
			if err != nil {
				log.Println("DeleteUserByToken error", err)
			}
		}

		token, err = db.CreateUser(command.email)
		if err != nil {
			log.Println("CreateUser", err)
		}

		fmt.Fprintln(out, command.email, token)
	} else {
		token, err := db.GetTokenByEmail(command.email)
		if err != nil {
			log.Println("GetTokenByEmail", err)
		}

		if token == "" {
			fmt.Fprintln(out, command.email+" not found")
		} else {
			fmt.Fprintln(out, command.email, token)
		}
	}
}
