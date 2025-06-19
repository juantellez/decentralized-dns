package dnsclient

import (
	"log"
	"time"

	"github.com/juantellez/dns-chain/config"
	"github.com/miekg/dns"
)

func ResolveWithFallback(domain string, qtype uint16) (*dns.Msg, error) {
	client := &dns.Client{Timeout: 3 * time.Second}
	msg := new(dns.Msg)
	msg.SetQuestion(domain, qtype)

	for _, server := range config.Conf.FallbackDNSServers {
		r, _, err := client.Exchange(msg, server)
		if err != nil {
			log.Printf("[WARN] Fallback DNS server %s failed: %v", server, err)
			continue
		}
		if r == nil || r.Rcode != dns.RcodeSuccess {
			log.Printf("[WARN] Fallback DNS server %s returned invalid response", server)
			continue
		}
		return r, nil
	}
	return nil, nil
}
