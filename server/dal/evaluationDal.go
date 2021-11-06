package dal

import "orderserve/model"

type EvaInterface interface {
	// 根据商品ID + 某字段查询 limit orderby time
	QueryEvaByKey(idName string, idVal int64,count int,start int) (evaData []model.Evaluation, err error)
}

type EvaDal struct {}

func NewEvaDal() EvaInterface {
	return &EvaDal{}
}

// 根据某个ID + 某字段 = value 查询 limit orderby time
func (ed *EvaDal) QueryEvaByKey(idName string, idVal int64,count int,start int) (evaData []model.Evaluation, err error) {
	err = DB.Where(idName + " = ?  and state = 1",idVal).Limit(count,start).OrderBy("create_time").Find(&evaData)
	return
}
