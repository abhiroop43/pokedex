package commands

import (
	"abhiroop43/pokedex/dto"
	"os"
)

func CommandExit(config *dto.Config) error {
	os.Exit(0)
	return nil
}
