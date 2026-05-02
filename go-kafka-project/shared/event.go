package shared

type Event struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Value string `json:"value"`
	Time  string `json:"time"`
}