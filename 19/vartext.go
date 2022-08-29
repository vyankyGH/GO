package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	//var ip net.IP
	fs := flag.NewFlagSet("ExampleTextVar", flag.ContinueOnError)
	fs.SetOutput(os.Stdout)
	// Type "net.IPv4" implements the interface "encoding.TextMarshaler",
	// specifically the method "MarshalText"
	var ip net.IP
	fs.TextVar(&ip, "ip", net.IPv4(192, 168, 0, 100), "`IP address` to parse")
	fs.Parse([]string{"-ip", "127.0.0.1"})
	fmt.Printf("{ip: %v}\n\n", ip)
}
