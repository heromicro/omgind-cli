package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/heromicro/omgind-cli/cmd/generate"
	"github.com/heromicro/omgind-cli/cmd/new"
	"github.com/urfave/cli"
)

// NewCommand 创建项目命令
func NewCommand() cli.Command {
	return cli.Command{
		Name:    "new",
		Aliases: []string{"n"},
		Usage:   "create new project",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "dir, d",
				Usage: "项目生成目录(默认GOPATH+包名)",
			},
			&cli.StringFlag{
				Name:  "pkg, p",
				Usage: "项目包名",
			},
			&cli.StringFlag{
				Name:  "branch, b",
				Usage: "指定分支(默认master)",
			},
			&cli.BoolFlag{
				Name:  "mirror, m",
				Usage: "使用国内镜像(gitee.com)",
			},
			&cli.BoolFlag{
				Name:  "web, w",
				Usage: "包含web项目",
			},
		},
		Action: func(c *cli.Context) error {
			cfg := new.Config{
				Dir:        c.String("dir"),
				PkgName:    c.String("pkg"),
				UseMirror:  c.Bool("mirror"),
				Branch:     c.String("branch"),
				IncludeWeb: c.Bool("web"),
			}

			if cfg.PkgName == "" {
				return errors.New("请指定包名")
			}

			if cfg.Dir == "" {
				vpath := os.Getenv("GOPATH")
				if vpath == "" {
					return errors.New("请指定dir或者设置GOPATH")
				}
				cfg.Dir = filepath.Join(vpath, "src", cfg.PkgName)
			}

			cfg.AppName = filepath.Base(cfg.PkgName)

			return new.Exec(cfg)
		},
	}
}

// GenerateCommand 生成项目模块命令
func GenerateCommand() cli.Command {
	return cli.Command{
		Name:    "generate",
		Aliases: []string{"g"},
		Usage:   "generate code for project",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "dir, d",
				Usage: "项目生成目录(默认GOPATH)",
			},
			&cli.StringFlag{
				Name:     "pkg, p",
				Usage:    "项目包名",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "name, n",
				Usage: "业务模块名称(结构体名称)",
			},
			&cli.StringFlag{
				Name:  "comment, c",
				Usage: "业务模块注释(结构体注释)",
			},
			&cli.StringFlag{
				Name:  "file, f",
				Usage: "指定模板文件(.yaml，模板配置可参考说明)",
			},
			&cli.StringFlag{
				Name:  "apiversion, v",
				Usage: "api 版本, 取值: v1或v2或v3, api的路径在:${project}/internal/api/${vx}/",
				Value: "v2",
			},
			&cli.StringFlag{
				Name:  "module, m",
				Usage: "指定生成模块（默认生成全部模块，以逗号分隔，支持：schema, repo, service, api,mock,router）",
			},
			&cli.StringFlag{
				Name:  "genpkg",
				Usage: "指定生成的包名最后一节, 比如完整包名为internal/gen/mainent, 则为mainent",
			},
			&cli.StringFlag{
				Name:  "storage, s",
				Usage: "指定model的实现存储（默认ent，支持：ent）",
			},
		},
		Action: func(c *cli.Context) error {
			cfg := generate.Config{
				Dir:        c.String("dir"),
				PkgName:    c.String("pkg"),
				Name:       c.String("name"),
				Comment:    c.String("comment"),
				File:       c.String("file"),
				Modules:    c.String("module"),
				Storage:    c.String("storage"),
				ApiVersion: c.String("apiversion"),
				GenPkg:     c.String("genpkg"),
			}

			if cfg.Dir == "" {
				vpath := os.Getenv("GOPATH")
				if vpath == "" {
					return errors.New("请指定dir或者设置GOPATH")
				}
				cfg.Dir = filepath.Join(vpath, "src", cfg.PkgName)
			}

			if cfg.PkgName == "" {
				fmt.Println("请指定包名")
				return nil
			} else if cfg.Name == "" && cfg.File == "" {
				fmt.Println("请指定模块名称或模板配置文件")
				return nil
			} else if cfg.Name != "" && cfg.Comment == "" {
				fmt.Println("请指定模块说明")
				return nil
			}

			if cfg.GenPkg == "" {
				// fmt.Println("请指定生成的包名最后一节")
				// return nil
				cfg.GenPkg = "mainent"
			}

			return generate.Exec(cfg)
		},
	}
}
