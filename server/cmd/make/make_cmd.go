package make

import (
	"fmt"

	"github.com/diy0663/gohub/pkg/console"
	"github.com/spf13/cobra"
)

var CmdMakeCMD = &cobra.Command{
	Use:   "cmd",
	Short: "Create a command, like: make cmd XX_cmd ",
	Run:   runMakeCMD,
	// 只接收一个参数
	Args: cobra.ExactArgs(1),
}

func runMakeCMD(cmd *cobra.Command, args []string) {
	model := makeModelFromString(args[0])
	filePath := fmt.Sprintf("cmd/%s.go", model.PackageName)

	createFileFromStub(filePath, "cmd", model)
	// 友好提示
	console.Success("command name:" + model.PackageName)
	console.Success("command variable name: cmd.Cmd" + model.StructName)
	console.Warning("please edit main.go's app.Commands slice to register command")
}
