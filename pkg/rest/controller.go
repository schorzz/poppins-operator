package rest

import (
	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/operator-framework/operator-sdk/pkg/util/k8sutil"
	"github.com/schorzz/poppins-operator/pkg/apis/schorzz/v1alpha"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

)

type RestController struct {
	ExcludedNamespaces	[]string
	SearchNameSpace string
}

func NewRestController() *RestController {
	controller := RestController{}
	controller.ExcludedNamespaces = []string{"default", "kube-public", "kube-system"}
	controller.SearchNameSpace, _ = k8sutil.GetWatchNamespace()
	//controller.SearchNameSpace = ""
	return &controller
}

func (rc *RestController) ListPodsInAllNamespaces() ([]string, error){

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
	err := sdk.List(rc.SearchNameSpace, namespaceList)

	if err != nil {
		return nil, err
		//panic(err.Error())
	}

	for _, elem := range namespaceList.Items{
		if ListContains(rc.ExcludedNamespaces, elem.Name) != true{
			list = append(list, elem.Name)
		}
	}
	return list, nil
}
func (rc *RestController) ListPoppinses() ([]string, error){
	var list []string;


	poppinsList := &v1alpha.PoppinsList{
		TypeMeta: metav1.TypeMeta{
			Kind: 		"Poppins",
			APIVersion: "v1alpha",
		},
	}

	// change namespace to operator namespace or ""

	err := sdk.List(rc.SearchNameSpace, poppinsList)

	if err != nil {
		return nil, err
		//panic(err.Error())
	}

	for _, elem := range poppinsList.Items{
		if !ListContains(rc.ExcludedNamespaces, elem.Name){
			list = append(list, elem.Name)
		}
	}
	return list, nil
}
func (rc RestController) CreatePoppins() {
	//labels := map[string]string{
	//	"poppins": "code-",
	//}
	poppins := &v1alpha.Poppins{
		TypeMeta: metav1.TypeMeta{
			Kind: 		"Poppins",
			APIVersion:	"v1alpha",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "code-created",
		},
	}
	err := sdk.Create(poppins)
	if err != nil {
		panic(err)
	}
}