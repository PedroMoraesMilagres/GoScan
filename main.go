package main

import("fmt"; "github.com/ArthurHydr/GoScan/port")

func main() {
  result := port.ScanPort("scanme.nmap.org", 80)
  fmt.Println(result)
}


