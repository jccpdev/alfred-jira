package jira

type Response struct {
	Key string `json:"key"`
	Fields     `json:"fields"`
}

type Fields struct {
	Assignee Assignee `json:"assignee"`
	Reporter Reporter `json:"reporter"`
	Summary  string   `json:"summary"`
	Status   Status   `json:"status"`
}

type Assignee struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

type Reporter struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

type Status struct {
	Name string `json:"name"`
}
