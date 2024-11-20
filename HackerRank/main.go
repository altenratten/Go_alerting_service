package main

import (
    "fmt"
    "math/big"
    "net"
)

func IpsBetween(start, end string) int {
    //your code goes here
    return 0
}

func IP4toInt(IPv4Addr string) int64 {
    IPv4Int := big.NewInt(0)
    IPv4Int.SetBytes(net.ParseIP(IPv4Addr).To4())
    return IPv4Int.Int64()
}

func main() {
    fmt.Println(IP4toInt("10.0.0.50") - IP4toInt("10.0.0.0"))
}