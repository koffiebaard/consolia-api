package models

import "time"


type Comic struct {
	ID        	uint `gorm:"id"`
	Slug 		string
	Title		string
	Sublog		string
	Tooltip		string
	Image		string
	Thumbnail	string
	ComicWidth	int
	ComicHeight	int
	PostOn 		time.Time
	PostedOn 	time.Time
	Enabled		bool
	RedditUrl	string
	RedditUpvotes	int
	TumblrUrl	string
	TumblrNotes	int
	//NineGagUrl	string `gorm:"9gag_url"`
	//NineGagUpvotes	int `gorm:"9gag_upvotes"`
	FacebookLikes	int
	CheezUrl	string `gorm:"cheez_url"`
	CheezUpvotes	int `gorm:"cheez_upvotes"`
}

func (c *Comic) Test () {

}