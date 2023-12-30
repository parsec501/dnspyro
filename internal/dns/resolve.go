package dns

import (
	"dnspyro/internal/cli"
	"fmt"
	"net"
)

type UserConfig struct {
	domainName string
	checkTxt   bool
}

func NewUserConfig(domainName string, checkTxt bool) *UserConfig {
	return &UserConfig{
		domainName: domainName,
		checkTxt:   checkTxt,
	}
}

func (d *UserConfig) GetTxtSetting() bool {
	return d.checkTxt
}

// LookupARecord performs a DNS lookup for A/AAAA records and
// calls ResolveHostAPI to get the hostname for each IP address
func (d *UserConfig) LookupARecord(r *cli.Resolved) error {
	aRecords, err := net.LookupIP(d.domainName)
	if err != nil {
		return err
	}

	for _, ip := range aRecords {
		host, err := ResolveHostAPI(ip.String())
		if err != nil {
			host = "N/A"
		}

		r.SetA(fmt.Sprintf("%s (%s)", ip.String(), host))
	}
	return nil
}

func (d *UserConfig) LookupMXRecord(r *cli.Resolved) error {
	mxRecords, err := net.LookupMX(d.domainName)
	if err != nil {
		return err
	}

	for _, mx := range mxRecords {
		r.SetMX(mx.Host[:len(mx.Host)-1])
	}
	return nil
}

func (d *UserConfig) LookupNSRecord(r *cli.Resolved) error {
	nsRecords, err := net.LookupNS(d.domainName)
	if err != nil {
		return err
	}

	for _, ns := range nsRecords {
		r.SetNS(ns.Host[:len(ns.Host)-1])
	}

	return nil
}

// LookupTXTRecord performs a DNS lookup for TXT records if the
// user has specified the --txt flag
func (d *UserConfig) LookupTXTRecord(r *cli.Resolved) error {
	txtRecords, err := net.LookupTXT(d.domainName)
	if err != nil {
		return err
	}

	for _, txt := range txtRecords {
		r.SetTXT(txt)
	}

	return nil
}
