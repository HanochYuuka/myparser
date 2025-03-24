package planner

import (
	"myparser/common"
)

type InsertPlan struct {
	Table  string
	Values []string
}

func (p *InsertPlan) Execute(ctx common.ExecutionContext) (common.Result, error) {
	print("excute Insert sql")

	// schema, err := ctx.Metadata.GetTableSchema(p.Table)
	// if err != nil {
	// 	return common.Result{}, err
	// }
	// columns := schema.Columns
	// if len(columns) != len(p.Values) {
	// 	return common.Result{}, fmt.Errorf("number of columns mismatch")
	// }
	// record := make(common.Record)
	// for i, col := range columns {
	// 	record[col.Name] = p.Values[i]
	// }
	// key := p.Table + ":" + record[schema.PrimaryKey] //step文件名
	// data, err := serializeRecord(record)             //转化为kv键值对中的value
	// if err != nil {
	// 	return common.Result{}, err
	// }
	// err = ctx.DB.Put(key, data)
	// if err != nil {
	// 	return common.Result{}, err
	// }
	return common.Result{RowsAffected: 1}, nil
}
