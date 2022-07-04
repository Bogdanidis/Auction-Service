package database

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func DropDatabase() {
	opt := options.Client().ApplyURI("mongodb://localhost:27017")
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
	tradesCollection := auctionDB.Collection("trades")
	bidsCollection := auctionDB.Collection("bids")
	auctionsCollection := auctionDB.Collection("auctions")
	productsCollection := auctionDB.Collection("products")
	assetsCollection := auctionDB.Collection("assets")
	blacklistsCollection := auctionDB.Collection("blacklists")
	sellersCollection := auctionDB.Collection("sellers")
	sellerUsersCollection := auctionDB.Collection("sellerUsers")
	biddersCollection := auctionDB.Collection("bidders")

	defer tradesCollection.Drop(ctx)
	defer bidsCollection.Drop(ctx)
	defer auctionsCollection.Drop(ctx)
	defer productsCollection.Drop(ctx)
	defer assetsCollection.Drop(ctx)
	defer blacklistsCollection.Drop(ctx)
	defer sellersCollection.Drop(ctx)
	defer sellerUsersCollection.Drop(ctx)
	defer biddersCollection.Drop(ctx)
}

func PingDatabase() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalln(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
func CreateDatabaseOld() {
	opt := options.Client().ApplyURI("mongodb://localhost:27017")
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
	tradesCollection := auctionDB.Collection("trades")
	bidsCollection := auctionDB.Collection("bids")
	auctionsCollection := auctionDB.Collection("auctions")
	productsCollection := auctionDB.Collection("products")
	assetsCollection := auctionDB.Collection("assets")
	blacklistsCollection := auctionDB.Collection("blacklists")
	sellersCollection := auctionDB.Collection("sellers")
	sellerUsersCollection := auctionDB.Collection("sellerUsers")
	biddersCollection := auctionDB.Collection("bidders")

	fmt.Println(tradesCollection)
	fmt.Println(bidsCollection)
	fmt.Println(auctionsCollection)
	fmt.Println(productsCollection)
	fmt.Println(assetsCollection)
	fmt.Println(blacklistsCollection)
	fmt.Println(sellersCollection)
	fmt.Println(sellerUsersCollection)
	fmt.Println(biddersCollection)

}
func CreateDatabase() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	auctionDB := client.Database("auctionDB")
	tradesCollection := auctionDB.Collection("trades")
	bidsCollection := auctionDB.Collection("bids")
	auctionsCollection := auctionDB.Collection("auctions")
	productsCollection := auctionDB.Collection("products")
	assetsCollection := auctionDB.Collection("assets")
	blacklistsCollection := auctionDB.Collection("blacklists")
	sellersCollection := auctionDB.Collection("sellers")
	sellerUsersCollection := auctionDB.Collection("sellerUsers")
	biddersCollection := auctionDB.Collection("bidders")

	fmt.Println(tradesCollection)
	fmt.Println(bidsCollection)
	fmt.Println(auctionsCollection)
	fmt.Println(productsCollection)
	fmt.Println(assetsCollection)
	fmt.Println(blacklistsCollection)
	fmt.Println(sellersCollection)
	fmt.Println(sellerUsersCollection)
	fmt.Println(biddersCollection)

}
func GetAuctions(c *gin.Context) {
	var auctions = []Auction{
		{Charge: 1.7}, {Charge: 1.5},
	}
	c.IndentedJSON(http.StatusOK, auctions)
}
func queryDatabase() {

}
