package main

import (
	"encoding/json"
	"net"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ipAddr := &IPInfo{
			IP: ip,
		}

		if ip := net.ParseIP(ip); ip == nil {
			ipAddr.Version = "Invalid IP"
		} else if ip.To4() != nil {
			ipAddr.Version = "IPv4"
		} else {
			ipAddr.Version = "IPv6"
		}

		ipAddr.IPv6 = getIPv6(ip)

		ipAddr.Proxy = r.Header.Get("X-Forwarded-For") != ""

		json.NewEncoder(w).Encode(ipAddr)

	})

	http.ListenAndServe(":8080", nil)
}

type IPInfo struct {
	IP      string `json:"ip"`
	Version string `json:"version"`
	IPv6    string `json:"ipv6"`
	Proxy   bool   `json:"isProxy"`
}

func getIPv6(ip string) string {

	// Lookup all IP addresses for host
	addrs, err := net.LookupIP(ip)
	if err != nil {
		return ""
	}

	// Find first IPv6 address
	for _, addr := range addrs {
		if ipv6 := addr.To16(); ipv6 != nil {
			return ipv6.String()
		}
	}
	return ""
}
