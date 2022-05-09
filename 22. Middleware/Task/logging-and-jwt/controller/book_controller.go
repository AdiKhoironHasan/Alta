package controller

import (
	"logging-and-jwt/config"
	"logging-and-jwt/model"
	"net/http"

	"github.com/labstack/echo"
)

func GetBooksController(c echo.Context) error {
	var books []model.Book

	err := config.DB.Find(&books).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    books,
	})
}

func CreateBookController(c echo.Context) error {
	book := model.Book{}

	c.Bind(&book)

	err := config.DB.Save(&book).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Book created",
		"data":    book,
	})
}

func GetBookController(c echo.Context) error {
	book := model.Book{}

	id := c.Param("id")
	err := config.DB.Where("id = ?", id).First(&book).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get Book by id",
		"data":    book,
	})
}

func DeleteBookController(c echo.Context) error {
	id := c.Param("id")

	err := config.DB.Delete(&model.Book{}, id).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete Book",
		"data":    nil,
	})
}

func UpdateBookController(c echo.Context) error {
	book := model.Book{}
	bookUpdate := model.Book{}
	id := c.Param("id")

	err := c.Bind(&bookUpdate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"data":    nil,
		})
	}

	res := config.DB.Where("id = ?", id).First(&book)
	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": res.Error,
			"data":    nil,
		})
	}

	// err := DB.Model(Book{}).Where("id = ?", id).Updates(BookUpdate).Error
	res = res.Updates(bookUpdate)
	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": res.Error,
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update Book",
		"data":    book,
	})
}
