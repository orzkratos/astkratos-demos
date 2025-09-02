package demo2kratos

import (
	"path/filepath"
	"testing"

	"github.com/yyle88/must"
	"github.com/yyle88/must/muststrings"
)

// TestSourceRoot 测试获取源代码根目录路径的功能
func TestSourceRoot(t *testing.T) {
	path := SourceRoot()

	must.Nice(path)
	muststrings.HasSuffix(path, "/demo2kratos")
	must.True(filepath.IsAbs(path)) // 路径是绝对路径
}
