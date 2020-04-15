package annotate

import (
	"fmt"
	"github.com/Pantani/logger"
	"github.com/Pantani/replacer/pkg/replacer"
	"regexp"
	"strings"
)

const (
	TAG_LINK    = "{{TAG_LINK_%d}}"
	TAG_GENERAL = "{{TAG_GENERAL_%d}}"
)

func ScrapeHtml(text string, replace replacer.Names) string {
	tagMap := make(map[string]string)
	text = strings.ReplaceAll(text, "><", "> <")
	reg := regexp.MustCompile("<a.*?a>")
	find := reg.FindAll([]byte(text), -1)
	for i, r := range find {
		tag := fmt.Sprintf(TAG_LINK, i)
		tagMap[tag] = string(r)
		text = strings.ReplaceAll(text, string(r), tag)
	}

	reg = regexp.MustCompile("<.*?>")
	find = reg.FindAll([]byte(text), -1)
	for i, r := range find {
		tag := fmt.Sprintf(TAG_GENERAL, i)
		tagMap[tag] = string(r)
		text = strings.ReplaceAll(text, string(r), tag)
	}

	result := replaceStrings(text, replace)
	for tag, value := range tagMap {
		result = strings.ReplaceAll(result, tag, value)
	}
	result = strings.ReplaceAll(result, "> <", "><")
	return result
}

func replaceStrings(text string, replace replacer.Names) string {
	result := text
	for _, name := range replace {
		err := name.GenerateSnippet()
		if err != nil {
			logger.Error(err)
			continue
		}
		regex := fmt.Sprintf(`\b%s\b`, name.Name)
		reg := regexp.MustCompile(regex)
		result = reg.ReplaceAllString(result, name.Snippet)
	}
	return result
}
