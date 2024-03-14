package golagen

type Config struct {
	Author      string            `mapstructure:"author"`
	Project     string            `mapstructure:"project"`
	Entries     []Entry           `mapstructure:"entries,omitempty"`
	Environment map[string]string `mapstructure:"environment,omitmepty"`
}
