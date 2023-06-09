package main

import (
	"github.com/gin-gonic/gin"
)

type Snaki struct {
	Sugar     float64 `json:"sugar"`
	Tapioca   float64 `json:"tapioca"`
	Chocolate float64 `json:"chocolate"`
	Packaging float64 `json:"packaging"`
}

type Cost struct {
	Sugar     float64 `json:"sugar"`
	Tapioca   float64 `json:"tapioca"`
	Chocolate float64 `json:"chocolate"`
	Packaging float64 `json:"packaging"`
}

func CalculateProductionCost(snaki Snaki, cost Cost) float64 {

	sugarKg := snaki.Sugar / 1000
	tapiocaKg := snaki.Tapioca / 1000
	chocolateKg := snaki.Chocolate / 1000

	productionCost := (sugarKg * cost.Sugar) +
		(tapiocaKg * cost.Tapioca) +
		(chocolateKg * cost.Chocolate) +
		(snaki.Packaging * cost.Packaging)

	return productionCost
}

type ProductionCostResponse struct {
	ProductionCost float64 `json:"production_cost"`
}

func CalculateProductionCostHandler(c *gin.Context) {
	var request struct {
		Snaki Snaki `json:"snaki"`
		Cost  Cost  `json:"cost"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	snaki := request.Snaki
	cost := request.Cost

	productionCost := CalculateProductionCost(snaki, cost)
	response := ProductionCostResponse{
		ProductionCost: productionCost,
	}

	c.JSON(200, response)
}

func main() {

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	r.POST("/production-cost", CalculateProductionCostHandler)
	r.Run(":8080")

}
