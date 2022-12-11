package main

import (
	"fmt"

	"sample/app/cmd"

	_ "sample/config"
	"sample/pkg/console"
)

func main() {
	if err := cmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("启动失败：%s", err.Error()))
	}
}
