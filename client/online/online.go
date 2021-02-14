package online

import (
	"bufio"
	"fmt"
	"os"

	"../Entity"
	"../onlineMan"
)

var MsgWriteChan = make(chan Entity.Message, 5)

func Online(user *Entity.User, exitChan chan bool) {
	fmt.Printf("online %v\n", user.Nickname)
	loopChan := make(chan bool, 1)
	loop := true
	for loop {
		showMenu()
		input := checkUserInput()
		handleInput(input, user, loopChan, exitChan)
		loop = <-loopChan
	}
}

func showMenu() {
	fmt.Println("------------welecome to the chat app------------")
	fmt.Println("1. show online user-----------------------------")
	fmt.Println("2. send message---------------------------------")
	fmt.Println("3. message list---------------------------------")
	fmt.Println("4. quit-----------------------------------------")
	fmt.Println("please enter (1-4)------------------------------")
}

func checkUserInput() int {
	var input int
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println(err)
	}
	return input
}
func handleInput(input int, user *Entity.User, loopChan chan bool, exitChan chan bool) {
	switch input {
	case 1:
		fmt.Println("--------------show online user----------------")
		fmt.Printf("total online user:%v\n", onlineMan.Manager.GetTotal())
		fmt.Println(onlineMan.Manager)
		loopChan <- true
	case 2:
		fmt.Println("---------------send message-------------------")
		fmt.Println("please enter your message")
		in := bufio.NewReader(os.Stdin)
		msgbody, err := in.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			loopChan <- true
			return
		}
		msg := Entity.Message{Type: Entity.MSG}
		msg.Data = fmt.Sprintf("(%v,%v): %v", user.Nickname, user.Username, msgbody)
		MsgWriteChan <- msg
		loopChan <- true

	case 3:
		fmt.Println("message list")
		onlineMan.Manager.ShowMsg()
		loopChan <- true
	case 4:
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
