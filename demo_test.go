package demokratos_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/orzkratos/demokratos"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/eroticgo"
	"github.com/yyle88/must"
	"github.com/yyle88/osexec"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/printgo"
	"github.com/yyle88/rese"
	"github.com/yyle88/runpath"
)

// GetDemo1Path 获取本地测试环境中的 demo1kratos 目录路径。用于与 demokratos 包中的路径进行比较，验证 fork 项目与源项目的差异
func GetDemo1Path() string {
	return runpath.PARENT.Join("demo1kratos")
}

// GetDemo2Path 获取本地测试环境中的 demo2kratos 目录路径。用于与 demokratos 包中的路径进行比较，验证 fork 项目与源项目的差异
func GetDemo2Path() string {
	return runpath.PARENT.Join("demo2kratos")
}

// TestShow1Changes 比较 demo1kratos 项目在源项目和当前项目中的差异
// 使用场景：
// - 在 fork 项目中运行此测试，可以看到相对于源项目 demokratos 的所有代码变更
// - 运行命令：go test -v -run TestShow1Changes
// - 如果有差异会显示具体的文件和代码变更内容
func TestShow1Changes(t *testing.T) {
	path0 := osmustexist.ROOT(demokratos.GetDemo1Path())
	path1 := osmustexist.ROOT(GetDemo1Path())
	comparePath(t, path0, path1)
}

// TestShow2Changes 比较 demo2kratos 项目在源项目和当前项目中的差异
// 使用场景：
// - 在 fork 项目中运行此测试，可以看到相对于源项目 demokratos 的所有代码变更
// - 运行命令：go test -v -run TestShow2Changes
// - 如果有差异会显示具体的文件和代码变更内容
func TestShow2Changes(t *testing.T) {
	path0 := osmustexist.ROOT(demokratos.GetDemo2Path())
	path2 := osmustexist.ROOT(GetDemo2Path())
	comparePath(t, path0, path2)
}

