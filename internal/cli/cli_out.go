package cli

import (
	"fmt"
	"os"
	"time"
)

const (
	USAGE_MSG   = "Usage: dnspyro <domain> [--txt] \nExample: dnspyro example.com --txt"
	SPINNER     = "⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏"
	LOADING_MSG = "🎆 Gathering info..."
)

func PrintLoading(done chan bool) {
	for {
		select {
		case <-done:
			fmt.Printf("\r\033[K")
			return
		default:
			for _, char := range SPINNER {
				fmt.Printf("\r\033[38;5;208m%s\033[0m %s", string(char), LOADING_MSG)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
}

func Println(recordType string, body string) {
	for len(recordType) < 8 {
		recordType += " "
	}

	if len(body) == 0 {
		body = "N/A"
	}

	fmt.Printf("\033[34m%s\033[0m %s\n", recordType, body)
}

// FmtErr formats the error message for the errList to be printed
// at the end of the program
func FmtErr(recordType string, err string) string {

	for len(recordType) < 8 {
		recordType += " "
	}

	return fmt.Sprintf("\033[33m%s\033[0m %s", recordType, err)
}

// PrintUsage prints the usage message and exits when the user
// provides invalid input
func PrintUsage() {
	fmt.Println(USAGE_MSG)
	os.Exit(1)
}

func (r *Resolved) PrintResults() {
	for _, a := range r.a {
		Println("A/AAAA", a)
	}

	for _, mx := range r.mx {
		Println("MX", mx)
	}

	for _, ns := range r.ns {
		Println("NS", ns)
	}

	if len(r.txt) > 0 {
		for _, txt := range r.txt {
			Println("TXT", txt)
		}
	} else {
		fmt.Println("\033[38;5;8mSkipping TXT records, use -txt to enable\033[0m")
	}
}
