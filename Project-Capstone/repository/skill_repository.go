package repository

import (
	"skillz_manager/model"
)

type ISkillRepository interface {
	CreateSkill(skill *model.Skill) error
	FindBySkill(skillName string) (model.Skill, error)
	AddEmployeeToSkill(skill model.Skill, employee model.Employee) (error, int)
	DeleteSkill(skill model.Skill) error
	DeleteEmployeeAssociation(skill model.Skill) error
}

type SkillRepository struct {
}

func (skillRepository SkillRepository) CreateSkill(skill *model.Skill) error {
	//implement code
	return nil
}
func (skillRepository SkillRepository) FindBySkill(skillName string) (model.Skill, error) {
	//implement code
	return model.Skill{}, nil
}

func (skillRepository SkillRepository) AddEmployeeToSkill(skill model.Skill, employee model.Employee) (error, int) {
	//implement code
	return nil, -1
}

func (skillRepository SkillRepository) DeleteSkill(skill model.Skill) error {
	//implement code
	return nil
}

func (skillRepository SkillRepository) DeleteEmployeeAssociation(skill model.Skill) error {
	//implement code
	return nil
}
