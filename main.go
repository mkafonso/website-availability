package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
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
		fmt.Println("Two")
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
		return
	}

	fmt.Println("Status Not OK", response.StatusCode)
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

func main() {
	for {
		presentation()

		userInput := selectOption()
		handleCommand(int(userInput))
		fmt.Printf("\n\n")
	}
}
