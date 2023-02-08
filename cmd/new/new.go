package new

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	githubSource    = "https://github.com/heromicro/omgind.git"
	giteeSource     = ""
	githubWebSource = "https://github.com/heromicro/omgind-react.git"
	giteeWebSource  = ""
	defaultPkgName  = "github.com/heromicro/omgind"
	defaultAppName  = "omgind"
)

// Config 配置参数
type Config struct {
	Dir        string
	PkgName    string
	AppName    string
	Branch     string
	UseMirror  bool
	IncludeWeb bool
}

// Exec 执行创建项目命令
func Exec(cfg Config) error {
	cmd := &Command{cfg: &cfg}
	return cmd.Exec()
}

// Command 创建项目命令
type Command struct {
	cfg *Config
}

// Exec 执行命令
func (a *Command) Exec() error {
	dir, err := filepath.Abs(a.cfg.Dir)
	if err != nil {
		return err
	}

	log.Printf("项目生成目录：%s", dir)

	notExist := false
	_, err = os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			notExist = true
		} else {
			return err
		}
	}

	if notExist {
		source := githubSource
		if a.cfg.UseMirror {
			source = giteeSource
		}
		err = a.gitClone(dir, source)
		if err != nil {
			return err
		}

		if a.cfg.IncludeWeb {
			ws := githubWebSource
			if a.cfg.UseMirror {
				ws = giteeWebSource
			}
			err = a.gitClone(filepath.Join(dir, "web"), ws)
			if err != nil {
				return err
			}
		}
	}

	if pkgName := a.cfg.PkgName; pkgName != "" && pkgName != defaultPkgName {
		err := a.changeDirPkgName(dir)
		if err != nil {
			return err
		}

		err = a.changeFilePkgName(fmt.Sprintf("%s/go.mod", a.cfg.Dir))
		if err != nil {
			return err
		}

		err = a.changeFileAppNames(
			fmt.Sprintf("%s/Makefile", a.cfg.Dir),
			fmt.Sprintf("%s/.gitignore", a.cfg.Dir),
			fmt.Sprintf("%s/.air.conf", a.cfg.Dir),

			fmt.Sprintf("%s/internal/app/ginx/ginx.go", a.cfg.Dir),
			fmt.Sprintf("%s/internal/app/swagger/docs.go", a.cfg.Dir),
			fmt.Sprintf("%s/internal/app/swagger/swagger.json", a.cfg.Dir),
			fmt.Sprintf("%s/internal/app/swagger/swagger.yaml", a.cfg.Dir),
			fmt.Sprintf("%s/pkg/auth/jwtauth/auth.go", a.cfg.Dir),

			fmt.Sprintf("%s/configs/config.toml", a.cfg.Dir),
			fmt.Sprintf("%s/configs/config.prod.toml", a.cfg.Dir),
			fmt.Sprintf("%s/configs/config.dev.toml", a.cfg.Dir),
			fmt.Sprintf("%s/configs/config.docker.toml", a.cfg.Dir),

			fmt.Sprintf("%s/deploy/config.dev.txt", a.cfg.Dir),
			fmt.Sprintf("%s/deploy/compose.dev/docker-compose.app.dev.yml", a.cfg.Dir),
			fmt.Sprintf("%s/deploy/compose.dev/docker-compose.dev.yml", a.cfg.Dir),
			fmt.Sprintf("%s/deploy/compose.dev/docker-compose.influxdb.dev.yml", a.cfg.Dir),
			fmt.Sprintf("%s/deploy/compose.dev/docker-compose.mysql.dev.yml", a.cfg.Dir),
			fmt.Sprintf("%s/deploy/compose.dev/docker-compose.pg.dev.yml", a.cfg.Dir),
			fmt.Sprintf("%s/deploy/compose.dev/docker-compose.rabbitmq.dev.yml", a.cfg.Dir),
			fmt.Sprintf("%s/deploy/compose.dev/docker-compose.redis.dev.yml", a.cfg.Dir),

			fmt.Sprintf("%s/deploy/dockerfile/run.sh", a.cfg.Dir),

			fmt.Sprintf("%s/scripts/config.dev.txt", a.cfg.Dir),
			fmt.Sprintf("%s/scripts/const.dev.sh", a.cfg.Dir),
			fmt.Sprintf("%s/scripts/build-image.sh", a.cfg.Dir),
			fmt.Sprintf("%s/scripts/ctrl.help.sh", a.cfg.Dir),

			fmt.Sprintf("%s/scripts/compose.dev/docker-compose.dev.yml", a.cfg.Dir),
			fmt.Sprintf("%s/scripts/compose.dev/docker-compose.app.yml", a.cfg.Dir),
			fmt.Sprintf("%s/scripts/compose.dev/docker-compose.influxdb.yml", a.cfg.Dir),
			fmt.Sprintf("%s/scripts/compose.dev/docker-compose.mysql.yml", a.cfg.Dir),
			fmt.Sprintf("%s/scripts/compose.dev/docker-compose.pg.yml", a.cfg.Dir),
			fmt.Sprintf("%s/scripts/compose.dev/docker-compose.rabbitmq.yml", a.cfg.Dir),
			fmt.Sprintf("%s/scripts/compose.dev/docker-compose.redis.yml", a.cfg.Dir),

			fmt.Sprintf("%s/scripts/simple/docker.golang.sh", a.cfg.Dir),
			fmt.Sprintf("%s/scripts/simple/docker.influxdb.sh", a.cfg.Dir),
			fmt.Sprintf("%s/scripts/simple/docker.rabbitmq.sh", a.cfg.Dir),

			fmt.Sprintf("%s/scripts/sql/init_mysql.sql", a.cfg.Dir),
			fmt.Sprintf("%s/scripts/sql/init_postgres.sql", a.cfg.Dir),
			fmt.Sprintf("%s/cmd/%s/main.go", a.cfg.Dir, defaultAppName),
		)
		if err != nil {
			return err
		}

		// change app name
		os.Rename(fmt.Sprintf("%s/cmd/%s", a.cfg.Dir, defaultAppName), fmt.Sprintf("%s/cmd/%s", a.cfg.Dir, a.cfg.AppName))
	}

	if notExist {
		err = a.gitInit(dir)
		if err != nil {
			return err
		}
	}

	fmt.Printf("\n项目创建成功：%s\n", dir)
	fmt.Println(TplProjectStructure)

	return nil
}

