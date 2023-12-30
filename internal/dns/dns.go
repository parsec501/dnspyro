package dns

import (
	"dnspyro/internal/cli"
)

// DNSHandler is the main function that handles the DNS queries
// through resolve.go
func DNSHandler(userConfig *UserConfig, resolvedStruct *cli.Resolved) []string {

	var errList []string

	err := userConfig.LookupARecord(resolvedStruct)
	if err != nil {
		errList = append(errList, cli.FmtErr("A/AAAA", err.Error()))
	}

	err = userConfig.LookupMXRecord(resolvedStruct)
	if err != nil {
		errList = append(errList, cli.FmtErr("MX", err.Error()))
	}

	err = userConfig.LookupNSRecord(resolvedStruct)
	if err != nil {
		errList = append(errList, cli.FmtErr("NS", err.Error()))
	}

	if userConfig.GetTxtSetting() {
		err = userConfig.LookupTXTRecord(resolvedStruct)
		if err != nil {
			errList = append(errList, cli.FmtErr("TXT", err.Error()))
		}
	}

	return errList
}
