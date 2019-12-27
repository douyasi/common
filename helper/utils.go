package helper

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

// format time
func FormatTime(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 15:04:05")
}

// md5(str)
func StrMD5(str string) string {
	md5 := md5.New()
	md5.Write([]byte(str))
	return hex.EncodeToString(md5.Sum(nil))
}

// PathExists
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// random string
// refer:
// https://colobu.com/2018/09/02/generate-random-string-in-Go/
// https://my.oschina.net/solate/blog/1793124
func RandomStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// get client ip
func ClientIP(r *http.Request) string {
	unknown := "unknown"
	proxyHeaders := []string{
		"CLIENT_IP",
		"FORWARDED",
		"FORWARDED_FOR",
		"FORWARDED_FOR_IP",
		"PROXY-CLIENT-IP",
		"HTTP_CLIENT_IP",
		"HTTP_FORWARDED",
		"HTTP_FORWARDED_FOR",
		"HTTP_FORWARDED_FOR_IP",
		"HTTP_VIA",
		"HTTP_X_FORWARDED",
		"HTTP_X_FORWARDED_FOR",
		"HTTP_X_FORWARDED_FOR_IP",
		"VIA",
		"X_FORWARDED",
		"X_FORWARDED_FOR",
		"X_REAL_IP",
	}
	ip, hv := "", ""
	for _, header := range proxyHeaders {
		hv = r.Header.Get(header)
		ip = strings.TrimSpace(strings.Split(hv, ",")[0])
		if ip != "" && ip != unknown {
			return ip
		}
	}
	rip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
	if err != nil {
		return "127.0.0.1"
	}
	return rip
}

// is or not in intranet environment by IP
func hasIntranetIP(ip net.IP) bool {
	if ip.IsLoopback() {
		return true
	}

	ip4 := ip.To4()
	if ip4 == nil {
		return false
	}

	return ip4[0] == 10 || // 10.0.0.0/8
		(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31) || // 172.16.0.0/12
		(ip4[0] == 169 && ip4[1] == 254) || // 169.254.0.0/16
		(ip4[0] == 192 && ip4[1] == 168) // 192.168.0.0/16
}

// is local ip address or not
// LAN IP address:
// 10.0.0.0/8, 172.16.0.0/12, 169.254.0.0/16, 192.168.0.0/16
func IsLocalIPAddress(ip string) bool {
	return hasIntranetIP(net.ParseIP(ip))
}