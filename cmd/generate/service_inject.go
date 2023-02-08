package generate

import (
	"context"
	"fmt"
	"strings"
)

func getBllInjectFileName(dir string) string {
	fullname := fmt.Sprintf("%s/internal/app/service/main.go", dir)
	return fullname
}

// 插入bll注入文件
func insertServiceInject(ctx context.Context, dir, name string) error {
	fullname := getBllInjectFileName(dir)

	injectContent := fmt.Sprintf("%sSet,", name)
	injectStart := 0
	insertFn := func(line string) (data string, flag int, ok bool) {
		srv_set := "var ServiceSet = wire.NewSet("
		if injectStart == 0 && strings.Contains(line, srv_set) {
			injectStart = 1
			return
		}

		if injectStart == 1 && strings.Contains(line, ")") {
			injectStart = -1
			data = injectContent
			flag = -1
			ok = true
			return
		}

		return "", 0, false
	}

	err := insertContent(fullname, insertFn)
	if err != nil {
		return err
	}

	fmt.Printf("文件[%s]写入成功\n", fullname)

	return execGoFmt(fullname)
}