// comparePath 使用 diff 命令比较两个路径的差异并输出结果
// 参数：
// - path0: 源路径（通常是 demokratos 包中的路径）
// - path1: 目标路径（通常是当前项目中的路径）
// 功能：
// - 忽略 go.mod 和 go.sum 文件的差异（因为依赖版本可能不同）
// - 如果没有差异，显示绿色的 "SAME" 消息
// - 如果有差异，显示黄色标记和具体的差异内容
func comparePath(t *testing.T, path0 string, path1 string) {
	t.Log("path0:", path0)
	t.Log("path1:", path1)
	output, err := osexec.NewOsCommand().WithDebugMode(osexec.SHOW_COMMAND).WithExpectExit(1, "DIFFERENCES FOUND").
		Exec(
			"diff",
			"-ruN",
			"--exclude=go.mod", // 忽略 go.mod 文件，避免依赖版本差异影响比较
			"--exclude=go.sum", // 忽略 go.sum 文件，避免依赖版本差异影响比较
			"--exclude=bin",    // 忽略 bin 目录，避免编译后的二进制文件差异影响比较
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

// TestCompare1Modules 专门比较 demo1kratos 项目的 go.mod 文件差异
// 使用场景：
// - 查看 demo1kratos 的依赖版本是否与源项目有差异
// - 运行命令：go test -v -run TestCompare1Modules
// - 常用于检查依赖升级或版本同步情况
func TestCompare1Modules(t *testing.T) {
	path0 := osmustexist.PATH(filepath.Join(demokratos.GetDemo1Path(), "go.mod"))
	path1 := osmustexist.PATH(filepath.Join(GetDemo1Path(), "go.mod"))
	comparePath(t, path0, path1)
}

// TestCompare2Modules 专门比较 demo2kratos 项目的 go.mod 文件差异
// 使用场景：
// - 查看 demo2kratos 的依赖版本是否与源项目有差异
// - 运行命令：go test -v -run TestCompare2Modules
// - 常用于检查依赖升级或版本同步情况
func TestCompare2Modules(t *testing.T) {
	path0 := osmustexist.PATH(filepath.Join(demokratos.GetDemo2Path(), "go.mod"))
	path2 := osmustexist.PATH(filepath.Join(GetDemo2Path(), "go.mod"))
	comparePath(t, path0, path2)
}

// TestShow1ReadableDiff 显示 Demo1 项目的易读代码差异
// 使用场景：
// - 在 fork 项目中运行此测试，可以看到相对于源项目 demokratos 的 Demo1 所有代码变更
// - 运行命令：go test -v -run TestShow1ReadableDiff
// - 如果有差异会显示彩色格式化的文件和代码变更内容，易于人眼识别
func TestShow1ReadableDiff(t *testing.T) {
	path0 := osmustexist.ROOT(demokratos.GetDemo1Path())
	path1 := osmustexist.ROOT(GetDemo1Path())

	showReadableDiff(t, path0, path1)
}

// TestShow2ReadableDiff 显示 Demo2 项目的易读代码差异
// 使用场景：
// - 在 fork 项目中运行此测试，可以看到相对于源项目 demokratos 的 Demo2 所有代码变更
// - 运行命令：go test -v -run TestShow2ReadableDiff
// - 如果有差异会显示彩色格式化的文件和代码变更内容，易于人眼识别
func TestShow2ReadableDiff(t *testing.T) {
	path0 := osmustexist.ROOT(demokratos.GetDemo2Path())
	path1 := osmustexist.ROOT(GetDemo2Path())

	showReadableDiff(t, path0, path1)
}

// showReadableDiff 显示格式化的易读 diff 结果
// 参数：
// - path0: 源路径（通常是 demokratos 包中的路径）
// - path1: 目标路径（通常是当前项目中的路径）
// 功能：
// - 忽略 go.mod 和 go.sum 文件的差异（因为依赖版本可能不同）
// - 如果没有差异，显示绿色的 "No changes" 消息
// - 如果有差异，显示彩色格式化的具体代码变更内容
// - 红色显示删除的代码行，绿色显示新增的代码行
// - 每个文件显示统计信息：文件名 (+新增行数 -删除行数)
func showReadableDiff(t *testing.T, path0, path1 string) {
	output, err := osexec.NewOsCommand().WithExpectExit(1, "DIFFERENCES FOUND").
		Exec(
			"diff",
			"-ruN",
			"--exclude=go.mod", // 忽略 go.mod 文件，避免依赖版本差异影响比较
			"--exclude=go.sum", // 忽略 go.sum 文件，避免依赖版本差异影响比较
			"--exclude=bin",    // 忽略 bin 目录，避免编译后的二进制文件差异影响比较
			path0,
			path1,
		)
	require.NoError(t, err)

	if len(output) == 0 {
		eroticgo.GREEN.ShowMessage("✅ NO CHANGES")
		return
	}
	eroticgo.AMBER.ShowMessage("📋 FOUND DIFFERENCES")

	var sourcePath string
	var adds, cuts []string

	printFile := func() {
		if sourcePath != "" && (len(adds) > 0 || len(cuts) > 0) {
			fmt.Printf("\n📄 %s (+%d -%d)\n", sourcePath, len(adds), len(cuts))
			for _, line := range cuts {
				fmt.Printf("  %s\n", eroticgo.RED.Sprint("- "+line))
			}
			for _, line := range adds {
				fmt.Printf("  %s\n", eroticgo.GREEN.Sprint("+ "+line))
			}
			fmt.Println()
		}
	}

	for _, line := range strings.Split(string(output), "\n") {
		switch {
		case strings.HasPrefix(line, "diff -ruN"):
			printFile() // 输出上一个文件
			sourcePath, adds, cuts = "", nil, nil

		case strings.HasPrefix(line, "---"):
			// 旧文件路径，跳过

		case strings.HasPrefix(line, "+++"):
			if parts := strings.Fields(line); len(parts) >= 2 {
				if strings.Contains(parts[1], path1+"/") {
					sourcePath = strings.TrimPrefix(parts[1], path1+"/")
				} else {
					sourcePath = filepath.Base(parts[1])
				}
			}

		case strings.HasPrefix(line, "+") && !strings.HasPrefix(line, "+++"):
			adds = append(adds, line[1:])

		case strings.HasPrefix(line, "-") && !strings.HasPrefix(line, "---"):
			cuts = append(cuts, line[1:])
		}
	}

	printFile() // 输出最后一个文件
}

// TestGenerate1Changes 生成 Demo1 项目的代码差异文件
// 使用场景：
// - 在 fork 项目中运行此测试，生成 changes/demo1.md 文件
// - 运行命令：go test -v -run TestGenerate1Changes
// - 文件内容使用 markdown 格式，GitHub 可直接查看
func TestGenerate1Changes(t *testing.T) {
	path0 := osmustexist.ROOT(demokratos.GetDemo1Path())
	path1 := osmustexist.ROOT(GetDemo1Path())
	t.Log(path1)
	// DIR must exist, please create if missing to avoid wrong location
	// 目录必须存在，如果缺失请创建，避免创建到错误位置
	outputRoot := osmustexist.ROOT(runpath.PARENT.Join("changes"))
	t.Log(outputRoot)
	outputPath := filepath.Join(outputRoot, "demo1.md")
	generateChangesFile(t, path0, path1, outputPath)
}

// TestGenerate2Changes 生成 Demo2 项目的代码差异文件
// 使用场景：
// - 在 fork 项目中运行此测试，生成 changes/demo2.md 文件
// - 运行命令：go test -v -run TestGenerate2Changes
// - 文件内容使用 markdown 格式，GitHub 可直接查看
func TestGenerate2Changes(t *testing.T) {
	path0 := osmustexist.ROOT(demokratos.GetDemo2Path())
	path1 := osmustexist.ROOT(GetDemo2Path())
	t.Log(path1)
	// DIR must exist, please create if missing to avoid wrong location
	// 目录必须存在，如果缺失请创建，避免创建到错误位置
	outputRoot := osmustexist.ROOT(runpath.PARENT.Join("changes"))
	t.Log(outputRoot)
	outputPath := filepath.Join(outputRoot, "demo2.md")
	generateChangesFile(t, path0, path1, outputPath)
}

// generateChangesFile 生成代码差异的 markdown 文件
// 参数：
// - path0: 源路径（通常是 demokratos 包中的路径）
// - path1: 目标路径（通常是当前项目中的路径）
// - outputFile: 输出文件路径（如 changes/demo1.md）
// 功能：
// - 忽略 go.mod 和 go.sum 文件的差异
// - 生成 markdown 格式的差异文件
// - 如果没有差异，生成包含 "No changes" 的文件
func generateChangesFile(t *testing.T, path0, path1, outputPath string) {
	output, err := osexec.NewOsCommand().WithExpectExit(1, "DIFFERENCES FOUND").
		Exec(
			"diff",
			"-ruN",
			"-U3",              // Show 3 lines of context around each change
			"--exclude=go.mod", // 忽略 go.mod 文件，避免依赖版本差异影响比较
			"--exclude=go.sum", // 忽略 go.sum 文件，避免依赖版本差异影响比较
			"--exclude=bin",    // 忽略 bin 目录，避免编译后的二进制文件差异影响比较
			path0,
			path1,
		)
	require.NoError(t, err)

	if len(output) == 0 {
		// Write empty result to file
		// 写入空结果到文件
		content := "# Changes\n\n✅ NO CHANGES\n"
		err := os.WriteFile(outputPath, []byte(content), 0644)
		require.NoError(t, err)
		t.Logf("Generated %s (no changes)", outputPath)
		return
	}

	var sourcePath string
	var diffLines []string
	var addCount, cutCount int

	ptx := printgo.NewPTX()
	ptx.Println("# Changes")
	ptx.Println()
	ptx.Println("Code differences compared to source project demokratos.")
	ptx.Println()

	processFile := func() {
		if sourcePath != "" && (addCount > 0 || cutCount > 0) {
			ptx.Printf("## %s (+%d -%d)\n\n", sourcePath, addCount, cutCount)
			ptx.Println("```diff")
			for _, line := range diffLines {
				ptx.Println(line)
			}
			ptx.Println("```")
			ptx.Println()
		}
	}

	for _, line := range strings.Split(string(output), "\n") {
		switch {
		case strings.HasPrefix(line, "diff -ruN"):
			processFile() // 处理上一个文件
			sourcePath, diffLines, addCount, cutCount = "", nil, 0, 0

		case strings.HasPrefix(line, "---"):
			// 旧文件路径，跳过

		case strings.HasPrefix(line, "+++"):
			if parts := strings.Fields(line); len(parts) >= 2 {
				if strings.Contains(parts[1], path1+"/") {
					sourcePath = strings.TrimPrefix(parts[1], path1+"/")
				} else {
					sourcePath = filepath.Base(parts[1])
				}
			}

		case strings.HasPrefix(line, "@@"):
			// Chunk heading, include it as context
			// Chunk 头部，包含作为上下文
			diffLines = append(diffLines, line)

		case strings.HasPrefix(line, "+") && !strings.HasPrefix(line, "+++"):
			diffLines = append(diffLines, line)
			addCount++

		case strings.HasPrefix(line, "-") && !strings.HasPrefix(line, "---"):
			diffLines = append(diffLines, line)
			cutCount++

		case strings.HasPrefix(line, " "):
			// Context line (space prefix indicates unchanged line)
			// 上下文行（空格前缀表示未改变的行）
			diffLines = append(diffLines, line)
		}
	}

	processFile() // 处理最后一个文件

	// Write to file
	// 写入文件
	require.NoError(t, os.WriteFile(outputPath, ptx.Bytes(), 0644))
	t.Logf("Generated %s with differences", outputPath)
}

func TestGenerateXChanges(t *testing.T) {
	{
		treePath, err := exec.LookPath("tree")
		if err != nil {
			t.Skip("tree is not available on this system, skipping test case")
		}
		t.Logf("Found tree at: %s", treePath)
	}

	root := runpath.PARENT.Path()
	t.Log(root)

	excludeSomeNames := []string{
		".git", // Can be omitted since hidden DIRs are skipped below // 这里可以省略，因为下面会跳过隐藏 DIR
		"changes",
		filepath.Base(GetDemo1Path()),
		filepath.Base(GetDemo2Path()),
	}
	t.Log("exclude:", excludeSomeNames)

	// List DIRs except known ones
	// 列举目录，排除已知的目录
	var matchNames []string
	for _, item := range rese.A1(os.ReadDir(root)) {
		if item.IsDir() {
			name := item.Name()
			t.Log("check:", name)
			if strings.HasPrefix(name, ".") {
				continue
			}
			if slices.Contains(excludeSomeNames, name) {
				continue
			}
			t.Log("match:", name)
			matchNames = append(matchNames, name)
		}
	}
	t.Log("match:", matchNames)

	//把其它目录的 tree 信息输出出来到文本里
	outputPath := osmustexist.FILE(filepath.Join(root, "changes", "demos-toolchain-trees.md"))
	t.Log(outputPath)

	if len(matchNames) == 0 {
		content := "# Changes\n\n✅ NO CHANGES\n"
		must.Done(os.WriteFile(outputPath, []byte(content), 0644))
		return
	}

	ptx := printgo.NewPTX()
	ptx.Println("# Changes")
	ptx.Println()

	// Overview section with sibling projects list
	// 概览章节，列出兄弟项目
	ptx.Println("## Overview")
	ptx.Println()
	ptx.Println("Sibling projects:")
	ptx.Println()
	for _, name := range matchNames {
		ptx.Fprintf("- [%s](#%s)", name, name)
		ptx.Println()
	}
	ptx.Println()

	// Detailed tree structure for each project
	// 每个项目的详细目录树结构
	ptx.Println("## Project Structures")
	ptx.Println()
	for idx, name := range matchNames {
		if idx > 0 {
			ptx.Println("---")
			ptx.Println()
		}

		ptx.Fprintf("### %s", name)
		ptx.Println()
		ptx.Println()
		ptx.Fprintf("**Location**: [%s](../%s)", name, name)
		ptx.Println()
		ptx.Println()
		ptx.Println("```bash")
		ptx.Fprintf("cd %s && tree --noreport", name)
		ptx.Println()
		ptx.Println("```")
		ptx.Println()

		subRoot := filepath.Join(root, name)
		t.Log(subRoot)
		treeOutput := rese.A1(osexec.ExecInPath(subRoot, "tree", "--noreport", "--charset=ascii", "--gitignore", "-I", "node_modules|.git|bin|.idea|.vscode"))
		t.Log(string(treeOutput))

		ptx.Println("```")
		ptx.Write(treeOutput)
		ptx.Println("```")
		ptx.Println()
	}
	must.Done(os.WriteFile(outputPath, ptx.Bytes(), 0644))
}
