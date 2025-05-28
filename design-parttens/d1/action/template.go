package action

import (
	"fmt"
	"time"
)

// 模版方法
type Opration func()

func TemplateWithLog(opName string, opt Opration) Opration {
	return func() {
		start := time.Now()
		fmt.Printf("开始 %s\n", opName)
		opt()
		fmt.Printf("完成 %s, 耗时: %v\n", opName, time.Since(start))
	}
}

func Template(setUp, process, cleanUp Opration) {
	fmt.Println("模版方法执行")
	setUp()
	process()
	cleanUp()
	fmt.Println("模版方法结束")
}
