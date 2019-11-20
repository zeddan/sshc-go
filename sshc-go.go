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
	out := fmt.Sprintf("%-"+widths[0]+"s %-"+widths[1]+"s %s", name, user, ip)
	fmt.Println(out)
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

func promptAdd(value string, str string) string {
	if len(str) == 0 {
		fmt.Print(fmt.Sprintf("%s: ", str))
	} else {
		fmt.Print(fmt.Sprintf("%s (%s): ", str, value))
	}
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	tmpValue := strings.TrimSuffix(text, "\n")
	if len(tmpValue) == 0 {
		return value
	} else {
		return tmpValue
	}
}

func addInstance(tmpName string, tmpUser string, tmpIP string) {
	name := ""
	if len(tmpName) != 0 {
		name = tmpName
	}
	name = promptAdd(name, "Name")

	user := ""
	if len(tmpUser) != 0 {
		user = tmpUser
	}
	user = promptAdd(user, "User")

	ip := ""
	if len(tmpIP) != 0 {
		ip = tmpIP
	}
	ip = promptAdd(ip, "IP")

	fmt.Print(fmt.Sprintf("Add %s for %s@%s? (Y/n) ", name, user, ip))

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	confirm := strings.TrimSuffix(text, "\n")

	if confirm == "y" || confirm == "Y" || confirm == "" {
		filename := "/Users/zeddan/.config/sshc/instances"
		f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
		check(err)
		defer f.Close()
		text = fmt.Sprintf("%s,%s,%s", name, user, ip)
		if _, err = f.WriteString(text); err != nil {
			panic(err)
		}
		fmt.Println("OK!")
	} else {
		addInstance(name, user, ip)
	}

}

func main() {
	//connect()
	//list()
	addInstance("", "", "")
}
