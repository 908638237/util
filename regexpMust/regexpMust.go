package regexpMust

import "regexp"

const (
	ID_CARD_15 = `^[1-6][1-9]\d{4}\d{2}((0[1-9])|10|11|12)(([0-2][1-9])|10|20|30|31)\d{3}$`
	ID_CARD_18 = `^([1-6][1-9]|50)\d{4}(18|19|20|21|22|23)\d{2}((0[1-9])|10|11|12)(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`
	MOBILE     = `^1(3[0-9]|4([0-1]|[4-9])|5([0-3]|[5-9])|6[2567]|7[0-8]|8[0-9]|9([0-3]|[5-9]))\d{8}$`
)

func NewMustCompile(str string) *regexp.Regexp {
	return regexp.MustCompile(str)
}
