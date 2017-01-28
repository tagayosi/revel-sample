package models

import (
	"fmt"
	"time"
)

type Episodes struct {
	Title				string		`json:"title"`
	OriginalAirDate		string		`json:"original_air_date"`
	Runtime				int			`json:"runtime"`
	GuestStaring		string		`json:"guest_staring"`
	GuestStaringRole	string		`json:"guest_staring_role"`
	DirectedBy			string		`json:"directed_by"`
	WrittenBy			[]string	`json:"written_by"`
	Teleplay			[]string	`json:"teleplay"`
	Season				int			`json:"season"`
	NoInSeason			int			`json:"no_in_season"`
	NoInSeries			int			`json:"no_in_series"`
	JapaneseTitle		string		`json:"japanese_title"`
	JapaneseAirDate		time.Time	`json:"japanese_air_date"`
}

func (c *Episodes) toString() string {

	return fmt.Sprintf("%+v", c)

}
