package main

type Task struct {
	Name     string   `json:"name"`
	Commands []string `json:"commands"`
	Params   []Param  `json:"params,omitempty"`
	Secret   string   `json:"secret,omitempty"`
}

type Param struct {
	TypeName     string `json:"type"`
	DefaultValue string `json:"default"`
}

type ParamType string

const (
	ParamString ParamType = "string"
	ParamInt              = "int"
	ParamFloat            = "float"
)

func (t *Task) render(paramValues map[string]string) []string {
	return t.Commands
}
