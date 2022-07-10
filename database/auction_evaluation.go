package database

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetWinningBid(c *gin.Context) {
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
	id := c.Param("id")
	objId, _ := primitive.ObjectIDFromHex(id)

	auctionDB := client.Database("auctionDB")
	bidsCollection := auctionDB.Collection("bids")

	results, err := bidsCollection.Find(ctx, bson.M{"_auctionid": objId})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Did not find bids in this auction."})
		return
	}
	fmt.Println(results)
	defer results.Close(ctx)
	//place the last one as winning considering we will check price in order to accept bids
	var winnerBid Bid
	var bids []Bid
	for results.Next(ctx) {
		var currentBid Bid
		if err = results.Decode(&currentBid); err != nil {
			c.JSON(http.StatusInternalServerError, nil)
		}
		bids = append(bids, currentBid)

		// the first iteration of for
		if winnerBid == (Bid{}) && currentBid.Accepted {
			copier.Copy(&winnerBid, &currentBid)
		} else if winnerBid == (Bid{}) && !currentBid.Accepted {
			continue
		}

		if currentBid.Timestamp.After(winnerBid.Timestamp) && currentBid.Accepted {
			//make a deep copy using Copier
			copier.Copy(&winnerBid, &currentBid)
		}
	}

	if winnerBid == (Bid{}) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Did not find valid bids in this auction."})
	} else {
		c.JSON(http.StatusOK, winnerBid)
	}

}
