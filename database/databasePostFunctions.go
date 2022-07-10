package database

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func PostAuction(c *gin.Context) {
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
	auctionsCollection := auctionDB.Collection("auctions")

	/*create the new auction and insert it in the database*/

	var newAuction Auction

	//validate the request body
	if err := c.BindJSON(&newAuction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Wrong request body."})
		return
	}

	newAuction.ID = primitive.NewObjectID()

	insertResult, err := auctionsCollection.InsertOne(ctx, newAuction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong with insertion"})
		return
	}
	fmt.Println("Inserted Auction:", insertResult.InsertedID)
	c.JSON(http.StatusCreated, gin.H{"message": "Succesfully added auction."})

}
func PostProduct(c *gin.Context) {
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
	productsCollection := auctionDB.Collection("products")

	/*create the new auction and insert it in the database*/

	var newProduct Product

	//validate the request body
	if err := c.BindJSON(&newProduct); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Wrong request body."})
		return
	}

	newProduct.ID = primitive.NewObjectID()

	insertResult, err := productsCollection.InsertOne(ctx, newProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong with insertion"})
		return
	}
	fmt.Println("Inserted Auction:", insertResult.InsertedID)
	c.JSON(http.StatusCreated, gin.H{"message": "Succesfully added product."})

}
func PostAsset(c *gin.Context) {
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
	assetsCollection := auctionDB.Collection("assets")

	/*get the product coresponding to productid */
	var product Product
	product_id := c.Param("productid")
	product_objId, _ := primitive.ObjectIDFromHex(product_id)

	err = client.Database("auctionDB").Collection("products").FindOne(ctx, bson.M{"_id": product_objId}).Decode(&product) //charge
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Did not find product"})
		return
	}

	/*get the auction coresponding to auxtionid */
	var auction Auction
	auction_id := c.Param("id")
	auction_objId, _ := primitive.ObjectIDFromHex(auction_id)

	err = client.Database("auctionDB").Collection("auctions").FindOne(ctx, bson.M{"_id": auction_objId}).Decode(&auction) //charge
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Did not find auction"})
		return
	}

	var newAsset Asset

	//validate the request body
	if err := c.BindJSON(&newAsset); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Wrong request body."})
		return
	}

	//make the new Asset
	newAsset.ID = primitive.NewObjectID()
	newAsset.ProductID = product.ID

	//add the new assetID to auction assetid table
	var table []primitive.ObjectID = auction.AssetIDs
	table = append(table, newAsset.ID)
	//update the assetIDs table of the auction
	filter := bson.M{"_id": auction_objId}
	update := bson.M{"$set": bson.M{"_assetids": table}}
	result, err := auctionDB.Collection("auctions").UpdateOne(ctx, filter, update)
	fmt.Println(result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong with insertion"})
		return
	}

	insertResult, err := assetsCollection.InsertOne(ctx, newAsset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong with insertion"})
		return
	}
	fmt.Println("Inserted Asset:", insertResult.InsertedID)
	c.JSON(http.StatusCreated, gin.H{"message": "Succesfully added asset."})

}

func PostBid(c *gin.Context) {
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
	bidsCollection := auctionDB.Collection("bids")

	/*get the auction corresponding to id */
	var auction Auction
	id := c.Param("id")
	objId, _ := primitive.ObjectIDFromHex(id)

	err = client.Database("auctionDB").Collection("auctions").FindOne(ctx, bson.M{"_id": objId}).Decode(&auction) //charge
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Did not find auction."})
		return
	}

	var newBid Bid

	//validate the request body
	if err := c.BindJSON(&newBid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Wrong request body."})
		return
	}

	//place the correct auctionID in the new bid
	newBid.ID = primitive.NewObjectID()
	newBid.AuctionID = auction.ID
	newBid.Timestamp = time.Now()

	//different determination for diferent auction/bid types
	switch newBid.Type {
	case 2:
		//get all bids to check the last accepted vs the new one
		results, err := bidsCollection.Find(ctx, bson.M{"_auctionid": objId})
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "Did not find bids in this auction."})
			return
		}
		defer results.Close(ctx)
		var lastBid Bid
		var bids []Bid
		for results.Next(ctx) {
			var currentBid Bid
			if err = results.Decode(&currentBid); err != nil {
				c.JSON(http.StatusInternalServerError, nil)
			}
			bids = append(bids, currentBid)

			// the first iteration of for
			if lastBid == (Bid{}) && currentBid.Accepted {
				copier.Copy(&lastBid, &currentBid)
			} else if lastBid == (Bid{}) && !currentBid.Accepted {
				continue
			}

			if currentBid.Timestamp.After(lastBid.Timestamp) && currentBid.Accepted {
				//make a deep copy using Copier
				copier.Copy(&lastBid, &currentBid)
			}
		}

		//accept or reject bid
		if lastBid == (Bid{}) || (newBid.Timestamp.After(lastBid.Timestamp) && newBid.Price > lastBid.Price) {
			newBid.Accepted = true
		} else {
			newBid.Accepted = false
		}
	}

	insertResult, err := bidsCollection.InsertOne(ctx, newBid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong with insertion"})
		return
	}
	fmt.Println("Inserted Auction:", insertResult.InsertedID)
	c.JSON(http.StatusCreated, gin.H{"message": "Succesfully added bid."})

}
