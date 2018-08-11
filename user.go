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
	"os/signal"
	"syscall"
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

	// Handle common process-killing signals so we can gracefully shut down:
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, os.Kill, syscall.SIGTERM)
	go func() {
		// Wait for a SIGINT or SIGKILL:
		<-sigc
		// Stop listening (and unlink the socket if unix type):
		l.Close()
		// And we're done:
		os.Exit(0)
	}()

	for {
		fd, err := l.Accept()
		if err != nil {
			log.Println("accept error", err)
			return
		}

		go ipcServer(db, fd)
	}
}

func listenIPC(db *BoltDB) {
	l, err := net.Listen("unix", executorSocket)
	if err != nil {
		log.Println("listen error", err)
		return
	}

	go acceptLoop(db, l)
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
		db.ListUsers(out)
		return
	}

	if len(command.email) == 0 {
		fmt.Fprintln(out, "user email is empty")
		return
	}

	if command.delete {
		token, err := db.GetTokenByEmail(command.email)
		if err != nil {
			log.Println(err)
		}

		if token != "" {
			err := db.DeleteUserByToken(token)
			if err != nil {
				log.Println(err)
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
				log.Println(err)
			}
		}

		token, err = db.CreateUser(command.email)
		if err != nil {
			log.Println(err)
		}

		fmt.Fprintln(out, command.email, token)
	} else {
		token, err := db.GetTokenByEmail(command.email)
		if err != nil {
			log.Println(err)
		}

		if token == "" {
			fmt.Fprintln(out, command.email+" not found")
		} else {
			fmt.Fprintln(out, command.email, token)
		}
	}
}
