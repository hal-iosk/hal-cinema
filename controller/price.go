package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hal-iosk/hal-cinema/model"
)

func GetPrice(c *gin.Context) {
	var Prices []model.Price
	Price := model.Price{
		CustomerType: "一般ピーポ〜",
		Price:        1800,
	}
	Price.ID = 1
	Prices = append(Prices, Price)

	Price = model.Price{
		CustomerType: "パリピ",
		Price:        1500,
	}
	Price.ID = 2
	Prices = append(Prices, Price)

	Price = model.Price{
		CustomerType: "厨二",
		Price:        1000,
	}
	Price.ID = 3
	Prices = append(Prices, Price)

	Price = model.Price{
		CustomerType: "幼稚園児ニア",
		Price:        800,
	}
	Price.ID = 4
	Prices = append(Prices, Price)

	c.JSON(http.StatusOK, gin.H{
		"prices": Prices,
	})
}
