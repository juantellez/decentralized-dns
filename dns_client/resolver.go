package dnsclient

import (
	"log"

	"github.com/miekg/dns"
)

// func handleDNSRequest(w dns.ResponseWriter, r *dns.Msg) {
func HandleDNSRequest(w dns.ResponseWriter, r *dns.Msg) {
	msg := new(dns.Msg)
	msg.SetReply(r)
	msg.Authoritative = true

	for _, q := range r.Question {
		domain := q.Name
		tld := getTLD(domain)

		log.Printf("[QUERY] Domain: %s | Type: %d", domain, q.Qtype)

		if reservedTLDs[tld] || !domainInBlockchain(domain) {
			log.Printf("[INFO] Using fallback resolver for domain: %s", domain)
			response, err := ResolveWithFallback(domain, q.Qtype)
			if err != nil || response == nil {
				log.Printf("[ERROR] Fallback resolver failed: %v", err)
				dns.HandleFailed(w, r)
				return
			}
			msg = new(dns.Msg)
			msg.SetReply(r)
			msg.Authoritative = false
			msg.RecursionAvailable = true
			msg.Answer = response.Answer
			msg.Ns = response.Ns
			msg.Extra = response.Extra
			log.Printf("[INFO] Fallback response successful for domain: %s", domain)
			break
		}

		record := blockchainRecords[domain]
		for _, ns := range record.NS {

			rr, err := dns.NewRR(domain + " 3600 IN NS " + ns)
			if err != nil {
				log.Printf("[ERROR] Failed to create RR for %s -> %s: %v", domain, ns, err)
				continue
			}
			msg.Ns = append(msg.Ns, rr)
		}
	}

	err := w.WriteMsg(msg)
	if err != nil {
		log.Printf("[ERROR] Failed to write DNS response: %v", err)
	} else {
		log.Printf("[INFO] Response successfully sent")
	}
}
