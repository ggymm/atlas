package resp

const (
	timeFormat = "2006-01-02 15:04:05"
)

type TreeNode[T any] struct {
	Key      int64  `json:"key"`
	Label    string `json:"label"`
	Children []*T   `json:"children"`
}

type Option struct {
	Value int64  `json:"value"`
	Label string `json:"label"`
}

type PageList struct {
	List  any   `json:"list"`
	Total int64 `json:"total"`
}
