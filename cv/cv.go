package cv

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
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

func ParseCV(CVPath string) CV {
	var cv CV
	_, err := toml.DecodeFile(CVPath, &cv)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cv.AboutMe.Name)
	return cv
}
