package config

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"os"
)

var (
	Conf = New()
)

func New() *toml.Tree {
	var url string
	if len(os.Args) > 1 && os.Args[1] == "production" {
		fmt.Println(os.Args)
		url = "/home/www/ccm/config/config.toml"
	} else {
		url = "./config/config.toml"
	}
	fmt.Println(url)
	config, err := toml.LoadFile(url)
	if err != nil {
		fmt.Println("Toml Loading Error", err.Error())
	}
	return config
}
