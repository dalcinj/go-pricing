package main

import (
	"fmt"
	"math/big"
	"net/http"

	"github.com/bojanz/currency"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

func main() {
	router := gin.Default()
	router.GET("calculate-pricing", calculatePricing)

	router.Run("localhost:3000")
}

func calculatePricing(context *gin.Context) {
	val1, convertDecimalErr := decimal.NewFromString(context.Query("val1"))
	val2, convertDecimalErr := decimal.NewFromString(context.Query("val2"))
	if convertDecimalErr != nil {
		panic(convertDecimalErr)
	}

	amount, _ := currency.NewAmount("275.98", "BRL")
	fmt.Println("amount: ", amount)
	fee, _ := currency.NewAmount("0.35", "BRL")
	fmt.Println("fee: ", fee)
	tax, mulErr := amount.Mul(fee.Number())
	if mulErr != nil {
		panic(mulErr)
	}
	fmt.Println("tax: ", tax.Number())

	a := big.NewInt(1199)
	b := big.NewInt(30)
	hundred := big.NewInt(100)
	c := new(big.Int)
	c.Mul(a, b)
	c.Div(c, hundred)
	fmt.Println("c: ", c)

	fmt.Println("val1: ", val1)
	fmt.Println("val2: ", val2)

	// fee, _ := decimal.NewFromString(".035")
	// taxRate, _ := decimal.NewFromString(".08875")

	// minOffer, _ := val1.Mul(fee).Round(2).Float64()
	// maxOffer, _ := val2.Mul(taxRate).Round(2).Float64()

	context.JSON(http.StatusOK, gin.H{
		"min_offer": "minOffer",
		"max_offer": "maxOffer",
	})
}
