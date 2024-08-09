package response

type Search struct {
	Count int64  `json:"count"`
	Rows  []Pool `json:"rows"`
}
