package main

import (
	"fmt"
	"github.com/edwin-Marrima/Pod-net-route-guard/internal/iptables"
	_ "github.com/edwin-Marrima/Pod-net-route-guard/internal/iptables"
)

func main() {
	a, err := iptables.New()
	chains, err := a.ListChains("nat")
	fmt.Println(chains, " === ", err)
}
