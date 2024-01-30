package handler

import (

	//user defined packages

	"article/models"
	"article/repository"
	"article/utilities"
	"fmt"

	"strconv"

	//built-in packages

	"time"

	//third-party package

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Repo *repository.DbConnection
}

func Newhandler(db *repository.DbConnection) *Handler {
	return &Handler{Repo: db}
}

//-------------------------------Article-------------------------------------------------------------------------------

// Handler function to add articles
func (h *Handler) AddArticle(c *fiber.Ctx) error {
	article := models.Article{}

	logrusEntry := logrus.WithFields(logrus.Fields{
		"handler": "AddArticle",
	})

	logrusEntry.Info("Message : 'AddArticle-API called'")
	if err := c.BodyParser(&article); err != nil {
		logrusEntry.Errorf("Error : %s, Status : 400", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.Repo.AddArticle(&article); err != nil {
		logrusEntry.Errorf("Error :%s  Status : 400", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Publish message to Kafka
	go utilities.PublishToKafka(fmt.Sprintf("%v", article), "articles.create", article.ID)

	logrusEntry.Info("Message : 'Article added successfully' Status : 201")
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Created Article!", "ArticleID": article.ID})
}

// Handler function to listall the aticles in golang
func (h *Handler) ListAllArticles(c *fiber.Ctx) error {
	articles := []models.Article{}

	logrusEntry := logrus.WithFields(logrus.Fields{
		"handler": "ListAllArticles",
	})

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}

	perPage := 20
	offset := (page - 1) * perPage

	logrusEntry.Info("Message : 'ListAllArticles-API called'")
	if err := h.Repo.ListAllArticles(perPage, offset, &articles); err != nil {
		logrusEntry.Errorf("Error : %s Status : 404", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Publish message to Kafka
	for _, article := range articles {
		go utilities.PublishToKafka(fmt.Sprintf("%v", article), "articles.list-all-articlesss", article.ID)
	}

	logrusEntry.Info("Message : 'Article(s) retrieved successfully' Status : 200")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "retrived successfully", "articles": articles})
}

// Handler function to list the article with ID
func (h *Handler) ListArticleByID(c *fiber.Ctx) error {
	article := models.Article{}
	logrusEntry := logrus.WithFields(logrus.Fields{
		"handler": "ListArticlesByID",
	})

	logrusEntry.Info("Message : 'ListArticleByID-API called'")
	if err := h.Repo.ListArticleByID(c.Query("article_id"), &article); err != nil {
		logrusEntry.Errorf("Error : %s Status : 400", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	go utilities.PublishToKafka(fmt.Sprintf("%v", article), "articles.list-article-by-id", article.ID)

	logrusEntry.Info("Message : 'Article retrieved successfully' Status : 200")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"article": article})
}

// -------------------------------------Comments------------------------------------------------------------------------

// Handler function to add comments
func (h *Handler) AddComment(c *fiber.Ctx) error {
	comment := models.Comment{}
	logrusEntry := logrus.WithFields(logrus.Fields{
		"handler": "AddComment",
	})

	logrusEntry.Info("Message : 'AddComment-API called'")

	//parsing comment data from request body
	if err := c.BodyParser(&comment); err != nil {
		logrusEntry.Errorf("Error : %s Status : 400", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.Repo.AddComment(c.Query("article_id"), &comment); err != nil {
		logrusEntry.Errorf("Error : %s Status : 400", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	go utilities.PublishToKafka(fmt.Sprintf("%v", comment), "comment.create", comment.ID)

	logrusEntry.Info("Message : 'Added comment successfully' Status : 201")
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"Message": "Added comment successfully", "CommentID": comment.ID})
}

// Handler function to list all comments
func (h *Handler) LisAllComments(c *fiber.Ctx) error {
	comments := []models.Comment{}
	logrusEntry := logrus.WithFields(logrus.Fields{
		"handler": "ListAllComment",
	})

	logrusEntry.Info("Message : 'LisAllComments-API called'")

	if err := h.Repo.ListAllComments(c.Query("article_id"), &comments); err != nil {
		logrusEntry.Errorf("Error : %s Status : 400", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	for _, comment := range comments {
		go utilities.PublishToKafka(fmt.Sprintf("%v", comment), "articles.get-comments-by-post-id", comment.ID)
	}

	logrusEntry.Info("Message : 'Comment(s) retrieved successfully' Status : 200")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "retrived all the comments", "comments": comments})
}

// Handler fucntion to comment on comment
func (h *Handler) AddCommentOnComment(c *fiber.Ctx) error {
	reply := models.Reply{}

	logrusEntry := logrus.WithFields(logrus.Fields{
		"handler": "AddCommentOnComment",
	})

	logrusEntry.Info("Message : 'AddCommentOnComment-API called'")

	if err := c.BodyParser(&reply); err != nil {
		logrusEntry.Errorf("Error : %s Status : 400", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	reply.CreationDate = time.Now()

	if err := h.Repo.AddCommentOnComment(c.Query("comment_id"), &reply); err != nil {
		logrusEntry.Errorf("Error : %s Status : 400", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	go utilities.PublishToKafka(fmt.Sprintf("%v", reply), "articles.post-comment-on-comment", reply.ID)

	logrusEntry.Info("Message : 'Added comment(s) on comment' Status : 200")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "Added comment on comment successfully"})
}
