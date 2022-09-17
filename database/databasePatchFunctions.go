package database

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func PatchAuctionAsset(c *gin.Context) {
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

	auctionDB := client.Database("auctionDB")
	//convert asset id to obj id
	asset_id, err := primitive.ObjectIDFromHex(c.Param("assetid"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Did not find auction"})
		return
	}

	/*get the auction coresponding to auctionid */
	var auction Auction
	auction_id := c.Param("id")
	auction_objId, _ := primitive.ObjectIDFromHex(auction_id)

	err = client.Database("auctionDB").Collection("auctions").FindOne(ctx, bson.M{"_id": auction_objId}).Decode(&auction) //charge
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Did not find auction"})
		return
	}

	//add the new assetID to auction assetid table
	var table []primitive.ObjectID = auction.AssetIDs
	table = append(table, asset_id)
	//update the assetIDs table of the auction
	filter := bson.M{"_id": auction_objId}
	update := bson.M{"$set": bson.M{"_assetids": table}}
	result, err := auctionDB.Collection("auctions").UpdateOne(ctx, filter, update)
	fmt.Println(result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong with insertion"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Succesfully updated auction assets."})

}
