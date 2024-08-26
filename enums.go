package main

import "fmt"

// iota 从0递增的枚举值
const (
	A = iota
	B
	C
	D
)

type Color int

const (
	RED Color = iota
	GREEN
	YELLOW
)

type SameSite int

// iota+1 代表从1递增
const (
	SameSiteDefaultMode SameSite = iota + 1
	SameSiteLaxMode
	SameSiteStrictMode
	SameSiteNoneMode
)

func main() {
	fmt.Println(A, B, C, D)
	fmt.Println(RED, GREEN, YELLOW)
	fmt.Println(SameSiteDefaultMode, SameSiteLaxMode, SameSiteStrictMode, SameSiteNoneMode)
}
