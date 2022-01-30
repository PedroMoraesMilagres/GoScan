package main

import(
  "fmt" 
  "flag" 
  "os" 
  "os/exec" 
  "github.com/ArthurHydr/GoScan/port"
)

func menu(){
  fmt.Println("  .d8888b.          .d8888b.                         ") 
  fmt.Println(" d88P  Y88b        d88P  Y88b                        ")
  fmt.Println(" 888    888        Y88b.                              ")
  fmt.Println(" 888        .d88b.  \"Y888b.   .d8888b 8888b. 88888b.  ")
  fmt.Println(" 888  88888d88\"\"88b    \"Y88b.d88P\"       \"88b888 \"88b ")
  fmt.Println(" 888    888888  888      \"888888     .d888888888  888 ")
  fmt.Println(" Y88b  d88PY88..88PY88b  d88PY88b.   888  888888  888 ")
  fmt.Println("  \"Y8888P88 \"Y88P\"  \"Y8888P\"  \"Y8888P\"Y888888888  888 ")
  fmt.Println("\nGoScan v1.0\n")
}

func ClearString(){
  c := exec.Command("clear")
  c.Stdout = os.Stdout
  c.Run()
}

func Scan(){
  host := flag.String("host", "", "Misuse! Correct way to use: ./main.go -host <address>")
  flag.Parse() 
  
  if *host == "" {
    fmt.Println("Misuse! Correct way to use: ./main.go -host <address>")
    os.Exit(1)
  } else {
	    port.GetPorts(*host, port.PortRange{Start: 1, End: 31337})
      println("\n")
  }
}

func main() {
  ClearString()
  menu()
  Scan()
 }
