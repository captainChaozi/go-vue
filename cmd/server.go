package main

import (
	"github.com/urfave/cli" // 一个非常好用的命令行工具
	"go-vue/utils/logger"
	"gopkg.in/macaron.v1"
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
		logger.Fatal(err)
	}

}

func getCommands() []cli.Command {
	command := cli.Command{
		Name: "run",
		Usage: "run server",
		Action: runServer,

	}
	return []cli.Command{command}
}

func  runServer(ctx *cli.Context){
	m := macaron.Classic()
	m.Get("/", func() string {
		return "Hello,World"
	})
	m.Run("0.0.0.0", 8000)


}