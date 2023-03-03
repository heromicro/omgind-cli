package new

// TplProjectStructure 项目结构
const TplProjectStructure = `
├── LICENSE
├── Makefile
├── README.md
├── cmd
│   └── omgind
│       └── main.go # 入口文件
├── configs
│   ├── config.dev.toml
│   ├── config.docker.toml
│   ├── config.prod.toml
│   ├── config.toml -> config.dev.toml
│   ├── menu.yaml
│   └── model.conf
├── deploy
│   ├── compose.dev
│   │   ├── docker-compose.app.dev.yml
│   │   ├── docker-compose.dev.yml
│   │   ├── docker-compose.influxdb.dev.yml
│   │   ├── docker-compose.mysql.dev.yml
│   │   ├── docker-compose.network.dev.yml
│   │   ├── docker-compose.pg.dev.yml
│   │   ├── docker-compose.rabbitmq.dev.yml
│   │   └── docker-compose.redis.dev.yml
│   ├── config.dev.txt
│   ├── const.dev.sh
│   └── dockerfile
│       ├── dockerfile.dev
│       └── run.sh
├── docs
│   └── data_model.md
├── go.mod
├── go.sum
├── internal
│   ├── api
│   │   └── v2
│   │       ├── demo.api.go
│   │       ├── dict.api.go
│   │       ├── main.go
│   │       ├── menu.api.go
│   │       ├── mock
│   │       │   ├── demo.mock.go
│   │       │   ├── dict.mock.go
│   │       │   ├── menu.mock.go
│   │       │   ├── mock.go
│   │       │   ├── role.mock.go
│   │       │   ├── signin.mock.go
│   │       │   └── user.mock.go
│   │       ├── role.api.go
│   │       ├── signin.api.go
│   │       └── user.api.go
│   ├── app
│   │   ├── app.go
│   │   ├── auth.go
│   │   ├── casbin.go
│   │   ├── contextx
│   │   │   └── contextx.go
│   │   ├── ent.go
│   │   ├── ginx
│   │   │   └── ginx.go
│   │   ├── influxdb.go
│   │   ├── injector.go
│   │   ├── logger.go
│   │   ├── middleware
│   │   │   ├── healthchek
│   │   │   │   └── mw_healthcheck.go
│   │   │   ├── middleware.go
│   │   │   ├── mw_auth.go
│   │   │   ├── mw_casbin.go
│   │   │   ├── mw_copy_body.go
│   │   │   ├── mw_cors.go
│   │   │   ├── mw_logger.go
│   │   │   ├── mw_rate_limiter.go
│   │   │   ├── mw_recover.go
│   │   │   ├── mw_trace.go
│   │   │   └── mw_www.go
│   │   ├── module
│   │   │   └── adapter
│   │   │       └── casbin.go
│   │   ├── rabbitmq.go
│   │   ├── redis.go
│   │   ├── resolver
│   │   │   └── resolver.go
│   │   ├── schema
│   │   │   ├── demo.sch.go
│   │   │   ├── dict.sch.go
│   │   │   ├── dict_item.sch.go
│   │   │   ├── main.go
│   │   │   ├── menu.sch.go
│   │   │   ├── menu_action.sch.go
│   │   │   ├── menu_action_resouce.sch.go
│   │   │   ├── role.sch.go
│   │   │   ├── signin.sch.go
│   │   │   └── user.sch.go
│   │   ├── service
│   │   │   ├── casbin.srv.go
│   │   │   ├── common.srv.go
│   │   │   ├── demo.srv.go
│   │   │   ├── dict.srv.go
│   │   │   ├── main.go
│   │   │   ├── menu.srv.go
│   │   │   ├── role.srv.go
│   │   │   ├── signin.srv.go
│   │   │   └── user.srv.go
│   │   ├── swagger
│   │   │   ├── docs.go
│   │   │   ├── swagger.json
│   │   │   └── swagger.yaml
│   │   ├── test
│   │   │   ├── doc.go
│   │   │   ├── t_demo_test.go
│   │   │   ├── t_init_test.go
│   │   │   ├── t_menu_test.go
│   │   │   ├── t_role_test.go
│   │   │   └── t_user_test.go
│   │   ├── vcode.go
│   │   ├── web.go
│   │   ├── wire.go
│   │   └── wire_gen.go
│   ├── router
│   │   ├── r_api.go
│   │   ├── router.go
│   │   ├── v2_demo.go
│   │   ├── v2_dict.go
│   │   ├── v2_menu.go
│   │   ├── v2_role.go
│   │   └── v2_user.go
│   └── schema
│       ├── entity
│       │   ├── entity.go
│       │   ├── sys_casbin_rule.entity.go
│       │   ├── sys_dict.entity.go
│       │   ├── sys_dict_item.entity.go
│       │   ├── sys_jwt_blocklist.go
│       │   ├── sys_logging.go
│       │   ├── sys_menu.entity.go
│       │   ├── sys_menu_action.entity.go
│       │   ├── sys_menu_action_resource.entity.go
│       │   ├── sys_role.entity.go
│       │   ├── sys_role_menu.entity.go
│       │   ├── sys_user.entity.go
│       │   ├── sys_user_role.entity.go
│       │   └── xxx_demo.go
│       ├── mixin
│       │   ├── active.go
│       │   ├── id.go
│       │   ├── meno.go
│       │   ├── sort.go
│       │   ├── status.go
│       │   └── time.go
│       ├── repo
│       │   ├── common.go
│       │   ├── demo.repo.go
│       │   ├── dict.repo.go
│       │   ├── dict_item.repo.go
│       │   ├── main.go
│       │   ├── menu.repo.go
│       │   ├── menu_action.repo.go
│       │   ├── menu_action_resource.repo.go
│       │   ├── role.repo.go
│       │   ├── role_menu.repo.go
│       │   ├── trans.repo.go
│       │   ├── user.repo.go
│       │   └── user_role.repo.go
│       └── template
│           └── mutation_input.tpl
├── pkg
│   ├── auth
│   │   ├── auth.go
│   │   └── jwtauth
│   │       ├── auth.go
│   │       ├── auth_test.go
│   │       ├── store
│   │       │   ├── buntdb
│   │       │   └── redis
│   │       ├── store.go
│   │       └── token.go
│   ├── config
│   │   ├── config.captcha.go
│   │   ├── config.casbin.go
│   │   ├── config.ent.go
│   │   ├── config.go
│   │   ├── config.http.go
│   │   ├── config.influxdb.go
│   │   ├── config.log.go
│   │   ├── config.menu.go
│   │   ├── config.moniter.go
│   │   ├── config.mysql.go
│   │   ├── config.postgres.go
│   │   ├── config.rabbit.go
│   │   ├── config.redis.go
│   │   ├── config.root.go
│   │   ├── config.sqlite.go
│   │   └── config.system.go
│   ├── errors
│   │   ├── errors.go
│   │   └── response.go
│   ├── global
│   │   └── global.go
│   ├── helper
│   │   ├── deepcopier
│   │   │   └── deepcopier.go
│   │   ├── hash
│   │   │   └── hash.go
│   │   ├── json
│   │   │   └── json.go
│   │   ├── pulid
│   │   │   └── pulid.go
│   │   ├── str
│   │   │   ├── conv.go
│   │   │   └── random.go
│   │   ├── structure
│   │   │   └── structure.go
│   │   ├── trace
│   │   │   └── trace.go
│   │   ├── uid
│   │   │   ├── id.go
│   │   │   ├── ulid
│   │   │   │   └── ulid.go
│   │   │   └── uuid
│   │   │       └── uuid.go
│   │   └── yaml
│   │       └── yaml.go
│   ├── logger
│   │   ├── hook
│   │   │   ├── gorm
│   │   │   │   └── gorm.go
│   │   │   └── hook.go
│   │   └── logger.go
│   └── vcode
│       ├── store
│       │   └── redis.go
│       └── vcode.go
└── scripts
`
