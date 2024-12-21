package request

type DateParams struct {
	BaseServiceParams BaseServiceParams
	Start             int64 // 今天0点间戳
	End               int64 // 今天结束时间戳
	YesterdayStart    int64 // 咋天0点时间戳
	YesterdayEnd      int64 // 咋天24点时间戳
}
