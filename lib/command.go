package lib

type commandID int

// Command ids
const (
	CmdSet commandID = iota
	CmdUse
	CmdStepto
	CmdHopto
	CmdList
	CmdSqueakto
	CmdExit
	CmdQuit
	Message
)

// Command model
type Command struct {
	ID      commandID
	Hamster *Hamster
	Args    []string
}
