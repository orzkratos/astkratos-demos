// Package generate_test provides comprehensive test cases demonstrating astkratos package capabilities
// This test suite covers gRPC component detection, AST-based struct extraction, Go module metadata parsing,
// and complete Kratos project structure scanning with debug mode enabled
//
// generate_test 包提供全面的测试用例，演示 astkratos 包的各项功能
// 测试套件涵盖 gRPC 组件检测、基于 AST 的结构体提取、Go 模块元数据解析，
// 以及完整的 Kratos 项目结构扫描，并启用调试模式
package generate_test

import (
	"path/filepath"
	"testing"

	"github.com/orzkratos/astkratos"
	"github.com/orzkratos/demokratos/demo2kratos"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/must"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/zaplog"
)

var projectPath string // Kratos project root DIR absolute path // Kratos 项目根目录绝对路径
var apiPath string     // API proto definitions DIR path // API proto 定义目录路径

// TestMain sets up the test environment, enables debug mode, and validates project paths
// It runs once at package test initialization and prepares shared path constants
//
// TestMain 设置测试环境，启用调试模式，并验证项目路径
// 在包测试初始化时运行一次，准备共享的路径常量
func TestMain(m *testing.M) {
	astkratos.SetDebugMode(true)

	projectPath = demo2kratos.SourceRoot()
	must.Nice(projectPath)
	osmustexist.ROOT(projectPath)
	zaplog.SUG.Debugln(projectPath)

	apiPath = filepath.Join(projectPath, "api")
	osmustexist.ROOT(apiPath)

	m.Run()
}

// TestListGrpcClients verifies gRPC client interface detection in _grpc.pb.go files
// Scans the api DIR and extracts types matching the "XxxClient interface" pattern
// Validates that GreeterClient is detected with the expected package name
//
// TestListGrpcClients 验证在 _grpc.pb.go 文件中检测 gRPC 客户端接口
// 扫描 api 目录并提取匹配 "XxxClient interface" 模式的类型
// 验证检测到 GreeterClient 且包名符合预期
func TestListGrpcClients(t *testing.T) {
	definitions := astkratos.ListGrpcClients(apiPath)
	t.Log(neatjsons.S(definitions))

	require.Len(t, definitions, 1)
	require.Equal(t, "GreeterClient", definitions[0].Name)
	require.Equal(t, "v1", definitions[0].Package)
}

// TestListGrpcServers verifies gRPC service interface detection in generated proto files
// Identifies types matching "XxxServer interface" pattern while excluding Unsafe prefixed ones
// Confirms GreeterServer detection with v1 package namespace
//
// TestListGrpcServers 验证在生成的 proto 文件中检测 gRPC 服务接口
// 识别匹配 "XxxServer interface" 模式的类型，同时排除 Unsafe 前缀的类型
// 确认检测到 GreeterServer 且命名空间是 v1
func TestListGrpcServers(t *testing.T) {
	definitions := astkratos.ListGrpcServers(apiPath)
	t.Log(neatjsons.S(definitions))

	require.Len(t, definitions, 1)
	require.Equal(t, "GreeterServer", definitions[0].Name)
	require.Equal(t, "v1", definitions[0].Package)
}

// TestListGrpcServices verifies gRPC service name detection without the Client/Server suffix
// Extracts base service names like "Greeter" from proto-generated functions
// Great when you need the core service name independent of the C/S mode
//
// TestListGrpcServices 验证不带 Client/Server 后缀的 gRPC 服务名称检测
// 从 proto 生成的函数中提取基础服务名称如 "Greeter"
// 当需要独立于客户端/服务端模式的核心服务名称时很实用
func TestListGrpcServices(t *testing.T) {
	definitions := astkratos.ListGrpcServices(apiPath)
	t.Log(neatjsons.S(definitions))

	require.Len(t, definitions, 1)
	require.Equal(t, "Greeter", definitions[0].Name)
	require.Equal(t, "v1", definitions[0].Package)
}

// TestListGrpcUnimplementedServers verifies detection of UnimplementedXxxServer stub types
// protoc-gen-go-grpc generates these stubs which must be embedded in service implementations
// Confirms UnimplementedGreeterServer is detected with expected source location
//
// TestListGrpcUnimplementedServers 验证检测 UnimplementedXxxServer 存根类型
// protoc-gen-go-grpc 生成这些存根，必须嵌入到服务实现中
// 确认检测到 UnimplementedGreeterServer 且源位置符合预期
func TestListGrpcUnimplementedServers(t *testing.T) {
	definitions := astkratos.ListGrpcUnimplementedServers(apiPath)
	t.Log(neatjsons.S(definitions))

	require.Len(t, definitions, 1)
	require.Equal(t, "UnimplementedGreeterServer", definitions[0].Name)
	require.Equal(t, "v1", definitions[0].Package)
}

// TestHasGrpcClients verifies the boolean check function that determines gRPC client presence
// Returns true when at least one gRPC client interface exists in the scanned path
// Quick validation without needing the complete definition list
//
// TestHasGrpcClients 验证判断 gRPC 客户端是否存在的布尔检查函数
// 当扫描路径中存在至少一个 gRPC 客户端接口时返回 true
// 快速验证，无需获取完整的定义列表
func TestHasGrpcClients(t *testing.T) {
	hasClients := astkratos.HasGrpcClients(apiPath)
	t.Log("HasGrpcClients:", hasClients)

	require.True(t, hasClients)
}