func (a *Command) execGit(dir string, args ...string) error {
	cmd := exec.Command("git", args...)
	if dir != "" {
		cmd.Dir = dir
	}

	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	err := cmd.Start()
	if err != nil {
		return err
	}

	go func() {
		io.Copy(os.Stdout, stdoutIn)
	}()

	go func() {
		io.Copy(os.Stderr, stderrIn)
	}()

	return cmd.Wait()
}

func (a *Command) gitClone(dir, source string) error {
	var args []string
	args = append(args, "clone")

	branch := "main"
	if v := a.cfg.Branch; v != "" {
		branch = v
	}
	args = append(args, "-b", branch)

	args = append(args, source)
	args = append(args, dir)

	log.Printf("执行命令：git %s", strings.Join(args, " "))
	return a.execGit("", args...)
}

func (a *Command) gitInit(dir string) error {
	os.RemoveAll(filepath.Join(dir, ".git"))
	if a.cfg.IncludeWeb {
		os.RemoveAll(filepath.Join(dir, "web", ".git"))
	}

	err := a.execGit(dir, "init")
	if err != nil {
		return err
	}

	err = a.execGit(dir, "add", "-A")
	if err != nil {
		return err
	}

	err = a.execGit(dir, "commit", "-m", "Initial commit")
	if err != nil {
		return err
	}

	return nil
}

func (a *Command) checkInDirs(dir, path string) bool {
	includeDirs := []string{"cmd", "internal", "pkg"}
	for _, d := range includeDirs {
		p := filepath.Join(dir, d)
		if strings.HasPrefix(path, p) {
			return true
		}
	}
	return false
}

func (a *Command) changeFileAppNames(names ...string) error {
	for _, name := range names {
		err := a.changeFileAppName(name)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *Command) changeFileAppName(name string) error {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		log.Printf("%s file not found", name)
		return nil
	} else {
		log.Printf("replace %s to %s in file %s", defaultAppName, a.cfg.AppName, name)
	}
	return a.readAndReplaceFile(name, func(line string) string {
		if strings.Contains(line, defaultAppName) {
			return strings.Replace(line, defaultAppName, a.cfg.AppName, -1)
		}
		return line
	})
}

func (a *Command) changeFilePkgName(name string) error {
	return a.readAndReplaceFile(name, func(line string) string {
		if strings.Contains(line, defaultPkgName) {
			return strings.Replace(line, defaultPkgName, a.cfg.PkgName, 1)
		}
		return line
	})
}

func (a *Command) changeDirPkgName(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(path) != ".go" || info.IsDir() || !a.checkInDirs(dir, path) {
			return nil
		}

		return a.readAndReplaceFile(path, func(line string) string {
			if strings.Contains(line, defaultPkgName) {
				return strings.Replace(line, defaultPkgName, a.cfg.PkgName, 1)
			}
			return line
		})
	})
}

func (a *Command) readAndReplaceFile(name string, call func(string) string) error {
	buf, err := a.readFileAndReplace(name, call)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(name, buf.Bytes(), 0644)
}

func (a *Command) readFileAndReplace(name string, call func(string) string) (*bytes.Buffer, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := call(scanner.Text())
		buf.WriteString(line)
		buf.WriteByte('\n')
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return buf, nil
}
