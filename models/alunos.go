package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	Nome string `json:"nome" validate:"nonzero"`
	Cpf  string `json:"cpf" validate:"len=11"`
	Rg   string `json:"rg" validate:"len=9"`
}

func ValidateAluno(a *Aluno) error {
	if err := validator.Validate(a); err != nil {
		return err
	}
	return nil
}
