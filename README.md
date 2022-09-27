## 快速命令行程序模版 fast-cli-template

### 依赖/特性 Dependencies/Feature
- Cobra: 命令行的解析等 - parse commandline 
- Viper: 配置文件的解析 - parse config files
- TableWriter: 渲染输出表格结构体等 - render tables structs
- ColorCobra: 彩色命令行输出 - color commandline output
- survey: 友好的命令行交互 - friendly commandline interaction
- logrus: 日志输出 - log output
- go doc: 支持 godoc - support godoc

### 项目结构 Structure

```
├── cmd                     // cmd 包 command package
│   ├── rootCmd.go          // 主命令 rootCmd 
│   └── version             // 子命令/模块 subcommand/submodules
│       └── version.go      // 子命令/模块实现 subcommand/submodules implementation
├── config                  // config 包 config package 
│   └── config.go           // 配置文件读取
├── main                    // main 包 创建应用二进制的地方 main package where building binary file
│   └── main.go             
└── utils                   // utils 包 工具函数 utils package tools functions
    ├── Ask
    │   └── confirm.go
    ├── Compare
    │   └── StringInStringArray.go
    ├── Error
    │   └── errorHandle.go
    ├── Print
    │   ├── StructPrinter.go
    │   └── TablePrinter.go
    └── log
      └── logger.go
```