package whois

import "errors"

const iana_url string = "whois.iana.org"

var (
	ErrConnectionIssue = errors.New("connection issue")
)

// TODO: optimise TCP connections as much as possible
func Query(query, server string) (err error, output string) {
	if server == "" {
		server = iana_url
	}
}
