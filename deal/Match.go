package deal

import (
	"GiveMeMail/src/global"
	"regexp"
)

// the func is used to match mail in the map
func MatchMail(ParseMap map[string]string) (match []string){
	//s := strings.Split(domain,".")
	if global.MatchRule=="all" {
		reg := regexp.MustCompile(`[A-Za-z0-9]+([_\.][A-Za-z0-9]+)*@([A-Za-z0-9\-]+\.)+[A-Za-z]{2,6}`)
		for _, content := range ParseMap {
			match = append(match, reg.FindAllString(content, -1)...)
		}
	}else if global.MatchRule=="strict"{
		reg := regexp.MustCompile(`[A-Za-z0-9]+([_\.][A-Za-z0-9]+)*@`+global.Domain)
		for _, content := range ParseMap {
			match = append(match, reg.FindAllString(content, -1)...)
		}
	}
	return match
}
