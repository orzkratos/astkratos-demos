# 这些是该组织内的其它工具，也都是服务于kratos项目的，能够方便你的开发和使用，当然主要是我自己用
init:
	# 这个工具能够让你方便的添加proto文件
	go install github.com/orzkratos/orzkratos/cmd/orzkratos-add-proto@latest
	# 这个工具能够让你的service实现和proto接口保持同步修改，使用前请先做好代码备份或者提交代码
	go install github.com/orzkratos/orzkratos/cmd/orzkratos-srv-proto@latest

# 当你的kratos项目在proto里新增接口时，通过这个命令能够在对应的服务里也增加函数逻辑，在删除接口时也能把服务代码改为非导出的，以下是使用的样例
orz:
	cd demo1kratos && orzkratos-srv-proto -auto
	cd demo2kratos && orzkratos-srv-proto -auto
