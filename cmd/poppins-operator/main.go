package main

import (
	"context"
	//"fmt"
	"github.com/gorilla/mux"
	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/operator-framework/operator-sdk/pkg/util/k8sutil"
	sdkVersion "github.com/operator-framework/operator-sdk/version"
	"github.com/schorzz/poppins-operator/pkg/rest"
	"github.com/schorzz/poppins-operator/pkg/stub"
	"github.com/sirupsen/logrus"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"net/http"
	"runtime"
	"time"
)

func printVersion() {
	logrus.Infof("Go Version: %s", runtime.Version())
	logrus.Infof("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH)
	logrus.Infof("operator-sdk Version: %v", sdkVersion.Version)
}

func main() {
	printVersion()

	sdk.ExposeMetricsPort()

	resource := "schorzz.poppins.com/v1alpha"
	kind := "Poppins"
	namespace, err := k8sutil.GetWatchNamespace()
	//namespace, err :=
	//err := nil
	//namespace = v1.NamespaceAll

	if err != nil {
		logrus.Fatalf("failed to get watch namespace: %v", err)
	}
	resyncPeriod := time.Duration(5) * time.Second
	logrus.Infof("Watching %s, %s, %s, %d", resource, kind, namespace, resyncPeriod)


	router := mux.NewRouter()//.StrictSlash(true)
	router.HandleFunc("/namespaces", rest.GetAllNamespaces).Methods("GET")
	router.HandleFunc("/namespaces/poppinses", rest.GetAllPoppinsNamespaces).Methods("GET")
	//router.HandleFunc("/namespaces/pods", rest.GetAllPodsNamespaces).Methods("GET")
	router.HandleFunc("/poppins", rest.CreatePoppins).Methods("POST")
	router.HandleFunc("/poppins", rest.UpdatePoppins).Methods("PUT")
	router.HandleFunc("/poppinses", rest.GetAllPoppinses).Methods("GET")//.Queries("expired_since","{expired_since}")//kk
	router.HandleFunc("/poppinses/expired", rest.GetAllExpiredPoppinses).Methods("GET")//.Queries("since", "{since}")
	router.HandleFunc("/poppinses/expired", rest.DeleteExpiredPoppins).Methods("DELETE")
	//router.Get("poppinses").Queries("expired_since","")
	go func() {
		http.ListenAndServe("0.0.0.0:8080", router)
	}()

	logrus.Infof("started webserver")

	//sdk.Watch(resource, kind, namespace, resyncPeriod)
	sdk.Handle(stub.NewHandler())
	sdk.Run(context.TODO())

}
