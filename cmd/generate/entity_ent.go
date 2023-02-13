package generate

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/heromicro/omgind-cli/helper"
)

type entityEntField struct {
	Name     string // 字段名
	Storage  string // 数据库存储字段名
	Comment  string // 字段注释
	Type     string // 字段类型
	Required bool   // 是否必须
	Max      int    // 最大
	Min      int    // 最小
	Default  string // 默认值
	Index    string // 索引
}

type fieldParts struct {
	ftype string
}

func getEntityEntFileName(dir, name string) string {
	fullname := fmt.Sprintf("%s/internal/schema/entity/%s.entity.go", dir, helper.ToLowerUnderlinedNamer(name))
	return fullname
}

// 生成entity文件
func genEntEntity(ctx context.Context, pkgName, dir, name, comment, mixin_parts string, fields ...entityEntField) error {
	var tfields []entityEntField

	mixins := strings.Split(mixin_parts, ",")

	mixin_items := make([]string, len(mixin_parts))

	tfields = append(tfields, fields...)
	tfields = append(tfields, entityEntField{Name: "Creator", Storage: "creator", Comment: "创建者", Type: "string"})

	field_buf := new(bytes.Buffer)
	indexes := make([]string, 0)

	for _, field := range tfields {

		// log.Printf(" --- --- field --- %+v ", field)
		// log.Printf(" --- --- field --- %d ", all_mixins.Size())
		// log.Printf(" --- --- field --- %t ", all_mixins.Empty())

		if all_mixins.Contains(strings.ToLower(field.Name)) {
			mixins = append(mixins, strings.ToLower(field.Name))
			continue
		}

		switch field.Type {
		case "string":
			field_buf.WriteString(fmt.Sprintf("field.String(\"%s\")", field.Storage))
		case "text":
			field_buf.WriteString(fmt.Sprintf("field.Text(\"%s\")", field.Storage))
		case "int":
			field_buf.WriteString(fmt.Sprintf("field.Int(\"%s\")", field.Storage))
		case "int64":
			field_buf.WriteString(fmt.Sprintf("field.Int64(\"%s\")", field.Storage))
		case "int32":
			field_buf.WriteString(fmt.Sprintf("field.Int32(\"%s\")", field.Storage))
		case "int16":
			field_buf.WriteString(fmt.Sprintf("field.Int16(\"%s\")", field.Storage))
		case "int8":
			field_buf.WriteString(fmt.Sprintf("field.Int8(\"%s\")", field.Storage))

		case "uint":
			field_buf.WriteString(fmt.Sprintf("field.Uint(\"%s\")", field.Storage))
		case "uint64":
			field_buf.WriteString(fmt.Sprintf("field.Uint64(\"%s\")", field.Storage))
		case "uint32":
			field_buf.WriteString(fmt.Sprintf("field.Uint32(\"%s\")", field.Storage))
		case "uint16":
			field_buf.WriteString(fmt.Sprintf("field.Uint16(\"%s\")", field.Storage))
		case "uint8":
			field_buf.WriteString(fmt.Sprintf("field.Uint8(\"%s\")", field.Storage))
		case "float":
			field_buf.WriteString(fmt.Sprintf("field.Float(\"%s\")", field.Storage))
		case "float32":
			field_buf.WriteString(fmt.Sprintf("field.Float32(\"%s\")", field.Storage))
		case "bool":
			field_buf.WriteString(fmt.Sprintf("field.Bool(\"%s\")", field.Storage))
		}

		switch field.Type {
		case "string":
			fallthrough
		case "text":
			if field.Max != 0 {
				field_buf.WriteString(fmt.Sprintf(".MaxLen(%d)", field.Max))
			}
			if field.Min != 0 {
				field_buf.WriteString(fmt.Sprintf(".MinLen(%d)", field.Min))
			}
			if field.Default != "" {
				field_buf.WriteString(fmt.Sprintf(".Default(\"%s\")", field.Default))
			}

		case "int":
			fallthrough
		case "int64":
			fallthrough
		case "int32":
			fallthrough
		case "int16":
			fallthrough
		case "int8":
			fallthrough
		case "uint":
			fallthrough
		case "uint64":
			fallthrough
		case "uint8":
			fallthrough
		case "uint16":
			fallthrough
		case "uint32":
			if field.Max != 0 {
				field_buf.WriteString(fmt.Sprintf(".Max(%d)", field.Max))
			}
			if field.Min != 0 {
				field_buf.WriteString(fmt.Sprintf(".Min(%d)", field.Min))
			}
			if field.Default != "" {
				field_buf.WriteString(fmt.Sprintf(".Default(%s)", field.Default))
			}

		case "float":
			fallthrough
		case "float32":
			if field.Max != 0 {
				field_buf.WriteString(fmt.Sprintf(".Max(%d)", field.Max))
			}
			if field.Min != 0 {
				field_buf.WriteString(fmt.Sprintf(".Min(%d)", field.Min))
			}
			if field.Default != "" {
				field_buf.WriteString(fmt.Sprintf(".Default(%s)", field.Default))
			}

		case "bool":
			if field.Default != "" {
				field_buf.WriteString(fmt.Sprintf(".Default(%s)", field.Default))
			}

		}

		if field.Storage != "" {
			field_buf.WriteString(fmt.Sprintf(".StorageKey(\"%s\")", field.Storage))
		}

		if field.Comment != "" {
			field_buf.WriteString(fmt.Sprintf(".Comment(\"%s\")", field.Comment))
		}

		idx := strings.ToLower(strings.TrimSpace(field.Index))

		if idx != "" && idx != "false" && idx != "no" {
			index_buf := new(bytes.Buffer)

			index_buf.WriteString(fmt.Sprintf("index.Fields(\"%s\")", field.Storage))
			if idx == "u" || idx == "unique" {
				index_buf.WriteString(".Unique(),\n")
			} else if idx != "" {
				index_buf.WriteString(",\n")
			}
			indexes = append(indexes, index_buf.String())
		}

		field_buf.WriteString(",")

		field_buf.WriteString(delimiter)
	}

	for i, item := range mixins {
		mixin_items[i] = fmt.Sprintf("mixin.%sMixin{},", strings.Title(strings.TrimSpace(item)))
	}

	tbuf, err := execParseTpl(entityEntTpl, map[string]interface{}{
		"PkgName":       pkgName,
		"Name":          name,
		"PluralName":    helper.ToPlural(name),
		"Fields":        field_buf.String(),
		"Indexes":       indexes,
		"Mixins":        mixin_items,
		"Comment":       comment,
		"UnderLineName": helper.ToLowerUnderlinedNamer(name),
		"BackQuote":     "`",
	})
	if err != nil {
		return err
	}

	filename := getEntityEntFileName(dir, name)
	err = createFile(ctx, filename, tbuf)
	if err != nil {
		return err
	}

	fmt.Printf("文件[%s]写入成功\n", filename)

	return execGoFmt(filename)
}

const entityEntTpl = `
package entity

import (
	"{{.PkgName}}/internal/schema/mixin"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	{{ $length := len .Indexes }}
	{{- if gt $length 0 }}
	// {{$length}}
	"entgo.io/ent/schema/index" 
	{{- end -}}
)

type {{.Name}} struct {
	ent.Schema
}

func ({{.Name}}) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
		{{- range .Mixins}}	
		{{- if . }}
			{{. }}
		{{- end }}
		{{- end }}
	}
}

func ({{.Name}}) Fields() []ent.Field {

	return []ent.Field{
		{{.Fields}}
	}
}

func ({{.Name}}) Edges() []ent.Edge {
	return []ent.Edge{}
}

func ({{.Name}}) Indexes() []ent.Index {
	return []ent.Index{
		{{- range .Indexes}}
		{{- if .}}
		{{. -}}
		{{- end }}
		{{- end }}
	}
}

func ({{.Name}}) Hooks() []ent.Hook {
	return []ent.Hook{
	}
}
`
