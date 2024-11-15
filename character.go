package astro

import "fmt"

type character struct {
	hour int
	*date
}

func (ew *character) Word() string {
	hGanZhi, _ := newGanZhi(ew.DayGanZhi[:3], ew.hour)
	return fmt.Sprintf("%s%s%s%s", ew.YearGanZhi, ew.MonthGanZhi, ew.DayGanZhi, hGanZhi)
}

func (ew *character) Words() []string {
	hGanZhi, _ := newGanZhi(ew.DayGanZhi[:3], ew.hour)
	words := []string{ew.YearGanZhi[:3], ew.YearGanZhi[3:], ew.MonthGanZhi[:3], ew.MonthGanZhi[3:], ew.DayGanZhi[:3], ew.DayGanZhi[3:], hGanZhi[:3], hGanZhi[3:]}
	return words
}

func (ew *character) HourGanZhi() string {
	hGanZhi, _ := newGanZhi(ew.DayGanZhi[:3], ew.hour)
	return hGanZhi
}

func (ew *character) WuXingAttr() []string {
	return wuXingAttrs(ew.Words())
}

func (ew *character) EWString() string {
	return fmt.Sprintf("%s %s %s %s月%s 八字：%v 五行：%v 缺：%v", ew.Solar(), ew.Constellation(), ew.Animal, ew.LunarMonth, ew.LunarDay, ew.Words(), ew.WuXingAttr(), ew.MissWuXingAttr())
}

func (ew *character) MissWuXingAttr() []string {
	return missWuXing(ew.WuXingAttr())
}
