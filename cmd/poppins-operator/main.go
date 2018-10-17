package main

import (
	//"fmt"
	"github.com/gorilla/mux"
	"github.com/operator-framework/operator-sdk/pkg/sdk"
	sdkVersion "github.com/operator-framework/operator-sdk/version"
	"github.com/schorzz/poppins-operator/internal/pkg/configurators"
	"github.com/schorzz/poppins-operator/internal/pkg/api/rest"
	"github.com/sirupsen/logrus"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"net/http"
	"runtime"
)

func printVersion() {
	logrus.Infof("Go Version: %s", runtime.Version())
	logrus.Infof("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH)
	logrus.Infof("operator-sdk Version: %v", sdkVersion.Version)
}

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter().StrictSlash(true)
	a.Router.HandleFunc("/namespaces", rest.GetAllNamespaces).Methods("GET")
	a.Router.HandleFunc("/namespaces/poppinses", rest.GetAllPoppinsNamespaces).Methods("GET")
	a.Router.HandleFunc("/poppins", rest.CreatePoppins).Methods("POST")
	a.Router.HandleFunc("/poppins", rest.UpdatePoppins).Methods("PUT")
	a.Router.HandleFunc("/poppinses", rest.GetAllPoppinses).Methods("GET")
	a.Router.HandleFunc("/poppinses/expired", rest.GetAllExpiredPoppinses).Methods("GET")
	a.Router.HandleFunc("/poppinses/expired", rest.DeleteExpiredPoppins).Methods("DELETE")
}
func (a *App) Run() {

	config := configurators.NewHTTPConfigurator()
	logrus.Infof("starting webserver on "+string(config.Listen))
	http.ListenAndServe(config.Listen, a.Router)
}


func main() {
	printVersion()

	sdk.ExposeMetricsPort()
	a := App{}
	a.Initialize()
	a.Run()

}
