package dto

type CliCommand struct {
	Name        string
	Description string
	Callback    func(config *Config, args ...string) error
}
