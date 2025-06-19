package dnsclient

import (
	"time"

	"github.com/miekg/dns"
)

func resolveWithFallback(domain string, qtype uint16) (*dns.Msg, error) {
	client := &dns.Client{Timeout: 3 * time.Second}
	msg := new(dns.Msg)
	msg.SetQuestion(domain, qtype)

	r, _, err := client.Exchange(msg, "8.8.8.8:53")
	if err != nil || r == nil || r.Rcode != dns.RcodeSuccess {
		return nil, err
	}
	return r, nil
}
