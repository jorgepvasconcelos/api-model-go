package endpoints

import (
	"api/src/services/book/bookService"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type ConsultBookInput struct {
	BookId int `form:"book_id" binding:"required"`
}

type ConsultBookOutput struct {
	BookId      int       `json:"book_id" binding:"required"`
	Isbn        string    `json:"isbn" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Author      string    `json:"author" binding:"required"`
	Publisher   string    `json:"publisher" binding:"required"`
	ReleaseDate time.Time `json:"release_date" binding:"required" time_format:"01/02/06"`
	Pages       int       `json:"pages" binding:"required"`
	Description string    `json:"description" binding:"required"`
}

// ConsultBook insere um usuário no banco de dados
func ConsultBook(c *gin.Context) {
	var params ConsultBookInput
	if err := c.ShouldBindQuery(&params); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	bookValue, err := bookService.BookService{}.FindBookById(params.BookId)
	if err != nil {
		log.Fatal(err)
	}

	response := ConsultBookOutput{
		BookId:      bookValue.BookId,
		Isbn:        bookValue.Isbn,
		Name:        bookValue.Name,
		Author:      bookValue.Author,
		Publisher:   bookValue.Publisher,
		ReleaseDate: bookValue.ReleaseDate,
		Pages:       bookValue.Pages,
		Description: bookValue.Description,
	}

	c.SecureJSON(http.StatusOK, response)
}
