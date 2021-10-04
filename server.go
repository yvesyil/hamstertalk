package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/HawkBrave/Hamstertalk/lib"
)

type server struct {
	houses   map[string]*lib.House
	tunnels  map[int]*lib.Tunnel
	commands chan lib.Command
}

func newServer() *server {
	return &server{
		houses:   make(map[string]*lib.House),
		commands: make(chan lib.Command),
	}
}

func (s *server) run() {
	for cmd := range s.commands {
		switch cmd.ID {
		case lib.CmdSet:
			s.cmdSet(cmd.Hamster, cmd.Args)
		case lib.CmdUse:
			s.cmdUse(cmd.Hamster, cmd.Args)
		case lib.CmdStepto:
			s.cmdStepto(cmd.Hamster, cmd.Args)
		case lib.CmdHopto:
			s.cmdHopto(cmd.Hamster, cmd.Args)
		case lib.CmdList:
			s.cmdList(cmd.Hamster, cmd.Args)
		case lib.CmdSqueakto:
			s.cmdSqueakto(cmd.Hamster, cmd.Args)
		case lib.CmdExit:
			s.cmdExit(cmd.Hamster, cmd.Args)
		case lib.CmdQuit:
			s.cmdQuit(cmd.Hamster)
		case lib.Message:
			s.message(cmd.Hamster, cmd.Args)
		}
	}
}

func (s *server) newHamster(conn net.Conn) {
	log.Printf("New hamster on the block! They connected with %s.", conn.RemoteAddr().String())

	h := &lib.Hamster{
		Conn:     conn,
		Nickname: "anon",
		Commands: s.commands,
	}

	for {
		h.Conn.Write([]byte(">> "))
		err := h.ReadInput()
		if err != nil {
			break
		}
	}
}

var errNotInHouse = errors.New("You are not currently in a house")

func (s *server) exitCurrentHouse(h *lib.Hamster) {
	if h.House != nil {
		delete(h.House.Members, h.Conn.RemoteAddr())
		h.House.Broadcast(h, fmt.Sprintf("%s has left the house", h.Nickname))
	}
}

func (s *server) exitCurrentTunnel(h *lib.Hamster) {
	if h.Tunnel != nil {
		delete(h.Tunnel.Users, h.Conn.RemoteAddr())
	}
}

func (s *server) cmdSet(h *lib.Hamster, args []string) {
	if args[1] == "nickname" {
		h.Nickname = args[2]
		if h.Nickname != "anon" {
			h.Msg(fmt.Sprintf("Now you shall be known as \"%s\".", h.Nickname))
		} else {
			h.Err(errors.New("You cannot take a preserved username"))
		}
	}
}

func (s *server) cmdUse(h *lib.Hamster, args []string) {
	if args[1] == "tunnel" {
		tid, err := strconv.Atoi(args[2])
		if err != nil {
			h.Err(fmt.Errorf("Couldn't convert tunnel id %s to int", err.Error()))
			return
		}
		for id := range s.tunnels {
			if tid == id {
				h.Tunnel = s.tunnels[id]
			}
		}
	}
}

func (s *server) cmdStepto(h *lib.Hamster, args []string) {

}

func (s *server) cmdHopto(h *lib.Hamster, args []string) {
	houseName := args[1]
	house, ok := s.houses[houseName]
	if !ok {
		house = &lib.House{
			Name:    houseName,
			Members: make(map[net.Addr]*lib.Hamster),
		}
		s.houses[houseName] = house
	}
	house.Members[h.Conn.RemoteAddr()] = h

	s.exitCurrentHouse(h)

	h.House = house

	house.Broadcast(h, fmt.Sprintf("%s has hopped in to the house", h.Nickname))
	h.Msg(fmt.Sprintf("Welcome to the house: %s", house.Name))
}

func (s *server) cmdList(h *lib.Hamster, args []string) {
	var list []string

	switch args[1] {
	case "tunnels":
		for id := range s.tunnels {
			list = append(list, fmt.Sprintf("tunnel#%d", id))
		}
	case "hamsters":
		if h.House != nil {
			for _, m := range h.House.Members {
				list = append(list, m.Nickname)
			}
		} else {
			h.Err(errNotInHouse)
			return
		}
	default:
		h.Err(errors.New("No such option"))
		return
	}

	h.Msg(fmt.Sprintf("List of available %s: %s", args[1], strings.Join(list, ", ")))
}

func (s *server) cmdSqueakto(h *lib.Hamster, args []string) {
	if h.House != nil {
		recvName := args[1]
		if recvName != "anon" {
			recv, err := h.House.Search(recvName)
			if err != nil {
				h.Err(err)
				return
			}
			recv.Msg("🐹" + h.Nickname + " (privately): " + strings.Join(args[2:], " "))
		} else {
			h.Err(errors.New("You cannot send private messages anonymous hamsters"))
		}
	} else {
		h.Err(errNotInHouse)
	}
}

func (s *server) cmdExit(h *lib.Hamster, args []string) {

	if len(args) <= 1 {
		s.exitCurrentHouse(h)
		return
	}

	where := args[1]
	if where != "" && where[:7] == "tunnel#" {
		s.exitCurrentTunnel(h)
		return
	}

	_, ok := s.houses[where]
	if !ok {
		h.Err(errors.New("No such house"))
	} else {
		s.exitCurrentHouse(h)
	}
}

func (s *server) cmdQuit(h *lib.Hamster) {
	log.Printf("Hamster with the address %s disconnected", h.Conn.RemoteAddr().String())
	if h.House != nil {
		h.House.Broadcast(h, fmt.Sprintf("%s falls asleep", h.Nickname))
	}
	h.Msg("Take care!")
	h.Conn.Close()
}

func (s *server) message(h *lib.Hamster, args []string) {
	if h.House != nil {
		h.House.Broadcast(h, "🐹 "+h.Nickname+": "+strings.Join(args[:], " "))
	} else {
		h.Err(errNotInHouse)
	}
}
