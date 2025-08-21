package cv

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

type TechnicalSkillCategory int

const (
	Language TechnicalSkillCategory = iota
	Framework
	Tool
	Other
)

type CV struct {
	Title          string
	Desc           string
	Lang           string
	AboutMe        Info
	WorkExperience []WorkExp
	Publications   []Publication
	Projects       []Project
	Education      []Edu
	Skills         []TechnicalSkill
	Languages      []SpokenLanguage
	References     []Reference
}

type Info struct {
	Selfie   string // path to selfie (optional)
	Name     string // full name
	Address  string
	Phone    string
	Email    string
	Website  string // personal website
	LinkedIn string // linkedin username
	Github   string // github username
	Summary  string // short summary about yourself
}

type WorkExp struct {
	Title       string // job title
	Company     string // name of the company
	Start       string // date you started
	End         string // date you quit
	Location    string // for example: city where company was located
	Description string // describe shortly what you did at [company name here]
}

type Publication struct {
	Title   string
	Link    string // diva link for example
	Year    int
	Summary string // short summary similar to abstract
}

type Project struct {
	Name         string // project name not your name
	Start        string // date when you began working on the project. for example: Aug 2025
	End          string
	Description  string
	Technologies []string // for example a list of programming languages
}

type Edu struct {
	Name   string
	Start  string
	End    string
	School string
	Desc   string
}

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

type SpokenLanguage struct {
	Name    string
	Summary string // how good you are at the language
}

type Reference struct {
	Name    string
	Company string
	Phone   string
	Email   string
}

func ParseCV(CVPath string) CV {
	var cv CV
	_, err := toml.DecodeFile(CVPath, &cv)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cv.AboutMe.Name)
	return cv
}
