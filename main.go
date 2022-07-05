package main

import (
	"example/Auction-Service/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.CreateDatabase()

	router := gin.Default()

	router.GET("/auctions", database.GetAuctions)
	router.GET("/auctions/:id", database.GetAuction)
	//router.GET("/auctions/:id/results", database.GetAuctionResutls(id))

	router.POST("/auctions", database.PostAuction)
	router.POST("/products", database.PostProduct)
	router.POST("/assets", database.PostAsset)
	router.POST("/auctions/:id/bids", database.PostBid)

	router.DELETE("/database", database.DeleteDatabase)
	router.DELETE("/auctions/:id", database.DeleteAuction)

	router.Run("localhost:8080")

}
