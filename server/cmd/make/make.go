package make

import (
	"embed"
	"fmt"
	"strings"

	"github.com/diy0663/go_project_packages/file"
	"github.com/diy0663/go_project_packages/str"
	"github.com/diy0663/gohub/pkg/console"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
)

type Model struct {
	//  数据表名  驼峰_复数
	TableName string
	// 结构名 大写驼峰单数
	StructName string
	// 大写驼峰复数
	StructNamePlural string
	// 变量名 驼峰单数
	VariableName string
	//
	VariableNamePlural string
	// 包名 下划线小写单数
	PackageName string
}

// 打包  embed.FS —— 以文件系统方式加载目录和子目录文件树
// 参考 https://learnku.com/courses/go-basic/1.19/packaging-go-programs-with-embedded/13445
// 在本项目中, 最终运行 只需要打包后的应用+ .env 即可
// embed 读取相对路径是相对于书写 //go:embed 指令的 .go 文件的。 注意只能在同级目录下进行

//go:embed stubs
var stubsFS embed.FS

var CmdMake = &cobra.Command{
	Use:   "make",
	Short: "Generate file and code",
	// 因为 make 其实要搭配后面的子命令, 例如 cmd, model,  api 之类的, 所以不需要Run

}

// 挂载 make 后面跟着的 一堆子命令
func init() {
	CmdMake.AddCommand(
		CmdMakeCMD,
	)

}

// 根据 子命令后面跟着的输入结果, 先整理一个处理好驼峰,大小写,下划线单复数等规范的struct
func makeModelFromString(name string) Model {
	model := Model{}
	model.StructName = str.Singular(strcase.ToCamel(name))
	model.StructNamePlural = str.Plural(model.StructName)
	model.TableName = str.Snake(model.StructNamePlural)
	model.VariableName = str.LowerCamel(model.StructName)
	model.PackageName = str.Snake(model.StructName)
	model.VariableNamePlural = str.LowerCamel(model.StructNamePlural)
	return model
}

// 根据某个指定模板去生成对应的代码文件到某个目录, 在生成的过程中可能需要涉及替换
// 最后一个参数是可选的 其实内部是当做一个slice , 里面存放 的是 map[string]string
func createFileFromStub(filePath string, stubName string, model Model, variables ...interface{}) {
	replaces := make(map[string]string)
	if len(variables) > 0 {
		// 有额外需要替换的变量,就得在最开始设置了, 这种只替换一次??
		replaces = variables[0].(map[string]string)
	}
	// 目标文件已存在,就报错,避免被覆盖
	if file.Exists(filePath) {
		console.Exit(filePath + " already exists")
	}
	// 读取模板文件 ,todo 这里面读取文件要  stubsFS , 是因为这些模板文件最后会被 stubsFS 打包,所以才得这样子读取
	stubFileData, err := stubsFS.ReadFile("stubs/" + stubName + ".stub")
	if err != nil {
		console.Exit(err.Error())
	}

	// 转为字符串
	stubFileString := string(stubFileData)

	// {{}} 这类变量都是在 stub 文件里面用来被替换的
	replaces["{{VariableName}}"] = model.VariableName
	replaces["{{VariableNamePlural}}"] = model.VariableNamePlural
	replaces["{{StructName}}"] = model.StructName
	replaces["{{StructNamePlural}}"] = model.StructNamePlural
	replaces["{{PackageName}}"] = model.PackageName
	replaces["{{TableName}}"] = model.TableName

	// 对模板文件做变量的替换
	for search, replace := range replaces {
		stubFileString = strings.ReplaceAll(stubFileString, search, replace)
	}

	if err := file.Put([]byte(stubFileString), filePath); err != nil {
		console.Exit("写入文件出错:" + err.Error())
	}
	console.Success(fmt.Sprintf("[%v] created.", filePath))

}
