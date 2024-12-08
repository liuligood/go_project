package service_data

type GetRealNameParams struct {
	BaseServiceParams
	UserId uint64
}

type Condition struct {
	UserID   string
	Nickname string
	Email    string
}
