package demokratos_test

import (
	"path/filepath"
	"testing"

	"github.com/orzkratos/demokratos"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/eroticgo"
	"github.com/yyle88/osexec"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/runpath"
)

func GetDemo1Path() string {
	return runpath.PARENT.Join("demo1kratos")
}

func GetDemo2Path() string {
	return runpath.PARENT.Join("demo2kratos")
}

func TestShow1Changes(t *testing.T) {
	path0 := osmustexist.ROOT(demokratos.GetDemo1Path())
	path1 := osmustexist.ROOT(GetDemo1Path())
	comparePath(t, path0, path1)
}

func TestShow2Changes(t *testing.T) {
	path0 := osmustexist.ROOT(demokratos.GetDemo2Path())
	path2 := osmustexist.ROOT(GetDemo2Path())
	comparePath(t, path0, path2)
}

func comparePath(t *testing.T, path0 string, path1 string) {
	t.Log("path0:", path0)
	t.Log("path1:", path1)
	output, err := osexec.NewOsCommand().WithDebugMode(osexec.SHOW_COMMAND).WithExpectExit(1, "DIFFERENCES FOUND").
		Exec(
			"diff",
			"-ruN",
			"--exclude=go.mod", // Ignore go.mod file, effective only when comparing directories // 忽略 go.mod 文件，仅在比较目录时生效
			"--exclude=go.sum", // Ignore go.sum file, effective only when comparing directories // 忽略 go.sum 文件，仅在比较目录时生效
			path0,
			path1,
		)
	require.NoError(t, err)
	if len(output) == 0 {
		eroticgo.GREEN.ShowMessage("SAME")
	} else {
		eroticgo.AMBER.ShowMessage("⬇⬇⬇")
		t.Log(string(output))
		eroticgo.AMBER.ShowMessage("⬆⬆⬆")
	}
}

func TestCompare1Modules(t *testing.T) {
	path0 := osmustexist.PATH(filepath.Join(demokratos.GetDemo1Path(), "go.mod"))
	path1 := osmustexist.PATH(filepath.Join(GetDemo1Path(), "go.mod"))
	comparePath(t, path0, path1)
}

func TestCompare2Modules(t *testing.T) {
	path0 := osmustexist.PATH(filepath.Join(demokratos.GetDemo2Path(), "go.mod"))
	path2 := osmustexist.PATH(filepath.Join(GetDemo2Path(), "go.mod"))
	comparePath(t, path0, path2)
}
