package rest

import (
	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/schorzz/poppins-operator/config"
	"github.com/schorzz/poppins-operator/pkg/apis/schorzz/v1alpha"
	"github.com/sirupsen/logrus"
	"k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

type RestController struct {
	ExcludedNamespaces	[]string
	SearchNameSpace string
}

func NewRestController() *RestController {
	controller := RestController{}
	controller.ExcludedNamespaces = []string{"default", "kube-public", "kube-system"}
	//controller.SearchNameSpace, _ = k8sutil.GetWatchNamespace()
	controller.SearchNameSpace = ""
	return &controller
}

type Deletable struct {
	Name 		string 	`json:"name"`
	Namespace 	string	`json:"namespace"`
	Kind 		string	`json:"kind"`
	APIVersion 	string	`json:"api_version"`
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
			APIVersion: "schorzz.poppins.com/v1alpha",

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
func (rc *RestController) CreatePoppins(namespace string, name string, expireDate time.Time) (*v1alpha.Poppins, error){
	//labels := map[string]string{
	//	"poppins": "code-",
	//}
	if expireDate.Before(time.Now()){
		expireDate = time.Now().UTC().Add(config.DEFAULTEXPIRETIME)
	}

	poppins := &v1alpha.Poppins{
		TypeMeta: metav1.TypeMeta{
			Kind: 		"Poppins",
			APIVersion:	"schorzz.poppins.com/v1alpha",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Namespace: namespace,
		},
		Spec:v1alpha.PoppinsSpec{
			ExpireDate: expireDate,
		},
	}
	err := sdk.Create(poppins)
	if err != nil {
		return nil, err
	}
	return poppins, nil
}

func (rc *RestController) UpdatePoppins(namespace string, name string, expireDate time.Time) (*v1alpha.Poppins, error){
	logrus.Infof("local from expiredate: %s",expireDate.Local())
	if expireDate.Before(time.Now()){
		expireDate = time.Now().UTC().Add(config.DEFAULTEXPIRETIME)
	}

	poppins := &v1alpha.Poppins{
		TypeMeta: metav1.TypeMeta{
			Kind: 		"Poppins",
			APIVersion:	"schorzz.poppins.com/v1alpha",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Namespace: namespace,
		},
	}
	err := sdk.Get(poppins)
	if err != nil {
		return nil, err
	}
	poppins.Spec.ExpireDate = expireDate
	err = sdk.Update(poppins)
	if err != nil {
		return nil, err
	}

	return poppins, nil
}

func (rc *RestController)GetPoppinses() ([]PoppinsListElementResponse, error){
	list := []PoppinsListElementResponse{}


	poppinsList := &v1alpha.PoppinsList{
		TypeMeta: metav1.TypeMeta{
			Kind: 		"Poppins",
			APIVersion: "schorzz.poppins.com/v1alpha",

		},

	}
	// change namespace to operator namespace or ""

	err := sdk.List(rc.SearchNameSpace, poppinsList)

	if err != nil {
		logrus.Error(err)
		return list, err
		//panic(err.Error())
	}
	for _, elem := range poppinsList.Items{

		poppins := PoppinsListElementResponse{}

		poppins.Namespace = elem.Namespace
		poppins.Name = elem.Name
		poppins.ExpireDate = elem.Spec.ExpireDate

		list = append(list, poppins)
	}
	return list, nil

}
func (rc *RestController) FilterExpiredPoppinsList(poppinslist []PoppinsListElementResponse, expireDate time.Time) []PoppinsListElementResponse {
	newList := []PoppinsListElementResponse{}

	for _, elem := range poppinslist{
		if elem.ExpireDate.Before(expireDate){
			newList = append(newList, elem)
			logrus.Infof("added %s to filtered list", elem)
		}
	}

	return newList
}
//func (rc RestController) DeleteFromPoppins(expiredPoppins []PoppinsListElementResponse, ) error{
//	for _, elem := range expiredPoppins{
//		elem
//		//delete everything here
//	}
//	return nil
//}
//func (rc RestController)DeleteDeploymentsFromNamespace(namespace string) {
//	deployments := v1beta1.DeploymentList{
//		TypeMeta: metav1.TypeMeta{
//			Kind:			"Deployment",
//			APIVersion: 	"apps/v1",
//		},
//	}
//
//	err := sdk.List(namespace, &deployments)
//
//	if err != nil{
//		logrus.Error(err)
//	}
//	for _, elem := range deployments.Items{
//		err = sdk.Delete(elem)
//	}
//}
func (rc *RestController) DeleteDeployments(expiredPoppins []PoppinsListElementResponse, deletables []Deletable) []Deletable{
	deployments := v1.DeploymentList{
		TypeMeta: metav1.TypeMeta{
			Kind:			"Deployment",
			APIVersion: 	"apps/v1",
		},
	}
	if deletables == nil {
		deletables = []Deletable{}
	}

	for _, elem := range expiredPoppins{
		err := sdk.List(elem.Namespace, &deployments)

		if err != nil{
			logrus.Error(err)
		}

		for _, elem := range deployments.Items{
			temp := Deletable{}
			temp.Name = elem.Name
			temp.Namespace = elem.Namespace
			temp.Kind = elem.Kind
			temp.APIVersion = elem.APIVersion
			logrus.Infof("Delete Deployment: {}", elem)
			deletables = append(deletables, temp)
			sdk.Delete(&elem)
		}
	}
	return deletables

}
func (rc *RestController) DeletePods(expiredPoppins []PoppinsListElementResponse, deletables []Deletable) []Deletable{
	pods := corev1.PodList{
		TypeMeta: metav1.TypeMeta{
			Kind: 		"Pod",
			APIVersion: "v1",
		},
	}
	if deletables == nil {
		deletables = []Deletable{}
	}

	for _, elem := range expiredPoppins{
		err := sdk.List(elem.Namespace, &pods)

		if err != nil{
			logrus.Error(err)
		}

		for _, elem := range pods.Items{
			temp := Deletable{}
			temp.Name = elem.Name
			temp.Namespace = elem.Namespace
			temp.Kind = elem.Kind
			temp.APIVersion = elem.APIVersion
			deletables = append(deletables, temp)
			logrus.Infof("Delete %s: %s", elem.Kind, elem)
			sdk.Delete(&elem)
		}
	}

	return deletables
}
func (rc *RestController) DeletePoppinses(expiredPoppins []PoppinsListElementResponse, deletables []Deletable) []Deletable{
	poppinses := v1alpha.PoppinsList{
		TypeMeta: metav1.TypeMeta{
			Kind: 		"Poppins",
			APIVersion: "schorzz.poppins.com/v1alpha",
		},
	}
	if deletables == nil {
		deletables = []Deletable{}
	}

	for _, elem := range expiredPoppins{
		err := sdk.List(elem.Namespace, &poppinses)

		if err != nil{
			logrus.Error(err)
		}

		for _, elem := range poppinses.Items{
			temp := Deletable{}
			temp.Name = elem.Name
			temp.Namespace = elem.Namespace
			temp.Kind = elem.Kind
			temp.APIVersion = elem.APIVersion
			deletables = append(deletables, temp)
			logrus.Infof("Delete %s: %s", elem.Kind, elem)
			sdk.Delete(&elem)
		}
	}

	return deletables
}



//func (rc RestController) delete(deletable Deletable){
//	switch deletable.Kind {
//	case "Deployment":
//		labels := map[string]string{
//			"name": deletable.Name,
//		}
//		deployment := v1.Deployment{
//			TypeMeta: 	metav1.TypeMeta{
//				Kind:			"Deployment",
//				APIVersion: 	deletable.APIVersion,
//			},
//			ObjectMeta: metav1.ObjectMeta{
//				Labels:			labels,
//			},
//		}
//		sdk.Delete(&deployment)
//		break
//
//	}
//}