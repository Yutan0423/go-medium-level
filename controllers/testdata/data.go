package testdata

import "github.com/Yutan0423/go-medium-level/models"

var articleTestData = []models.Article{
	{
		ArticleID:   1,
		Title:       "firstPost",
		Contents:    "This is my first blog",
		UserName:    "saki",
		NiceNum:     2,
		CommentList: commentTestData,
	},
	{
		ArticleID: 2,
		Title:     "2nd",
		Contents:  "Second blog post",
		UserName:  "saki",
		NiceNum:   4,
	},
}

var commentTestData = []models.Comment{
	{
		CommentID: 1,
		ArticleID: 1,
		Message:   "1st comment yeah",
	},
	{
		CommentID: 2,
		ArticleID: 1,
		Message:   "welcome",
	},
}
