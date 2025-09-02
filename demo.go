package demokratos

import (
	"github.com/orzkratos/demokratos/demo1kratos"
	"github.com/orzkratos/demokratos/demo2kratos"
)

// GetDemo1Path 获取 demo1kratos 项目的源代码根目录路径
func GetDemo1Path() string {
	return demo1kratos.SourceRoot()
}

// GetDemo2Path 获取 demo2kratos 项目的源代码根目录路径
func GetDemo2Path() string {
	return demo2kratos.SourceRoot()
}
