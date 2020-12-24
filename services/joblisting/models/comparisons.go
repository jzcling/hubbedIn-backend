package models

import "in-backend/helpers"

// IsEqual checks the equivalence of two JobPost objects
func (m1 *JobPost) IsEqual(m2 interface{}) bool {
	convertedM2 := m2.(*JobPost)
	isNil, resolve := helpers.CheckNil(m1, m2)
	if resolve {
		return isNil
	}

	if (m1.CreatedAt == nil) != (convertedM2.CreatedAt == nil) ||
		(m1.UpdatedAt == nil) != (convertedM2.UpdatedAt == nil) ||
		(m1.StartAt == nil) != (convertedM2.StartAt == nil) ||
		(m1.ExpireAt == nil) != (convertedM2.ExpireAt == nil) {
		return false
	}

	if m1.CompanyID != convertedM2.CompanyID ||
		m1.HRContactID != convertedM2.HRContactID ||
		m1.HiringManagerID != convertedM2.HiringManagerID ||
		m1.JobPlatformID != convertedM2.JobPlatformID ||
		m1.Title != convertedM2.Title ||
		m1.Description != convertedM2.Description ||
		m1.SeniorityLevel != convertedM2.SeniorityLevel ||
		m1.YearsExperience != convertedM2.YearsExperience ||
		m1.EmploymentType != convertedM2.EmploymentType ||
		m1.FunctionID != convertedM2.FunctionID ||
		m1.IndustryID != convertedM2.IndustryID ||
		m1.Location != convertedM2.Location ||
		m1.Remote != convertedM2.Remote ||
		m1.SalaryCurrency != convertedM2.SalaryCurrency ||
		m1.MinSalary != convertedM2.MinSalary ||
		m1.MaxSalary != convertedM2.MaxSalary ||
		*m1.CreatedAt != *convertedM2.CreatedAt ||
		*m1.UpdatedAt != *convertedM2.UpdatedAt ||
		*m1.StartAt != *convertedM2.StartAt ||
		*m1.ExpireAt != *convertedM2.ExpireAt ||
		helpers.IsUint64SliceEqual(m1.SkillID, convertedM2.SkillID) {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two Company objects
func (m1 *Company) IsEqual(m2 interface{}) bool {
	convertedM2 := m2.(*Company)
	isNil, resolve := helpers.CheckNil(m1, m2)
	if resolve {
		return isNil
	}

	if m1.Name != convertedM2.Name ||
		m1.LogoURL != convertedM2.LogoURL ||
		m1.Size != convertedM2.Size {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two Industry objects
func (m1 *Industry) IsEqual(m2 interface{}) bool {
	convertedM2 := m2.(*Industry)
	isNil, resolve := helpers.CheckNil(m1, m2)
	if resolve {
		return isNil
	}

	if m1.Name != convertedM2.Name {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two JobFunction objects
func (m1 *JobFunction) IsEqual(m2 interface{}) bool {
	convertedM2 := m2.(*JobFunction)
	isNil, resolve := helpers.CheckNil(m1, m2)
	if resolve {
		return isNil
	}

	if m1.Name != convertedM2.Name {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two KeyPerson objects
func (m1 *KeyPerson) IsEqual(m2 interface{}) bool {
	convertedM2 := m2.(*KeyPerson)
	isNil, resolve := helpers.CheckNil(m1, m2)
	if resolve {
		return isNil
	}

	if m1.CompanyID != convertedM2.CompanyID ||
		m1.Name != convertedM2.Name ||
		m1.ContactNumber != convertedM2.ContactNumber ||
		m1.Email != convertedM2.Email ||
		m1.JobTitle != convertedM2.JobTitle ||
		m1.UpdatedAt != convertedM2.UpdatedAt {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two JobPlatform objects
func (m1 *JobPlatform) IsEqual(m2 interface{}) bool {
	convertedM2 := m2.(*JobPlatform)
	isNil, resolve := helpers.CheckNil(m1, m2)
	if resolve {
		return isNil
	}

	if m1.Name != convertedM2.Name ||
		m1.BaseURL != convertedM2.BaseURL {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two CompanyIndustry objects
func (m1 *CompanyIndustry) IsEqual(m2 interface{}) bool {
	convertedM2 := m2.(*CompanyIndustry)
	isNil, resolve := helpers.CheckNil(m1, m2)
	if resolve {
		return isNil
	}

	if m1.CompanyID != convertedM2.CompanyID ||
		m1.IndustryID != convertedM2.IndustryID {
		return false
	}
	return true
}
