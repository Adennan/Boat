package cmd

import (
	"Boat/internal/word"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

const (
	Upper = iota + 1
	Lower
	UnderlineToUpperCamel
	UnderlineToLowerCamel
	CamelToUnderline
)

var longDesc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下：",
	"1：全部转大写",
	"2：全部转小写",
	"3：下划线转大写驼峰",
	"4：下划线转小写驼峰",
	"5：驼峰转下划线",
}, "\n")

var (
	str  string
	mode int
)

// wordCmd represents the word command
var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  longDesc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case Upper:
			content = word.ToUpper(str)
		case Lower:
			content = word.ToLower(str)
		case UnderlineToUpperCamel:
			content = word.UnderlineToUpperCamel(str)
		case UnderlineToLowerCamel:
			content = word.UnderlineToLowerCamel(str)
		case CamelToUnderline:
			content = word.CamelToUnderline(str)
		default:
			log.Fatalln("暂不支持该转换模式，请执行 help 帮助文档")
		}
		log.Printf("输出结果: %s \n", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入要转换的单词")
	wordCmd.Flags().IntVarP(&mode, "mode", "m", 0, "请输入单词转换的模式")
}
