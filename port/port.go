package port

import("fmt"; "net"; "strconv"; "sync"; "time" )

var commom = map[init]string {
  7:    "echo",
	20:   "ftp",
	21:   "ftp",
	22:   "ssh",
	23:   "telnet",
	25:   "smtp",
	43:   "whois",
	53:   "dns",
	67:   "dhcp",
	68:   "dhcp",
	80:   "http",
	110:  "pop3",
	123:  "ntp",
	137:  "netbios",
	138:  "netbios",
	139:  "netbios",
	143:  "imap4",
	443:  "https",
	513:  "rlogin",
	540:  "uucp",
	554:  "rtsp",
	587:  "smtp",
	873:  "rsync",
	902:  "vmware",
	989:  "ftps",
	990:  "ftps",
	1194: "openvpn",
	3306: "mysql",
	5000: "unpn",
	8080: "https-proxy",
	8443: "https-alt",
}

func ScanPort(hostname string, port int) bool {
  host := hostname + ":" + strconv.Itoa(port)
  conn, err := net.DialTimeout("tcp", host, 60*time.Second)
  
  if err != nil {
    return false
  }
  defer conn.Close()
  return true

}

