package main

import("flag"; "os"; "os/exec"; "github.com/ArthurHydr/GoScan/port")

func menu(){
  println("  .d8888b.          .d8888b.                         ") 
  println(" d88P  Y88b        d88P  Y88b                        ")
  println(" 888    888        Y88b.                              ")
  println(" 888        .d88b.  \"Y888b.   .d8888b 8888b. 88888b.  ")
  println(" 888  88888d88\"\"88b    \"Y88b.d88P\"       \"88b888 \"88b ")
  println(" 888    888888  888      \"888888     .d888888888  888 ")
  println(" Y88b  d88PY88..88PY88b  d88PY88b.   888  888888  888 ")
  println("  \"Y8888P88 \"Y88P\"  \"Y8888P\"  \"Y8888P\"Y888888888  888 ")
  println("\nGoScan v1.0\n")
}


func ClearString(){
  c := exec.Command("clear")
  c.Stdout = os.Stdout
  c.Run()
}

func main() {
  ClearString()
  menu()
 
  host := flag.String("host", "default value", "string")

  flag.Parse() 

	port.GetPorts(*host, port.PortRange{Start: 1, End: 31337})
}
