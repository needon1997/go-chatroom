package main

import (
	"fmt"

	"../Entity"
	"../online"
	"../userprocess"
)

func login() (*Entity.User, bool) {
	var username, password string
	fmt.Println("please enter your username")
	_, err := fmt.Scanln(&username)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("please enter your password")
	_, err1 := fmt.Scanln(&password)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(username, password)
	user, err := userprocess.Login(username, password)
	if err != nil {
		fmt.Println(err)
	}
	if user != nil {
		loop := false
		return user, loop
	} else {
		fmt.Println("username or password error, return to main menu")
		loop := true
		return user, loop
	}
}
func register() {
	var username, password, nickname string
	fmt.Println("please enter your username")
	_, err := fmt.Scanln(&username)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("please enter your password")
	_, err = fmt.Scanln(&password)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("please enter your nickname")
	_, err = fmt.Scanln(&nickname)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(username, password, nickname)
	err = userprocess.Register(username, password, nickname)
	if err != nil {
		fmt.Println(err)
	}
}

func showMenu() {
	fmt.Println("-----------welecome to the chat app------------")
	fmt.Println("1.register-------------------------------------")
	fmt.Println("2.login----------------------------------------")
	fmt.Println("3.quit-----------------------------------------")
	fmt.Println("please enter (1-3)-----------------------------")
}
func checkUserInput() int {
	var input int
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println(err)
	}
	return input
}
func handleInput(input int, loopChan chan bool, exitChan chan bool) {
	switch input {
	case 1:
		register()
		loopChan <- true

	case 2:
		fmt.Println("login")
		user, loop := login()
		loopChan <- loop
		if !loop {
			close(loopChan)
			go online.Online(user, exitChan)
		}
	case 3:
		fmt.Println("exit")
		loopChan <- false
		exitChan <- true
		close(loopChan)
		close(exitChan)
	default:
		fmt.Println("please select 1-3")
		loopChan <- true
	}

}
func main() {
	var input int
	var loop bool = true
	loopChan := make(chan bool, 1)
	exitChan := make(chan bool, 1)
	for loop {
		showMenu()
		input = checkUserInput()
		handleInput(input, loopChan, exitChan)
		loop = <-loopChan
	}
	<-exitChan
	fmt.Println("bye")
}
