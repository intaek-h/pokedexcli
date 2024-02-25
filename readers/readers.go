package readers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadString() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Simple Shell")
	fmt.Println("------------")

	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("end", text) == 0 {
			fmt.Println("Goodbye!")
			break
		}

		if strings.Compare("hi", text) == 0 {
			fmt.Println("Hello, Yourself")
		}
	}
}

func ReadRune() {
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(char)

	switch char {
	case 'A':
		fmt.Println("A Key Pressed")
	case 'B':
		fmt.Println("B Key Pressed")
	}
}

func Scanner() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")

	for scanner.Scan() {
		text := scanner.Text()

		switch text {
		case "help":
			fmt.Println("Available commands: help, hi, end")
		default:
			fmt.Print("Pokedex > ")
		}
	}
}
