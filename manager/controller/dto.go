package controller

type mineral struct {
	ID        string `json:"id"`
	ClientID  string `json:"clientId"`
	Name      string `json:"name"`
	State     string `json:"state"`
	Fractures int    `json:"fractures"`
}
