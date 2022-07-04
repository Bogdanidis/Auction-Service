package main

import (
	"example/Auction-Service/database"
)

func main() {
	database.PingDatabase()
	//database.CreateDatabase()
	//database.DropDatabase()

	/*
		router := gin.Default()
		router.GET("/auctions", database.GetAuctions)/
		router.GET("/auctions/:id/results", database.GetAuctionResutls(id))
		router.POST("/auctions", database.PostAuction)
		router.POST("/products", database.PostProduct)
		router.POST("/assets", database.PostAsset)
		router.POST("/auctions/:id/bids", database.PostBid(id))

		router.Run("localhost:8080")
	*/
}
