package cv

type TechnicalSkillCategory int

const (
	Language TechnicalSkillCategory = iota
	Framework
	Tool
	Other
)

type TechnicalSkill struct {
	Name     string
	Category TechnicalSkillCategory
}

func (cv CV) GetTechnicalSkillsByCategory(category TechnicalSkillCategory) []string {
	var skills []string
	for _, skill := range cv.Skills {
		if skill.Category == category {
			skills = append(skills, skill.Name)
		}
	}
	return skills
}
