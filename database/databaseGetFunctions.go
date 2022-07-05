package database

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func GetAuctions(c *gin.Context) {
	/* establish connection to the database*/
	opt := options.Client().ApplyURI("mongodb://localhost:27017")
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

	var auctions []Auction
	auctionDB := client.Database("auctionDB")
	auctionsCollection := auctionDB.Collection("auctions")

	results, err := auctionsCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Empty auction collection"})
		panic(err)
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleAuction Auction
		if err = results.Decode(&singleAuction); err != nil {
			c.JSON(http.StatusInternalServerError, nil)
		}

		auctions = append(auctions, singleAuction)
	}

	c.JSON(http.StatusOK, auctions)

}

func GetAuction(c *gin.Context) {
	/* establish connection to the database*/
	opt := options.Client().ApplyURI("mongodb://localhost:27017")
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

	var auction Auction
	id := c.Param("id")
	objId, _ := primitive.ObjectIDFromHex(id)

	err = client.Database("auctionDB").Collection("auctions").FindOne(ctx, bson.M{"_id": objId}).Decode(&auction) //charge
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Empty auction collection"})
		panic(err)
	}

	c.JSON(http.StatusOK, auction)

}
