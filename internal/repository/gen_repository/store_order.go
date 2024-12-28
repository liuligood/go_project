package gen_repository

import (
	"crmeb_go/internal/model/model_data"
)

type StoreOrderQuerier interface {
	//	SELECT
	//	 DATE(FROM_UNIXTIME(created_at)) AS every_date,
	//   COUNT(id) AS id,
	//   SUM(pay_price) AS pay_price
	//  FROM
	//		store_order
	// 	{{where}}
	// 		{{if condition.Start !=0}}
	// 			store_order.created_at >= @condition.Start AND
	// 		{{end}}
	// 		{{if condition.End !=0}}
	// 			store_order.created_at <  @condition.End AND
	// 		{{end}}
	// 		store_order.deleted_at = 0
	// 	{{end}}
	// GROUP BY every_date
	// ORDER BY every_date
	QueryOrderGroupByDate(condition *model_data.DateCondition) ([]*model_data.StoreOrderDateResp, error)
}
