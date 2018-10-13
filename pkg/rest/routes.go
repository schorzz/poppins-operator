package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ListNamespaceResponse struct {
	Namespaces []string `json:"namespaces"`
}

type ListPoppinsesResponse struct {
	Poppinses []string `json:"poppinses"`
}

type ListPodsAllNamespaces struct {
	Pods []string `json:"pods"`
}
type ListResponse struct {
	Type string `json:"type"`
	Data []string `json:"data"`
}

func GetAllNamespaces(w http.ResponseWriter, r *http.Request){
	callable := RestController{}
	GetList(w, r, "namespaces", callable.ListNamespaces)
	//w.Header().Add("Content-Type", "application/json")
	//controller := RestController{}
	//response := ListNamespaceResponse{}
	//
	//list, err := controller.ListNamespaces()
	//if err != nil {
	//	fmt.Fprintf(w, "%s", err)
	//	http.Error(w, err.Error(), 400)
	//	panic(err)
	//}
	//if list == nil{
	//	list = []string{}
	//}
	//response.Namespaces = list
	//jsonResponse, err := json.Marshal(response)
	//
	//if err != nil{
	//	panic(err)
	//}
	//fmt.Fprintf(w, string(jsonResponse))
}
func GetAllPoppinses(w http.ResponseWriter, r *http.Request)  {
	callable := RestController{}
	GetList(w, r, "poppinses", callable.ListPoppinses)
	//w.Header().Add("Content-Type", "application/json")
	//controller := RestController{}
	//response := ListPoppinsesResponse{}
	//
	//list, err := controller.ListPoppinses()
	//if err != nil {
	//	fmt.Fprintf(w, "%s", err)
	//	http.Error(w, err.Error(), 400)
	//	panic(err)
	//}
	//if list == nil{
	//	list = []string{}
	//}
	//response.Poppinses = list
	//jsonResponse, err := json.Marshal(response)
	//
	//if err != nil{
	//	panic(err)
	//}
	//fmt.Fprintf(w, string(jsonResponse))
}
func GetAllPodsAllNamespaces(w http.ResponseWriter, r *http.Request)  {
	callable := RestController{}//.ListPodsInAllNamespaces

	GetList(w, r, "pods", callable.ListPodsInAllNamespaces)
	//w.Header().Add("Content-Type", "application/json")
	//controller := RestController{}
	//response := ListPodsAllNamespaces{}
	//
	//list, err := controller.ListPodsInAllNamespaces()
	//if err != nil {
	//	fmt.Fprintf(w, "%s", err)
	//	http.Error(w, err.Error(), 400)
	//	panic(err)
	//}
	//if list == nil{
	//	list = []string{}
	//}
	//response.Pods = list
	//jsonResponse, err := json.Marshal(response)
	//
	//if err != nil{
	//	panic(err)
	//}
	//fmt.Fprintf(w, string(jsonResponse))
}

func GetList(w http.ResponseWriter, r *http.Request,dataname string, f func()([]string, error))  {
	w.Header().Add("Content-Type", "application/json")
	response := ListResponse{}

	list, err := f()
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		http.Error(w, err.Error(), 400)

		panic(err)
	}
	if list == nil{
		list = []string{}
	}
	response.Type = dataname
	response.Data = list
	jsonResponse, err := json.Marshal(response)

	if err != nil{
		panic(err)
	}
	fmt.Fprintf(w, string(jsonResponse))
}
func CreatePoppins(w http.ResponseWriter, r *http.Request) {
	controller := RestController{}

	controller.CreatePoppins()
}