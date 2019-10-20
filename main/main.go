package main

import "data-structures/questions"

func main() {

	//f := &linked_list.Feed{}
	//f2 := &linked_list.Feed_d{}
	//p1 := linked_list.Post{
	//	Body:        "First Value",
	//	PublishDate: time.Now().Unix(),
	//}
	//t := time.Now().Unix()
	//f2.Append(&p1)
	//p2 := linked_list.Post{
	//	Body:        "Second Value",
	//	PublishDate: t,
	//}
	//p3 := linked_list.Post{
	//	Body:        "Third Value",
	//	PublishDate: time.Now().Unix(),
	//}
	//f2.Append(&p2)
	//f2.InsertWithSort(&p3)
	//
	//l := list.New()
	//l.PushBack("First")
	//l.PushFront("Second")
	//println(l.Len())
	//
	//s := stack.New()
	//s.Push(f2)
	//s.Push("1")
	//var val = s.Pop()

	//p := trees.Node{
	//	Tag:  "p",
	//	Text: "And this is some text in a paragraph. And next to it there's an image.",
	//}
	//
	//span := trees.Node{
	//	Tag:  "span",
	//	Id:   "copyright",
	//	Text: "2019 &copy; Ilija Eftimov",
	//}
	//
	//div := trees.Node{
	//	Tag:      "div",
	//	Class:    "footer",
	//	Text:     "This is the footer of the page.",
	//	Children: []*trees.Node{&span},
	//}
	//
	//h1 := trees.Node{
	//	Tag:  "h1",
	//	Text: "This is a H1",
	//}
	//
	//body := trees.Node{
	//	Tag:      "body",
	//	Children: []*trees.Node{&h1, &p, &div},
	//}
	//
	//html := trees.Node{
	//	Tag:      "html",
	//	Children: []*trees.Node{&body},
	//}

	//questions.AsteroidCollision([]int{5,10,-5})
	//questions.AsteroidCollision([]int{-2,-1,1,2})
	//questions.ProductExceptSelf([]int{1, 2, 3, 4})
	println(questions.TwoSumLessThanK([]int {100,200,360},60))
}
