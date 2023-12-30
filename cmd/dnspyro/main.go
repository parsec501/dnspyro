package main

import (
	"dnspyro/internal/cli"
	"dnspyro/internal/dns"
	"fmt"
)

func main() {
	// userConfig contains the domain name and the checkTxt setting
	userConfig := dns.NewUserConfig(cli.GetUserInput())
	// resolvedStruct contains the results of the DNS queries
	resolvedStruct := cli.NewResolved()

	// doneLoading is a channel that is used to stop the loading animation
	doneLoading := make(chan bool)
	go cli.PrintLoading(doneLoading)

	// errList contains any errors that occurred during the DNS queries
	errList := dns.DNSHandler(userConfig, resolvedStruct)

	doneLoading <- true

	resolvedStruct.PrintResults()

	if len(errList) > 0 {
		for _, err := range errList {
			fmt.Println(err)
		}
	}
}
