package types

type ToDo struct {
	Title    string `json:"title"`
	Note     string `json:"note"`
	Complete bool   `json:"complete"`
}
