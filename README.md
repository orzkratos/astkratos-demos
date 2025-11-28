[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/orzkratos/demokratos/release.yml?branch=main&label=BUILD)](https://github.com/orzkratos/demokratos/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/orzkratos/demokratos)](https://pkg.go.dev/github.com/orzkratos/demokratos)
[![Coverage Status](https://img.shields.io/coveralls/github/orzkratos/demokratos/main.svg)](https://coveralls.io/github/orzkratos/demokratos?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25+-lightgrey.svg)](https://github.com/orzkratos/demokratos)
[![GitHub Release](https://img.shields.io/github/release/orzkratos/demokratos.svg)](https://github.com/orzkratos/demokratos/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/orzkratos/demokratos)](https://goreportcard.com/report/github.com/orzkratos/demokratos)

# demokratos

Demo projects built with the Go-Kratos framework, serving as the foundation of the orzkratos ecosystem.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->
## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Introduction

**demokratos** is a reference implementation demonstrating best practices when building microservices with the [Go-Kratos](https://go-kratos.dev) framework. It serves as:

- 🎯 **Foundation Project** - The base template of 15+ specialized demo projects in the orzkratos ecosystem
- 🛠️ **Toolchain Integration Example** - Showcasing orzkratos development tools in action
- 📚 **Learning Resource** - Complete microservice structure following Kratos conventions
- ⚡ **Fast Development** - Auto-sync proto and code through magic commands like make orz

## About Go-Kratos

[Go-Kratos](https://go-kratos.dev) is a concise and efficient microservice framework that provides:

- Clean architecture with distinct separation of concerns
- Comprehensive middleware and plugin ecosystem
- Built-in support for gRPC and HTTP protocols
- Excellent documentation and active ecosystem

**demokratos builds upon this solid foundation**, adding enhanced tooling and automation to streamline the development workflow.

## Core Features

### 🚀 orzkratos Toolchain Integration

Provides robust orzkratos toolchain:

- **orzkratos-add-proto** - Simplifies adding proto files to the Kratos project
- **orzkratos-srv-proto** - Auto syncs service implementations with proto definitions

Setup tools:
```bash
make init
```

### ⚡ Magic Command: `make orz`

The core feature - auto synchronization between proto files and service code:

```bash
make orz
```

**What it does:**
- ✅ New methods in proto → Auto generates service method stubs
- ✅ Deleted methods → Converts to unexported functions (preserves logic)
- ✅ Reordered methods → Auto rearranges service code to match proto sequence

**Workflow example:**
1. Add `CreateUser` method to the proto file
2. Run `make orz`
3. Service generates `CreateUser` method stub
4. Implement the business logic

### 🔀 Fork Project Synchronization

Provides complete automated workflow to sync fork projects with upstream changes.

Through `make merge-stepN` series commands, auto handles upstream code merging, conflict resolution, dependencies upgrades, test validation, and more.

See [Makefile](./Makefile) with detailed workflow and usage instructions.

## Project Structure

### Demo Projects

Provides two demos to showcase the usage of various features:

- [demo1kratos](./demo1kratos) - Basic Kratos microservice example
- [demo2kratos](./demo2kratos) - Advanced features and integrations

Both demos follow standard Kratos project structure.

We provide a code comparison between Demo1 and Demo2, highlighting the changed code blocks.

When this project is forked, you can also compare it with the source to see the differences.

### Changes DIR

The [changes/](./changes) DIR contains markdown files documenting code differences:

- [changes/demo1.md](./changes/demo1.md) - Demo1 changes compared to source
- [changes/demo2.md](./changes/demo2.md) - Demo2 changes compared to source
- [changes/demos-toolchain-trees.md](./changes/demos-toolchain-trees.md) - Sibling projects and toolchain structures

Tests auto-generate these files:

```bash
go test -v -run TestGenerate1Changes # Generate demo1.md
go test -v -run TestGenerate2Changes # Generate demo2.md
go test -v -run TestGenerateXChanges # Generate demos-toolchain-trees.md
```

**In source project:** Files show `✅ NO CHANGES`

**In fork projects:** Files show code differences with syntax highlighting, making it simple to track customizations on GitHub.

## Forks

|    demo     |                        repo                         |
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

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-25 03:52:28.131064 +0000 UTC -->

## 📄 License

MIT License - see [LICENSE](LICENSE).

---

## 💬 Contact & Feedback

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Mistake reports?** Open an issue on GitHub with reproduction steps
- 💡 **Fresh ideas?** Create an issue to discuss
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/orzkratos/demokratos.svg?variant=adaptive)](https://starchart.cc/orzkratos/demokratos)
