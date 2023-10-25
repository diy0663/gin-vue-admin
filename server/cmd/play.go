package cmd

import (
	"context"
	"fmt"

	"github.com/diy0663/gohub/pkg/console"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/spf13/cobra"
)

// todo 最后要保存一次才能自动import ,import 的时候也检查一下是否包引入是正确的

// todo 这个生成的命令 ,还得记得挂到上层命令那里面去

var CmdPlay = &cobra.Command{
	Use:   "play",
	Short: "HERE PUTS THE COMMAND DESCRIPTION",
	Run:   runPlay,
	Args:  cobra.NoArgs, // 只允许且必须传 1 个参数
}

func runPlay(cmd *cobra.Command, args []string) {

	console.Warning("在runPlay 这里里面写你要测试验证的代码 ")
	// todo 在这里写代码
	pingResult, err := global.GVA_REDIS.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(pingResult)
	}

}
