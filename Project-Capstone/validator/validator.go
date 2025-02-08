package validator

import (
	"skillz_manager/model"
	"skillz_manager/utility"
)

var (
	log             = utility.Log
	resourceManager = utility.ResourceManager{}
)

type IValidator interface {
	ValidateSkill(skill model.Skill) error
	ValidateEmployee(employee model.Employee) error
	ValidateSkillName(skillName string) error
	ValidateSkillCategory(category string) error
	ValidateEmployeeLevel(level string) error
	ValidateEmployeeName(name string) error
	ValidateEmployeeAge(age int64) error
}

type Validator struct {
}

func (validator Validator) ValidateSkill(skill model.Skill) error {
	//implement code
	return nil
}

func (validator Validator) ValidateEmployee(employee model.Employee) error {
	//implement code
	return nil
}

func (validator Validator) ValidateSkillName(skillName string) error {
	//implement code
	return nil
}

func (validator Validator) ValidateSkillCategory(category string) error {
	//implement code
	return nil
}

func (validator Validator) ValidateEmployeeLevel(level string) error {
	//implement code
	return nil
}

func (validator Validator) ValidateEmployeeName(name string) error {
	//implement code
	return nil
}

func (validator Validator) ValidateEmployeeAge(age int64) error {
	//implement code
	return nil
}
