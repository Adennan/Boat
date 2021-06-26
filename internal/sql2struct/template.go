package sql2struct

import (
	"Boat/internal/word"
	"fmt"
	"os"
	"text/template"
)

const structTpl = `type {{.TableName | ToCamelCase}} struct {
	{{range .Columns}}	{{ $length := len .Comment}} {{ if gt $length 0 }}// {{.Comment}} {{else}}// {{.Name}} {{ end }}
		{{ $typeLen := len .Type }} {{ if gt $typeLen 0 }}{{.Name | ToCamelCase}}	{{.Type}}	{{.Tag}}{{ else }}{{.Name}}{{ end }}
	{{end}}}
	func (model {{.TableName | ToCamelCase}}) TableName() string {
		return "{{.TableName}}"
	}`

type StructTemplate struct {
	structTpl string
}

type StructColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

type StructTemplateDB struct {
	TableName string
	Columns   []*StructColumn
}

func NewStructTeplate() *StructTemplate {
	return &StructTemplate{
		structTpl: structTpl,
	}
}

func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		tplColumns = append(tplColumns, &StructColumn{
			Name:    column.ColumnName,
			Type:    DBTypeToStructType[column.ColumnType],
			Tag:     fmt.Sprintf("`json:"+"%s"+"`", column.ColumnName),
			Comment: column.ColumnComment,
		})
	}
	return tplColumns
}

func (t *StructTemplate) Generate(tableName string, tplColumns []*StructColumn) error {
	tpl := template.Must(template.New("sql2struct").Funcs(template.FuncMap{
		"ToCamelCase": word.UnderlineToUpperCamel,
	}).Parse(t.structTpl))

	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns:   tplColumns,
	}
	if err := tpl.Execute(os.Stdout, tplDB); err != nil {
		return err
	}
	return nil
}
