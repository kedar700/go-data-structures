package linked_list

import "errors"

type Feed_d struct {
	Length int
	Start  *Post
	End    *Post
}

func (f *Feed_d) Append(newPost *Post) {
	if f.Length == 0 {
		f.Start = newPost
		f.End = newPost
	} else {
		lastPost := f.End
		lastPost.Next = newPost
		f.End = newPost
	}
	f.Length++
}

func (f *Feed_d) Remove(publishDate int64) {
	if f.Length == 0 {
		panic("The feed is empty")
	}
	var previousPost *Post = new(Post)
	currentPost := f.Start
	for currentPost.PublishDate != publishDate {
		if currentPost.Next == nil {
			panic(errors.New("No such Post found."))
		}
		previousPost = currentPost
		currentPost = currentPost.Next
	}
	previousPost.Next = currentPost.Next

	f.Length--
}

func (f *Feed_d) InsertWithSort(newPost *Post) {
	if f.Length == 0 {
		f.Start = newPost
		f.End = newPost
	} else {
		var previousPost *Post = new(Post)
		currentPost := f.Start
		for currentPost.PublishDate < newPost.PublishDate {
			previousPost = currentPost
			currentPost = previousPost.Next
		}
		previousPost.Next = newPost
		newPost.Next = currentPost
	}
	f.Length++
}
