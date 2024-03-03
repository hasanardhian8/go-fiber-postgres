package models

import "time"

type Bayars struct {
	Id         int       `json:"id" gorm:"primary_key"`
	Idpenduduk int       `json:"idPenduduk" gorm:"Foreignkey:PenduduksRefer; AssociationForeignKey:Id"`
	Idpetugas  int       `json:"idPetugas" gorm:"Foreignkey:PetugasesRefer; AssociationForeignKey:Id"`
	Nominal    int       `json:"nominal"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
