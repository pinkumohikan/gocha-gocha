package api

type GreetRequest struct {
	Name string `json:"name"`
}

type GreetResponse struct {
	Greeting string `json:"greeting"`
}
