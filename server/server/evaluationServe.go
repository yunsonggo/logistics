package server

import (
	"orderserve/dal"
	"orderserve/model"
)

type IEvaluation interface {
	// 根据商品ID + 某字段查询 limit orderby time
	GetEvaByKey(idName string, idVal int64,count int,start int) (evaData []model.Evaluation, err error)
}

type EvaServe struct {}

func NewEvaServe() IEvaluation {
	return &EvaServe{}
}
var evaDal = dal.NewEvaDal()
// 根据商品ID + 某字段查询 limit orderby time
func (es *EvaServe) GetEvaByKey(idName string, idVal int64,count int,start int) (evaData []model.Evaluation, err error) {
	return evaDal.QueryEvaByKey(idName,idVal,count,start)
}
