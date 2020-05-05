package main

import (
	"fmt"
	"github.com/urfave/cli" // 一个非常好用的命令行工具
	"os"
)
func main() {
	cliApp := cli.NewApp()
	cliApp.Name = "go-vue"
	cliApp.Usage = "快速搭建go和Vue的生产环境"
	cliApp.Version = "V1.0"
	cliApp.Commands = getCommands()
	//cliApp.Flags =append()
	err := cliApp.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}

}

func getCommands() []cli.Command {
	return []cli.Command{}
}