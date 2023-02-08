package generate

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

// Config 配置参数
type Config struct {
	Dir        string
	PkgName    string
	Name       string
	Comment    string
	File       string
	Storage    string
	Modules    string
	ApiVersion string
}

// Exec 执行生成模块命令
func Exec(cfg Config) error {
	cmd := &Command{cfg: &cfg}
	return cmd.Exec()
}

// Command 生成命令
type Command struct {
	cfg *Config
}

func (a *Command) hasModule(m string) bool {
	if v := a.cfg.Modules; v == "" || v == "all" {
		return true
	}

	for _, s := range strings.Split(a.cfg.Modules, ",") {
		if s == m {
			return true
		}
	}

	return false
}

func (a *Command) handleError(err error, desc string) {
	if err != nil {
		fmt.Printf("%s:%s", desc, err.Error())
	}
}

// Exec 执行命令
func (a *Command) Exec() error {
	var item TplItem

	if a.cfg.File != "" {
		// 模版文件不为空,则读取
		b, err := readFile(a.cfg.File)
		if err != nil {
			return err
		}

		err = yaml.Unmarshal(b.Bytes(), &item)
		if err != nil {
			return err
		}
	} else {
		// 模板文件为空, 使用 结构体名称及结构体注释
		item.StructName = a.cfg.Name
		item.Comment = a.cfg.Comment
	}

	dir, err := filepath.Abs(a.cfg.Dir)
	if err != nil {
		return err
	}

	pkgName := a.cfg.PkgName
	ctx := context.Background()

	if a.hasModule("schema") {
		err = genSchema(ctx, pkgName, dir, item.StructName, item.Comment, item.toSchemaFields()...)
		if err != nil {
			return err
		}
	}

	if a.hasModule("repo") {
		// err = genRepo(ctx, pkgName, dir, item.StructName, item.Comment)
		// a.handleError(err, "Generate model interface")
		switch a.cfg.Storage {
		// case "mongo":
		// 	err = genMongoEntity(ctx, pkgName, dir, item.StructName, item.Comment, item.toEntityMongoFields()...)
		// 	a.handleError(err, "Generate mongo entity")

		// 	err = insertEntityInjectMongo(ctx, dir, item.StructName)
		// 	a.handleError(err, "Insert mongo entity inject")

		// 	err = genModelImplMongo(ctx, pkgName, dir, item.StructName, item.Comment)
		// 	a.handleError(err, "Generate mongo model")

		// 	err = insertRepoInjectMongo(ctx, dir, item.StructName)
		// 	a.handleError(err, "Insert mongo model inject")
		case "gorm":
			err = genGormEntity(ctx, pkgName, dir, item.StructName, item.Comment, item.toEntityGormFields()...)
			a.handleError(err, "Generate gorm entity")

			err = insertEntityInjectGorm(ctx, dir, item.StructName)
			a.handleError(err, "Insert gorm entity inject")

			err = genRepoImplGorm(ctx, pkgName, dir, item.StructName, item.Comment)
			a.handleError(err, "Generate gorm model")

			err = insertRepoInjectGorm(ctx, dir, item.StructName)
			a.handleError(err, "Insert gorm model inject")
		default:
			err = genEntEntity(ctx, pkgName, dir, item.StructName, item.Comment, item.Mixin, item.
				toEntityEntFields()...)
			a.handleError(err, "Generate ent entity")

			err = genRepoImplEnt(ctx, pkgName, dir, item.StructName, item.Comment)
			a.handleError(err, "Generate ent entity")
			err = insertRepoInjectEnt(ctx, dir, item.StructName)
			a.handleError(err, "Insert ent model inject")

		}
	}

	if a.hasModule("service") {
		//err = genBll(ctx, pkgName, dir, item.StructName, item.Comment)
		//a.handleError(err, "Generate service interface")

		err = genServiceImpl(ctx, pkgName, dir, item.StructName, item.Comment)
		a.handleError(err, "Generate service impl")

		err = insertServiceInject(ctx, dir, item.StructName)
		a.handleError(err, "Insert service inject")
	}

	if a.hasModule("api") {
		err = genAPI(ctx, pkgName, dir, item.StructName, item.Comment, a.cfg.ApiVersion)
		a.handleError(err, "Generate api")

		err = insertAPIInject(ctx, dir, item.StructName, a.cfg.ApiVersion)
		a.handleError(err, "Insert api inject")
	}

	if a.hasModule("mock") {
		err = genAPIMock(ctx, pkgName, dir, item.StructName, item.Comment, a.cfg.ApiVersion)
		a.handleError(err, "Generate api mock")

		err = insertAPIMockInject(ctx, dir, item.StructName, a.cfg.ApiVersion)
		a.handleError(err, "Insert api mock inject")
	}

	if a.hasModule("router") {
		err = genRouterImpl(ctx, pkgName, dir, item.StructName, a.cfg.ApiVersion)
		err = insertRouterAPI(ctx, dir, item.StructName, a.cfg.ApiVersion)
		a.handleError(err, "Insert router api")

		err = insertRouterInject(ctx, dir, item.StructName, a.cfg.ApiVersion)
		a.handleError(err, "Insert router inject")
	}

	return nil
}
