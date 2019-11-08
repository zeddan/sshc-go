package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
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

	var servers []string

	scanner := bufio.NewScanner(f)
	i := 1
	for scanner.Scan() {
		servers = append(servers, scanner.Text())
		s := strings.Split(scanner.Text(), ",")
		fmt.Printf("%d) %s\n", i, s[0])
		i++
	}

	fmt.Println()
	fmt.Print("> ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")

	choice, err := strconv.Atoi(text)
	check(err)
	choice--

	server := strings.Split(servers[choice], ",")

	fmt.Printf("ssh %s@%s\n", server[1], server[2])

	cmd := exec.Command("ssh", fmt.Sprintf("%s@%s", server[1], server[2]))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	check(err)

}

func main() {
	connect()
}
