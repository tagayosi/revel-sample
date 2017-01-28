package forms

import (
	"regexp"
	"strconv"
	"strings"
	"github.com/revel/revel"
)

type SearchForm struct {
	Title			string //タイトル
	Runtime			string //放送時間
	OriginalAirDate	string //放送日
	GuestStaring	string //ゲスト出演
}

func (s SearchForm) Validate(v *revel.Validation, locale string) {

	if str := strings.Trim(s.Title, " "); str != " " {
		v.MinSize(str, 6).Message(revel.Message(locale,"search.form.validate.title.min", 6)).Key("s.Title")
	}
	
	if str := strings.Trim(s.Runtime, " "); str != " " {
		rt, err := strconv.Atoi(str)
		if err != nil {
			v.Error(revel.Message(locale, "search.form.validate.runtime.number")).Key("s.Runtime")
		} else {
			v.Range(rt, 50, 120).Message(revel.Message(locale, "search.form.validate.runtime.range")).Key("s.Runtime")
		}
	}
	
	if str := strings.Trim(s.OriginalAirDate, " "); str != " " {
		v.Match(str, regexp.MustCompile("^(January|February|March|April|May|June|July|August|September|October|November|December).*$")).Message(revel.Message(locale, "search.form.validate.originalairdate.match")).Key("s.OriginalAirDate")
	}

}
