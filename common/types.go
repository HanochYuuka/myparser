package common

type Plan interface {
	Execute(ctx ExecutionContext) (Result, error)
}

type ExecutionContext struct {
	DB       DB
	Metadata Metadata
	// 其他可能需要的上下文信息
}

type Result struct {
	RowsAffected int
	Records      []Record
	// 其他可能需要的字段
}

type Record map[string]interface{}

type Metadata interface {
	GetTableSchema(table string) (TableSchema, error)
}

type TableSchema struct {
	Columns []ColumnInfo
	// 其他表结构信息
}

type ColumnInfo struct {
	Name string
	Type string
	// 其他列信息
}
