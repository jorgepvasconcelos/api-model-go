package endpoints

import (
	"api/src/presentation/schemas"
	"api/src/services/book/bookError"
	"api/src/services/book/bookService"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ConsultBook get a book from database
func ConsultBook(c *gin.Context) {
	var params schemas.ConsultBookInput
	if err := c.ShouldBindQuery(&params); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	bookValue, err := bookService.FindBookById(params.BookId)
	if err != nil {
		if errors.Is(err, bookError.NotFound) {
			c.SecureJSON(http.StatusNotFound, schemas.Message{Message: "Book not found"})
			return
		}
	}

	response := schemas.ConsultBookOutput{
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
