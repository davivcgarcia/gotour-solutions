/*
Copyright 2017 Davi Garcia (davivcgarcia@gmail.com)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	wordMap := make(map[string]int)
	inputWordList := strings.Fields(s)

	for _, word := range inputWordList {
		if _, ok := wordMap[word]; ok {
			wordMap[word] += 1
		} else {
			wordMap[word] = 1
		}
	}

	return wordMap
}

func main() {
	wc.Test(WordCount)
}
