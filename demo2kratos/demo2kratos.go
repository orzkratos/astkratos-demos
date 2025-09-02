package demo2kratos

import "github.com/yyle88/runpath"

// SourceRoot 返回 demo2kratos 项目的源代码根目录路径
func SourceRoot() string {
	return runpath.PARENT.Path()
}
