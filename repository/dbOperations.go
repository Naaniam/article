package repository

import (

	// user defined package
	"article/models"
	"article/utilities"
	"time"

	// built-in packages
	"fmt"

	// third party package
	"gorm.io/gorm"
)

type DbConnection struct {
	DB *gorm.DB
	//Logger *log.Logger
}

func NewDbConnection(db *gorm.DB) *DbConnection {
	return &DbConnection{DB: db}
}

type Operations interface {
	AddArticle(post *models.Article) error
	ListAllArticles(perPage, offSet int, articles *[]models.Article) error
	ListArticleByID(articleID string, article *models.Article) error
	AddComment(articleID string, comment *models.Comment) error
	ListAllComments(articleID string, comments *[]models.Comment) error
	CommentOnComment(commentID string, reply *models.Reply) error
}

// Method to add article to the database
func (db *DbConnection) AddArticle(article *models.Article) error {

	//validating the article struct
	err := utilities.ValidateStruct(article)
	if err != nil {
		return err
	}

	//adding articles details to the article table
	if err := db.DB.Create(&article).Error; err != nil {
		return err
	}
	return nil
}

// Method to get all the articles from the database
func (db *DbConnection) ListAllArticles(perPage, offSet int, articles *[]models.Article) error {
	//retriving all the articles from the articles table
	if err := db.DB.Debug().Limit(perPage).Offset(offSet).Find(&articles).Error; err != nil {
		return err
	}

	//checking if the artilces in the articles table are empty o
	if len(*articles) == 0 {
		return fmt.Errorf("no articles yet!")
	}

	return nil
}

// Method to get the article by ID
func (db *DbConnection) ListArticleByID(articleID string, article *models.Article) error {

	//checking if the article id is empty string or not
	if articleID == "" {
		return fmt.Errorf("article ID can not be empty")
	}

	//retriving the article with the given artcile ID
	if err := db.DB.Debug().Preload("Comments.Replies").First(&article, articleID).Error; err != nil {
		return err
	}
	return nil
}

// ----------------------------------------------Comment---------------------------------------------------------------

// Method to add comment
func (db *DbConnection) AddComment(articleID string, comment *models.Comment) error {
	article := models.Article{}
	//checking if the article id is empty string or not
	if articleID == "" {
		return fmt.Errorf("articleID can not be empty")
	}

	//checking if the article with the given id is present in the articles table
	if err := db.DB.Debug().First(&article, articleID).Error; err != nil {
		return fmt.Errorf("Article not found")
	}

	//validating the comment struct
	err := utilities.ValidateStruct(comment)
	if err != nil {
		return err
	}

	comment.ArticleID = article.ID
	comment.CreationDate = time.Now()

	//adding the comment details to the comments table
	if err := db.DB.Create(&comment).Error; err != nil {
		return err
	}
	return nil
}

// Method to get all the comments from the database based on article id
func (db *DbConnection) ListAllComments(articleID string, comments *[]models.Comment) error {

	//checking if the article id is empty string or not
	if articleID == "" {
		return fmt.Errorf("articleID can not be empty")
	}

	//retriving all the comments fot the given article id from the comments table
	if err := db.DB.Debug().Where("article_id = ?", articleID).Preload("Replies").Find(&comments).Error; err != nil {
		return err
	}

	//checking if the comments in the comments are empty or not for the given post
	if len(*comments) == 0 {
		return fmt.Errorf("no comments yet")
	}

	return nil
}

// Method to add comment reply to the comment (comment on comment)
func (db *DbConnection) AddCommentOnComment(commentID string, reply *models.Reply) error {
	var parentComment models.Comment

	//checking if the comment id is empty string or not
	if commentID == "" {
		return fmt.Errorf("articleID can not be empty")
	}

	//check if the article is present or not in the article table
	if err := db.DB.Debug().First(&parentComment, commentID).Error; err != nil {
		return fmt.Errorf("comment ID not found")
	}

	reply.CommentID = parentComment.ID

	//validating the reply struct
	err := utilities.ValidateStruct(reply)
	if err != nil {
		return err
	}

	//adding the reply details to the replies table
	if err := db.DB.Create(&reply).Error; err != nil {
		return err
	}
	return nil
}
