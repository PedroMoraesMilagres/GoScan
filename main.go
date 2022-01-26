package main

import("fmt"; "github.com/ArthurHydr")

func main() {
  result := port.ScanPort("scanme.nmap.org", 80)
  fmt.Println(result)
}


