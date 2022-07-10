package database

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateDatabase() {
	opt := options.Client().ApplyURI(databaseURI)
	client, err := mongo.NewClient(opt)
	if err != nil {
		panic(err)
	}

	ctx := context.TODO()

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	defer client.Disconnect(ctx)

	auctionDB := client.Database("auctionDB")
	auctionDB.Collection("trades")
	auctionDB.Collection("bids")
	auctionDB.Collection("auctions")
	auctionDB.Collection("products")
	auctionDB.Collection("assets")
	auctionDB.Collection("blacklists")
	auctionDB.Collection("sellers")
	auctionDB.Collection("sellerUsers")
	auctionDB.Collection("bidders")

}
func DeleteDatabase(c *gin.Context) {
	/*establish connection*/
	opt := options.Client().ApplyURI(databaseURI)
	client, err := mongo.NewClient(opt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong with database connection"})
		panic(err)
	}

	ctx := context.TODO()

	err = client.Connect(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong with database connection"})
		panic(err)
	}

	defer client.Disconnect(ctx)

	client.Database("auctionDB").Collection("trades").Drop(ctx)
	client.Database("auctionDB").Collection("bids").Drop(ctx)
	client.Database("auctionDB").Collection("auctions").Drop(ctx)
	client.Database("auctionDB").Collection("products").Drop(ctx)
	client.Database("auctionDB").Collection("assets").Drop(ctx)
	client.Database("auctionDB").Collection("blacklists").Drop(ctx)
	client.Database("auctionDB").Collection("sellers").Drop(ctx)
	client.Database("auctionDB").Collection("sellerUsers").Drop(ctx)
	client.Database("auctionDB").Collection("bidders").Drop(ctx)

	c.JSON(http.StatusOK, nil)

}
