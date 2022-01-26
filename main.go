package main

import("flag"; "os"; "os/exec"; "github.com/ArthurHydr/GoScan/port")

func ClearString(){
  c := exec.Command("clear")
  c.Stdout = os.Stdout
  c.Run()
}

func main() {
  ClearString()
 
  host := flag.String("host", "default value", "string")
  flag.Parse() 

	port.GetPorts(*host, port.PortRange{Start: 1, End: 10000})
}
