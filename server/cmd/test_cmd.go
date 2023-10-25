package cmd

import (
	"errors"

	"github.com/diy0663/gohub/pkg/console"
	"github.com/spf13/cobra"
)

// todo 最后要保存一次才能自动import ,import 的时候也检查一下是否包引入是正确的

// todo 这个生成的命令 ,还得记得挂到上层命令那里面去

var CmdTestCmd = &cobra.Command{
	Use:   "test_cmd",
	Short: "HERE PUTS THE COMMAND DESCRIPTION",
	Run:   runTestCmd,
	Args:  cobra.NoArgs, // 只允许且必须传 0 个参数
}

func runTestCmd(cmd *cobra.Command, args []string) {

	console.Success("这是一条成功的提示")
	console.Warning("这是一条提示")
	console.Error("这是一条错误信息")
	console.Warning("终端输出最好使用英文，这样兼容性会更好~")
	console.Exit("exit 方法可以用来打印消息并中断程序！")
	console.ExitIf(errors.New("在 err != nil 的时候打印并退出"))
}
