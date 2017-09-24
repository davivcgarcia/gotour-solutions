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

import "fmt"

func fibonacci() func() int {
	current := 1
	previous := 0
	index := 0
	return func() int {
		switch index {
		case 0:
			index += 1
			return 0
		case 1:
			index += 1
			return 1
		default:
			tmp := current + previous
			previous = current
			current = tmp
			index += 1
			return tmp
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
