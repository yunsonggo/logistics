package backParam

import (
	"orderserve/model"
	"orderserve/model/backendModel"
)

type JwtBackendParam struct {
	ManagerToken string `json:"manager_token"`
	ManagerId string `json:"manager_id"`
	AddMenuData backendModel.BackendMenu `json:"add_menudata"`
	ManagerData backendModel.Manager `json:"manager_data"`
	SupplierData model.Supplier `json:"supplier_data"`
	CategoryData model.Cates `json:"category_data"`
	GiveData model.Give `json:"give_data"`
	SalePromoteData model.Sale `json:"promote_data"`
	OrderInfo model.Orders `json:"order_info"`
}
