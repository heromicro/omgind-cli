package generate

// TplItem 模板项
type TplItem struct {
	StructName string         `yaml:"name"`    // 结构体名称
	Comment    string         `yaml:"comment"` // 注释
	Mixin      string         `yaml:"mixin"`   // memo, sort, time, status, active
	Fields     []TplFieldItem `yaml:"fields"`  // 字段项
}

func (t TplItem) toSchemaFields() []schemaField {
	var items []schemaField
	for _, f := range t.Fields {
		items = append(items, schemaField{
			Name:           f.StructFieldName,
			Comment:        f.Comment,
			Type:           f.StructFieldType,
			IsRequired:     f.StructFieldRequired,
			BindingOptions: f.BindingOptions,
		})
	}
	return items
}

func (t TplItem) toEntityEntFields() []entityEntField {
	var items []entityEntField
	for _, f := range t.Fields {
		items = append(items, entityEntField{
			Name:    f.StructFieldName,
			Storage: f.StorageKey,
			Comment: f.Comment,
			Type:    f.StructFieldType,
			Max:     f.Max,
			Min:     f.Min,
			Default: f.Default,
			Index:   f.Index,
		})
	}
	return items
}

func (t TplItem) toEntityGormFields() []entityGormField {
	var items []entityGormField
	for _, f := range t.Fields {
		items = append(items, entityGormField{
			Name:        f.StructFieldName,
			Comment:     f.Comment,
			Type:        f.StructFieldType,
			GormOptions: f.GormOptions,
		})
	}
	return items
}

func (t TplItem) toEntityMongoFields() []entityMongoField {
	var items []entityMongoField
	for _, f := range t.Fields {
		items = append(items, entityMongoField{
			Name:    f.StructFieldName,
			Comment: f.Comment,
			Type:    f.StructFieldType,
		})
	}
	return items
}

// TplFieldItem 模板字段项
type TplFieldItem struct {
	StructFieldName     string `yaml:"name"`            // 结构体字段名称
	StorageKey          string `yaml:"storage"`         // 数据库存储字段名称
	StructFieldRequired bool   `yaml:"required"`        // 结构字段必选项
	Comment             string `yaml:"comment"`         // 注释
	StructFieldType     string `yaml:"type"`            // 结构体字段类型
	GormOptions         string `yaml:"gorm_options"`    // gorm配置项
	BindingOptions      string `yaml:"binding_options"` // binding配置项
	Max                 int    `yaml:"max"`             // 字符串最长, 数字类型最大
	Min                 int    `yaml:"min"`             // 字符串最短, 数字类型最小
	Default             string `yaml:"default"`         // 默认值
	Index               string `yaml:"index"`           // 数据库索引
}
