package ipcheker

import (
	"net"
	"net/http"
	"strings"
)

// CheckIP -
func CheckIP(r *http.Request, trustedSubnet string) bool {
	ip := strings.Split(r.Header.Get("X-Real-IP"), ":")[0]
	_, ipNet, err := net.ParseCIDR(trustedSubnet)
	if err != nil {
		return false
	}
	return !ipNet.Contains(net.ParseIP(ip))
}
