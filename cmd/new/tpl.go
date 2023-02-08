package new

// TplProjectStructure 项目结构
const TplProjectStructure = `
├── LICENSE
├── Makefile
├── README.md
├── cmd
│   └── omgind
│       └── main.go # 入口文件
├── configs # 配置文件
│   ├── config.toml
│   ├── menu.yaml
│   └── model.conf
├── docs
│   └── data_model.md
├── go.mod
├── go.sum
├── internal
│   ├── api
│	│	└──	v2
│	│		└── mock
│	└──	app
│       ├── service
│       ├── config
│       ├── contextx
│       ├── ginx
│       ├── middleware
│       ├── module
│       │   └── adapter
│       ├── router
│       ├── schema
│       ├── swagger
│       │   ├── docs.go
│       │   ├── swagger.json
│       │   └── swagger.yaml
│       ├── test
├── pkg
│   ├── auth
│   │   └── jwtauth
│   ├── errors
│   ├── logger
│   │   ├── hook
│   │   │   ├── gorm
│   └── helper
│       ├── hash
│       ├── json
│       ├── structure
│       ├── trace
│       ├── uuid
│       └── yaml
└── scripts
`
