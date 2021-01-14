package lib

import "net"

// Tunnel model
type Tunnel struct {
	ID       int
	Users    map[net.Addr]*Hamster
	Current  *House
	Next     *House
	Previous *House
}