// TestHasGrpcServers verifies the boolean check function that determines gRPC service presence
// Returns true when at least one gRPC service interface exists in the scanned path
// Enables fast existence validation without loading complete definitions
//
// TestHasGrpcServers 验证判断 gRPC 服务接口是否存在的布尔检查函数
// 当扫描路径中存在至少一个 gRPC 服务接口时返回 true
// 支持快速存在性验证，无需加载完整定义
func TestHasGrpcServers(t *testing.T) {
	hasServers := astkratos.HasGrpcServers(apiPath)
	t.Log("HasGrpcServers:", hasServers)

	require.True(t, hasServers)
}

// TestCountGrpcServices verifies the counting function that sums up gRPC services
// Gives statistics about the gRPC service count in the project
// Confirms the demo project has just one Greeter service defined
//
// TestCountGrpcServices 验证统计 gRPC 服务的计数函数
// 提供项目中 gRPC 服务数量的统计信息
// 确认演示项目中定义了唯一一个 Greeter 服务
func TestCountGrpcServices(t *testing.T) {
	count := astkratos.CountGrpcServices(apiPath)
	t.Log("CountGrpcServices:", count)

	require.Equal(t, 1, count)
}

// TestGetStructsMap verifies AST-based struct extraction from protobuf generated .pb.go files
// Parses Go source code and builds a map of struct names to complete definitions
// Confirms HelloRequest and HelloReply message structs have complete source code extracted
//
// TestGetStructsMap 验证从 protobuf 生成的 .pb.go 文件中基于 AST 提取结构体
// 解析 Go 源代码并构建结构体名称到完整定义的映射
// 确认提取了 HelloRequest 和 HelloReply 消息结构体的完整源码
func TestGetStructsMap(t *testing.T) {
	structMap := astkratos.GetStructsMap(filepath.Join(apiPath, "helloworld/v1/greeter.pb.go"))
	t.Log("struct count:", len(structMap))

	for name, definition := range structMap {
		t.Log(name, "->", definition.Name)
	}

	require.Len(t, structMap, 2)
	require.Contains(t, structMap, "HelloRequest")
	require.Contains(t, structMap, "HelloReply")
}

// TestGetModuleInfo verifies Go module metadata extraction via go mod edit -json command
// Parses go.mod and returns module path, Go version, toolchain settings, and dependencies
// Confirms the demo project has expected module path and valid Go version
//
// TestGetModuleInfo 验证通过 go mod edit -json 命令提取 Go 模块元数据
// 解析 go.mod 并返回模块路径、Go 版本、工具链设置和依赖信息
// 确认演示项目的模块路径符合预期且 Go 版本有效
func TestGetModuleInfo(t *testing.T) {
	moduleInfo, err := astkratos.GetModuleInfo(projectPath)
	require.NoError(t, err)

	t.Log("Module Path:", moduleInfo.Module.Path)
	t.Log("Go Version:", moduleInfo.Go)
	t.Log("Toolchain:", moduleInfo.Toolchain)
	t.Log("Toolchain Version:", moduleInfo.GetToolchainVersion())
	t.Log("Dependencies count:", len(moduleInfo.Require))

	require.Equal(t, "github.com/orzkratos/demokratos/demo2kratos", moduleInfo.Module.Path)
	require.NotEmpty(t, moduleInfo.Go)
}

// TestAnalyzeProject verifies the comprehensive Kratos project scanning function
// Combines module info extraction with gRPC component detection in a single operation
// Returns ProjectReport containing ModuleInfo, Clients, Servers, and Services lists
//
// TestAnalyzeProject 验证全面的 Kratos 项目扫描函数
// 在单次操作中组合模块信息提取和 gRPC 组件检测
// 返回包含 ModuleInfo、Clients、Servers 和 Services 列表的 ProjectReport
func TestAnalyzeProject(t *testing.T) {
	report := astkratos.AnalyzeProject(projectPath)

	t.Log(neatjsons.S(report))

	require.NotNil(t, report.ModuleInfo)
	require.Equal(t, "github.com/orzkratos/demokratos/demo2kratos", report.ModuleInfo.Module.Path)

	require.Len(t, report.Clients, 1)
	require.Equal(t, "GreeterClient", report.Clients[0].Name)

	require.Len(t, report.Servers, 1)
	require.Equal(t, "GreeterServer", report.Servers[0].Name)

	require.Len(t, report.Services, 1)
	require.Equal(t, "Greeter", report.Services[0].Name)
}

// TestDebugMode verifies debug mode is enabled as expected from TestMain initialization
// The debug mode was set to true at the beginning, this test confirms the setting persists
//
// TestDebugMode 验证调试模式按预期从 TestMain 初始化中启用
// 调试模式在开始时设置为 true，此测试确认设置持续有效
func TestDebugMode(t *testing.T) {
	t.Log("Debug mode:", astkratos.IsDebugMode())

	require.True(t, astkratos.IsDebugMode())
}
