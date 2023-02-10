package generate

import (
	"context"
	"fmt"
	"strings"

	"github.com/heromicro/omgind-cli/helper"
)

func getRouterAPIFileName(dir, name, apiv string) string {
	fullname := fmt.Sprintf("%s/internal/router/%s_%s.go", dir, apiv, helper.ToLowerUnderlinedNamer(name))
	return fullname
}

// 生成model实现文件
func genRouterImpl(ctx context.Context, pkgName, dir, name, apiv string) error {

	data := map[string]interface{}{
		"PkgName":    pkgName,
		"Name":       name,
		"PluralName": helper.ToPlural(name),
		"EntPackage": strings.ToLower(name),
		"Apiv":       apiv,
	}

	buf, err := execParseTpl(routerImplTpl, data)
	if err != nil {
		return err
	}

	fullname := getRouterAPIFileName(dir, name, apiv)
	err = createFile(ctx, fullname, buf)
	if err != nil {
		return err
	}

	fmt.Printf("文件[%s]写入成功\n", fullname)

	return execGoFmt(fullname)
}

const routerImplTpl = `
package router

import (
	"github.com/gin-gonic/gin"
	api_{{.Apiv}} "{{.PkgName}}/internal/api/{{.Apiv}}"
)

func (r *Router) init{{.Name}}Router{{.Apiv | ToUpper}}(urg *gin.RouterGroup, api *api_{{.Apiv}}.{{.Name}}, pathcomp string) {
	g{{.Name}} := urg.Group(pathcomp)
	{
		g{{.Name}}.GET("", api.Query)
		g{{.Name}}.GET(":id", api.Get)
		g{{.Name}}.POST("", api.Create)
		g{{.Name}}.PUT(":id", api.Update)
		g{{.Name}}.DELETE(":id", api.Delete)
		g{{.Name}}.PATCH(":id/enable", api.Enable)
		g{{.Name}}.PATCH(":id/disable", api.Disable)
	}
}

`
