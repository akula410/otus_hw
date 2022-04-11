package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var pattern = regexp.MustCompile(`[^а-яА-Яa-zA-Z-\d]`)
var size = 10

func Top10(text string) []string {
	var result = make([]string, 0)

	var tmpSort = map[string]int64{}

	var key string
	var maxKey int64

	for _, item := range pattern.Split(text, -1) {
		key = strings.ToLower(item)
		if key == "-" || len(key) == 0 {
			continue
		}

		if _, ok := tmpSort[key]; ok {
			tmpSort[key]++
		} else {
			tmpSort[key] = 1
		}

		if tmpSort[key] > maxKey {
			maxKey = tmpSort[key]
		}
	}

	var tmpResult = make([][]string, maxKey+1)

	for item, index := range tmpSort {
		if tmpResult[index] == nil {
			tmpResult[index] = make([]string, 0)
		}

		tmpResult[index] = append(tmpResult[index], item)
	}


stop:
	for i := len(tmpResult) - 1; i > 0; i-- {
		if tmpResult[i] != nil {
			sort.Strings(tmpResult[i])
			for _, value := range tmpResult[i] {
				if len(result) >= size {
					break stop
				}
				result = append(result, value)
			}
		}
	}

	return result
}
