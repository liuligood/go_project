package gen_repository

import (
	"crmeb_go/internal/model"
	"crmeb_go/internal/model/model_data"
)

type QueryOrderGroupByDate interface {
	//	SELECT
	//	 DATE(FROM_UNIXTIME(created_at)) AS orderId,
	//   COUNT(id) AS id,
	//   SUM(pay_price) AS pay_price
	//  FROM
	//		store_order
	// 	{{where}}
	// 		{{if condition.Start !=0}}
	// 			store_order.created_at >= @condition.Start
	// 		{{end}}
	// 		{{if condition.End !=0}}
	// 			store_order.created_at <  @condition.End AND
	// 		{{end}}
	// 		store_order.deleted_at = 0
	// 	{{end}}
	// GROUP BY orderId
	// ORDER BY orderId
	QueryOrderGroupByDate(condition *model_data.DateCondition) ([]*model.StoreOrder, error)
}
