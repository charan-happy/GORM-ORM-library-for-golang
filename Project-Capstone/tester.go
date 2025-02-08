package main

import (
	"skillz_manager/repository"
	"skillz_manager/service"
	"skillz_manager/utility"
	"skillz_manager/validator"
)

func main() {
	utility.DbConnect()
	// createSkill()
	// findBySkillName()
	// addEmployeeToSkill()
	// deleteEmployeeFromSkill()
	// deleteSkill()
}

var (
	resourceManager = utility.ResourceManager{}
	skillService    = service.NewSkillService(repository.SkillRepository{}, repository.EmployeeRepository{}, validator.Validator{})
)

//Uncomment the methods after implementing all the files

// func createSkill() {

// 	skill := model.Skill{Skill: "Golang", Category: "Backend", Employee: []model.Employee{{Name: "Willam Jess", Age: 25, Level: "BEGINNER"}, {Name: "Jone Caryy", Age: 28, Level: "ADVANCED"}}}

// 	error := skillService.CreateSkill(&skill)
// 	if error != nil {
// 		msg, _ := resourceManager.GetProperty("Tester.ERROR")
// 		fmt.Println(msg)
// 		fmt.Println(error.Error())
// 	} else {
// 		message, _ := resourceManager.GetProperty("Tester.CREATE_SKILL_SUCCESS")
// 		fmt.Println(skill.Skill, message)
// 	}
// }

// func findBySkillName() {
// 	skillName := "Java"
// 	skill, err := skillService.FindBySkill(skillName)
// 	if err != nil {
// 		msg, _ := resourceManager.GetProperty("Tester.ERROR")
// 		fmt.Println(msg)
// 		fmt.Println(err.Error())
// 	} else {
// 		fmt.Println("SKILL DETAILS")
// 		fmt.Println("Skill Name: ", skill.Skill)
// 		if len(skill.Employee) != 0 {
// 			fmt.Println("Employees skilled in  ", skill.Skill)

// 			for _, employee := range skill.Employee {
// 				fmt.Println("Employee Name: ", employee.Name, " Age: ", employee.Age, "Skill: ", employee.Level)
// 			}
// 		}

// 	}
// }
// func addEmployeeToSkill() {
// 	skillName := "Java"
// 	employee := model.Employee{Name: "Joseph K", Age: 26, Level: "BEGINNER"}
// 	err := skillService.AddEmployeeToSkill(skillName, employee)
// 	if err != nil {
// 		msg, _ := resourceManager.GetProperty("Tester.ERROR")
// 		fmt.Println(msg)
// 		fmt.Println(err.Error())
// 	} else {
// 		message, _ := resourceManager.GetProperty("Tester.ADD_EMPLOYEE")
// 		fmt.Println(employee.Name, message, skillName)
// 	}
// }

// func deleteEmployeeFromSkill() {
// 	skillName := "Java"
// 	name := "Eunice Reneau"
// 	err := skillService.DeleteEmployeeFromSkill(skillName, name)
// 	if err != nil {
// 		msg, _ := resourceManager.GetProperty("Tester.ERROR")
// 		fmt.Println(msg)
// 		fmt.Println(err.Error())
// 	} else {
// 		message, _ := resourceManager.GetProperty("Tester.DELETE_EMPLOYEE")
// 		fmt.Println(name, message, skillName)
// 	}
// }

// func deleteSkill() {
// 	skillName := "Java"
// 	err := skillService.DeleteSkill(skillName)
// 	if err != nil {
// 		msg, _ := resourceManager.GetProperty("Tester.ERROR")
// 		fmt.Println(msg)
// 		fmt.Println(err.Error())
// 	} else {
// 		message, _ := resourceManager.GetProperty("Tester.DELETE_SKILL")
// 		fmt.Println(skillName, message)
// 	}
// }
