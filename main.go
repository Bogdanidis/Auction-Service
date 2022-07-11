package main

import (
	"example/Auction-Service/database"

	"github.com/gin-gonic/gin"
)

var serverURI = "localhost:8080"

/*
				TODOs
-assets are only linked to auctions right?
-(make get winning-bid request a service?)
-(do we need bid validation service?)
-change post asset to not connect asset with auction
-enrich yaml
-check heroku or render.com for deployment
-parameterize and enrich testing
-set ap auction ending time and service that terminates auction and declares winner, post trade
-configure bidder, seller ... how to distinguish request sources, block features depending on user type
-enrich get auctions request, search filters etc

*/
func main() {
	database.CreateDatabase()

	router := gin.Default()

	router.GET("/auctions", database.GetAuctions)
	router.GET("/auctions/:id", database.GetAuction)
	router.GET("/auctions/:id/bids", database.GetBids)
	router.GET("/auctions/:id/winning-bid", database.GetWinningBid)

	router.POST("/auctions", database.PostAuction)
	router.POST("/products", database.PostProduct)
	//router.POST("/auctions/:id/products/:productid/assets", database.PostAsset)
	router.POST("/products/:id/assets", database.PostAsset)
	//asociates asset with an auction
	router.PATCH("/auctions/:id/assets/:assetid", database.PostAuctionAsset)
	router.POST("/auctions/:id/bids", database.PostBid)

	router.DELETE("/database", database.DeleteDatabase)
	router.DELETE("/auctions/:id", database.DeleteAuction)

	router.Run(serverURI)

}
