package api

// Response is an Response
type Response struct {
	Status	 string 		  `json:"status"`
	Data		 interface{}  `json:"data"`
	Message	 string 			`json:"message"`
}
