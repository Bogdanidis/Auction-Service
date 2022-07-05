package database

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func PostAuction(c *gin.Context) {
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

	auctionDB := client.Database("auctionDB")
	auctionsCollection := auctionDB.Collection("auctions")

	/*create the new auction and insert it in the database*/

	var newAuction Auction

	//validate the request body
	if err := c.BindJSON(&newAuction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Wrong request body."})
		return
	}
	//validate the required fields
	if validationErr := validator.New().Struct(&newAuction); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Request body did not validate "})
		return
	}

	newAuction.ID = primitive.NewObjectID()

	insertResult, err := auctionsCollection.InsertOne(ctx, newAuction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong with insertion"})
		panic(err)
	}
	fmt.Println("Inserted Auction:", insertResult.InsertedID)
	c.JSON(http.StatusOK, gin.H{"message": "Succesfully added auction."})

}
func PostProduct(c *gin.Context) {
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

	auctionDB := client.Database("auctionDB")
	productsCollection := auctionDB.Collection("products")

	/*create the new auction and insert it in the database*/

	var newProduct Product

	//validate the request body
	if err := c.BindJSON(&newProduct); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Wrong request body."})
		return
	}
	//validate the required fields
	if validationErr := validator.New().Struct(&newProduct); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Request body did not validate "})
		return
	}

	newProduct.ID = primitive.NewObjectID()

	insertResult, err := productsCollection.InsertOne(ctx, newProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong with insertion"})
		panic(err)
	}
	fmt.Println("Inserted Auction:", insertResult.InsertedID)
	c.JSON(http.StatusOK, gin.H{"message": "Succesfully added product."})

}
func PostAsset(c *gin.Context) {
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

	auctionDB := client.Database("auctionDB")
	assetsCollection := auctionDB.Collection("assets")

	/*create the new auction and insert it in the database*/

	var newAsset Asset

	//validate the request body
	if err := c.BindJSON(&newAsset); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Wrong request body."})
		return
	}
	//validate the required fields
	if validationErr := validator.New().Struct(&newAsset); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Request body did not validate "})
		return
	}

	newAsset.ID = primitive.NewObjectID()

	insertResult, err := assetsCollection.InsertOne(ctx, newAsset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong with insertion"})
		panic(err)
	}
	fmt.Println("Inserted Auction:", insertResult.InsertedID)
	c.JSON(http.StatusOK, gin.H{"message": "Succesfully added asset."})

}
func PostBid(c *gin.Context) {
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

	auctionDB := client.Database("auctionDB")
	bidsCollection := auctionDB.Collection("bids")

	/*get the auction coreesponding to id */
	var auction Auction
	id := c.Param("id")
	objId, _ := primitive.ObjectIDFromHex(id)

	err = client.Database("auctionDB").Collection("auctions").FindOne(ctx, bson.M{"_id": objId}).Decode(&auction) //charge
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Empty auction collection"})
		panic(err)
	}

	var newBid Bid

	//validate the request body
	if err := c.BindJSON(&newBid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Wrong request body."})
		return
	}
	//validate the required fields
	if validationErr := validator.New().Struct(&newBid); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Request body did not validate "})
		return
	}

	newBid.ID = primitive.NewObjectID()
	//place the correct auctionID in the bid
	newBid.AuctionID = auction.ID

	insertResult, err := bidsCollection.InsertOne(ctx, newBid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong with insertion"})
		panic(err)
	}
	fmt.Println("Inserted Auction:", insertResult.InsertedID)
	c.JSON(http.StatusOK, gin.H{"message": "Succesfully added bid."})

}

