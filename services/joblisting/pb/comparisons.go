package pb

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

	if m1.CompanyId != convertedM2.CompanyId ||
		m1.HrContactId != convertedM2.HrContactId ||
		m1.HiringManagerId != convertedM2.HiringManagerId ||
		m1.JobPlatformId != convertedM2.JobPlatformId ||
		m1.Title != convertedM2.Title ||
		m1.Description != convertedM2.Description ||
		m1.SeniorityLevel != convertedM2.SeniorityLevel ||
		m1.YearsExperience != convertedM2.YearsExperience ||
		m1.EmploymentType != convertedM2.EmploymentType ||
		m1.FunctionId != convertedM2.FunctionId ||
		m1.IndustryId != convertedM2.IndustryId ||
		m1.Location != convertedM2.Location ||
		m1.Remote != convertedM2.Remote ||
		m1.SalaryCurrency != convertedM2.SalaryCurrency ||
		m1.MinSalary != convertedM2.MinSalary ||
		m1.MaxSalary != convertedM2.MaxSalary ||
		m1.CreatedAt.AsTime() != convertedM2.CreatedAt.AsTime() ||
		m1.UpdatedAt.AsTime() != convertedM2.UpdatedAt.AsTime() ||
		m1.StartAt.AsTime() != convertedM2.StartAt.AsTime() ||
		m1.ExpireAt.AsTime() != convertedM2.ExpireAt.AsTime() ||
		helpers.IsUint64SliceEqual(m1.SkillId, convertedM2.SkillId) {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two Company objects
func (m1 *JobCompany) IsEqual(m2 interface{}) bool {
	convertedM2 := m2.(*JobCompany)
	isNil, resolve := helpers.CheckNil(m1, m2)
	if resolve {
		return isNil
	}

	if m1.Name != convertedM2.Name ||
		m1.LogoUrl != convertedM2.LogoUrl ||
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

// IsEqual checks the equivalence of two KeyPerson objects
func (m1 *KeyPerson) IsEqual(m2 interface{}) bool {
	convertedM2 := m2.(*KeyPerson)
	isNil, resolve := helpers.CheckNil(m1, m2)
	if resolve {
		return isNil
	}

	if m1.CompanyId != convertedM2.CompanyId ||
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
		m1.BaseUrl != convertedM2.BaseUrl {
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

	if m1.CompanyId != convertedM2.CompanyId ||
		m1.IndustryId != convertedM2.IndustryId {
		return false
	}
	return true
}

// IsEqual checks the equivalence of two ProfileSkill objects
func (m1 *ProfileSkill) IsEqual(m2 interface{}) bool {
	convertedM2 := m2.(*ProfileSkill)
	isNil, resolve := helpers.CheckNil(m1, m2)
	if resolve {
		return isNil
	}

	if m1.Name != convertedM2.Name {
		return false
	}
	return true
}
