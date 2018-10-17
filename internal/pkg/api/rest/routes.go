package rest

import (
	"encoding/json"
	"fmt"
	models "github.com/schorzz/poppins-operator/internal/pkg/api/rest/models"
	conf "github.com/schorzz/poppins-operator/internal/pkg/configurators"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)
//----------------------------------------------------
var(
	Config = conf.New(conf.Config{})
)
//----------------------------------------------------
type RestController struct {
	Writer http.ResponseWriter
	Reader *http.Request
}

func (rc *RestController) jsonResponse(body string, statusCode int){
	rc.Writer.Header().Add("Content-Type", "application/json")
	rc.Writer.WriteHeader(statusCode)
	fmt.Fprintf(rc.Writer, body)
}
//----------------------------------------------------
func GetAllNamespaces(w http.ResponseWriter, r *http.Request){
	callable := K8sController{}
	GetList(w, r, "namespaces", callable.ListNamespaces)
}
func GetAllPoppinsNamespaces(w http.ResponseWriter, r *http.Request)  {
	callable := K8sController{}
	GetList(w, r, "poppinses", callable.ListPoppinses)

}

func GetList(w http.ResponseWriter, r *http.Request, dataname string, f func()([]string, error))  {
	hh := HttpHelper{}
	response := models.ListDTO{}

	list, err := f()
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if list == nil{
		list = []string{}
	}
	response.Type = dataname
	response.Data = list
	jsonResponse, err := json.Marshal(response)

	if err != nil{
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	hh.Jsonresponse(w, string(jsonResponse), http.StatusOK)
}

func CreatePoppins(w http.ResponseWriter, r *http.Request) {
	controller := NewK8sController()
	hh := HttpHelper{}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Unmarshal
	var request CreatePoppinsRequest
	err = json.Unmarshal(b, &request)
	if err != nil {
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	poppins, k8serror := controller.CreatePoppins(request.Namespace, request.Name, request.ExpireDate)

	if k8serror != nil{
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonresponse, err := json.Marshal(poppins)

	if err != nil{
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusCreated)
		return
	}


	hh.Jsonresponse(w, string(jsonresponse),http.StatusCreated)
}

func UpdatePoppins(w http.ResponseWriter, r *http.Request) {
	controller := NewK8sController()
	hh := HttpHelper{}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Unmarshal
	var request CreatePoppinsRequest
	err = json.Unmarshal(b, &request)
	if err != nil {
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	poppins, k8serror := controller.UpdatePoppins(request.Namespace, request.Name, request.ExpireDate)

	if k8serror != nil{
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonresponse, err := json.Marshal(poppins)

	if err != nil{
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusCreated)
		return
	}


	hh.Jsonresponse(w, string(jsonresponse),http.StatusCreated)
}

func GetAllPoppinses(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	rc := NewK8sController()
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
	sinceQuery := r.FormValue("since")

	since, err := time.Parse(Config.HttpTimeQueryLayout, sinceQuery)
	if err != nil{
		logrus.Error(err)
		since = time.Now()
	}

	rc := NewK8sController()
	list, err := rc.GetPoppinses()

	if err != nil{
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	filteredList := []PoppinsListElementResponse{}
	filteredList = rc.FilterExpiredPoppinsList(list, since)

	jsonResponse, err := json.Marshal(filteredList)

	if err != nil{
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	hh := HttpHelper{}
	hh.Jsonresponse(w, string(jsonResponse), http.StatusOK)

}

func DeleteAllExpiredPoppinses(w http.ResponseWriter, r *http.Request) {
	hh := HttpHelper{}
	since:= hh.getQueryExpiredSince(r)

	rc := NewK8sController()
	list, err := rc.GetPoppinses()

	if err != nil{
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filteredList := []PoppinsListElementResponse{}
	filteredList = rc.FilterExpiredPoppinsList(list, since)

	jsonResponse, err := json.Marshal(filteredList)

	if err != nil{
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hh.Jsonresponse(w, string(jsonResponse),http.StatusOK)

}
func DeleteExpiredPoppins(w http.ResponseWriter, r *http.Request)  {
	httpHelper := HttpHelper{}
	rc := NewK8sController()

	poppinses, err := rc.GetPoppinses()
	if err != nil{
		logrus.Error(err)
	}

	since := httpHelper.getQueryExpiredSince(r)
	expiredPoppinses := rc.FilterExpiredPoppinsList(poppinses, since)

	deleted := rc.DeleteDeployments(expiredPoppinses, nil)
	deleted = rc.DeletePods(expiredPoppinses, deleted)
	deleted = rc.DeletePoppinses(expiredPoppinses, deleted)


	jsonResponse, err := json.Marshal(deleted)
	if err != nil{
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpHelper.Jsonresponse(w, string(jsonResponse), 200)

}
//TODO REFACTOR!!
func (h *HttpHelper) getQueryExpiredSince(r *http.Request)  time.Time{
	sinceQuery := r.FormValue("since")
	//sinceQuery = sinc

	since, err := time.Parse(Config.HttpTimeQueryLayout, sinceQuery)
	if err != nil{
		logrus.Error(err)
		since = time.Now()
	}
	return since
}

