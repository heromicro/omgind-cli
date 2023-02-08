package generate

import (
	"context"
	"fmt"
	"strings"

	"github.com/heromicro/omgind-cli/helper"
)

func getRepoImplEntFileName(dir, name string) string {
	fullname := fmt.Sprintf("%s/internal/schema/repo/%s.repo.go", dir, helper.ToLowerUnderlinedNamer(name))
	return fullname
}

// 生成model实现文件
func genRepoImplEnt(ctx context.Context, pkgName, dir, name, comment string) error {
	data := map[string]interface{}{
		"PkgName":    pkgName,
		"Name":       name,
		"PluralName": helper.ToPlural(name),
		"Comment":    comment,
		"EntPackage": strings.ToLower(name),
	}

	buf, err := execParseTpl(modelImplEntTpl, data)
	if err != nil {
		return err
	}

	fullname := getRepoImplEntFileName(dir, name)
	err = createFile(ctx, fullname, buf)
	if err != nil {
		return err
	}

	fmt.Printf("文件[%s]写入成功\n", fullname)

	return execGoFmt(fullname)
}

const modelImplEntTpl = `
package repo

import (
	"context"
	"time"

	"{{.PkgName}}/internal/app/schema"
	"{{.PkgName}}/internal/gen/ent"
	"{{.PkgName}}/internal/gen/ent/{{.EntPackage}}"
	"{{.PkgName}}/pkg/errors"
	"{{.PkgName}}/pkg/helper/structure"

	"github.com/google/wire"
)


// {{.Name}}Set 注入{{.Name}}
var {{.Name}}Set = wire.NewSet(wire.Struct(new({{.Name}}), "*"))

// {{.Name}} {{.Comment}}存储
type {{.Name}} struct {
	EntCli *ent.Client
}


// ToSchema{{.Name}} 转换为
func (a *{{.Name}}) toSchema{{.Name}}(et *ent.{{.Name}}) *schema.{{.Name}} {
	item := new(schema.{{.Name}})
	structure.Copy(et, item)
	return item
}

func (a *{{.Name}}) toSchema{{.PluralName}}(ets ent.{{.PluralName}}) []*schema.{{.Name}} {
	list := make([]*schema.{{.Name}}, len(ets))
	for i, item := range ets {
		list[i] = a.toSchema{{.Name}}(item)
	}
	return list
}

func (a *{{.Name}}) ToEntCreate{{.Name}}Input(sch *schema.{{.Name}}) *ent.Create{{.Name}}Input {
	createinput := new(ent.Create{{.Name}}Input)
	structure.Copy(sch, &createinput)

	return createinput
}

func (a *{{.Name}}) ToEntUpdate{{.Name}}Input(sch *schema.{{.Name}}) *ent.Update{{.Name}}Input {
	updateinput := new(ent.Update{{.Name}}Input)
	structure.Copy(sch, &updateinput)

	return updateinput
}

func (a *{{.Name}}) getQueryOption(opts ...schema.{{.Name}}QueryOptions) schema.{{.Name}}QueryOptions {
	var opt schema.{{.Name}}QueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *{{.Name}}) Query(ctx context.Context, params schema.{{.Name}}QueryParam, opts ...schema.{{.Name}}QueryOptions) (*schema.{{.Name}}QueryResult, error) {
	opt := a.getQueryOption(opts...)

	query := a.EntCli.{{.Name}}.Query()

	query = query.Where({{.EntPackage}}.DeletedAtIsNil())
	// TODO: 查询条件


	count, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// get total
	pr := &schema.PaginationResult{Total: count}
	if params.PaginationParam.OnlyCount {
		return &schema.{{.Name}}QueryResult{PageResult: pr}, nil
	}

	opt.OrderFields = append(opt.OrderFields, schema.NewOrderField("id", schema.OrderByDESC))
	query = query.Order(ParseOrder(opt.OrderFields)...)

	pr.Current = params.PaginationParam.GetCurrent()
	pr.PageSize = params.PaginationParam.GetPageSize()
	if params.Offset() > count {
		return &schema.{{.Name}}QueryResult{PageResult: pr}, nil
	}
	query = query.Limit(params.Limit()).Offset(params.Offset())

	list, err1 := query.All(ctx)
	if err1 != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.{{.Name}}QueryResult{
		PageResult: pr,
		Data:       a.toSchema{{.PluralName}}(list),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *{{.Name}}) Get(ctx context.Context, id string, opts ...schema.{{.Name}}QueryOptions) (*schema.{{.Name}}, error) {
	
	r_{{.Name | ToLower}}, err := a.EntCli.{{.Name}}.Query().Where({{.EntPackage}}.IDEQ(id)).Only(ctx)
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	return a.toSchema{{.Name}}(r_{{.Name | ToLower}}), nil
}

// Create 创建数据
func (a *{{.Name}}) Create(ctx context.Context, item schema.{{.Name}}) (*schema.{{.Name}}, error) {
	
	iteminput := a.ToEntCreate{{.Name}}Input(&item)
	r_{{.Name | ToLower}}, err := a.EntCli.{{.Name}}.Create().SetInput(*iteminput).Save(ctx)

	if err != nil {
		return nil, err
	}
	sch_{{.Name | ToLower}} := a.toSchema{{.Name}}(r_{{.Name | ToLower}})
	return sch_{{.Name | ToLower}}, nil
}

// Update 更新数据
func (a *{{.Name}}) Update(ctx context.Context, id string, item schema.{{.Name}}) (*schema.{{.Name}}, error) {
	
	oitem, err := a.EntCli.{{.Name}}.Query().Where({{.Name | ToLower}}.IDEQ(id)).Only(ctx)
	if err != nil {
		return nil, err
	}

	iteminput := a.ToEntUpdate{{.Name}}Input(&item)

	r_{{.Name | ToLower}}, err := oitem.Update().SetInput(*iteminput).Save(ctx)
	sch_{{.Name | ToLower}} := a.toSchema{{.Name}}(r_{{.Name | ToLower}})

	return sch_{{.Name | ToLower}}, nil
}

// Delete 删除数据
func (a *{{.Name}}) Delete(ctx context.Context, id string) error {

	r_{{.Name | ToLower}}, err := a.EntCli.{{.Name}}.Query().Where({{.Name | ToLower}}.IDEQ(id)).Only(ctx)

	if err != nil {
		return err
	}

	_, err1 := r_{{.Name | ToLower}}.Update().SetDeletedAt(time.Now()).Save(ctx)

	return errors.WithStack(err1)
}

// UpdateActive 更新状态
func (a *{{.Name}}) UpdateActive(ctx context.Context, id string, active bool) error {

	_, err1 := a.EntCli.{{.Name}}.Update().Where({{.Name | ToLower}}.IDEQ(id)).SetIsActive(active).Save(ctx)

	return errors.WithStack(err1)
}

`
