package service

import (
	"skillz_manager/model"
	"skillz_manager/repository"
	"skillz_manager/utility"
	"skillz_manager/validator"
)

type ISkillService interface {
	CreateSkill(skill *model.Skill) error
	FindBySkill(skillName string) (model.Skill, error)
	AddEmployeeToSkill(skillName string, employee model.Employee) (error, int)
	DeleteEmployeeFromSkill(skillName string, name string) (error, int)
	DeleteSkill(skillName string) error
}

type SkillService struct {
	skillRepo    repository.ISkillRepository
	employeeRepo repository.IEmployeeRepository
	validator    validator.IValidator
}

func NewSkillService(skillRepo repository.ISkillRepository, employeeRepo repository.IEmployeeRepository, validator validator.IValidator) ISkillService {
	return SkillService{
		skillRepo:    skillRepo,
		employeeRepo: employeeRepo,
		validator:    validator,
	}
}

var (
	log             = utility.Log
	resourceManager = utility.ResourceManager{}
)

func (skillService SkillService) CreateSkill(skill *model.Skill) error {
	//implement code
	return nil
}

func (skillService SkillService) FindBySkill(skillName string) (model.Skill, error) {
	//implement code
	return model.Skill{}, nil
}

func (skillService SkillService) AddEmployeeToSkill(skillName string, employee model.Employee) (error, int) {
	//implement code
	return nil, -1
}

func (skillService SkillService) DeleteEmployeeFromSkill(skillName string, name string) (error, int) {
	//implement code
	return nil, -1
}

func (skillService SkillService) DeleteSkill(skillName string) error {
	//implement code
	return nil
}
