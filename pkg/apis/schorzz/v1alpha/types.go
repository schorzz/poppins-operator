package v1alpha

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type PoppinsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Poppins `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Poppins struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              PoppinsSpec   `json:"spec"`
}

type PoppinsSpec struct {
	NamespaceExpDate 	time.Time `json:"namespace_expire_date"`
	PodExpDate			time.Time `json:"pod_expire_date"`
	DeploymentExpDate 	time.Time `json:"deployment_expire_date"`
	StatefullSetExpDate	time.Time `json:"statefull_set_exp_date"`
	ReplicaSetExpDate	time.Time `json:"replica_set_exp_date"`
	IngressExpDate		time.Time `json:"ingress_exp_date"`
}
