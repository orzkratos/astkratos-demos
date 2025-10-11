<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT BADGES -->
[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/orzkratos/demokratos/release.yml?branch=main&label=BUILD)](https://github.com/orzkratos/demokratos/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/orzkratos/demokratos)](https://pkg.go.dev/github.com/orzkratos/demokratos)
[![Coverage Status](https://img.shields.io/coveralls/github/orzkratos/demokratos/main.svg)](https://coveralls.io/github/orzkratos/demokratos?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25+-lightgrey.svg)](https://github.com/orzkratos/demokratos)
[![GitHub Release](https://img.shields.io/github/release/orzkratos/demokratos.svg)](https://github.com/orzkratos/demokratos/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/orzkratos/demokratos)](https://goreportcard.com/report/github.com/orzkratos/demokratos)
<!-- TEMPLATE (ZH) END: STANDARD PROJECT BADGES -->

# demokratos

使用 Go-Kratos 框架的一些演示项目

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->
## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## Go-Kratos

[Go-Kratos](https://go-kratos.dev) 是一个简洁的微服务框架

## 演示项目

本项目包含两个演示项目，展示各种功能的使用：

- [Demo1](./demo1kratos)
- [Demo2](./demo2kratos)

我们提供 Demo1 和 Demo2 的代码比较，突出显示改动的代码块。

当此项目被 fork 时，你也可以将其与源项目进行比较，查看差异。

### Changes DIR

[changes/](./changes) DIR 包含记录代码差异的 markdown 文件：

- [changes/demo1.md](./changes/demo1.md) - Demo1 相对源项目的变更
- [changes/demo2.md](./changes/demo2.md) - Demo2 相对源项目的变更

这些文件通过测试自动生成：

```bash
go test -v -run TestGenerate1Changes  # 生成 demo1.md
go test -v -run TestGenerate2Changes  # 生成 demo2.md
```

**在源项目中：** 文件显示 `✅ NO CHANGES`

**在 fork 项目中：** 文件显示代码差异并带语法高亮，简单直观地在 GitHub 上查看定制化的改动。

## Fork 项目列表

|    演示     |                       仓库                       |
|:-----------:|:-----------------------------------------------:|
|   ast-go    |  https://github.com/orzkratos/astkratos-demos   |
| spf13/cobra | https://github.com/orzkratos/cobrakratos-demos  |
|  http-cors  | https://github.com/orzkratos/cors-kratos-demos  |
|    i18n     |  https://github.com/orzkratos/i18nkratos-demos  |
|    nacos    | https://github.com/orzkratos/nacos-kratos-demos |
| swagger-doc | https://github.com/orzkratos/swaggokratos-demos |
|    trace    | https://github.com/orzkratos/tracekratos-demos  |
|  unittest   |  https://github.com/orzkratos/testkratos-demos  |
| vue3-client |  https://github.com/orzkratos/vue3kratos-demos  |
|    wire     | https://github.com/orzkratos/wire2kratos-demos  |
|     zap     |  https://github.com/orzkratos/zapkratos-demos   |
| zap-zh-hans | https://github.com/orzkratos/zapzhkratos-demos  |

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-26 07:39:27.188023 +0000 UTC -->

## 📄 许可证类型

MIT 许可证。详见 [LICENSE](LICENSE)。

---

## 🤝 项目贡献

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **发现问题？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **功能建议？** 创建 issue 讨论您的想法
- 📖 **文档疑惑？** 报告问题，帮助我们改进文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，帮助我们优化性能
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
7. **文档**：为面向用户的更改更新文档，并使用有意义的提交消息
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来为此项目做出贡献。

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
