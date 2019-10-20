package linked_list

type Post struct {
	Body        string
	PublishDate int64
	Next        *Post
}

type Feed struct {
	Length int
	Start  *Post
}

func (f *Feed) Append(newPost *Post)  {
	if f.Length == 0 {
		f.Start = newPost
	} else {
		currentPost := f.Start // setting the Starting element
		for currentPost.Next != nil { // traversing till the last element
			currentPost = currentPost.Next
		}
		//appending to the last element.
		currentPost.Next = newPost
	}
	f.Length++
}

