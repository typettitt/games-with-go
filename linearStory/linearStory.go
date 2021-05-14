package main

import "fmt"

/*implementation of linkedlist */

type storyPage struct {
	text     string
	nextPage *storyPage
}

func (page *storyPage) playStory() {
	for page != nil {
		fmt.Println(page.text)
		page = page.nextPage
	}
}

func (page *storyPage) addToEnd(text string) {
	pageToAdd := &storyPage{text, nil}
	for page.nextPage != nil {
		page = page.nextPage
	}
	page.nextPage = pageToAdd
}

func (page *storyPage) addAfter(text string) {
	newPage := &storyPage{text, page.nextPage}
	page.nextPage = newPage
}

func main() {

	page1 := storyPage{"You are standing in a field west of white house.", nil}
	page1.addToEnd("You climb into the attic, it is pitch black, you can't see a thing!")
	page1.addToEnd("You are eaten by a Grue")
	page1.addAfter("Testing add after")

	page1.playStory()

	//Functions - has a return value
	//Procedures - does not have a return value, just executes and does things
	//Methods - functions that are attached to struct or object

}
