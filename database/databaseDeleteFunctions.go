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

func DeleteAuction(c *gin.Context) {
	/*establish connection*/
	opt := options.Client().ApplyURI(databaseURI)
	client, err := mongo.NewClient(opt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not conect to database."})
		return
	}

	ctx := context.TODO()

	err = client.Connect(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not conect to database."})
		return
	}

	defer client.Disconnect(ctx)

	id := c.Param("id")
	objId, _ := primitive.ObjectIDFromHex(id)
	result, err := client.Database("auctionDB").Collection("auctions").DeleteOne(ctx, bson.M{"_id": objId}) //charge

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong with deletion"})
		return
	}

	if result.DeletedCount < 1 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Did not found particular auction", "id": id, "objId": objId})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Succesfully deleted the auction", "deletedID": objId})

}
