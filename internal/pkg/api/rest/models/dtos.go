package models

import (
	"github.com/schorzz/poppins-operator/pkg/apis/schorzz/v1alpha"
)

type PoppinsDTO struct {
	Name 		string				`json:"name"`
	Namespace 	string 				`json:"namespace"`
	ExpireDates v1alpha.PoppinsSpec	`json:"expire_dates"`
}

func (dto *PoppinsDTO) initialze() {

}
