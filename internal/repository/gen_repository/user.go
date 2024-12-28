package gen_repository

import (
	"crmeb_go/internal/model/model_data"
	"gorm.io/gen"
)

type UserQuerier interface {
	//	SELECT id,user_id,nickname,email FROM users
	// 		{{where}}
	// 			{{if condition.Nickname !=""}}
	// 				nickname = @condition.Nickname AND
	// 			{{end}}
	// 			{{if condition.UserID !=""}}
	// 				user_id = @condition.UserID AND
	// 			{{end}}
	// 			{{if condition.Email !=""}}
	// 				email = @condition.Email
	// 			{{end}}
	// 		{{end}}
	GetUserByCondition(condition model_data.Condition) (*gen.T, error)

	//	SELECT
	//	 DATE(FROM_UNIXTIME(created_at)) AS every_date,
	//   COUNT(uid) AS uid
	//  FROM
	//		user
	// 	{{where}}
	// 		{{if condition.Start !=0}}
	// 			user.created_at >= @condition.Start AND
	// 		{{end}}
	// 		{{if condition.End !=0}}
	// 			user.created_at <  @condition.End AND
	// 		{{end}}
	// 		user.deleted_at = 0
	// 	{{end}}
	// GROUP BY every_date
	// ORDER BY every_date ASC
	GetAddUserCountGroupDate(condition *model_data.DateCondition) ([]*model_data.UserDateResp, error)
}
