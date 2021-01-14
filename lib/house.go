package lib

import (
	"fmt"
	"net"
)

// House model
type House struct {
	Name    string
	Members map[net.Addr]*Hamster
}

// Broadcast broadcasts a message to all house members
func (h *House) Broadcast(sender *Hamster, msg string) {
	for addr, m := range h.Members {
		if addr != sender.Conn.RemoteAddr() {
			m.Msg(msg)
		} else {
			sender.Conn.Write([]byte("   "))
		}
	}
}

// Search searches for a hamster by nickname if found, returns the hamster, otherwise nil
func (h *House) Search(search string) (*Hamster, error) {
	for _, m := range h.Members {
		if m.Nickname == search {
			return m, nil
		}
	}
	return nil, fmt.Errorf("Could not find a hamster called %s", search)
}
