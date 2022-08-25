package robotname

import (
	"errors"
	"math/rand"
	"time"
)

// Define the Robot type here.
type Robot struct {
	name string
}

var nameSpace = map[string]int{}

const last = 676000 // 26*26*10*10*10

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		r.name = randomString(2, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") + randomString(3, "0123456789")

		for {
			_, chk := nameSpace[r.name]
			if chk {
				r.name = randomString(2, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") + randomString(3, "0123456789")
			} else {
				break
			}
		}
	}
	if len(nameSpace) >= last {
		return "", errors.New("namespace is exhausted")
	}
	nameSpace[r.name] = 1

	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}

func randomString(n int, allowed string) string {
	letters := []rune(allowed)
	rand.Seed(time.Now().UnixNano())
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
