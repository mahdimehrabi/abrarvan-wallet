package infrastracture

import (
	"math/rand"
	"time"
)

var letterBytes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

type Random struct {
	Rand *rand.Rand
}

func NewRandom() Random {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return Random{
		Rand: r,
	}
}

//RefreshSeed => refresh random source seed ,
//please use this func before generate random important
func (r *Random) RefreshSeed() {
	s := rand.NewSource(time.Now().UnixNano())
	r.Rand = rand.New(s)
}

//RandomNumber => generate random nubmer
func (r *Random) RandomNumber(n int) int {
	return r.Rand.Intn(n)
}

//GenerateRandomStr
func (r *Random) GenerateRandomStr(ln int) string {
	b := make([]rune, ln)
	for i := 0; i < ln; i++ {
		b[i] = letterBytes[r.RandomNumber(len(letterBytes))]
	}
	return string(b)
}
