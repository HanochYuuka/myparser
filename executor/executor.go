package executor

import (
	"myparser/common"
)

type Executor interface {
	Execute(plan common.Plan, ctx common.ExecutionContext) (common.Result, error)
}

type DefaultExecutor struct{}

func NewDefaultExecutor() *DefaultExecutor {
	return &DefaultExecutor{}
}

func (e *DefaultExecutor) Execute(plan common.Plan, ctx common.ExecutionContext) (common.Result, error) {
	return plan.Execute(ctx)
}
