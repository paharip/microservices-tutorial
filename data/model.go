package data

type helloWorldResponse struct {
	Message string `json:"message"`
	Author  string `jason:"-"`
	Date    string `json:",omitempty"`
	Id      int    `json:"id,string"`
}

type helloWorldRequest struct {
	Name string `json:"name"`
}
