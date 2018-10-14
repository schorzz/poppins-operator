package rest

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const TIMEQUERYLAYOUT  = "2006-01-02"

type ListNamespaceResponse struct {
	Namespaces []string `json:"namespaces"`
}

type ListPoppinsesResponse struct {
	Poppinses []string `json:"poppinses"`
}

type ListPodsAllNamespaces struct {
	Pods 	[]string 	`json:"pods"`
}
type ListResponse struct {
	Type 	string 		`json:"type"`
	Data 	[]string 	`json:"data"`
}
type PoppinsListElementResponse struct {
	Name 		string 		`json:"name"`
	Namespace 	string 		`json:"namespace"`
	ExpireDate	time.Time 	`json:"expire_date,omitempty"`
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
func GetAllPoppinsNamespaces(w http.ResponseWriter, r *http.Request)  {
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
func GetAllPodsNamespaces(w http.ResponseWriter, r *http.Request)  {
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

func GetList(w http.ResponseWriter, r *http.Request, dataname string, f func()([]string, error))  {
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

	controller.CreatePoppins("default")
}
func GetAllPoppinses(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	rc := RestController{}
	list, err := rc.GetPoppinses()

	if err != nil{
		logrus.Error(err)
		http.Error(w, err.Error(), 400)
		return
	}
	jsonResponse, err := json.Marshal(list)

	if err != nil{
		logrus.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintf(w, string(jsonResponse))

}
func GetAllExpiredPoppinses(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	sinceQuery := r.FormValue("since")

	since, err := time.Parse(TIMEQUERYLAYOUT, sinceQuery)
	if err != nil{
		logrus.Error(err)
		since = time.Now()
	}

	rc := RestController{}
	list, err := rc.GetPoppinses()

	if err != nil{
		logrus.Error(err)
		http.Error(w, err.Error(), 400)
		return
	}

	filteredList := []PoppinsListElementResponse{}
	filteredList = rc.FilterExpiredPoppinsList(list, since)

	jsonResponse, err := json.Marshal(filteredList)

	if err != nil{
		logrus.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintf(w, string(jsonResponse))

}
