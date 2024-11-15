package astro

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strconv"
)

//go:embed date.json
var data []byte

var dateMap map[string]*date

func init() {
	var dates []*date
	err := json.Unmarshal(data, &dates)
	if err != nil {
		panic(err)
	}
	dateMap = make(map[string]*date)
	for _, d := range dates {
		dateMap[d.Solar()] = d
	}
}

type Date interface {
	Solar() string
	Lunar() string
	Week() string
	Constellation() string
	YearCharacter() []string
	MonthCharacter() []string
	DayCharacter() []string
	HourCharacter(hour int) []string
	AllCharacter(hour int) []string
	FiveElementsAttr(hour int) []string
	Zodiac() string
	Festivals() []string
	Suitable() []string
	Avoid() []string
}

// GetDate 根据日期获取黄历数据，参数date格式： 2006-01-02
func GetDate(date string) Date {
	if d, ok := dateMap[date]; ok {
		return d
	}
	return nil
}

type date struct {
	Year          int      `json:"year"`
	Month         int      `json:"month"`
	Day           int      `json:"day"`
	LunarMonth    string   `json:"lunar_month"`
	LunarDay      string   `json:"lunar_day"`
	WeekNo        int      `json:"week"`
	IsLargeMonth  bool     `json:"is_large_month"`
	Animal        string   `json:"animal"`
	YearGanZhi    string   `json:"year_gan_zhi"`
	MonthGanZhi   string   `json:"month_gan_zhi"`
	DayGanZhi     string   `json:"day_gan_zhi"`
	FestivalsList []string `json:"festivals"`
	SuitableList  []string `json:"suitable"`
	AvoidList     []string `json:"avoid"`
}

func (d *date) Solar() string {
	var mm, dd string
	if d.Day < 10 {
		dd = fmt.Sprintf("0%d", d.Day)
	} else {
		dd = strconv.Itoa(d.Day)
	}
	if d.Month < 10 {
		mm = fmt.Sprintf("0%d", d.Month)
	} else {
		mm = strconv.Itoa(d.Month)
	}
	return fmt.Sprintf("%d-%s-%s", d.Year, mm, dd)
}

func (d *date) Lunar() string {
	return fmt.Sprintf("%s年%s月%s日", d.YearGanZhi, d.LunarMonth, d.LunarDay)
}

func (d *date) Week() string {
	switch d.WeekNo {
	case 1:
		return "星期一"
	case 2:
		return "星期二"
	case 3:
		return "星期三"
	case 4:
		return "星期四"
	case 5:
		return "星期五"
	case 6:
		return "星期六"
	default:
		return "星期日"
	}
}

func (d *date) Constellation() string {
	return GetConstellation(d.Solar())
}

func (d *date) NewCharacter(hour int) *character {
	return &character{
		hour: hour,
		date: d,
	}
}

func (d *date) AllCharacter(hour int) []string {
	c := &character{
		hour: hour,
		date: d,
	}
	return c.Words()
}

func (d *date) FiveElementsAttr(hour int) []string {
	c := &character{
		hour: hour,
		date: d,
	}
	return c.WuXingAttr()
}

func (d *date) YearCharacter() []string {
	return []string{d.YearGanZhi[:3], d.YearGanZhi[3:]}
}

func (d *date) MonthCharacter() []string {
	return []string{d.MonthGanZhi[:3], d.MonthGanZhi[3:]}
}

func (d *date) DayCharacter() []string {
	return []string{d.DayGanZhi[:3], d.DayGanZhi[3:]}
}

func (d *date) HourCharacter(hour int) []string {
	c := &character{
		hour: hour,
		date: d,
	}
	hGanZhi := c.HourGanZhi()
	return []string{hGanZhi[:3], hGanZhi[3:]}
}

func (d *date) Zodiac() string {
	return d.Animal
}

func (d *date) Festivals() []string {
	return d.FestivalsList
}

func (d *date) Suitable() []string {
	return d.SuitableList
}

func (d *date) Avoid() []string {
	return d.AvoidList
}
