# ========================================
# 开始你的 Kratos 开发之旅吧
# ========================================

# 这些是我专为 Kratos 项目打造的效率工具，让开发变得更轻松愉快！请运行 make init 安装
init:
	@echo "正在安装 Kratos 相关的开发工具链..."
	# orzkratos-add-proto: 告别繁琐的 proto 文件手动创建！
	# 使用方法: 在 api/helloworld/ 目录下运行 orzkratos-add-proto demo.proto
	go install github.com/orzkratos/orzkratos/cmd/orzkratos-add-proto@latest
	# orzkratos-srv-proto: 自动同步神器，让你的服务实现与 proto 接口始终保持一致
	# 重要提醒: 使用前请务必备份代码或提交到 git，因为会直接修改源文件
	go install github.com/orzkratos/orzkratos/cmd/orzkratos-srv-proto@latest
	# depbump: 一键升级所有依赖包，告别版本管理烦恼
	go install github.com/go-mate/depbump/cmd/depbump@latest
	# go-lint: 代码质量守护者，自动格式化 + 静态检查
	go install github.com/go-mate/go-lint/cmd/go-lint@latest
	@echo "✅ 工具安装完成！现在可以开始愉快地开发啦"

# 构建所有演示项目，包括 proto 生成、配置文件处理、代码生成等
all:
	@echo "开始构建所有演示项目..."
	cd demo1kratos && make all
	cd demo2kratos && make all
	@echo "✅ 所有项目构建完成！"

# ========================================
# 魔法命令 make orz - 自动同步 Proto 代码与服务
# ========================================
# 这是最强大的功能！当你修改 proto 文件后，运行这个命令：
# 新增接口 → 自动在服务层添加对应的函数实现（新增个空函数，需要您实现函数内部逻辑）
# 删除接口 → 自动将服务代码中对应的方法改为非导出的（避免编译错误）
# 函数排序 → 根据你 proto 里定义的函数顺序重新排列服务里的实现代码
# 使用场景举例:
#   1. 在 proto 文件中新增了 CreateUser 接口
#   2. 运行 make orz
#   3. 服务层自动生成 CreateUser 方法框架，你只需要填充业务逻辑！
orz:
	@echo "开始执行 Proto-Service 自动同步函数..."
	cd demo1kratos && make all && orzkratos-srv-proto -auto
	cd demo2kratos && make all && orzkratos-srv-proto -auto
	@echo "✅ 同步完成！请检查生成的代码并完善业务逻辑"

# ========================================
# 同步上游仓库最新修改到 fork 项目的完整流程
# ========================================
# 背景说明：
# 1. 这些项目都是由 demokratos fork 来的，而每个fork都会展示一种特殊的使用技巧，这样方便新手查看具体如何使用。
# 2. 在新 fork 项目里，还贴心的提供了和源项目 demokratos 代码的比较的测试函数
# 3. 当源项目 demokratos 修改了东西，或者使用了更新的 kratos 框架版本时，还可以在 fork 项目里同步修改。
# 4. 因此，这些 fork 的项目，基本都不会再合并到 demokratos 里，而是作为单独的样例长期存在。
# 5. 项目提供了同步源代码修改的逻辑，这个逻辑只能在 fork 项目里执行。
# 使用说明：
# 1. 首先检查工作区状态 git status，如有未提交的修改需要处理：
#    - 仅包含 go.mod/go.sum 的变化：git stash (依赖升级可稍后再合)
#    - 有业务代码变化：需要先提交代码，避免混合提交历史
# 2. 按顺序执行 merge-step1 到 merge-step12 完成同步的操作
# 3. 如果有冲突，自行解决 (常见于 go.mod/go.sum 文件)
# 4. 若任何步骤出现错误需要再次修改代码/依赖时，改完都要再次运行测试和代码静态检查，避免引入新问题

merge-step1:
	# 添加上游仓库为远程源
	# 注意: 如果 upstream 远程源已存在，而且是同名仓库，就忽略重复的错误，因为这不是问题，但是假如指向其他仓库，就报错，而且不往下执行
	git remote add upstream git@github.com:orzkratos/demokratos.git
	@echo "✅ 已添加上游仓库远程源"

merge-step2:
	# 获取上游仓库的最新代码，不获取标签以避免冲突
	git fetch --no-tags upstream main
	@echo "✅ 已获取上游仓库最新代码"

