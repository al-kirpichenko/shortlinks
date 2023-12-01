package ipcheker

import (
	"net"
	"net/http"
)

// CheckIP -
func CheckIP(r *http.Request, trustedSubnet string) bool {
	ip := r.Header.Get("X-Real-IP")
	_, ipNet, err := net.ParseCIDR(trustedSubnet)
	if err != nil {
		return false
	}
	return !ipNet.Contains(net.ParseIP(ip))
}
