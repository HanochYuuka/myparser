package planner

import (
	"myparser/ast"
	"myparser/common"
)

type Planner struct{}

func NewPlanner() *Planner {
	return &Planner{}
}

func (p *Planner) GeneratePlan(astNode ast.IAST) common.Plan {
	if astNode == nil {
		return nil
	}
	switch node := astNode.(type) {
	case *ast.ASTSelectQuery:
		return p.generateSelectPlan(node)
	case *ast.ASTInsertQuery:
		return p.generateInsertPlan(node)
	case *ast.ASTUpdateQuery:
		return p.generateUpdatePlan(node)
	case *ast.ASTDeleteQuery:
		return p.generateDeletePlan(node)
	default:
		panic("Unexpected type")
	}
}

func (p *Planner) generateSelectPlan(query *ast.ASTSelectQuery) *SelectPlan {
	return &SelectPlan{
		From:    query.From,
		Columns: query.Columns,
		Where:   query.Where,
	}
}

func (p *Planner) generateInsertPlan(query *ast.ASTInsertQuery) *InsertPlan {
	return &InsertPlan{
		Table:  query.Table,
		Values: query.Values,
	}
}

func (p *Planner) generateUpdatePlan(query *ast.ASTUpdateQuery) *UpdatePlan {
	setMap := make(map[string]interface{})
	for k, v := range query.Set {
		setMap[k] = v
	}
	return &UpdatePlan{
		Table: query.Table,
		Set:   setMap,
		Where: query.Where,
	}
}

func (p *Planner) generateDeletePlan(query *ast.ASTDeleteQuery) *DeletePlan {
	return &DeletePlan{
		Table: query.Table,
		Where: query.Where,
	}
}
