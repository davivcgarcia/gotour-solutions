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
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	for i := range b {
		b[i] = rot13(b[i])
	}
	return n, err
}

func rot13(x byte) byte {
	var a, z byte
	switch {
	case x >= 'a' && x <= 'z':
		a, z = 'a', 'z'
	case x >= 'A' && x <= 'Z':
		a, z = 'A', 'Z'
	default:
		return x
	}
	return (((x + 13) - a) % (z - a + 1)) + a
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
