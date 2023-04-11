package benchjsonallocate

type StaticExtra struct {
	PoolId           string   `json:"poolId"`
	LpToken          string   `json:"lpToken"`
	Type             string   `json:"type"`
	Tokens           []string `json:"tokens"`
	DodoV1SellHelper string   `json:"dodoV1SellHelper"`
}
