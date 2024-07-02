package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandInrは最小値から最大値の中からランダムな整数を返す
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) //min->maxの範囲でランダムな値を返す
}

// RandomStringはn文字のランダムな文字列を返す
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwnerはランダムな所有者名を返す
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoneyはランダムな金額を返す
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrencyはランダムな通貨を返す
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
