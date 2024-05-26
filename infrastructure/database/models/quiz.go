package models

type Quiz struct {
	Id        string     `json:"id"`
	Questions []Question `json:"questions"`
}

func (q Quiz) Key() string {
	return q.Id
}
