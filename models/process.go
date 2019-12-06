package models

// Process struct
type Process struct {
	CPU     string `json:"cpu"`
	Pid     string `json:"pid"`
	User    string `json:"user"`
	Command string `json:"command"`
}

// Init Process
func (p *Process) Init(cp, pi, u, c string) {
	p.CPU = cp
	p.Pid = pi
	p.User = u
	p.Command = c
}
