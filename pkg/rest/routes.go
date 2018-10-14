package rest

import (
	"encoding/json"
	"fmt"
	"github.com/schorzz/poppins-operator/config"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var(
	TIMEQUERYLAYOUT  = config.TIMEQUERYLAYOUT
)

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
type HttpHelper struct {
	//empty
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
	hh := HttpHelper{}

	poppins := controller.CreatePoppins("default")
	jsonresponse, err := json.Marshal(poppins)

	if err == nil{
		hh.Jsonresponse(w,"{}", 201)
	}


	hh.Jsonresponse(w, string(jsonresponse),200)
	//TODO: return created poppins
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

func DeleteAllExpiredPoppinses(w http.ResponseWriter, r *http.Request) {
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
func DeleteExpiredPoppins(w http.ResponseWriter, r *http.Request)  {
	httpHelper := HttpHelper{}
	rc := RestController{}

	poppinses, err := rc.GetPoppinses()
	if err != nil{
		logrus.Error(err)
	}

	since := httpHelper.getHeaderExpiredSince(r)
	expiredPoppinses := rc.FilterExpiredPoppinsList(poppinses, since)

	deleted := rc.DeleteDeployments(expiredPoppinses, nil)
	deleted = rc.DeletePods(expiredPoppinses, deleted)
	deleted = rc.DeletePoppinses(expiredPoppinses, deleted)


	jsonResponse, err := json.Marshal(deleted)
	if err != nil{
		logrus.Error(err)
	}

	httpHelper.Jsonresponse(w, string(jsonResponse), 200)
	//w.Header().Add("Content-Type", "application/json")
	//w.WriteHeader(200)
	//fmt.Fprintf(w, string(jsonResponse))

}
func (h *HttpHelper) getHeaderExpiredSince(r *http.Request)  time.Time{
	sinceQuery := r.FormValue("since")
	//sinceQuery = sinc

	since, err := time.Parse(TIMEQUERYLAYOUT, sinceQuery)
	if err != nil{
		logrus.Error(err)
		since = time.Now()
	}
	return since
}
func (h *HttpHelper) Jsonresponse(w http.ResponseWriter, body string, status int){
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	fmt.Fprintf(w, body)
}
