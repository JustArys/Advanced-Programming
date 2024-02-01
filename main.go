package main

import (
	"Music_Instrument_Shop/models"
	"fmt"
	"github.com/bxcodec/faker"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"math"
	"os"
	"strconv"
)

func main() {
	err := run()

	if err != nil {
		panic(err)
	}
}

func run() error {
	dsn := "host=localhost user=postgres password=Just_arys7 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}
	errr := db.AutoMigrate(&models.Book{})
	if errr != nil {
		return errr
	}

	app := fiber.New()

	app.Use(cors.New())

	app.Post("/api/products/populate", func(ctx *fiber.Ctx) error {
		for i := 0; i < 50; i++ {
			db.Create(&models.Book{
				Title:  faker.WORD,
				Author: faker.FirstName + " " + faker.LastName,
				Year:   faker.BaseDate,
			})
		}
		return ctx.JSON(fiber.Map{
			"message": "success",
		})
	})

	app.Get("api/books/frontend", func(ctx *fiber.Ctx) error {
		var books []models.Book

		sql := "SELECT * FROM books"

		if s := ctx.Query("s"); s != "" {
			sql = fmt.Sprintf("%s WHERE title LIKE '%%%s%%'", sql, s)
		}

		if sort := ctx.Query("sort"); sort != "" {
			sql += " ORDER BY " + sort
		}

		page, _ := strconv.Atoi(ctx.Query("page", "1"))
		perPage := 9
		var total int64

		db.Raw(sql).Count(&total)
		sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, perPage, (page-1)*perPage)

		db.Raw(sql).Scan(&books)

		return ctx.JSON(fiber.Map{
			"data":      books,
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(total / int64(page))),
		})
	})
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}
	app.Listen(":" + port)

	return nil
}
