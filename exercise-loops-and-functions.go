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
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(x / 2)

	for z*z > x*1.0000000000000005 {
		z -= (z*z - x) / (2 * z)
	}

	return z
}

func main() {
	test := float64(13214)
	fmt.Println(Sqrt(test), math.Sqrt(test))
}
