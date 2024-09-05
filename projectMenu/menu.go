package main

import (
	"fmt"
	"os/exec"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

func playSnake() {
	snakeApp := "../snakeGame/game"

	startCmd := exec.Command(snakeApp)

	startGame, startGameErr := startCmd.Output()

	if startGameErr != nil {
		fmt.Println(startGameErr.Error())
		return
	}

	fmt.Println(string(startGame))
}

func menuListener(key keys.Key) (stop bool, err error) {
	switch key.Code {
	case keys.RuneKey:
		if key.String() == "1" {
			playSnake()
			return true, nil
		}

	default:
		fmt.Printf("\rYou pressed: %s, please press a key listed above.", key)
	}
	return false, nil
}

func main() {
	availablePrograms := [1]string{"snakeGame"}

	fmt.Println("Select the option youd like to try!")

	for i := 0; i < len(availablePrograms); i++ {
		fmt.Printf("Press: %v for %v\n", i+1, availablePrograms[i])
	}

	keyboard.Listen(menuListener)
}
