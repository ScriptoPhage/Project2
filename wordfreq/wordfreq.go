package wordfreq

import (
	"encoding/json"
	"log"
	"regexp"
	"sort"
	"strings"
)

type wordFreqPair struct {
	Word  string `json:"Word"`
	Count int    `json:"Times of Occurrence"`
}

type wordFrequencies []wordFreqPair

func (a wordFrequencies) Len() int {
	return len(a)
}
func (a wordFrequencies) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a wordFrequencies) Less(i, j int) bool {
	return a[i].Count <= a[j].Count
}

func mapToSlice(m map[string]int) wordFrequencies {
	var wf wordFrequencies
	for k, v := range m {
		wf = append(wf, wordFreqPair{Word: k, Count: v})
	}
	return wf
}

func WordCountService(text string) []byte {
	lowerCasedText := strings.ToLower(text)
	newLineLessText := strings.ReplaceAll(lowerCasedText, "\n", " ")
	punctuationLessText := depunctuator(newLineLessText)
	words := strings.Fields(punctuationLessText)
	wordCounts := make(map[string]int)
	for _, word := range words {
		wordCounts[word]++
	}
	wf := mapToSlice(wordCounts)
	sort.Sort(sort.Reverse(wf)) //do sort in descending order
	var top10Words wordFrequencies
	if len(wf) < 10 {
		top10Words = wf
	} else {
		top10Words = wf[:10]
	}
	jsonData, _ := json.MarshalIndent(top10Words, "", " ")
	return jsonData

}

func depunctuator(text string) string {
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	return re.ReplaceAllString(text, " ")
}
