package planner

import (
	"myparser/common"
)

type DeletePlan struct {
	Table string
	Where string
}

func (p *DeletePlan) Execute(ctx common.ExecutionContext) (common.Result, error) {
	print("excute Delete sql")
	// keyPrefix := p.Table + ":" //文件名_子文件名_id
	// iterator := ctx.DB.NewIterator()
	// iterator.Seek(keyPrefix)
	// var rowsAffected int
	// for iterator.Valid() {
	// 	key := iterator.Key()
	// 	if strings.HasPrefix(string(key), keyPrefix) {
	// 		err := ctx.DB.Delete(key)
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
