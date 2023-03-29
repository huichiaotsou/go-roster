package types

type Functions struct {
	// TeamID    int64    `json:"team_id"`
	FuncNames []string `json:"funcNames"`
}

type FuncIDs struct {
	FuncIDs []int64 `json:"funcIDs"`
}
