package main

import (
	"fmt"
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/schema"
)

type PricingInput struct {
	ValueOne int64 `schema:"value_one"`
	ValueTwo int64 `schema:"value_two"`
}

var decoder = schema.NewDecoder()

func main() {
	router := gin.Default()
	router.GET("calculate-pricing", calculatePricing)

	router.Run("localhost:3000")
}

func calculatePricing(context *gin.Context) {

	var pricingInput PricingInput
	err := decoder.Decode(&pricingInput, context.Request.URL.Query())
	if err != nil {
		fmt.Println("Error in GET parameters : ", err)
	} else {
		fmt.Println("GET parameters : ", pricingInput)
		fmt.Println("ValueOne: ", pricingInput.ValueOne)
		fmt.Println("ValueTwo: ", pricingInput.ValueTwo)
	}

	minOffer := pricingInput.ValueOne + pricingInput.ValueTwo
	maxOffer := pricingInput.ValueOne * pricingInput.ValueTwo

	n := new(big.Int)
	n, ok := n.SetString("10", 10)
	if !ok {
		fmt.Println("SetString: error")
		return
	}
	fmt.Println("n: ", n)

	a := big.NewInt(1199)
	b := big.NewInt(30)
	hundred := big.NewInt(100)
	c := new(big.Int)
	c.Mul(a, b)
	c.Div(c, hundred)
	fmt.Println("c: ", c)

	context.JSON(http.StatusOK, gin.H{
		"min_offer": minOffer,
		"max_offer": maxOffer,
	})
}
