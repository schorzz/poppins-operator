package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ListNamespaceResponse struct {
	Namespaces []string `json:"namespaces"`
}

func GetAllNamespaces(w http.ResponseWriter, r *http.Request){
	controller := RestController{}
	response := ListNamespaceResponse{}

	list, err := controller.ListNamespaces()
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		http.Error(w, err.Error(), 400)
	}

	response.Namespaces = list
	jsonResponse, err := json.Marshal(response)

	if err != nil{
		panic(err)
	}
	fmt.Fprintf(w, string(jsonResponse))
}
