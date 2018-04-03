package servactive

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func isStop(word string) bool {
	word = strings.TrimSpace(strings.ToLower(word))
	switch word {
	case "n", "no", "nope", "stop", "q", "quit", "exit", "e":
		return true
	default:
		return false
	}
}

func withPromt(promt string) (string, bool) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("%s", promt)
	var input string
	for scanner.Scan() {
		input = scanner.Text()
		break
	}
	return input, scanner.Err() == io.EOF
}

func yes(message string) (bool, string) {
	fmt.Printf("%s [Y/N]: ", message)
	scanner := bufio.NewScanner(os.Stdin)
	answer := "N"
	for scanner.Scan() {
		answer = strings.TrimSpace(scanner.Text())
		break
	}
	return strings.ToLower(answer) == "y", answer
}

func askWord(message string) (string, bool) {
	fmt.Printf("%s : ", message)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		return scanner.Text(), true
	}
	return "", false
}

func askFieldToChange(fields []string) (int, bool) {
	for i, field := range fields {
		fmt.Printf("%d) %s\n", i+1, field)
	}
	for {
		answer, ok := askWord(fmt.Sprintf("Which field do you want to change? (print no to stop)"))
		if !ok {
			return -1, false
		}
		switch strings.ToLower(answer) {
		case "no", "n", "-", "skip", "exit", "q", "quit":
			return -1, false
		}
		field, err := strconv.Atoi(answer)
		if err != nil {
			fmt.Printf("Please, type number!\n")
			continue
		}
		if field < 1 || field > len(fields) {
			fmt.Printf("Please, value have to be in range 1-%d\n", len(fields))
		}
		return field - 1, true
	}
}
