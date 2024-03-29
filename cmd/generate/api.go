package generate

import (
	"context"
	"fmt"
	"strings"

	"github.com/heromicro/omgind-cli/helper"
)

func getAPIFileName(dir, name, apiv string) string {
	fullname := fmt.Sprintf("%s/internal/api/%s/%s.api.go", dir, apiv, helper.ToLowerUnderlinedNamer(name))
	return fullname
}

func genAPI(ctx context.Context, pkgName, dir, name, comment, apiv string) error {
	pname := helper.ToPlural(helper.ToLowerUnderlinedNamer(name))
	pname = strings.Replace(pname, "_", "-", -1)

	data := map[string]interface{}{
		"PkgName": pkgName,
		"Name":    name,
		"Comment": comment,
		"Apiv":    apiv,
	}

	buf, err := execParseTpl(apiTpl, data)
	if err != nil {
		return err
	}

	fullname := getAPIFileName(dir, name, apiv)
	err = createFile(ctx, fullname, buf)
	if err != nil {
		return err
	}

	fmt.Printf("文件[%s]写入成功\n", fullname)

	return execGoFmt(fullname)
}

const apiTpl = `
package api_{{.Apiv}}

import (
	"{{.PkgName}}/internal/app/service"
	"{{.PkgName}}/internal/app/ginx"
	"{{.PkgName}}/internal/app/schema"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// {{.Name}}Set 注入{{.Name}}
var {{.Name}}Set = wire.NewSet(wire.Struct(new({{.Name}}), "*"))

// {{.Name}} {{.Comment}}
type {{.Name}} struct {
	{{.Name}}Srv *service.{{.Name}}
}

// Query 查询数据
func (a *{{.Name}}) Query(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.{{.Name}}QueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResErrorCode(c, -1100, err)
		return
	}

	params.Pagination = true
	result, err := a.{{.Name}}Srv.Query(ctx, params)
	if err != nil {
		ginx.ResErrorCode(c, -1110, err)
		return
	}
	
	ginx.ResPage(c, result.Data, result.PageResult)
}

// QuerySelectPage 查询数据
func (a *{{.Name}}) QuerySelectPage(c *gin.Context) {
	ctx := c.Request.Context()
	var params schema.{{.Name}}QueryParam
	if err := ginx.ParseQuery(c, &params); err != nil {
		ginx.ResErrorCode(c, -1100, err)
		return
	}

	params.Pagination = true
	result, err := a.{{.Name}}Srv.QuerySelectPage(ctx, params)
	if err != nil {
		ginx.ResErrorCode(c, -1120, err)
		return
	}
	
	ginx.ResPage(c, result.Data, result.PageResult)
}

// Get 查询指定数据
func (a *{{.Name}}) Get(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.{{.Name}}Srv.Get(ctx, c.Param("id"))
	if err != nil {
		ginx.ResErrorCode(c, -1130, err)
		return
	}
	ginx.ResSuccess(c, item)
}

// View 查询指定数据
func (a *{{.Name}}) View(c *gin.Context) {
	ctx := c.Request.Context()
	item, err := a.{{.Name}}Srv.View(ctx, c.Param("id"))
	if err != nil {
		ginx.ResErrorCode(c, -1140, err)
		return
	}
	ginx.ResSuccess(c, item)
}

// Create 创建数据
func (a *{{.Name}}) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.{{.Name}}
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResErrorCode(c, -1100, err)
		return
	}

	// item.Creator = ginx.GetUserID(c)
	result, err := a.{{.Name}}Srv.Create(ctx, item)
	if err != nil {
		ginx.ResErrorCode(c, -1150, err)
		return
	}
	ginx.ResSuccess(c, result)
}

// Update 更新数据
func (a *{{.Name}}) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var item schema.{{.Name}}
	if err := ginx.ParseJSON(c, &item); err != nil {
		ginx.ResErrorCode(c, -1100, err)
		return
	}

	result, err := a.{{.Name}}Srv.Update(ctx, c.Param("id"), item)
	if err != nil {
		ginx.ResErrorCode(c, -1160, err)
		return
	}
	ginx.ResSuccess(c, result)
}

// Delete 删除数据
func (a *{{.Name}}) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.{{.Name}}Srv.Delete(ctx, c.Param("id"))
	if err != nil {
		ginx.ResErrorCode(c, -1170, err)
		return
	}
	ginx.ResOK(c, "成功删除数据")
}


// Enable 启用数据
func (a *{{.Name}}) Enable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.{{.Name}}Srv.UpdateActive(ctx, c.Param("id"), true)
	if err != nil {
		ginx.ResErrorCode(c, -1180, err)
		return
	}
	ginx.ResOK(c, "启用成功")
}

// Disable 禁用数据
func (a *{{.Name}}) Disable(c *gin.Context) {
	ctx := c.Request.Context()
	err := a.{{.Name}}Srv.UpdateActive(ctx, c.Param("id"), false)
	if err != nil {
		ginx.ResErrorCode(c, -1190, err)
		return
	}
	ginx.ResOK(c, "启用成功")
}

`
