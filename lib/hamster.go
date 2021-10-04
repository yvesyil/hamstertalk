package lib

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// Hamster model
type Hamster struct {
	Conn     net.Conn
	Nickname string
	House    *House
	Tunnel   *Tunnel
	Commands chan<- Command
}

// ReadInput reads the input of the hamster
func (h *Hamster) ReadInput() error {
	msg, err := bufio.NewReader(h.Conn).ReadString('\n')

	if err != nil {
		return err
	}

	msg = strings.Trim(msg, "\r\n")

	args := strings.Split(msg, " ")
	inp := strings.TrimSpace(args[0])

	if len(inp) == 0 {
		return nil
	}

	if inp[0] != '!' {
		h.Commands <- Command{
			ID:      Message,
			Hamster: h,
			Args:    args,
		}
		return nil
	}

	switch inp {
	case "!set":
		h.Commands <- Command{
			ID:      CmdSet,
			Hamster: h,
			Args:    args,
		}
	case "!use":
		h.Commands <- Command{
			ID:      CmdUse,
			Hamster: h,
			Args:    args,
		}
	case "!stepto":
		h.Commands <- Command{
			ID:      CmdStepto,
			Hamster: h,
			Args:    args,
		}
	case "!hopto":
		h.Commands <- Command{
			ID:      CmdHopto,
			Hamster: h,
			Args:    args,
		}
	case "!list":
		h.Commands <- Command{
			ID:      CmdList,
			Hamster: h,
			Args:    args,
		}
	case "!squeakto":
		h.Commands <- Command{
			ID:      CmdSqueakto,
			Hamster: h,
			Args:    args,
		}
	case "!exit":
		h.Commands <- Command{
			ID:      CmdExit,
			Hamster: h,
			Args:    args,
		}
	case "!quit":
		h.Commands <- Command{
			ID:      CmdQuit,
			Hamster: h,
			Args:    args,
		}
	default:
		h.Err(fmt.Errorf("No such action as %s", inp))
	}
	return nil
}

// Err sends an error message to the hamster
func (h *Hamster) Err(err error) {
	h.Conn.Write([]byte("\r⚠️ Error️: " + err.Error() + "\n>> "))
}

// Msg sends a message to the hamster
func (h *Hamster) Msg(msg string) {
	h.Conn.Write([]byte("\r " + msg + "\n>> "))
}
