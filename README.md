# go-lua-syntax-checker
用go来实现的一个lua语法检查工具

开发工具 - goland

应用场景：

1，在自动构建项目的时候，对项目中的lua文件进行批量语法检查

2，在策划提交lua配置表的时候，在hook里对语法进行检查

3，定时对项目中的lua文件进行检查并进行报警

用法 ./lua-syntax-checker [path] [-R] [logfilePath]

1  ./lua-syntax-checker .  校验当前目录下的lua文件

2  ./lua-syntax-checker . -R 校验当前目录及其子目录下的lua文件

3  ./lua-syntax-checker abc.lua  校验当前目录下的abc.lua文件

4  ./lua-syntax-checker /Users/mac/abc.lua  校验/Users/mac/abc.lua文件

5  ./lua-syntax-checker . -R 校验当前目录及其子目录下的lua文件

6  ./lua-syntax-checker /Users/mac/ -R 校验/Users/mac/及其子目录下的lua文件

7  ./lua-syntax-checker . -R  /Users/mac/result.log 校验/Users/mac/及其子目录下的lua文件,并输出日志到/Users/mac/result.log文件
