package controllers

import (
	"strconv"
	"time"
	"strings"
	"github.com/revel/revel"
	"revel-sample/app/forms"
	"revel-sample/app/models"
	"revel-sample/app/routes"
)

type Columbo struct {
	*revel.Controller
}

var episodes []*models.Episodes

func init() {
	lc, _ := time.LoadLocation("Asia/Tokyo")
	
	episodes = []*models.Episodes{
		&models.Episodes{Title: "Prescription: Murder", OriginalAirDate: "February 20, 1968", Runtime: 98, GuestStaring: "Gene Barry", GuestStaringRole: "Dr. Ray Fleming (Gene Barry), a psychiatrist", DirectedBy: "Richard Irving", WrittenBy: []string{"Richard Levinson & William Link"}, Teleplay: []string{""}, Season: 0, NoInSeason: 1, NoInSeries: 1, JapaneseTitle: "殺人処方箋", JapaneseAirDate: time.Date(1972, 8, 27, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "Ransom for a Dead Man", OriginalAirDate: "March 1, 1971", Runtime: 98, GuestStaring: "Lee Grant", GuestStaringRole: "Leslie Williams, a brilliant lawyer and pilot", DirectedBy: "Richard Irving", WrittenBy: []string{"Richard Levinson & William Link"}, Teleplay: []string{"Dean Hargrove"}, Season: 0, NoInSeason: 2, NoInSeries: 2, JapaneseTitle: "死者の身代金", JapaneseAirDate: time.Date(1973, 4, 22, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "Murder by the Book", OriginalAirDate: "September 15, 1971", Runtime: 73, GuestStaring: "Jack Cassidy", GuestStaringRole: "Ken Franklin is one half of a mystery writing team", DirectedBy: "Steven Spielberg", WrittenBy: []string{"Steven Bochco"}, Teleplay: []string{""}, Season: 1, NoInSeason: 1, NoInSeries: 3, JapaneseTitle: "構想の死角", JapaneseAirDate: time.Date(1972, 11, 26, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "Death Lends a Hand", OriginalAirDate: "October 6, 1971", Runtime: 73, GuestStaring: "Robert Culp", GuestStaringRole: "Carl Brimmer, The head of a private detective agency", DirectedBy: "Bernard L. Kowalski", WrittenBy: []string{"RRichard Levinson & William Link"}, Teleplay: []string{""}, Season: 1, NoInSeason: 2, NoInSeries: 4, JapaneseTitle: "指輪の爪あと", JapaneseAirDate: time.Date(1973, 1, 21, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "Dead Weight", OriginalAirDate: "October 27, 1971", Runtime: 73, GuestStaring: "Eddie Albert", GuestStaringRole: "Major General Martin Hollister, a retired Marine Corps war hero", DirectedBy: "Jack Smight", WrittenBy: []string{"John T. Dugan"}, Teleplay: []string{""}, Season: 1, NoInSeason: 3, NoInSeries: 5, JapaneseTitle: "ホリスター将軍のコレクション", JapaneseAirDate: time.Date(1972, 9, 24, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "Suitable for Framing", OriginalAirDate: "November 17, 1971", Runtime: 73, GuestStaring: "Ross Martin", GuestStaringRole: "Dale Kingston, Art critic", DirectedBy: "Hy Averback", WrittenBy: []string{"Jackson Gillis"}, Teleplay: []string{""}, Season: 1, NoInSeason: 4, NoInSeries: 6, JapaneseTitle: "二枚のドガの絵", JapaneseAirDate: time.Date(1972, 10, 22, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "Lady in Waiting", OriginalAirDate: "December 15, 1971", Runtime: 73, GuestStaring: "Susan Clark", GuestStaringRole: "Beth Chadwick, sister of domineering, Bryce", DirectedBy: "Norman Lloyd", WrittenBy: []string{"Barney Slater"}, Teleplay: []string{"Steven Bochco"}, Season: 1, NoInSeason: 5, NoInSeries: 7, JapaneseTitle: "もう一つの鍵", JapaneseAirDate: time.Date(1972, 12, 17, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "Short Fuse", OriginalAirDate: "January 19, 1972", Runtime: 73, GuestStaring: "Roddy McDowall", GuestStaringRole: "Roger Stanford, A chemist", DirectedBy: "Edward M. Abrams", WrittenBy: []string{"Lester & Tina Pine", "Jackson Gillis"}, Teleplay: []string{"Jackson Gillis"}, Season: 1, NoInSeason: 6, NoInSeries: 8, JapaneseTitle: "死の方程式", JapaneseAirDate: time.Date(1973, 3, 18, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "Blueprint for Murder", OriginalAirDate: "February 9, 1972", Runtime: 73, GuestStaring: "Patrick O'Neal", GuestStaringRole: "Elliot Markham, An architect", DirectedBy: "Peter Falk", WrittenBy: []string{"William Kelley"}, Teleplay: []string{"Steven Bochco"}, Season: 1, NoInSeason: 7, NoInSeries: 9, JapaneseTitle: "パイルD-3の壁", JapaneseAirDate: time.Date(1973, 2, 25, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "Étude in Black", OriginalAirDate: "September 17, 1972", Runtime: 98, GuestStaring: "John Cassavetes", GuestStaringRole: "Alex Benedict, The conductor of the Los Angeles Philharmonic Orchestra", DirectedBy: "Nicholas Colasanto", WrittenBy: []string{"Richard Levinson & William Link"}, Teleplay: []string{"Steven Bochco"}, Season: 2, NoInSeason: 1, NoInSeries: 10, JapaneseTitle: "黒のエチュード", JapaneseAirDate: time.Date(1973, 9, 30, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "The Greenhouse Jungle", OriginalAirDate: "October 15, 1972", Runtime: 73, GuestStaring: "Ray Milland", GuestStaringRole: "Jarvis Goodland, An expert in orchids", DirectedBy: "Boris Sagal", WrittenBy: []string{"Jonathan Latimer"}, Teleplay: []string{""}, Season: 2, NoInSeason: 2, NoInSeries: 11, JapaneseTitle: "悪の温室", JapaneseAirDate: time.Date(1973, 5, 27, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "The Most Crucial Game", OriginalAirDate: "November 5, 1972", Runtime: 73, GuestStaring: "Robert Culp", GuestStaringRole: "Paul Hanlon, The general manager of the Los Angeles Rockets football team", DirectedBy: "Jeremy Kagan", WrittenBy: []string{"John T. Dugan"}, Teleplay: []string{""}, Season: 2, NoInSeason: 3, NoInSeries: 12, JapaneseTitle: "アリバイのダイヤル", JapaneseAirDate: time.Date(1973, 6, 24, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "Dagger of the Mind", OriginalAirDate: "November 26, 1972", Runtime: 98, GuestStaring: "Richard Basehart", GuestStaringRole: "Actors Nicholas Framer and his wife, Lillian Stanhope", DirectedBy: "Richard Quine", WrittenBy: []string{"Richard Levinson & William Link"}, Teleplay: []string{"Jackson Gillis"}, Season: 2, NoInSeason: 4, NoInSeries: 13, JapaneseTitle: "ロンドンの傘", JapaneseAirDate: time.Date(1973, 7, 29, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "Requiem for a Falling Star", OriginalAirDate: "January 21, 1973", Runtime: 73, GuestStaring: "Anne Baxter", GuestStaringRole: "movie star Nora Chandler", DirectedBy: "Richard Quine", WrittenBy: []string{"Jackson Gillis"}, Teleplay: []string{""}, Season: 2, NoInSeason: 5, NoInSeries: 14, JapaneseTitle: "偶像のレクイエム", JapaneseAirDate: time.Date(1973, 8, 26, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "A Stitch in Crime", OriginalAirDate: "February 11, 1973", Runtime: 73, GuestStaring: "Leonard Nimoy", GuestStaringRole: "Cardiac surgeon Dr. Barry Mayfield", DirectedBy: "Hy Averback", WrittenBy: []string{"Shirl Hendryx"}, Teleplay: []string{""}, Season: 2, NoInSeason: 6, NoInSeries: 15, JapaneseTitle: "溶ける糸", JapaneseAirDate: time.Date(1973, 10, 28, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "The Most Dangerous Match", OriginalAirDate: "March 4, 1973", Runtime: 73, GuestStaring: "Laurence Harvey", GuestStaringRole: "Chess Grandmaster Emmett Clayton", DirectedBy: "Edward M. Abroms", WrittenBy: []string{"Jackson Gillis", "Richard Levinson & William Link"}, Teleplay: []string{"Jackson Gillis"}, Season: 2, NoInSeason: 7, NoInSeries: 16, JapaneseTitle: "断たれた音", JapaneseAirDate: time.Date(1973, 11, 25, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "Double Shock", OriginalAirDate: "March 25, 1973", Runtime: 73, GuestStaring: "Martin Landau", GuestStaringRole: "Flamboyant television chef Dexter Paris and his twin brother, conservative banker Norman", DirectedBy: "Robert Butler", WrittenBy: []string{"Jackson Gillis", "Richard Levinson & William Link"}, Teleplay: []string{"Steven Bochco"}, Season: 2, NoInSeason: 8, NoInSeries: 17, JapaneseTitle: "二つの顔", JapaneseAirDate: time.Date(1973, 12, 23, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "Lovely But Lethal", OriginalAirDate: "September 23, 1973", Runtime: 73, GuestStaring: "Vera Miles", GuestStaringRole: "Cosmetics queen Viveca Scott", DirectedBy: "Jeannot Szwarc", WrittenBy: []string{"Myrna Bercovici"}, Teleplay: []string{"Jackson Gillis"}, Season: 3, NoInSeason: 1, NoInSeries: 18, JapaneseTitle: "毒のある花", JapaneseAirDate: time.Date(1974, 9, 14, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "Any Old Port in a Storm", OriginalAirDate: "October 7, 1973", Runtime: 98, GuestStaring: "Donald Pleasence", GuestStaringRole: "Wine connoisseur Adrian Carsini", DirectedBy: "Leo Penn", WrittenBy: []string{"Larry Cohen"}, Teleplay: []string{"Stanley Ralph Ross"}, Season: 3, NoInSeason: 2, NoInSeries: 19, JapaneseTitle: "別れのワイン", JapaneseAirDate: time.Date(1974, 6, 29, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "Candidate for Crime", OriginalAirDate: "November 4, 1973", Runtime: 98, GuestStaring: "Jackie Cooper", GuestStaringRole: "Nelson Hayward, is coercing the womanizing senatorial candidate", DirectedBy: "Boris Sagal", WrittenBy: []string{"Larry Cohen"}, Teleplay: []string{"Irving Pearlberg & Alvin R. Friedman", "Roland Kibbee & Dean Hargrove"}, Season: 3, NoInSeason: 3, NoInSeries: 20, JapaneseTitle: "野望の果て", JapaneseAirDate: time.Date(1974, 8, 17, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "Double Exposure", OriginalAirDate: "December 16, 1973", Runtime: 73, GuestStaring: "Robert Culp", GuestStaringRole: "Dr. Bart Keppel, A motivation research specialist", DirectedBy: "Richard Quine", WrittenBy: []string{"Stephen J. Cannell"}, Teleplay: []string{""}, Season: 3, NoInSeason: 4, NoInSeries: 21, JapaneseTitle: "意識の下の映像", JapaneseAirDate: time.Date(1974, 8, 10, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "Publish or Perish", OriginalAirDate: "January 18, 1974", Runtime: 73, GuestStaring: "Jack Cassidy", GuestStaringRole: "Riley Greenleaf, Publisher", DirectedBy: "Robert Butler", WrittenBy: []string{"Peter S. Fischer"}, Teleplay: []string{""}, Season: 3, NoInSeason: 5, NoInSeries: 22, JapaneseTitle: "第三の終章", JapaneseAirDate: time.Date(1974, 12, 14, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "Mind Over Mayhem", OriginalAirDate: "February 10, 1974", Runtime: 73, GuestStaring: "José Ferrer", GuestStaringRole: "Dr. Marshall Cahill, director of a high-tech Pentagon think tank", DirectedBy: "Alf Kjellin", WrittenBy: []string{"Robert Specht"}, Teleplay: []string{"Steven Bochco", "Dean Hargrove & Roland Kibbee"}, Season: 3, NoInSeason: 6, NoInSeries: 23, JapaneseTitle: "愛情の計算", JapaneseAirDate: time.Date(1974, 8, 31, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "Swan Song", OriginalAirDate: "March 3, 1974", Runtime: 98, GuestStaring: "Johnny Cash", GuestStaringRole: "Gospel-singing superstar Tommy Brown", DirectedBy: "Nicholas Colasanto", WrittenBy: []string{"Stanley Ralph Ross"}, Teleplay: []string{"David Rayfiel"}, Season: 3, NoInSeason: 7, NoInSeries: 24, JapaneseTitle: "白鳥の歌", JapaneseAirDate: time.Date(1974, 9, 21, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "A Friend in Deed", OriginalAirDate: "May 5, 1974", Runtime: 98, GuestStaring: "Richard Kiley", GuestStaringRole: "Deputy police commissioner Mark Halperin", DirectedBy: "Ben Gazzara", WrittenBy: []string{"Peter S. Fischer"}, Teleplay: []string{""}, Season: 3, NoInSeason: 8, NoInSeries: 25, JapaneseTitle: "権力の墓穴", JapaneseAirDate: time.Date(1974, 10, 5, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "An Exercise in Fatality", OriginalAirDate: "September 15, 1974", Runtime: 98, GuestStaring: "Robert Conrad", GuestStaringRole: "Renowned exercise guru Milo Janus", DirectedBy: "Bernard L. Kowalski", WrittenBy: []string{"Larry Cohen"}, Teleplay: []string{"Peter S. Fischer"}, Season: 4, NoInSeason: 1, NoInSeries: 26, JapaneseTitle: "自縛の紐", JapaneseAirDate: time.Date(1975, 12, 27, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "Negative Reaction", OriginalAirDate: "October 6, 1974", Runtime: 98, GuestStaring: "Dick Van Dyke", GuestStaringRole: "professional photographer Paul Galesko", DirectedBy: "Alf Kjellin", WrittenBy: []string{"Peter S. Fischer"}, Teleplay: []string{""}, Season: 4, NoInSeason: 2, NoInSeries: 27, JapaneseTitle: "逆転の構図", JapaneseAirDate: time.Date(1975, 12, 20, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "By Dawn's Early Light", OriginalAirDate: "October 27, 1974", Runtime: 98, GuestStaring: "Patrick McGoohan", GuestStaringRole: "Colonel Lyle C. Rumford, head of the Haynes Military Academy", DirectedBy: "Harvey Hart", WrittenBy: []string{"Howard Berk"}, Teleplay: []string{""}, Season: 4, NoInSeason: 3, NoInSeries: 28, JapaneseTitle: "祝砲の挽歌", JapaneseAirDate: time.Date(1976, 1, 10, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "Troubled Waters", OriginalAirDate: "February 9, 1975", Runtime: 98, GuestStaring: "Robert Vaughn", GuestStaringRole: "Auto executive Hayden Danziger", DirectedBy: "Ben Gazzara", WrittenBy: []string{"Jackson Gillis", "William Driskill"}, Teleplay: []string{"William Driskill"}, Season: 4, NoInSeason: 4, NoInSeries: 29, JapaneseTitle: "歌声の消えた海", JapaneseAirDate: time.Date(1976, 1, 3, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "Playback", OriginalAirDate: "March 2, 1975", Runtime: 73, GuestStaring: "Oskar Werner", GuestStaringRole: "Harold Van Wick, the gadget-obsessed president of Midas Electronics", DirectedBy: "Bernard L. Kowalski", WrittenBy: []string{"David P. Lewis & Booker T. Bradshaw"}, Teleplay: []string{""}, Season: 4, NoInSeason: 5, NoInSeries: 30, JapaneseTitle: "ビデオテープの証言", JapaneseAirDate: time.Date(1976, 12, 11, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "A Deadly State of Mind", OriginalAirDate: "April 27, 1975", Runtime: 73, GuestStaring: "George Hamilton", GuestStaringRole: "Psychiatrist Dr. Mark Collier", DirectedBy: "Harvey Hart", WrittenBy: []string{"Peter S. Fischer"}, Teleplay: []string{""}, Season: 4, NoInSeason: 6, NoInSeries: 31, JapaneseTitle: "5時30分の目撃者", JapaneseAirDate: time.Date(1976, 12, 18, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "Forgotten Lady", OriginalAirDate: "September 14, 1975", Runtime: 85, GuestStaring: "Janet Leigh", GuestStaringRole: "Aging former movie star Grace Wheeler", DirectedBy: "Harvey Hart", WrittenBy: []string{"Bill Driskill"}, Teleplay: []string{""}, Season: 5, NoInSeason: 1, NoInSeries: 32, JapaneseTitle: "忘れられたスター", JapaneseAirDate: time.Date(1977, 1, 3, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "A Case of Immunity", OriginalAirDate: "October 12, 1975", Runtime: 73, GuestStaring: "Héctor Elizondo", GuestStaringRole: "Hassan Salah, chief diplomat of the Legation of Swahari", DirectedBy: "Ted Post", WrittenBy: []string{"James Menzies"}, Teleplay: []string{"Lou Shaw"}, Season: 5, NoInSeason: 2, NoInSeries: 33, JapaneseTitle: "ハッサン・サラーの反逆", JapaneseAirDate: time.Date(1976, 12, 25, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "Identity Crisis", OriginalAirDate: "November 2, 1975", Runtime: 98, GuestStaring: "Patrick McGoohan", GuestStaringRole: "speech-writing consultant Nelson Brenner", DirectedBy: "Patrick McGoohan", WrittenBy: []string{"Bill Driskill"}, Teleplay: []string{""}, Season: 5, NoInSeason: 3, NoInSeries: 34, JapaneseTitle: "仮面の男", JapaneseAirDate: time.Date(1977, 9, 24, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "A Matter of Honor", OriginalAirDate: "February 1, 1976", Runtime: 73, GuestStaring: "Ricardo Montalban", GuestStaringRole: "Luis Montoya, A Mexican national hero", DirectedBy: "Ted Post", WrittenBy: []string{"Brad Radnitz"}, Teleplay: []string{""}, Season: 5, NoInSeason: 4, NoInSeries: 35, JapaneseTitle: "闘牛士の栄光", JapaneseAirDate: time.Date(1977, 10, 1, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "Now You See Him...", OriginalAirDate: "February 29, 1976", Runtime: 85, GuestStaring: "Jack Cassidy", GuestStaringRole: "Great Santini, a magician extraordinaire", DirectedBy: "Harvey Hart", WrittenBy: []string{"Michael Sloan"}, Teleplay: []string{""}, Season: 5, NoInSeason: 5, NoInSeries: 36, JapaneseTitle: "魔術師の幻想", JapaneseAirDate: time.Date(1977, 12, 31, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "Last Salute to the Commodore", OriginalAirDate: "May 2, 1976", Runtime: 98, GuestStaring: "Robert Vaughn", GuestStaringRole: "Son-in-law Charles Clay", DirectedBy: "Patrick McGoohan", WrittenBy: []string{"Jackson Gillis"}, Teleplay: []string{""}, Season: 5, NoInSeason: 6, NoInSeries: 37, JapaneseTitle: "さらば提督", JapaneseAirDate: time.Date(1977, 10, 8, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "Fade in to Murder", OriginalAirDate: "October 10, 1976", Runtime: 73, GuestStaring: "William Shatner", GuestStaringRole: "Egocentric actor Ward Fowler", DirectedBy: "Bernard L. Kowalski", WrittenBy: []string{"Henry Garson"}, Teleplay: []string{"Lou Shaw", "Peter S. Feibleman"}, Season: 6, NoInSeason: 1, NoInSeries: 38, JapaneseTitle: "ルーサン警部の犯罪", JapaneseAirDate: time.Date(1977, 12, 17, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "Old Fashioned Murder", OriginalAirDate: "November 28, 1976", Runtime: 73, GuestStaring: "Joyce Van Patten", GuestStaringRole: "Ruth Lytton, Owner of the Lytton Museum", DirectedBy: "Robert Douglas", WrittenBy: []string{"Lawrence Vail"}, Teleplay: []string{"Peter S. Feibleman"}, Season: 6, NoInSeason: 2, NoInSeries: 39, JapaneseTitle: "黄金のバックル", JapaneseAirDate: time.Date(1977, 12, 24, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "The Bye-Bye Sky High IQ Murder Case", OriginalAirDate: "May 22, 1977", Runtime: 73, GuestStaring: "Theodore Bikel", GuestStaringRole: "Oliver Brandt, a senior partner in an accounting firm", DirectedBy: "Sam Wanamaker", WrittenBy: []string{"Robert Malcolm Young"}, Teleplay: []string{""}, Season: 6, NoInSeason: 3, NoInSeries: 40, JapaneseTitle: "殺しの序曲", JapaneseAirDate: time.Date(1978, 5, 20, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "Try and Catch Me", OriginalAirDate: "November 21, 1977", Runtime: 73, GuestStaring: "Ruth Gordon", GuestStaringRole: "Mystery author Abigail Mitchell", DirectedBy: "James Frawley", WrittenBy: []string{"Gene Thompson"}, Teleplay: []string{"Gene Thompson & Paul Tuckahoe"}, Season: 7, NoInSeason: 1, NoInSeries: 41, JapaneseTitle: "死者のメッセージ", JapaneseAirDate: time.Date(1978, 4, 8, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "Murder Under Glass", OriginalAirDate: "January 30, 1978", Runtime: 73, GuestStaring: "Louis Jourdan", GuestStaringRole: "Renowned restaurant critic Paul Gerard", DirectedBy: "Jonathan Demme", WrittenBy: []string{"Robert van Scoyk"}, Teleplay: []string{""}, Season: 7, NoInSeason: 2, NoInSeries: 42, JapaneseTitle: "美食の報酬", JapaneseAirDate: time.Date(1978, 5, 27, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "Make Me a Perfect Murder", OriginalAirDate: "February 28, 1978", Runtime: 98, GuestStaring: "Trish Van Devere", GuestStaringRole: "TV programmer Kay Freestone", DirectedBy: "James Frawley", WrittenBy: []string{"Robert Blees"}, Teleplay: []string{""}, Season: 7, NoInSeason: 3, NoInSeries: 43, JapaneseTitle: "秒読みの殺人", JapaneseAirDate: time.Date(1979, 1, 2, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "How to Dial a Murder", OriginalAirDate: "April 15, 1978", Runtime: 73, GuestStaring: "Nicol Williamson", GuestStaringRole: "Mind control seminar guru Dr. Eric Mason", DirectedBy: "James Frawley", WrittenBy: []string{"Anthony Lawrence"}, Teleplay: []string{"Tom Lazarus"}, Season: 7, NoInSeason: 4, NoInSeries: 44, JapaneseTitle: "攻撃命令", JapaneseAirDate: time.Date(1979, 1, 4, 0, 0, 0, 0, lc)},
		&models.Episodes{Title: "The Conspirators", OriginalAirDate: "May 13, 1978", Runtime: 98, GuestStaring: "Clive Revill", GuestStaringRole: "Famous Irish poet and author Joe Devlin", DirectedBy: "Leo Penn", WrittenBy: []string{"Howard Berk"}, Teleplay: []string{""}, Season: 7, NoInSeason: 5, NoInSeries: 45, JapaneseTitle: "策謀の結末", JapaneseAirDate: time.Date(1979, 1, 3, 0, 0, 0, 0, lc)},
	}
}

func (c Columbo) Search(s forms.SearchForm) revel.Result {
	message := "これはStep3の実装です。"
	return c.Render(s, message)

}

func (c Columbo) Confirm(s forms.SearchForm) revel.Result {
	s.Validate(c.Validation, c.Request.Locale)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Columbo.Search(s))
	}
	return c.Redirect(routes.Columbo.Result(s))
}

func (c Columbo) Result(s forms.SearchForm) revel.Result {
	lists := []*models.Episodes{}
	unmatch := false
	
	for _, v := range episodes {
		unmatch = false
		if s.Title != " " {
			if !isContains(v.Title , s.Title) {
				unmatch = true
			}
		}
		if s.Runtime != " " {
			runtime, err := strconv.Atoi(s.Runtime)
			if err != nil {
				unmatch = true
			} else {
				if v.Runtime != runtime {
					unmatch = true
				}
			}
		}
		if s.GuestStaring != " " {
			if !isContains(v.GuestStaring, s.GuestStaring) {
				unmatch = true
			}
		}
		if !unmatch {
			lists = append(lists, v)
		}
	
	}
	
	return c.Render(s, lists)

}

func isContains(v string, s string) bool {
	if s == " " {
		return false
	}
	
	if v == " " {
		return false
	}
	return strings.Contains(strings.ToLower(v),strings.ToLower(s))

}