package rest

import (
	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type RestController struct {
	ExcludedNamespaces	[]string
}

func NewRestController() *RestController {
	controller := RestController{}
	controller.ExcludedNamespaces = []string{"default", "kube-public", "kube-system"}
	//return &RestController{}
	return &controller
}

func (rc *RestController)  ListPodsInAllNamespaces() ([]string, error){

	var list []string;

	podList := &corev1.PodList{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
	}

	// change namespace to operator namespace or ""
	_ = sdk.List("", podList)
	for _, elem := range podList.Items{
		list = append(list, elem.Name)
	}
	return list, nil
}
func (rc *RestController)  ListNamespaces() ([]string, error){

	var list []string;

	namespaceList := &corev1.NamespaceList{
		TypeMeta: metav1.TypeMeta{
			Kind: "Namespace",
			APIVersion: "v1",
		},
	}


	// change namespace to operator namespace or ""
	err := sdk.List("", namespaceList)

	if err != nil {
		panic(err.Error())
	}

	logrus.Infof("got namespaces: %s", namespaceList)
	for _, elem := range namespaceList.Items{
		if !ListContains(rc.ExcludedNamespaces, elem.Name){
			list = append(list, elem.Name)
		}
	}
	return list, nil
}