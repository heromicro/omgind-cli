package generate

import (
	"context"
	"errors"
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
	GenPkg     string
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
	switch a.cfg.Storage {
	case "", "ent", "entgo", "gorm", "mongo", "mongodb":
	default:
		return fmt.Errorf("unsupported storage %q: choose entgo, gorm or mongo per model", a.cfg.Storage)
	}

	var generationErrors []error
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

	if a.cfg.Modules == "entity" {
		switch a.cfg.Storage {
		case "mongo", "mongodb":
			return genMongoEntity(ctx, pkgName, dir, item.StructName, item.Comment, item.toEntityMongoFields()...)
		case "gorm":
			return genGormEntity(ctx, pkgName, dir, item.StructName, item.Comment, item.toEntityGormFields()...)
		default:
			return genEntEntity(ctx, pkgName, dir, item.StructName, item.Comment, item.Mixin, item.toEntityEntFields()...)
		}
	}

	if a.hasModule("repo") {
		// err = genRepo(ctx, pkgName, dir, item.StructName, item.Comment)
		// a.handleError(err, "Generate model interface")
		switch a.cfg.Storage {
		case "mongo", "mongodb":
			err = genMongoEntity(ctx, pkgName, dir, item.StructName, item.Comment, item.toEntityMongoFields()...)
			if err != nil {
				generationErrors = append(generationErrors, fmt.Errorf("generate mongo entity: %w", err))
			}

			err = insertEntityInjectMongo(ctx, dir, item.StructName)
			if err != nil {
				generationErrors = append(generationErrors, fmt.Errorf("insert mongo entity inject: %w", err))
			}

			err = genModelImplMongo(ctx, pkgName, dir, item.StructName, item.Comment)
			if err != nil {
				generationErrors = append(generationErrors, fmt.Errorf("generate mongo model: %w", err))
			}

			err = insertRepoInjectMongo(ctx, dir, item.StructName)
			if err != nil {
				generationErrors = append(generationErrors, fmt.Errorf("insert mongo model inject: %w", err))
			}
		case "gorm":
			err = genGormEntity(ctx, pkgName, dir, item.StructName, item.Comment, item.toEntityGormFields()...)
			if err != nil {
				generationErrors = append(generationErrors, fmt.Errorf("generate gorm entity: %w", err))
			}

			err = insertEntityInjectGorm(ctx, dir, item.StructName)
			if err != nil {
				generationErrors = append(generationErrors, fmt.Errorf("insert gorm entity inject: %w", err))
			}

			err = genRepoImplGorm(ctx, pkgName, dir, item.StructName, item.Comment)
			if err != nil {
				generationErrors = append(generationErrors, fmt.Errorf("generate gorm model: %w", err))
			}

			err = insertRepoInjectGorm(ctx, dir, item.StructName)
			if err != nil {
				generationErrors = append(generationErrors, fmt.Errorf("insert gorm model inject: %w", err))
			}
		default:
			err = genEntEntity(ctx, pkgName, dir, item.StructName, item.Comment, item.Mixin, item.
				toEntityEntFields()...)
			a.handleError(err, "Generate ent entity")

			err = genRepoImplEnt(ctx, pkgName, dir, a.cfg.GenPkg, item.StructName, item.Comment)
			a.handleError(err, "Generate ent entity")
			err = insertRepoInjectEnt(ctx, dir, item.StructName)
			a.handleError(err, "Insert ent model inject")

		}
	}

	if a.hasModule("service") {
		//err = genBll(ctx, pkgName, dir, item.StructName, item.Comment)
		//a.handleError(err, "Generate service interface")

		err = genServiceImpl(ctx, pkgName, dir, item.StructName, item.Comment, a.cfg.Storage)
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

	return errors.Join(generationErrors...)
}
