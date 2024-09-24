package internal

type CmdFlags struct {
	Mode   string
	Config string
}

var (
	Flags = CmdFlags{}
)
