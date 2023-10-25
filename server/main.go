package main

import (
	"fmt"
	"os"

	"go.uber.org/zap"

	"github.com/diy0663/gohub/pkg/console"
	"github.com/flipped-aurora/gin-vue-admin/server/cmd"
	"github.com/flipped-aurora/gin-vue-admin/server/cmd/make"
	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"

	"github.com/spf13/cobra"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title                       Swagger Example API
// @version                     0.0.1
// @description                 This is a sample Server pets
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func init() {
	global.GVA_VP = core.Viper() // 初始化Viper
	initialize.OtherInit()
	global.GVA_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	initialize.DBList()
}
func main() {

	var rootCmd = &cobra.Command{
		Use:     "gva",
		Short:   "top cmd : gva ",
		Long:    `Default will run "gva" command, you can use "-h" flag to see all subcommands`,
		Example: "gva serve",
		// rootCmd 的所有子命令都会执行以下代码
		PersistentPreRun: func(command *cobra.Command, args []string) {
			// 其实, root 以及其子命令也依赖main , main 执行前 也有init方法要执行
			// 所以  某种层面这个 PersistentPreRun 达到的效果也跟 init 一样
		},
	}
	rootCmd.AddCommand(
		cmd.CmdServe,
		make.CmdMake,
		cmd.CmdTestCmd,
	)
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)

	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}

}
