package dal

import "orderserve/model"

type GoodsInterface interface {
	InsertGoodOne(goodInfo *model.Goods) (err error)
	DelGoodById(id int64) (err error)
	UpdateGoodById(id int64,goodInfo *model.Goods) (err error)
	QueryGoodById(id int64) (goodData *model.Goods,err error)
	QueryGoodsAll() (goodsData []model.Goods,err error)
	QueryGoodsByAny(key string,data interface{}) (goodsData []model.Goods,err error)
	QueryGoodsByKeywords(key string,data string) (goodsData []model.Goods,err error)
	//根据某字段 升 / 降序 排列
	QueryGoodsByKeyOrder(key string,order string,count int,start int) (goodsData []model.Goods,err error)
	QueryByArrayAndPrice(giveIds []interface{},saleIds []interface{},minPrice float32,maxPrice float32,count int,start int) (goodsData []model.Goods,err error)
}

type GoodsDal struct {}

func NewGoodsDal() GoodsInterface {
	return &GoodsDal{}
}

// 添加 商品
func (gd *GoodsDal) InsertGoodOne(goodInfo *model.Goods) (err error) {
	_,err = DB.InsertOne(goodInfo)
	return
}

// 根据ID 删除商品
func (gd *GoodsDal) DelGoodById(id int64) (err error) {
	good := new(model.Goods)
	_,err = DB.ID(id).Delete(good)
	return
}
// 更新商品
func (gd *GoodsDal) UpdateGoodById(id int64,goodInfo *model.Goods) (err error) {
	_,err = DB.ID(id).Update(goodInfo)
	return
}
// 根据ID 查商品
func (gd *GoodsDal) QueryGoodById(id int64) (goodData *model.Goods,err error) {
	var good model.Goods
	_,err = DB.ID(id).Get(&good) 
	goodData = &good
	return
}
// 查询所有
func (gd *GoodsDal) QueryGoodsAll() (goodsData []model.Goods,err error) {
	err = DB.Find(&goodsData)
	return
}
// 根据字段 查询
func (gd *GoodsDal) QueryGoodsByAny(key string,data interface{}) (goodsData []model.Goods,err error) {
	err = DB.Where(key + " = ? and state = 1",data).OrderBy("sell_count desc").Find(&goodsData)
	return
}
// 根据关键词 模糊查询
func (gd *GoodsDal) QueryGoodsByKeywords(key string,data string) (goodsData []model.Goods,err error) {
	err = DB.Where(key + " like ? and state = 1","%"+data+"%").Find(&goodsData)
	return
}

// 根据某字段 升 / 降排序
func (gd *GoodsDal) QueryGoodsByKeyOrder(key string,order string,count int ,start int) (goodsData []model.Goods,err error) {
	err = DB.Where("state = 1").Limit(count,start).OrderBy(key + " " + order).Find(&goodsData)
	return
}

func (gd *GoodsDal) QueryByArrayAndPrice(giveIds []interface{},saleIds []interface{},minPrice float32,maxPrice float32,count int,start int) (goodsData []model.Goods,err error) {
	err = DB.Cols("id", "shop_id","give_id","sale_id","price","state").In("give_id",giveIds).In("sale_id",saleIds).Or("price >= ? and price <= ? and state = 1",minPrice,maxPrice).Limit(count,start).Find(&goodsData)
	//err = DB.Cols("id", "shop_id","give_id","sale_id","price","state").In("sale_id",saleIds).In("give_id",giveIds).In("sale_id",saleIds).Find(&goodsData)
	return
}
