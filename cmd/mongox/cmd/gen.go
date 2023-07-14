/*
Copyright © 2023 go-kenka <1107015496@qq.com>
*/
package cmd

import (
	"errors"
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/go-kenka/mongox/internal/code/gen"
	"github.com/go-kenka/mongox/internal/code/load"
	"github.com/go-kenka/mongox/utils"
	"github.com/spf13/cobra"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "通过DSL生成CRUD代码模型",
	Long:  `通过DSL生成CRUD代码模型`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("正在生成代码 %+v\n", args)

		// ast 路径
		schemaPath := args[0]
		if len(schemaPath) == 0 {
			schemaPath = "./schema"
		}

		// 代码生成路径
		targetPath, _ := cmd.Flags().GetString("target")

		// 读取schema定义
		cfg := &load.Config{Path: schemaPath}

		spec, err := cfg.Load()
		if err != nil {
			panic(err)
		}

		// 回退到上一目录
		fullPath := path.Join(spec.PkgPath, "../")

		relaPath := strings.Replace(fullPath, spec.Module.Path, ".", -1)

		// 自定义地址的场合
		if len(targetPath) > 0 {
			utils.Mkdir(targetPath)
			if filepath.IsAbs(targetPath) {
				panic(errors.New("target参数只支持相对路径"))
			}
			relaPath = targetPath
			fullPath = path.Join(spec.Module.Path, targetPath)
		}

		// 更新pkg路径为相对路径
		spec.PkgPath = relaPath

		// 开始生成代码
		err = gen.GenClient(fullPath, spec)
		if err != nil {
			panic(err)
		}

		fmt.Println("代码生成完成")
		fmt.Println("正在使用gofmt格式化代码")
		err = utils.GoFmt(relaPath)
		if err != nil {
			fmt.Println("格式化出错了，请安装gofmt,并将bin目录设置到环境变量中", err)
		}
		fmt.Println("格式化完成")
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	genCmd.Flags().StringP("target", "t", ".", "生成代码文件路径")
}
