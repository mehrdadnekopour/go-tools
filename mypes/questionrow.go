package mypes

// QuestionRow ...
type QuestionRow struct {
	ID         int          `json:"id"`
	Title      string       `json:"title"`
	Alias      string       `json:"alias"`
	Rows       QuestionRows `json:"rows"`
	Properties Properties   `json:"properties"`
	Attributes Properties   `json:"attribiutes"`
}

// QuestionRows ...
type QuestionRows []*QuestionRow

// Property ...
type Property struct {
	ID    int         `json:"id"`
	Title string      `json:"title"`
	Alias string      `json:"alias"`
	Value interface{} `json:"value"`
}

// Properties ...
type Properties []*Property
