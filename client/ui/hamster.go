package ui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

type TextArea struct {
	name string
	Body string
}

func NewTextArea(name, body string) *TextArea {
	return &TextArea{
		name: name,
		Body: body,
	}
}

func (ta *TextArea) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	midX := maxX / 2
	midY := maxY / 2
	x0, y0 := midX-(3*midX)/4, midY-(3*midY)/4
	x1, y1 := midX+(3*midX)/4, midY+(3*midY)/4
	if v, err := g.SetView("textarea", x0, y0, x1, y1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = true
		fmt.Fprintf(v, ta.Body)
	}
	return nil
}

type Prompt struct {
	name      string
	maxLength int
	Cmd       string
}

func NewPrompt(name string, maxLength int) *Prompt {
	return &Prompt{
		name:      name,
		maxLength: maxLength,
	}
}

func (p *Prompt) Edit(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	cx, _ := v.Cursor()
	ox, _ := v.Origin()
	limit := ox+cx+1 > p.maxLength
	switch {
	case ch != 0 && mod == 0 && !limit:
		v.EditWrite(ch)
	case key == gocui.KeySpace && !limit:
		v.EditWrite(' ')
	case key == gocui.KeyBackspace || key == gocui.KeyBackspace2:
		v.EditDelete(true)
	}
}

func (p *Prompt) Layout(g *gocui.Gui) error {
	x0, _, x1, y1, err := g.ViewPosition("textarea")
	if err != nil {
		return err
	}
	if v, err := g.SetView(p.name, x0, y1+1, x1, y1+3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editor = p
		v.Editable = true
	}
	return nil
}

func SetFocus(name string) func(g *gocui.Gui) error {
	return func(g *gocui.Gui) error {
		_, err := g.SetCurrentView(name)
		return err
	}
}

func Quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func (p *Prompt) UpdateTextAreaView(g *gocui.Gui, v *gocui.View) error {
	ox, oy := v.Origin()
	tv, err := g.View("textarea")
	if err != nil {
		return err
	}
	p.Cmd = v.Buffer()
	fmt.Fprintln(tv, p.Cmd)
	v.Clear()
	v.SetCursor(ox, oy)
	return nil
}

func (ta *TextArea) UpdateTextArea(g *gocui.Gui) error {
	tv, err := g.View("textarea")
	if err != nil {
		return err
	}
	fmt.Fprintln(tv, ta.Body)
	return nil
}
