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

func findWidths(servers [][]string) []string {
	var widths []string
	longest := 0
	for i := range servers {
		length := len(servers[i][0])
		if length > longest {
			longest = length
		}
	}
	widths = append(widths, strconv.Itoa(longest+2))
	longest = 0
	for i := range servers {
		length := len(servers[i][1])
		if length > longest {
			longest = length
		}
	}
	widths = append(widths, strconv.Itoa(longest+2))
	return widths
}

func prettyPrint(name string, user string, ip string, widths []string) {
	col1, _ := strconv.Atoi(widths[0])
	col2, _ := strconv.Atoi(widths[1])
	x := fmt.Sprintf("%-"+fmt.Sprintf("%d", col1)+"s %-"+fmt.Sprintf("%d", col2)+"s %s", name, user, ip)
	fmt.Println(x)
}

func list() {
	f, err := os.Open("/Users/zeddan/.config/sshc/instances")
	check(err)
	defer f.Close()

	var servers [][]string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		server := strings.Split(scanner.Text(), ",")
		servers = append(servers, server)
	}

	widths := findWidths(servers)
	prettyPrint("NAME", "USER", "IP", widths)
	for i := range servers {
		name := servers[i][0]
		user := servers[i][1]
		ip := servers[i][2]
		prettyPrint(name, user, ip, widths)
	}
}

func main() {
	//connect()
	list()
}
