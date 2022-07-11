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
	opt := options.Client().ApplyURI(databaseURI)
	client, err := mongo.NewClient(opt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong with database connection"})
		return
	}

	ctx := context.TODO()

	err = client.Connect(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong with database connection"})
		return
	}

	defer client.Disconnect(ctx)

	var auctions []Auction
	auctionDB := client.Database("auctionDB")
	auctionsCollection := auctionDB.Collection("auctions")

	results, err := auctionsCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Empty auction collection"})
		return
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
	opt := options.Client().ApplyURI(databaseURI)
	client, err := mongo.NewClient(opt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong with database connection"})
		return
	}

	ctx := context.TODO()

	err = client.Connect(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong with database connection"})
		return
	}

	defer client.Disconnect(ctx)

	var auction Auction
	id := c.Param("id")
	objId, _ := primitive.ObjectIDFromHex(id)

	err = client.Database("auctionDB").Collection("auctions").FindOne(ctx, bson.M{"_id": objId}).Decode(&auction) //charge
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Empty auction collection"})
		return
	}

	c.JSON(http.StatusOK, auction)

}

func GetBids(c *gin.Context) {
	/* establish connection to the database*/
	opt := options.Client().ApplyURI(databaseURI)
	client, err := mongo.NewClient(opt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong with database connection"})
		return
	}

	ctx := context.TODO()

	err = client.Connect(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong with database connection"})
		return
	}

	defer client.Disconnect(ctx)
	/*get the auction coreesponding to id */
	var auction Auction
	id := c.Param("id")
	objId, _ := primitive.ObjectIDFromHex(id)

	err = client.Database("auctionDB").Collection("auctions").FindOne(ctx, bson.M{"_id": objId}).Decode(&auction) //charge
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Did not find auction."})
		return
	}

	var bids []Bid
	auctionDB := client.Database("auctionDB")
	bidsCollection := auctionDB.Collection("bids")

	results, err := bidsCollection.Find(ctx, bson.M{"_auctionid": objId})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Did not find bids in this auction."})
		return
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleBid Bid
		if err = results.Decode(&singleBid); err != nil {
			c.JSON(http.StatusInternalServerError, nil)
		}

		bids = append(bids, singleBid)
	}

	c.JSON(http.StatusOK, bids)

}