merge-step3:
	# 确保当前在 main 分支里
	git checkout main
	@echo "✅ 已切换到 main 分支里"

merge-step4:
	# 检查当前是否有未提交的代码，如果有变动则暂存起来
	git status
	# 如果有未提交的变动，暂存到 stash（会自动检查是否有变动）
	git diff --quiet || git stash push -m "临时保存：merge 前的未提交变动"
	git status
	@echo "✅ 已检查并暂存未提交的代码（如果有的话）"

merge-step5:
	# 合并上游仓库的 main 分支，使用 --no-edit 避免弹出编辑器，这样适合在脚本里执行
	git merge upstream/main --no-edit
	git status
	# 【重要提醒】假如出现冲突，请严格按照以下步骤操作：
	# 1. 编辑冲突文件，逐个解决所有 <<<<<<< ======= >>>>>>> 标记的冲突
	# 【技巧策略】假如是go.sum有冲突，也可以不手动改，而是在解决完 go.mod 的冲突后执行 go mod tidy 即可解决
	# 2. 使用 git add <文件名> 将解决后的文件标记为已解决
	# 3. 继续合并流程：git merge --continue（绝对不要使用 git commit）
	# 【助手注意】在 merge 冲突场景下必须使用 git merge --continue 而非 git commit
	# 【再次强调】解决冲突后不要直接使用 git commit，这会破坏 merge 流程的状态管理
	@echo "✅ 已合并上游代码-请检查是否有冲突需要解决"

merge-step6:
	# 假如 merge 无冲突则跳过这步就行
	# 当解决完所有冲突时，需要执行 git merge --continue
	# 检查是否还在 merge 状态，并判断冲突是否已全部解决
	@if [ -f .git/MERGE_HEAD ]; then \
		echo "检测到 merge 状态，检查冲突解决情况"; \
		if git diff --name-only --diff-filter=U | grep -q .; then \
			echo "⚠️  发现未解决的冲突文件："; \
			git diff --name-only --diff-filter=U; \
			echo "请手动解决所有冲突后再执行此步骤！"; \
			echo "解决完冲突以后请手动添加这些文件："; \
			for file in $$(git diff --name-only --diff-filter=U); do \
				echo "  git add $$file"; \
			done; \
			echo "或者添加全部：git add -A"; \
			exit 1; \
		else \
			echo "已解决所有冲突，继续合并代码"; \
			git merge --continue; \
		fi; \
	else \
		echo "merge 已完成，无需继续"; \
	fi
	git status
	@echo "✅ 已完成 merge 流程"

merge-step7:
	# 升级所有项目的依赖包到最新版本
	# depbump: 完整升级根目录依赖
	# depbump directs: 依次升级子项目的直接依赖
	depbump
	# 在项目根目录里进第1个项目
	cd demo1kratos && depbump directs
	# 在项目根目录里进第2个项目
	cd demo2kratos && depbump directs
	@echo "✅ 已升级所有依赖包"

merge-step8:
	# 检查是否有依赖升级的变动，如果有则单独提交
	git diff --quiet || (git add -A && git commit -m "简单升级依赖包")
	git status
	@echo "✅ 已提交依赖升级变动（如果有的话）"

merge-step9:
	# 运行所有的测试确保代码正常工作
	# 清理测试缓存避免旧缓存影响结果
	go clean -testcache
	go test -v ./...
	# 在项目根目录里进第1个项目
	cd demo1kratos && go test -v ./...
	# 在项目根目录里进第2个项目
	cd demo2kratos && go test -v ./...
	@echo "✅ 已进行单元测试"

merge-step10:
	# 整理 go.mod 和 go.sum 文件
	# -e 参数允许在有错误时继续执行
	go mod tidy -e
	# 在项目根目录里进第1个项目
	cd demo1kratos && go mod tidy -e
	# 在项目根目录里进第2个项目
	cd demo2kratos && go mod tidy -e
	@echo "✅ 已整理所有依赖"

merge-step11:
	# 运行代码静态检查和格式化
	go-lint
	@echo "✅ 已进行代码检查"

merge-step12:
	# 恢复之前暂存的代码（如果有的话）
	# 检查是否有 stash 存在，如果有则恢复
	git stash list
	git stash list | grep -q "临时保存：merge 前的未提交变动" && git stash pop || echo "没有找到需要恢复的 stash"
	git status
	@echo "✅ 已恢复之前暂存的代码（如果有的话）"
