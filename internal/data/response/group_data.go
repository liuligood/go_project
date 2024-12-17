package response

type GetGetValueListResult struct {
	Data map[int64][]ValueData `json:"data"`
}

type ValueData struct {
	ID     int      `json:"id"`
	Sort   int      `json:"sort"`
	Fields []Fields `json:"fields"`
	Status bool     `json:"status"`
}

type Fields struct {
	Name  string `json:"name"`
	Title string `json:"title"`
	Value string `json:"value"`
}
