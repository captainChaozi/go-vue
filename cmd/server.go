package main

import (
	"github.com/urfave/cli" // 一个非常好用的命令行工具
	"go-vue/api/middleware"
	"go-vue/utils/logger"
	"gopkg.in/macaron.v1"
	"os"
)

const (
	DefaultPort=8000
)

// 用cli的方式来启动程序,可以充分
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

// 这是进入程序的命令,可以注册多个命令给程序,每个命令中通过Action和相应的flags来触发相应的函数,
// 运行相应的代码块,然后可以通过命令行传入相应的参数
func getCommands() []cli.Command {
	command := cli.Command{
		Name: "run",
		Usage: "run server",
		Action: runServer,
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "port,p",
				Value: DefaultPort,
				Usage: "bind port",
			},
			cli.StringFlag{
				Name:  "env,e",
				Value: "prod",
				Usage: "runtime environment, dev|test|prod",
			},

		},

	}
	return []cli.Command{command}
}

// 解析启动端口
func parsePort (ctx *cli.Context) int{
	port := DefaultPort
	if ctx.IsSet("port") {
		port = ctx.Int("port")
	}
	if port <=0 || port >= 65535 {
		port = DefaultPort
	}
	return port
}

func setEnv(ctx *cli.Context) {
	env := "prod" // 同名变量的同类型的赋值覆盖
	if ctx.IsSet("env") {
		env  = ctx.String("env")
	}
	switch env {
	case "prod":
		macaron.Env = macaron.PROD
	case "dev":
		macaron.Env = macaron.DEV
	case "test":
		macaron.Env = macaron.TEST
	}
}

// 让web应用跑起来
func  runServer(ctx *cli.Context){
	setEnv(ctx) // 设置DEV PROD还是其他的
	m := macaron.Classic()
	// 设置路由

	// 设置中间件
	middleware.RegisterMiddleware(m)
	m.Get("/", func() string {
		return "Hello,World"
	})
	port := parsePort(ctx)
	m.Run("0.0.0.0",port )// 其中的环境变量

}