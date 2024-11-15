package astro

import (
	"fmt"
	"testing"
	"time"
)

func Test_Lunar(t *testing.T) {
	now := time.Now()
	d := GetDate(now.Format(time.DateOnly))
	fmt.Println("公历：", d.Solar(), d.Week())
	fmt.Println("农历：", d.Lunar())
	fmt.Println("星座：", d.Constellation())
	fmt.Println("生肖：", d.Zodiac())
	festivals := d.Festivals()
	if len(festivals) > 0 {
		fmt.Println("节日：", festivals)
	}
	fmt.Println("宜：", d.Suitable())
	fmt.Println("忌：", d.Avoid())
	fmt.Println("八字：", d.AllCharacter(now.Hour()))
	fmt.Println("五行：", d.FiveElementsAttr(now.Hour()))
}
