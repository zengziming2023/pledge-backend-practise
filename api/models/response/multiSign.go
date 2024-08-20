package response

type MultiSign struct {
	SpName           string   `json:"sp_name"`
	SpToken          string   `json:"_spToken"`
	JpName           string   `json:"jp_name"`
	JpToken          string   `json:"_jpToken"`
	SpAddress        string   `json:"sp_address"`
	JpAddress        string   `json:"jp_address"`
	SpHash           string   `json:"spHash"`
	JpHash           string   `json:"jpHash"`
	MultiSignAccount []string `json:"multi_sign_account"`
}

// DbMultiSign multi-sign signature
type DbMultiSign struct {
	Id               int32  `gorm:"column:id;primaryKey"`
	SpName           string `json:"sp_name" gorm:"column:sp_name"`
	ChainId          int    `json:"chain_id" gorm:"column:chain_id"`
	SpToken          string `json:"_spToken" gorm:"column:sp_token"`
	JpName           string `json:"jp_name" gorm:"column:jp_name"`
	JpToken          string `json:"_jpToken" gorm:"column:jp_token"`
	SpAddress        string `json:"sp_address" gorm:"column:sp_address"`
	JpAddress        string `json:"jp_address" gorm:"column:jp_address"`
	SpHash           string `json:"spHash" gorm:"column:sp_hash"`
	JpHash           string `json:"jpHash" gorm:"column:jp_hash"`
	MultiSignAccount string `json:"multi_sign_account" gorm:"column:multi_sign_account"`
}
