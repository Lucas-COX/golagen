package internal

type Rule struct {
	Name        string    `mapstructure:"name"`
	Command     string    `mapstructure:"command"`
	Prerules    *[]string `mapstructure:"prerules,omitempty"`
	Environment *[]string `mapstructure:"environment,omitempty"`
}

type Entry struct {
	Name        string             `mapstructure:"name"`
	Methods     []string           `mapstructure:"methods"`
	Route       string             `mapstructure:"route"`
	Mods        *map[string]string `mapstructure:"mods,omitempty"`
	Environment *[]string          `mapstructure:"environment,omitempty"`
	Rules       *[]Rule            `mapstructure:"rules,omitempty"`
}

type Config struct {
	Project     string             `mapstructure:"project"`
	Author      string             `mapstructure:"author"`
	Entries     *[]Entry           `mapstructure:"entries,omitempty"`
	Environment *map[string]string `mapstructure:"environment,omitempty"`
	Rules       *map[string]string `mapstructure:"rules,omitempty"`
}
