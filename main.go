package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const MONITORING_TIMES = 3
const DELAY_BETWEEN_MONITORING = 3 * time.Second // 3 seconds

func presentation() {
	fmt.Println("Check Website Availability")
	fmt.Println("..........................")

	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("3 - Exit")
	fmt.Println()
	fmt.Println()
}

func selectOption() int8 {
	var input int8
	fmt.Print("Select an option: ")
	fmt.Scanf("%d", &input)

	return input
}

func handleCommand(input int) {
	switch input {
	case 1:
		startMonitoring()
	case 2:
		showLogs()
	case 3:
		os.Exit(0)
	default:
		fmt.Println("Invalid input. Please select a number between 1 and 3")
		os.Exit(-1)
	}
}

func checkWebsiteStatus(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Printf("Error checking the website")
		return
	}

	if response.StatusCode == http.StatusOK {
		fmt.Printf("website: %-40v  StatusOK: %v \n", site, http.StatusOK)
		handleLogFile(site, true)
		return
	}

	fmt.Println("Status Not OK", response.StatusCode)
	handleLogFile(site, false)
}

func startMonitoring() {
	sites := readFromFile()

	for i := 0; i < MONITORING_TIMES; i++ {
		for _, site := range sites {
			checkWebsiteStatus(site)
		}

		time.Sleep(DELAY_BETWEEN_MONITORING)
	}
}

func showLogs() {
	file, err := ioutil.ReadFile("logs.txt")

	if err != nil {
		fmt.Printf("Error opening logs.txt")
	}

	clearScreen()

	fmt.Println(string(file))
}

func readFromFile() []string {
	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Printf("Error opening sites.txt")
		return nil
	}

	reader := bufio.NewReader(file)

	for {
		row, err := reader.ReadString('\n')
		sites = append(sites, strings.TrimSpace(row))
		if err == io.EOF {
			break
		}
	}

	return sites
}

func handleLogFile(site string, isAvailable bool) {
	file, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Printf("Error opening logs.txt")
	}

	date := time.Now().Format("01/02/2006 15:04:05PM")
	log := fmt.Sprintf("%-25v website: %-40v  IsOnline: %v \n", date, site, isAvailable)

	file.WriteString(log)

	file.Close()
}

func clearScreen() {
	var clearCmd *exec.Cmd

	if runtime.GOOS == "windows" {
		clearCmd = exec.Command("cmd", "/c", "cls")
	} else {
		clearCmd = exec.Command("clear")
	}

	clearCmd.Stdout = os.Stdout
	clearCmd.Run()
}

func main() {
	for {
		presentation()

		userInput := selectOption()
		handleCommand(int(userInput))
		fmt.Printf("\n\n")
	}
}
