package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, X interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		fmt.Println("utils:", body)
		if err := json.Unmarshal(body, X); err != nil {
			fmt.Println("error: ", err)
			fmt.Println(X)
			return
		}
	}
}
