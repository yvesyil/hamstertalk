package lib

import "net"

type TunnelNode struct {
	House    *House
	Next     *TunnelNode
	Previous *TunnelNode
}

// Tunnel model
type Tunnel struct {
	ID    int
	Users map[net.Addr]*Hamster
	Head  *TunnelNode
}
