package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func connect() {
	f, err := os.Open("/Users/zeddan/.config/sshc/instances")
	check(err)
	defer f.Close()

	fmt.Println("Select instance:")
	fmt.Println()

	scanner := bufio.NewScanner(f)
	i := 1
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ",")
		fmt.Printf("%d) %s\n", i, s[0])
		i++
	}

	fmt.Println()
	fmt.Print("> ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}

func main() {
	connect()
}
