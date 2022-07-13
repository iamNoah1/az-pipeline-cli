package internal

type Project struct {
	Name string `json:name`
}

type ProjectResponse struct {
	Value []Project `json:value`
}
