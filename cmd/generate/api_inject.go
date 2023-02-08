package generate

import (
	"context"
	"fmt"
	"strings"
)

func getAPIInjectFileName(dir, apiv string) string {
	fullname := fmt.Sprintf("%s/internal/api/%s/main.go", dir, apiv)
	return fullname
}

func insertAPIInject(ctx context.Context, dir, name, apiv string) error {
	injectContent := fmt.Sprintf("%sSet,", name)
	injectStart := 0

	insertFn := func(line string) (data string, flag int, ok bool) {

		setvar := fmt.Sprintf("var API%sSet = wire.NewSet(", strings.ToUpper(apiv))

		if injectStart == 0 && strings.Contains(line, setvar) {
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

	filename := getAPIInjectFileName(dir, apiv)
	err := insertContent(filename, insertFn)
	if err != nil {
		return err
	}

	fmt.Printf("文件[%s]写入成功\n", filename)

	return execGoFmt(filename)
}
