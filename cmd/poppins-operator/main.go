package main

import (
	//"fmt"
	"github.com/gorilla/mux"
	"github.com/operator-framework/operator-sdk/pkg/sdk"
	sdkVersion "github.com/operator-framework/operator-sdk/version"
	config2 "github.com/schorzz/poppins-operator/pkg/config"
	"github.com/schorzz/poppins-operator/pkg/rest"
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

func main() {
	printVersion()

	sdk.ExposeMetricsPort()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/namespaces", rest.GetAllNamespaces).Methods("GET")
	router.HandleFunc("/namespaces/poppinses", rest.GetAllPoppinsNamespaces).Methods("GET")
	router.HandleFunc("/poppins", rest.CreatePoppins).Methods("POST")
	router.HandleFunc("/poppins", rest.UpdatePoppins).Methods("PUT")
	router.HandleFunc("/poppinses", rest.GetAllPoppinses).Methods("GET")
	router.HandleFunc("/poppinses/expired", rest.GetAllExpiredPoppinses).Methods("GET")
	router.HandleFunc("/poppinses/expired", rest.DeleteExpiredPoppins).Methods("DELETE")

	config := config2.NewHTTPConfigurator()
	logrus.Infof("starting webserver on "+string(config.Listen))
	http.ListenAndServe(config.Listen, router)

}
