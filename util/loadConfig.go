package util

import (
	"os"

	"github.com/spf13/viper"
)

func LoadConfig() {

	file, err := os.Open("conf.json")
	if err != nil {
		return
	}
	viper.ReadConfig(file)
}
