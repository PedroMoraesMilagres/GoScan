package main

import("fmt"; "time"; "github.com/ArthurHydr/GoScan/port")

func main() {
	start := time.Now()
	port.GetOpenPorts("Businesscorp.com.br", port.PortRange{Start: 1, End: 10000})

	elapsed := time.Since(start)
	fmt.Printf("Scan duration: %s", elapsed)
}
