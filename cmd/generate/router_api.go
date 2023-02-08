package generate

import (
	"context"
	"fmt"
	"strings"

	"github.com/heromicro/omgind-cli/helper"
)

func getGeneralRouterAPIFileName(dir string) string {
	fullname := fmt.Sprintf("%s/internal/router/r_api.go", dir)
	return fullname
}

func insertRouterAPI(ctx context.Context, dir, name, apiv string) error {
	fullname := getGeneralRouterAPIFileName(dir)

	pname := helper.ToPlural(helper.ToLowerUnderlinedNamer(name))
	pname = strings.Replace(pname, "_", "-", -1)
	apiContent, err := execParseTpl(routerAPITpl, map[string]string{
		"Name":       name,
		"PluralName": pname,
		"ApiV":       apiv,
	})
	if err != nil {
		return err
	}

	var apiStart int
	apiStack := 0
	insertFn := func(line string) (data string, flag int, ok bool) {
		if apiStart == 0 && strings.Contains(line, fmt.Sprintf("%s := g.Group", apiv)) {
			apiStart = 1
			return
		}

		if apiStart == 1 {
			if v := strings.TrimSpace(line); v == "{" {
				apiStack++
			} else if v == "}" {
				apiStack--
			}

			if apiStack == 0 {
				apiStart = -1
				data = apiContent.String()
				flag = -1
				ok = true
				return
			}
		}
		return "", 0, false
	}

	err = insertContent(fullname, insertFn)
	if err != nil {
		return err
	}

	fmt.Printf("文件[%s]写入成功\n", fullname)

	return execGoFmt(fullname)
}

const routerAPITpl = `
a.init{{.Name}}Router{{.ApiV | ToUpper}}(v2, a.{{.Name}}API{{.ApiV | ToUpper}}, "{{.Name | ToLower}}")

`
