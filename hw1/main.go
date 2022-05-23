package hw1

import "io/ioutil"

func Read(filename string) (error, string) {
	b, e := ioutil.ReadFile(filename)
	if e != nil {
		return e, ""
	}
	return nil, string(b)
}
