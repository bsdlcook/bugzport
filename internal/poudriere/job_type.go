package poudriere

import (
	"gitlab.com/lcook/bugzport/internal/jail"
	"gitlab.com/lcook/bugzport/internal/port"
)

type Options struct {
	Output      bool
	Report      bool
	Interactive bool
	Config      bool
}

type Job struct {
	Jail    *jail.Jail
	Port    *port.Port
	Tree    string
	WorkDir string
	Options *Options
}
