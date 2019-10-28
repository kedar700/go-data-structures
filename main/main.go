package main

import (
	"data-structures/questions"
)

func main() {

	questions.AsteroidCollision([]int{5, 10, -5})
	questions.AsteroidCollision([]int{-2, -1, 1, 2})
	questions.ProductExceptSelf([]int{1, 2, 3, 4})
	println(questions.TwoSumLessThanK([]int{100, 200, 360}, 60))
	println(questions.LongestCommonPrefix([]string{"", ""}))

}
