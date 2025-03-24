package planner

import (
	"myparser/common"
)

type SelectPlan struct {
	From    string
	Columns []string
	Where   string
}

func (p *SelectPlan) Execute(ctx common.ExecutionContext) (common.Result, error) {
	print("excute Select sql")
	// keyPrefix := p.From + ":" //根据具体的key前缀来设置
	// iterator := ctx.DB.NewIterator()
	// iterator.Seek(keyPrefix)
	// var records []common.Record
	// for iterator.Valid() {
	// 	key := iterator.Key()
	// 	if strings.HasPrefix(string(key), keyPrefix) {
	// 		value := iterator.Value()
	// 		record, err := deserializeRecord(value)
	// 		if err != nil {
	// 			iterator.Close()
	// 			return common.Result{}, err
	// 		}
	// 		records = append(records, record)
	// 	}
	// 	iterator.Next()
	// }
	// // 处理列选择
	// iterator.Close()
	var records []common.Record
	records = append(records, common.Record{
		"name": "zhangsan",
		"age":  18,
	})
	return common.Result{Records: records}, nil
}
