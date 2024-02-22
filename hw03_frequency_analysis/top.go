package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

func Top10(text string) []string {
	const top10 = 10
	strs := strings.Fields(text)
	glossary := map[string]int{}

	re := regexp.MustCompile(`^\p{P}{0,1}(\p{L}+|\p{L}+\p{P}+\p{L}+|\-{2,})\p{P}{0,1}$`)

	for _, s := range strs {
		if matched := re.FindAllStringSubmatch(s, -1); matched != nil {
			glossary[strings.ToLower(matched[0][1])]++
		}
	}
	words := make([]string, 0, len(glossary))

	for k := range glossary {
		words = append(words, k)
	}

	if len(words) == 0 {
		return words
	}

	sort.Strings(words)
	sort.SliceStable(words, func(i, j int) bool {
		return glossary[words[i]] > glossary[words[j]]
	})

	return words[:top10]
}
