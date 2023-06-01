## 目录结构


```shell
|   .env.example  # env 示例文件
|   .gitignore
|   Dockerfile
|   go.mod        # 包管理
|   main.go       # 入口文件
|   README.md
|
+---api          # 服务包
|   \---v1         # v1 版本入口
|       \---example          # 示例
|           +---controller     # 控制器
|           +---dto            # 模型
|           +---repository     # 数据层
|           \---service        # 业务层
|       \---...              # 其他
|           +---controller     # 控制器
|           +---dto            # 模型
|           +---repository     # 数据层
|           \---service        # 业务层
|   \---v2        # v2 版本入口
+---initialize      # 初始化目录
+---common          # 通用模块
+---config          # 配置
+---router          # 路由
+---utils           # 工具
```


## 依赖服务包
+ gin web框架

+ gorm orm包

+ viper 配置服务

+ zip 日志
