package experience

import (
	"company-profile-api/src/experience/skils"
	"company-profile-api/src/experience/tools"
	"time"
)

type Service interface {
	CreateExperienceServ(reqExperience CreateExperienceRequest) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateExperienceServ(reqExperience CreateExperienceRequest) (bool, error) {
	var keyExperience Experience

	keyExperience.AccountID = reqExperience.AccountID
	keyExperience.Position = reqExperience.Position
	keyExperience.EmploymentType = reqExperience.EmploymentType // Full-time, Part-time, Freelance, Contract, Internship
	keyExperience.CompanyName = reqExperience.CompanyName
	keyExperience.Location = reqExperience.Location
	startDate, err := time.Parse("2006-01-02", reqExperience.StartDate)
	if err != nil {
		return false, err
	}
	keyExperience.StartDate = startDate
	endTime, err := time.Parse("2006-01-02", reqExperience.EndDate)
	if err != nil {
		return false, err
	}
	keyExperience.EndDate = endTime
	keyExperience.Description = reqExperience.Description
	createExperience, err := s.repository.CreateExperienceRep(keyExperience)
	if err != nil {
		return false, err
	}

	// cara pertama
	// var keySkils skils.Skils
	// for _, keySkil := range reqExperience.Skils {
	// 	fmt.Println(keySkil)
	// 	keySkils.Title = keySkil
	// 	keySkils.ExperienceID = createExperience.ID
	// 	s.repository.CreateSkilsRep(keySkils)
	// }
	// cara kedua
	var keySkils skils.Skils
	var skilsArray []skils.Skils
	for _, keySkil := range reqExperience.Skils {
		keySkils.Title = keySkil
		keySkils.ExperienceID = createExperience.ID
		skilsArray = append(skilsArray, keySkils)
	}

	_, err = s.repository.CreateSkilsRep(skilsArray)
	if err != nil {
		return false, err
	}

	var keyTools tools.Tools
	var toolsArray []tools.Tools
	for _, keyTool := range reqExperience.Tools {
		keyTools.Title = keyTool
		keyTools.ExperienceID = createExperience.ID
		toolsArray = append(toolsArray, keyTools)
	}

	_, err = s.repository.CreateToolsRep(toolsArray)
	if err != nil {
		return false, err
	}

	return true, nil
}
