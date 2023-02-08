package generate

import (
	"context"
	"fmt"

	"github.com/heromicro/omgind-cli/helper"
)

func getServiceImplFileName(dir, name string) string {
	fullname := fmt.Sprintf("%s/internal/app/service/%s.srv.go", dir, helper.ToLowerUnderlinedNamer(name))
	return fullname
}

// 生成bll实现文件
func genServiceImpl(ctx context.Context, pkgName, dir, name, comment string) error {
	data := map[string]interface{}{
		"PkgName": pkgName,
		"Name":    name,
		"Comment": comment,
	}

	buf, err := execParseTpl(bllImplTpl, data)
	if err != nil {
		return err
	}

	fullname := getServiceImplFileName(dir, name)
	err = createFile(ctx, fullname, buf)
	if err != nil {
		return err
	}

	fmt.Printf("文件[%s]写入成功\n", fullname)

	return execGoFmt(fullname)
}

const bllImplTpl = `
package service

import (
	"context"

	"github.com/google/wire"

	"{{.PkgName}}/internal/schema/repo"
	"{{.PkgName}}/internal/app/schema"
	"{{.PkgName}}/pkg/errors"
)

// {{.Name}}Set 注入{{.Name}}
var {{.Name}}Set = wire.NewSet(wire.Struct(new({{.Name}}), "*"))

// {{.Name}} {{.Comment}}
type {{.Name}} struct {
	{{.Name}}Repo *repo.{{.Name}}
}

// Query 查询数据
func (a *{{.Name}}) Query(ctx context.Context, params schema.{{.Name}}QueryParam, opts ...schema.{{.Name}}QueryOptions) (*schema.{{.Name}}QueryResult, error) {
	return a.{{.Name}}Repo.Query(ctx, params, opts...)
}

// Get 查询指定数据
func (a *{{.Name}}) Get(ctx context.Context, id string, opts ...schema.{{.Name}}QueryOptions) (*schema.{{.Name}}, error) {
	item, err := a.{{.Name}}Repo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

// Create 创建数据
func (a *{{.Name}}) Create(ctx context.Context, item schema.{{.Name}}) (*schema.{{.Name}}, error) {
	// TODO: check?

	sch_{{.Name | ToLower}}, err := a.{{.Name}}Repo.Create(ctx, item)
	if err != nil {
		return nil, err
	}

	return sch_{{.Name | ToLower}}, nil
}

// Update 更新数据
func (a *{{.Name}}) Update(ctx context.Context, id string, item schema.{{.Name}}) (*schema.{{.Name}}, error) {
	
	oitem, err := a.{{.Name}}Repo.Get(ctx, id)
	if err != nil {
		return nil, err
	} else if oitem == nil {
		return nil, errors.ErrNotFound
	}

	item.ID = oitem.ID

	return a.{{.Name}}Repo.Update(ctx, id, item)
}

// Delete 删除数据
func (a *{{.Name}}) Delete(ctx context.Context, id string) error {
	oldItem, err := a.{{.Name}}Repo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.{{.Name}}Repo.Delete(ctx, id)
}

// UpdateActive 更新状态
func (a *{{.Name}}) UpdateActive(ctx context.Context, id string, active bool) error {
	oldItem, err := a.{{.Name}}Repo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.{{.Name}}Repo.UpdateActive(ctx, id, active)
}

`
