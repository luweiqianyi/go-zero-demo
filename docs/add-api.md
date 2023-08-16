# 新增一个api接口
步骤如下：
1. 修改`account.api`，新增请求定义、响应定义、`handler`定义
2. `cmd`进入`account`目录，执行命令`goctl api go -api user.api -dir . -style gozero`,该命令会自动生成对应go代码