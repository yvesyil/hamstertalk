package main

import (
	"bufio"
	"fmt"
	"net"

	"github.com/HawkBrave/Hamstertalk/client/ui"
	"github.com/jroimartin/gocui"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		panic(err)
	}
	defer g.Close()

	g.Cursor = true

	textarea := ui.NewTextArea("textarea", "")
	prompt := ui.NewPrompt("prompt", 30)
	focus := gocui.ManagerFunc(ui.SetFocus("prompt"))
	g.SetManager(textarea, prompt, focus)

	go func() {
		conn, err := net.Dial("tcp", "localhost:8888")
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		p, _ := g.View("prompt")
		ta, _ := g.View("textarea")
		for {
			reader := bufio.NewReader(p)
			text, _ := reader.ReadString('\n')
			fmt.Fprintf(conn, text)
			message, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Fprintln(ta, message)
		}

	}()

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, ui.Quit); err != nil {
		panic(err)
	}

	if err := g.SetKeybinding("prompt", gocui.KeyEnter, gocui.ModNone, prompt.UpdateTextAreaView); err != nil {
		panic(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		panic(err)
	}

}
