[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/orzkratos/demokratos/release.yml?branch=main&label=BUILD)](https://github.com/orzkratos/demokratos/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/orzkratos/demokratos)](https://pkg.go.dev/github.com/orzkratos/demokratos)
[![Coverage Status](https://img.shields.io/coveralls/github/orzkratos/demokratos/main.svg)](https://coveralls.io/github/orzkratos/demokratos?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25+-lightgrey.svg)](https://github.com/orzkratos/demokratos)
[![GitHub Release](https://img.shields.io/github/release/orzkratos/demokratos.svg)](https://github.com/orzkratos/demokratos/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/orzkratos/demokratos)](https://goreportcard.com/report/github.com/orzkratos/demokratos)

# demokratos

基于 Go-Kratos 框架构建的演示项目，orzkratos 生态系统的基石

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->
## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## 项目简介

**demokratos** 是使用 [Go-Kratos](https://go-kratos.dev) 框架构建微服务的最佳实践参考实现。它的作用是：

- 🎯 **基础项目** - orzkratos 生态系统中 15+ 专用演示项目的基础模板
- 🛠️ **工具链集成示例** - 展示 orzkratos 开发工具的实践应用
- 📚 **学习资源** - 完整的微服务架构，遵循 Kratos 规范
- ⚡ **快速开发** - 通过 make orz 等魔法命令实现 proto 和代码自动同步

## 关于 Go-Kratos

[Go-Kratos](https://go-kratos.dev) 是一个简洁高效的微服务框架，它提供：

- 清晰的架构和关注点分离
- 完善的中间件和插件生态
- 内置 gRPC 和 HTTP 协议支持
- 优秀的文档和活跃的生态

**demokratos 在这个坚实的基础上**，添加了增强的工具链和自动化能力，简化开发工作流程。

## 核心功能

### 🚀 orzkratos 工具链集成

提供成熟的 orzkratos 工具链：

- **orzkratos-add-proto** - 简化向 Kratos 项目添加 proto 文件的过程
- **orzkratos-srv-proto** - 自动同步服务实现与 proto 定义

安装工具：
```bash
make init
```

### ⚡ 魔法命令：`make orz`

最强大的功能 - proto 文件与服务代码之间的自动同步：

```bash
make orz
```

**它的作用：**
- ✅ proto 中新增方法 → 自动生成服务方法框架
- ✅ 删除的方法 → 转换成非导出函数（保留逻辑）
- ✅ 重排序方法 → 自动调整服务代码顺序以匹配 proto

**工作流示例：**
1. 在 proto 文件中添加 `CreateUser` 方法
2. 运行 `make orz`
3. 服务自动生成 `CreateUser` 方法框架
4. 实现业务逻辑

### 🔀 Fork 项目同步

提供完整的自动化工作流，用于同步 fork 项目与上游变更。

通过 `make merge-stepN` 系列命令，自动处理上游代码合并、冲突解决、依赖升级、测试验证等流程。

详细工作流程和使用说明请查看 [Makefile](./Makefile)。

## 项目结构

### 演示项目

提供两个演示项目，展示各种功能的使用：

- [demo1kratos](./demo1kratos) - Kratos 微服务基础示例
- [demo2kratos](./demo2kratos) - 高级功能和集成

两个演示都遵循标准的 Kratos 项目结构。

我们提供 Demo1 和 Demo2 的代码比较，突出显示改动的代码块。

当此项目被 fork 时，你也可以将其与源项目进行比较，查看差异。

### Changes DIR

[changes/](./changes) DIR 包含记录代码差异的 markdown 文件：

- [changes/demo1.md](./changes/demo1.md) - Demo1 相对源项目的变更
- [changes/demo2.md](./changes/demo2.md) - Demo2 相对源项目的变更
- [changes/demos-toolchain-trees.md](./changes/demos-toolchain-trees.md) - 兄弟项目和工具链结构

这些文件通过测试自动生成：

```bash
go test -v -run TestGenerate1Changes # 生成 demo1.md
go test -v -run TestGenerate2Changes # 生成 demo2.md
go test -v -run TestGenerateXChanges # 生成 demos-toolchain-trees.md
```

**在源项目中：** 文件显示 `✅ NO CHANGES`

**在 fork 项目中：** 文件显示代码差异并带语法高亮，简单直观地在 GitHub 上查看定制化的改动。

## Fork 项目列表

|     演示      |                         仓库                          |
|:-----------:|:---------------------------------------------------:|
|   ast-go    |    https://github.com/orzkratos/astkratos-demos     |
|    auth     |    https://github.com/orzkratos/authkratos-demos    |
|   config    |   https://github.com/orzkratos/configkratos-demos   |
|  ebz-must   |    https://github.com/orzkratos/ebzkratos-demos     |
| spf13/cobra |   https://github.com/orzkratos/cobrakratos-demos    |
|    gorm     |    https://github.com/orzkratos/gormkratos-demos    |
|  http-cors  |   https://github.com/orzkratos/cors-kratos-demos    |
|    i18n     |    https://github.com/orzkratos/i18nkratos-demos    |
|    nacos    |   https://github.com/orzkratos/nacos-kratos-demos   |
| swagger-doc |   https://github.com/orzkratos/swaggokratos-demos   |
|    trace    |   https://github.com/orzkratos/tracekratos-demos    |
|  unittest   |    https://github.com/orzkratos/testkratos-demos    |
| vue3-client |    https://github.com/orzkratos/vue3kratos-demos    |
|    wire     |   https://github.com/orzkratos/wirekratos-demos    |
|     zap     |    https://github.com/orzkratos/zapkratos-demos     |
| zap-zh-hans |   https://github.com/orzkratos/zapzhkratos-demos    |
|   migrate   |  https://github.com/orzkratos/migratekratos-demos   |
|    ping     |    https://github.com/orzkratos/pingkratos-demos    |
| supervisors | https://github.com/orzkratos/supervisorkratos-demos |

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-25 03:52:28.131064 +0000 UTC -->

## 📄 许可证类型

MIT 许可证 - 详见 [LICENSE](LICENSE)。

---

## 💬 联系与反馈

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **问题报告？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **新颖思路？** 创建 issue 讨论
- 📖 **文档疑惑？** 报告问题，帮助我们完善文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，协助解决性能问题
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：面向用户的更改需要更新文档
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来贡献此项目。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## GitHub 标星点赞

[![Stargazers](https://starchart.cc/orzkratos/demokratos.svg?variant=adaptive)](https://starchart.cc/orzkratos/demokratos)
