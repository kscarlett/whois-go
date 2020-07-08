package whois

import "io"

const iana_url string = "whois.iana.org"

// TODO: optimise TCP connections as much as possible
func Query(query, server string, w *io.Writer) {
	if server == "" {
		server = iana_url
	}
}
