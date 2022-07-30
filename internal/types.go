package internal

type Project struct {
	Name string `json:name`
}

type ProjectResponse struct {
	Value []Project `json:value`
}

type Pipeline struct {
	Name string `json:name`
}

type PipelineResponse struct {
	Value []Pipeline `json:value`
}
