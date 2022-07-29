package wordsFilter

import (
	wf "github.com/pides/gwordsfilter"
)

var Generate = wf.New()
var ReplaceStr = "*"

func init() {
	var err error
	err = Generate.Read("./gamedata/sys_filter_words.txt")
	if err != nil {
		panic(err)
	}
}

func Contains(text string) bool {
	flag, _ := Generate.CheckWord(text)
	return !flag
}

func Replace(text string) string {
	return Generate.Replace(text)
}
