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
	//fill with default values
}
//-------------------------------------------------------

type ListDTO struct {
	Type 	string 		`json:"type"`
	Data 	[]string 	`json:"data"`
}
