# go-lua-syntax-checker
用go来实现的一个lua语法检查工具

引用 go-lua库  https://github.com/Azure/golua.git

开发工具 - intelliJ idea+go插件  GOPATH一定要设置对

目前我只编译的mac版本，需要windows版本的，请自行编译


应用场景：

1，在自动构建项目的时候，对项目中的lua文件进行批量语法检查

2，在策划提交lua配置表的时候，在hook里对语法进行检查

3，定时对项目中的lua文件进行检查并进行报警




用法 ./lua-syntax-checker [path] [logfilePath]（必须是完整路径）

1  ./lua-syntax-checker .  校验当前目录下的lua文件

2  ./lua-syntax-checker abc.lua  校验当前目录下的abc.lua文件

3  ./lua-syntax-checker /Users/mac/  校验/Users/mac/及其子目录下的lua文件

4  ./lua-syntax-checker .   /Users/mac/result.log 校验当前目录下的lua文件输出日志到/Users/mac/result.log
