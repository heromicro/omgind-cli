package generate

import (
	"context"
	"fmt"
	"strings"

	"github.com/heromicro/omgind-cli/helper"
)

func getAPIMockFileName(dir, name, apiv string) string {
	fullname := fmt.Sprintf("%s/internal/api/%s/mock/%s.mock.go", dir, apiv, helper.ToLowerUnderlinedNamer(name))
	return fullname
}

func genAPIMock(ctx context.Context, pkgName, dir, name, comment, apiv string) error {
	pname := helper.ToPlural(helper.ToLowerUnderlinedNamer(name))
	pname = strings.Replace(pname, "_", "-", -1)

	data := map[string]interface{}{
		"PkgName":    pkgName,
		"Name":       name,
		"Comment":    comment,
		"PluralName": strings.Replace(helper.ToPlural(pname), "-", "", -1),
		"Apiv":       apiv,
	}

	buf, err := execParseTpl(apiMockTpl, data)
	if err != nil {
		return err
	}

	fullname := getAPIMockFileName(dir, name, apiv)
	err = createFile(ctx, fullname, buf)
	if err != nil {
		return err
	}

	fmt.Printf("文件[%s]写入成功\n", fullname)

	return execGoFmt(fullname)
}

const apiMockTpl = `
package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// {{.Name}}Set 注入{{.Name}}
var {{.Name}}Set = wire.NewSet(wire.Struct(new({{.Name}}), "*"))

// {{.Name}} {{.Comment}}
type {{.Name}} struct {
}

// Query 查询数据
// @Tags {{.Comment}}
// @Summary 查询数据
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param current query int true "分页索引" default(1)
// @Param pageSize query int true "分页大小" default(10)
// @Param queryValue query string false "查询值"
// @Success 200 {array} schema.ListResult{list=schema.{{.PluralName | Title}},pagination=schema.PaginationResult}
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/{{.Apiv}}/{{.PluralName}} [get]
func (a *{{.Name}}) Query(c *gin.Context) {
}

// Get 查询指定数据
// @Tags {{.Comment}}
// @Summary 查询指定数据
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "唯一标识"
// @Success 200 {object} schema.{{.Name}}
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 404 {object} schema.ErrorResult "{error:{code:0,message:资源不存在}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/{{.Apiv}}/{{.PluralName}}/{id} [get]
func (a *{{.Name}}) Get(c *gin.Context) {
}

// View 查询指定数据
// @Tags {{.Comment}}
// @Summary 查询指定数据
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "唯一标识"
// @Success 200 {object} schema.{{.Name}}
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 404 {object} schema.ErrorResult "{error:{code:0,message:资源不存在}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/{{.Apiv}}/{{.PluralName}}/{id}/view [get]
func (a *{{.Name}}) View(c *gin.Context) {
}

// Create 创建数据
// @Tags {{.Comment}}
// @Summary 创建数据
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param body body schema.{{.Name}} true "创建数据"
// @Success 200 {object} schema.IDResult
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:无效的请求参数}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/{{.Apiv}}/{{.PluralName}} [post]
func (a *{{.Name}}) Create(c *gin.Context) {
}

// Update 更新数据
// @Tags {{.Comment}}
// @Summary 更新数据
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "唯一标识"
// @Param body body schema.{{.Name}} true "更新数据"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 400 {object} schema.ErrorResult "{error:{code:0,message:无效的请求参数}}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/{{.Apiv}}/{{.PluralName}}/{id} [put]
func (a *{{.Name}}) Update(c *gin.Context) {
}

// Delete 删除数据
// @Tags {{.Comment}}
// @Summary 删除数据
// @Security ApiKeyAuth
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "唯一标识"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/{{.Apiv}}/{{.PluralName}}/{id} [delete]
func (a *{{.Name}}) Delete(c *gin.Context) {
}

// Enable 启用数据
// @Tags {{.Comment}}
// @Summary 启用数据
// @Security ApiKeyAuth
// @Param id path string true "唯一标识"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/{{.Apiv}}/{{.PluralName}}/{id}/enable [patch]
func (a *{{.Name}}) Enable(c *gin.Context) {
}

// Disable 禁用数据
// @Tags {{.Comment}}
// @Summary 禁用数据
// @Security ApiKeyAuth
// @Param id path string true "唯一标识"
// @Success 200 {object} schema.StatusResult "{status:OK}"
// @Failure 401 {object} schema.ErrorResult "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.ErrorResult "{error:{code:0,message:服务器错误}}"
// @Router /api/{{.Apiv}}/{{.PluralName}}/{id}/disable [patch]
func (a *{{.Name}}) Disable(c *gin.Context) {
}

`
