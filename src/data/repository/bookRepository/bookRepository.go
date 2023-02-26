package bookRepository

import (
	"api/src/data/db_orm/sessions"
	"api/src/data/db_orm/tables"
	"api/src/data/dto"
	"api/src/data/errors/sqlError"
	"api/src/data/redisClient"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log"
)

func FindBookById(bookId int) (dto.BookDTO, sqlError.SqlError) {
	client := redisClient.GetRedisClient()
	defer client.Close()
	cacheKey := fmt.Sprintf("book-id-%d", bookId)

	if cacheResponse, err := client.Get(redisClient.Ctx, cacheKey).Bytes(); err != nil {
		response := dto.BookDTO{}
		json.Unmarshal(cacheResponse, &response)
		return response, nil
	}

	session, _ := sessions.OpenSession()
	var tblBook tables.TblBooks
	result := session.First(&tblBook, "id = ?", rune(bookId))

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return dto.BookDTO{}, sqlError.NotFound
	}

	response := dto.BookDTO{
		Id:          int(tblBook.ID),
		Isbn:        tblBook.Isbn,
		Name:        tblBook.Name,
		Author:      tblBook.Author,
		Publisher:   tblBook.Publisher,
		ReleaseDate: tblBook.ReleaseDate,
		Pages:       tblBook.Pages,
		Description: tblBook.Description,
	}

	serializedResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Set(redisClient.Ctx, cacheKey, serializedResponse, redis.KeepTTL).Err()
	if err != nil {
		log.Fatal(err)
	}

	return response, nil
}
