package generate

import (
	"bytes"
	"context"
	"fmt"

	"github.com/heromicro/omgind-cli/helper"
)

type schemaField struct {
	Name           string // 字段名
	Comment        string // 字段注释
	Type           string // 字段类型
	IsRequired     bool   // 是否必选项
	BindingOptions string // binding配置项(不包含required，required由IsRequired控制)
}

func getSchemaFileName(dir, name string) string {
	fullname := fmt.Sprintf("%s/internal/app/schema/%s.sch.go", dir, helper.ToLowerUnderlinedNamer(name))
	return fullname
}

// 生成schema文件
func genSchema(ctx context.Context, pkgName, dir, name, comment string, fields ...schemaField) error {
	var tfields []schemaField

	tfields = append(tfields, schemaField{Name: "ID", Comment: "唯一标识", Type: "string"})
	tfields = append(tfields, fields...)
	tfields = append(tfields, schemaField{Name: "CreatedAt", Comment: "创建时间", Type: "*time.Time"})
	tfields = append(tfields, schemaField{Name: "UpdatedAt", Comment: "更新时间", Type: "*time.Time"})

	buf := new(bytes.Buffer)
	for _, field := range tfields {
		switch field.Type {
		case "time":
			buf.WriteString(fmt.Sprintf("%s \t %s.Time \t", field.Name, field.Type))

		default:
			buf.WriteString(fmt.Sprintf("%s \t %s \t", field.Name, field.Type))
		}

		buf.WriteByte('`')
		buf.WriteString(fmt.Sprintf(`json:"%s"`, helper.ToLowerUnderlinedNamer(field.Name)))

		bindingOpts := ""
		if field.IsRequired {
			bindingOpts = "required"
		}
		if v := field.BindingOptions; v != "" {
			if bindingOpts != "" {
				bindingOpts += ","
			}
			bindingOpts = bindingOpts + v
		}
		if bindingOpts != "" {
			buf.WriteByte(' ')
			buf.WriteString(fmt.Sprintf(`binding:"%s"`, bindingOpts))
		}

		buf.WriteByte('`')

		if field.Comment != "" {
			buf.WriteString(fmt.Sprintf("// %s", field.Comment))
		}
		buf.WriteString(delimiter)
	}

	tbuf, err := execParseTpl(schemaTpl, map[string]interface{}{
		"PkgName":    pkgName,
		"Name":       name,
		"PluralName": helper.ToPlural(name),
		"Fields":     buf.String(),
		"Comment":    comment,
	})
	if err != nil {
		return err
	}

	fullname := getSchemaFileName(dir, name)
	err = createFile(ctx, fullname, tbuf)
	if err != nil {
		return err
	}

	fmt.Printf("文件[%s]写入成功\n", fullname)

	return execGoFmt(fullname)
}

const schemaTpl = `
package schema

import (
	"time"

	"{{.PkgName}}/pkg/helper/json"
)

// {{.Name}} {{.Comment}}对象
type {{.Name}} struct {
	{{.Fields}}
}


func (a *{{.Name}}) String() string {
	return json.MarshalToString(a)
}

// {{.Name}}QueryParam 查询条件
type {{.Name}}QueryParam struct {
	PaginationParam
	q string ` + "`" + `form:"q"` + "`" + ` // 模糊查询

	BasicOrderParam
	TimeOrderParam
}

// {{.Name}}QueryOptions 查询可选参数项
type {{.Name}}QueryOptions struct {
	OrderFields []*OrderField // 排序字段

	FieldsAll      bool  		 // all fields
	FieldsIncludes []string      // includes fields
	FieldsExcludes []string      // exlcudes fields
}

// {{.Name}}QueryResult 查询结果
type {{.Name}}QueryResult struct {
	Data       {{.PluralName}}
	PageResult *PaginationResult
}

// {{.PluralName}} {{.Comment}}列表
type {{.PluralName}} []*{{.Name}}

`
