package new

// TplProjectStructure 项目结构
const TplProjectStructure = `
.
├── LICENSE
├── Makefile
├── README.md
├── cmd
│   ├── generator
│   │   └── entgo
│   │       └── main.go
│   └── omgind
│       ├── common
│       │   └── ent_client.go
│       ├── enter
│       │   └── root.go
│       ├── main.go
│       ├── migrate
│       │   ├── district.go
│       │   ├── dump.go
│       │   ├── load.go
│       │   └── migrate.go
│       ├── version
│       │   └── version.go
│       └── web
│           └── web.go
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
│   │       ├── main.go
│   │       ├── mock
│   │       │   ├── demo.mock.go
│   │       │   ├── mock.go
│   │       │   ├── signin.mock.go
│   │       │   ├── sys_address.mock.go
│   │       │   ├── sys_dict.mock.go
│   │       │   ├── sys_district.mock.go
│   │       │   ├── sys_menu.mock.go
│   │       │   ├── sys_role.mock.go
│   │       │   └── sys_user.mock.go
│   │       ├── signin.api.go
│   │       ├── sys_address.api.go
│   │       ├── sys_dict.api.go
│   │       ├── sys_district.api.go
│   │       ├── sys_menu.api.go
│   │       ├── sys_role.api.go
│   │       └── sys_user.api.go
│   ├── app
│   │   ├── app.go
│   │   ├── auth.go
│   │   ├── casbin.go
│   │   ├── contextx
│   │   │   └── contextx.go
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
│   │   │   ├── main.go
│   │   │   ├── signin.sch.go
│   │   │   ├── sys_address.sch.go
│   │   │   ├── sys_dict.sch.go
│   │   │   ├── sys_dict_item.sch.go
│   │   │   ├── sys_district.sch.go
│   │   │   ├── sys_menu.sch.go
│   │   │   ├── sys_menu_action.sch.go
│   │   │   ├── sys_menu_action_resouce.sch.go
│   │   │   ├── sys_role.sch.go
│   │   │   └── sys_user.sch.go
│   │   ├── service
│   │   │   ├── casbin.srv.go
│   │   │   ├── common.srv.go
│   │   │   ├── demo.srv.go
│   │   │   ├── main.go
│   │   │   ├── signin.srv.go
│   │   │   ├── sys_address.srv.go
│   │   │   ├── sys_dict.srv.go
│   │   │   ├── sys_district.srv.go
│   │   │   ├── sys_menu.srv.go
│   │   │   ├── sys_role.srv.go
│   │   │   └── sys_user.srv.go
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
│   ├── gen
│   │   └── ent
│   │       ├── client.go
│   │       ├── debug.go
│   │       ├── ent.go
│   │       ├── enttest
│   │       │   └── enttest.go
│   │       ├── entviz.go
│   │       ├── hook
│   │       │   └── hook.go
│   │       ├── internal
│   │       │   ├── schema.go
│   │       │   └── schemaconfig.go
│   │       ├── migrate
│   │       │   ├── migrate.go
│   │       │   └── schema.go
│   │       ├── mutation.go
│   │       ├── mutation_input.go
│   │       ├── predicate
│   │       │   └── predicate.go
│   │       ├── runtime
│   │       │   └── runtime.go
│   │       ├── runtime.go
│   │       ├── schema-viz.html
│   │       ├── stringer.go
│   │       ├── sysaddress
│   │       │   ├── sysaddress.go
│   │       │   └── where.go
│   │       ├── sysaddress.go
│   │       ├── sysaddress_create.go
│   │       ├── sysaddress_delete.go
│   │       ├── sysaddress_query.go
│   │       ├── sysaddress_update.go
│   │       ├── sysdict
│   │       │   ├── sysdict.go
│   │       │   └── where.go
│   │       ├── sysdict.go
│   │       ├── sysdict_create.go
│   │       ├── sysdict_delete.go
│   │       ├── sysdict_query.go
│   │       ├── sysdict_update.go
│   │       ├── sysdictitem
│   │       │   ├── sysdictitem.go
│   │       │   └── where.go
│   │       ├── sysdictitem.go
│   │       ├── sysdictitem_create.go
│   │       ├── sysdictitem_delete.go
│   │       ├── sysdictitem_query.go
│   │       ├── sysdictitem_update.go
│   │       ├── sysdistrict
│   │       │   ├── sysdistrict.go
│   │       │   └── where.go
│   │       ├── sysdistrict.go
│   │       ├── sysdistrict_create.go
│   │       ├── sysdistrict_delete.go
│   │       ├── sysdistrict_query.go
│   │       ├── sysdistrict_update.go
│   │       ├── sysjwtblock
│   │       │   ├── sysjwtblock.go
│   │       │   └── where.go
│   │       ├── sysjwtblock.go
│   │       ├── sysjwtblock_create.go
│   │       ├── sysjwtblock_delete.go
│   │       ├── sysjwtblock_query.go
│   │       ├── sysjwtblock_update.go
│   │       ├── syslogging
│   │       │   ├── syslogging.go
│   │       │   └── where.go
│   │       ├── syslogging.go
│   │       ├── syslogging_create.go
│   │       ├── syslogging_delete.go
│   │       ├── syslogging_query.go
│   │       ├── syslogging_update.go
│   │       ├── sysmenu
│   │       │   ├── sysmenu.go
│   │       │   └── where.go
│   │       ├── sysmenu.go
│   │       ├── sysmenu_create.go
│   │       ├── sysmenu_delete.go
│   │       ├── sysmenu_query.go
│   │       ├── sysmenu_update.go
│   │       ├── sysmenuaction
│   │       │   ├── sysmenuaction.go
│   │       │   └── where.go
│   │       ├── sysmenuaction.go
│   │       ├── sysmenuaction_create.go
│   │       ├── sysmenuaction_delete.go
│   │       ├── sysmenuaction_query.go
│   │       ├── sysmenuaction_update.go
│   │       ├── sysmenuactionresource
│   │       │   ├── sysmenuactionresource.go
│   │       │   └── where.go
│   │       ├── sysmenuactionresource.go
│   │       ├── sysmenuactionresource_create.go
│   │       ├── sysmenuactionresource_delete.go
│   │       ├── sysmenuactionresource_query.go
│   │       ├── sysmenuactionresource_update.go
│   │       ├── sysrole
│   │       │   ├── sysrole.go
│   │       │   └── where.go
│   │       ├── sysrole.go
│   │       ├── sysrole_create.go
│   │       ├── sysrole_delete.go
│   │       ├── sysrole_query.go
│   │       ├── sysrole_update.go
│   │       ├── sysrolemenu
│   │       │   ├── sysrolemenu.go
│   │       │   └── where.go
│   │       ├── sysrolemenu.go
│   │       ├── sysrolemenu_create.go
│   │       ├── sysrolemenu_delete.go
│   │       ├── sysrolemenu_query.go
│   │       ├── sysrolemenu_update.go
│   │       ├── sysuser
│   │       │   ├── sysuser.go
│   │       │   └── where.go
│   │       ├── sysuser.go
│   │       ├── sysuser_create.go
│   │       ├── sysuser_delete.go
│   │       ├── sysuser_query.go
│   │       ├── sysuser_update.go
│   │       ├── sysuserrole
│   │       │   ├── sysuserrole.go
│   │       │   └── where.go
│   │       ├── sysuserrole.go
│   │       ├── sysuserrole_create.go
│   │       ├── sysuserrole_delete.go
│   │       ├── sysuserrole_query.go
│   │       ├── sysuserrole_update.go
│   │       ├── tx.go
│   │       ├── xxxdemo
│   │       │   ├── where.go
│   │       │   └── xxxdemo.go
│   │       ├── xxxdemo.go
│   │       ├── xxxdemo_create.go
│   │       ├── xxxdemo_delete.go
│   │       ├── xxxdemo_query.go
│   │       └── xxxdemo_update.go
│   ├── router
│   │   ├── r_api.go
│   │   ├── router.go
│   │   ├── v2_demo.go
│   │   ├── v2_sys_address.go
│   │   ├── v2_sys_dict.go
│   │   ├── v2_sys_district.go
│   │   ├── v2_sys_menu.go
│   │   ├── v2_sys_role.go
│   │   └── v2_sys_user.go
│   └── scheme
│       ├── client.go
│       ├── entity
│       │   ├── entity.go
│       │   ├── hooks.go
│       │   ├── sys_address.entity.go
│       │   ├── sys_casbin_rule.entity.go
│       │   ├── sys_dict.entity.go
│       │   ├── sys_dict_item.entity.go
│       │   ├── sys_district.entity.go
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
│       │   ├── creator.go
│       │   ├── hooks.go
│       │   ├── id.go
│       │   ├── meno.go
│       │   ├── mptt_tree.go
│       │   ├── organ.go
│       │   ├── owner.go
│       │   ├── softdel.go
│       │   ├── sort.go
│       │   ├── time.go
│       │   └── tree.go
│       ├── repo
│       │   ├── common.go
│       │   ├── demo.repo.go
│       │   ├── main.go
│       │   ├── sys_address.repo.go
│       │   ├── sys_dict.repo.go
│       │   ├── sys_dict_item.repo.go
│       │   ├── sys_district.repo.go
│       │   ├── sys_menu.repo.go
│       │   ├── sys_menu_action.repo.go
│       │   ├── sys_menu_action_resource.repo.go
│       │   ├── sys_role.repo.go
│       │   ├── sys_role_menu.repo.go
│       │   ├── sys_user.repo.go
│       │   ├── sys_user_role.repo.go
│       │   └── trans.repo.go
│       └── template
│           ├── debug.tpl
│           ├── mutation_input.tpl
│           └── stringer.tpl
├── pkg
│   ├── auth
│   │   ├── auth.go
│   │   └── jwtauth
│   │       ├── auth.go
│   │       ├── auth_test.go
│   │       ├── store
│   │       │   ├── buntdb
│   │       │   │   ├── buntdb.go
│   │       │   │   └── buntdb_test.go
│   │       │   └── redis
│   │       │       ├── redis.go
│   │       │       └── redis_test.go
│   │       ├── store.go
│   │       └── token.go
│   ├── config
│   │   ├── config.go
│   │   └── option
│   │       ├── config.captcha.go
│   │       ├── config.casbin.go
│   │       ├── config.ent.go
│   │       ├── config.http.go
│   │       ├── config.influxdb.go
│   │       ├── config.log.go
│   │       ├── config.menu.go
│   │       ├── config.moniter.go
│   │       ├── config.mysql.go
│   │       ├── config.postgres.go
│   │       ├── config.rabbit.go
│   │       ├── config.redis.go
│   │       ├── config.root.go
│   │       ├── config.sqlite.go
│   │       └── config.system.go
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
│   │   │   ├── entgo
│   │   │   │   └── entgo.go
│   │   │   ├── gorm
│   │   │   │   └── gorm.go
│   │   │   └── hook.go
│   │   └── logger.go
│   └── vcode
│       ├── store
│       │   └── redis.go
│       └── vcode.go
└── scripts
    ├── build-image.sh
    ├── compose.dev
    │   ├── docker-compose.app.yml
    │   ├── docker-compose.dev.yml
    │   ├── docker-compose.external.yml
    │   ├── docker-compose.influxdb.yml
    │   ├── docker-compose.mysql.yml
    │   ├── docker-compose.network.ipv4.yml
    │   ├── docker-compose.network.ipv6.yml
    │   ├── docker-compose.pg.yml
    │   ├── docker-compose.rabbitmq.yml
    │   └── docker-compose.redis.yml
    ├── config.dev.txt
    ├── config_init
    │   ├── mysql
    │   │   └── my.cnf
    │   └── redis
    │       └── redis.conf
    ├── const.dev.sh
    ├── ctrl.dev.sh
    ├── ctrl.help.sh
    ├── g
    │   ├── address.yaml
    │   ├── district.yaml
    │   └── organ.yaml
    ├── simple
    │   ├── docker.golang.sh
    │   ├── docker.influxdb.sh
    │   └── docker.rabbitmq.sh
    ├── sql
    │   ├── district.csv
    │   ├── init_mysql.sql
    │   └── init_postgres.sql
    └── utils.sh
`
