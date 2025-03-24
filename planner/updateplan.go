package planner

import (
	"myparser/common"
)

type UpdatePlan struct {
	Table string
	Set   map[string]interface{}
	Where string
}

func (p *UpdatePlan) Execute(ctx common.ExecutionContext) (common.Result, error) {
	print("excute update sql")
	// keyPrefix := p.Table + ":"
	// iterator := ctx.DB.NewIterator()
	// iterator.Seek(keyPrefix)
	// var rowsAffected int
	// for iterator.Valid() {
	// 	key := iterator.Key()
	// 	if strings.HasPrefix(string(key), keyPrefix) {
	// 		value := iterator.Value()
	// 		record, err := deserializeRecord(value)
	// 		if err != nil {
	// 			iterator.Close()
	// 			return common.Result{}, err
	// 		}
	// 		for k, v := range p.Set {
	// 			record[k] = v
	// 		}
	// 		updatedValue, err := serializeRecord(record)
	// 		if err != nil {
	// 			iterator.Close()
	// 			return common.Result{}, err
	// 		}
	// 		err = ctx.DB.Put(key, updatedValue)
	// 		if err != nil {
	// 			iterator.Close()
	// 			return common.Result{}, err
	// 		}
	// 		rowsAffected++
	// 	}
	// 	iterator.Next()
	// }
	// iterator.Close()
	rowsAffected := 1
	return common.Result{RowsAffected: rowsAffected}, nil
}
