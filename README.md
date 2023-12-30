## dnspyro 🎆🎇

Simple DNS analysis tool for the CLI, written entirely in the Go stdlib. Can resolve all major DNS record types and also resolves hostnames for IP adresses directly from the [RIPE](https://en.wikipedia.org/wiki/RIPE_NCC) API.

![image](https://github.com/parsec501/dnspyro/assets/105080989/f0886c88-359c-400d-be54-05c4ccf1a69f)

### Usage

`dnspyro <domain> -txt`

All you have to supply is a domain name. Resolving TXT records can optionally be enabled by including the --txt or -txt flag.

To build the project yourself, with the Go toolchain installed: `go build cmd/dnspyro/main.go`

