package golagen

type Rule struct {
	Name        string    `mapstructure:"name"`
	Command     string    `mapstructure:"command"`
	Prerules    *[]string `mapstructure:"prerules,omitempty"`
	Environment *[]string `mapstructure:"environment,omitempty"`
}
