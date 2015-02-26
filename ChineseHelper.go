package gpinyin

import (
	"regexp"
)

const (
	data_Chinese_tas  = "db/chinese.txt"
	data_pinyin       = "db/pinyin.txt"
	data_multi_pinyin = "db/multi_pinyin.txt"
)

var traditionalChinese map[string]string
var simplifiedChinese map[string]string
var chineseRegex *regexp.Regexp

func init() {
	chineseRegex = regexp.MustCompile("[\u4e00-\u9fa5]")
}

func ConvertToSimplifiedChinese(source string) string {
	result := ""
	for _, runeValue := range source {
		result += toSimplifiedChinese(string(runeValue))
	}
	return result
}

func ConvertToTraditionalChinese(source string) string {
	result := ""
	for _, runeValue := range source {
		result += toTraditionalChinese(string(runeValue))
	}
	return result
}

func loadMapFromResource(resourceName string, reverse bool) map[string]string {
	v := make(map[string]string)
	err := loadResource(resourceName, v, reverse)
	if err != nil {
		panic(err)
	}
	return v
}

func toSimplifiedChinese(source string) string {
	if simplifiedChinese == nil {
		simplifiedChinese = loadMapFromResource(data_Chinese_tas, false)
	}
	v := simplifiedChinese[source]
	if len(v) == 0 {
		return source
	}
	return v
}

func toTraditionalChinese(source string) string {
	if traditionalChinese == nil {
		traditionalChinese = loadMapFromResource(data_Chinese_tas, true)
	}
	v := traditionalChinese[source]
	if len(v) == 0 {
		return source
	}
	return v
}

func isChinese(char string) bool {
	return chineseRegex.MatchString(char)
}