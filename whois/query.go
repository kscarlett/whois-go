package whois

import "io"

const iana_url string = "whois.iana.org"

// TODO: optimise TCP connections as much as possible
func Query(query string, w *io.Writer) {
	//
}
