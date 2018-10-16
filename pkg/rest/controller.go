package rest

import (
	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/schorzz/poppins-operator/pkg/config"
	"github.com/schorzz/poppins-operator/pkg/apis/schorzz/v1alpha"
	"github.com/sirupsen/logrus"
	"k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

var (
	EXPIRY_TIME = config.NewPoppinsConfigurator().Expiretime
)

type RestController struct {
	ExcludedNamespaces	[]string
	SearchNameSpace string
}

func NewRestController() *RestController {
	controller := RestController{}
	controller.ExcludedNamespaces = []string{"default", "kube-public", "kube-system"}	//controller.SearchNameSpace, _ = k8sutil.GetWatchNamespace()
	controller.SearchNameSpace = ""
	return &controller
}

type Deletable struct {
	Name 		string 	`json:"name"`
	Namespace 	string	`json:"namespace"`
	Kind 		string	`json:"kind"`
	APIVersion 	string	`json:"api_version"`
}
type RessourceListI interface {
	getPodList() *corev1.PodList
	getNamespaceList() *corev1.NamespaceList
	getPoppinsList() *v1alpha.PoppinsList
}
type RessourceList struct {
	PodList 	corev1.PodList
	APIVersion 	string
}

func (rl RessourceList) getPodList() *corev1.PodList{
	list := &corev1.PodList{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
	}
	return list
}
func (rl RessourceList) getNamespaceList() *corev1.NamespaceList{
	list := &corev1.NamespaceList{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Namespace",
			APIVersion: "v1",
		},
	}
	return list
}
func (rl RessourceList) getPoppinsList() *v1alpha.PoppinsList{
	list := &v1alpha.PoppinsList{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Poppins",
			APIVersion: "schorzz.poppins.com/v1alpha",
		},
	}
	return list
}
func (rc *RestController) ListNamespaces() ([]string, error){
	var list []string;
	rl := RessourceList{}
	namespaceList := rl.getNamespaceList()
	err := sdk.List(rc.SearchNameSpace, namespaceList)

	if err != nil {
		return nil, err
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
	rl := RessourceList{}
	poppinsList := rl.getPoppinsList()

	err := sdk.List(rc.SearchNameSpace, poppinsList)

	if err != nil {
		return nil, err
	}

	for _, elem := range poppinsList.Items{
		if !ListContains(rc.ExcludedNamespaces, elem.Name){
			list = append(list, elem.Name)
		}
	}
	return list, nil
}
func (rc *RestController) CreatePoppins(namespace string, name string, expireDate time.Time) (*v1alpha.Poppins, error){
	if expireDate.Before(time.Now()){
		expireDate = time.Now().UTC().Add(EXPIRY_TIME)
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
		expireDate = time.Now().UTC().Add(EXPIRY_TIME)
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
	rl := RessourceList{}
	poppinsList := rl.getPoppinsList()

	err := sdk.List(rc.SearchNameSpace, poppinsList)

	if err != nil {
		logrus.Error(err)
		return list, err
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