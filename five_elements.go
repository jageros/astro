package astro

var combineAttr = map[string]string{
	// 五合
	"甲己": "土",
	"乙庚": "金",
	"丙辛": "水",
	"丁壬": "木",
	"戊癸": "火",
	// 六合
	"子丑": "土",
	"寅亥": "木",
	"卯戌": "火",
	"辰酉": "金",
	"巳申": "水",
	"午未": "土",
}

// wuXingAttr 五行属性
func wuXingAttr(word string) string {
	switch word {
	case "壬", "癸", "亥", "子":
		return "水"
	case "甲", "乙", "寅", "卯":
		return "木"
	case "丙", "丁", "巳", "午":
		return "火"
	case "庚", "辛", "申", "酉":
		return "金"
	case "戊", "己", "辰", "戌", "丑", "未":
		return "土"
	default:
		return ""
	}
}

// 八字的五行属性
func wuXingAttrs(words []string) []string {
	var wuxings []string
	for _, w := range words {
		wuxings = append(wuxings, wuXingAttr(w))
	}
	return wuxings
}

// missWuXing 缺失的属性
func missWuXing(wuxings []string) []string {
	var all = map[string]bool{"金": true, "木": true, "水": true, "火": true, "土": true}
	var miss []string
	for _, wx := range wuxings {
		if all[wx] {
			delete(all, wx)
		}
	}
	for wx := range all {
		miss = append(miss, wx)
	}
	return miss
}

func direction(attr string) string {
	switch attr {
	case "金":
		return "西"
	case "木":
		return "东"
	case "水":
		return "北"
	case "火":
		return "南"
	case "土":
		return "中"
	default:
		return ""
	}
}

// Combine 五合，六合
func Combine(words1, words2 []string) map[string]string {
	var result = map[string]string{}
	for _, w1 := range words1 {
		for _, w2 := range words2 {
			attr1 := combineAttr[w1+w2]
			attr2 := combineAttr[w2+w1]
			if attr1 != "" {
				result[w1+w2] = attr1
			}
			if attr2 != "" {
				result[w1+w2] = attr2
			}
		}
	}
	return result
}
