package io

import "encoding/json"

//Todo struct 
type Todo struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Complete bool `json:"complete"`
}

func (t Todo) String() string{
	b, err := json.Marshal(t)
	if err != nil{
		return "unsupported value type"
	}
	return string(b)
}