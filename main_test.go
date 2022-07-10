package main

import (
	"bytes"
	"encoding/json"
	"example/Auction-Service/database"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestPostAuction(t *testing.T) {
	router := SetUpRouter()
	router.POST("/auctions", database.PostAuction)
	auction := database.Auction{
		Charge: 123,
	}

	jsonValue, _ := json.Marshal(auction)
	req, _ := http.NewRequest("POST", "/auctions", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}
func TestPostAsset(t *testing.T) {
	router := SetUpRouter()
	router.POST("/auctions/:id/products/:productid/assets", database.PostAsset)
	asset := database.Asset{
		Volume:       123123,
		ShippingFrom: "Heraklion",
	}
	auctionId := "62c3b3350ff66c5091bf94bc"
	productId := "62c3ba6edbbe9579da8c2984"

	jsonValue, _ := json.Marshal(asset)
	req, _ := http.NewRequest("POST", "/auctions/"+auctionId+"/products/"+productId+"/assets", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
	//wrong auction id
	bad_req, _ := http.NewRequest("POST", "/auctions/1/products/"+productId+"/assets", bytes.NewBuffer(jsonValue))
	w = httptest.NewRecorder()
	router.ServeHTTP(w, bad_req)
	assert.Equal(t, http.StatusNotFound, w.Code)
	//wrong product id
	bad_req, _ = http.NewRequest("POST", "/auctions/"+auctionId+"/products/1/assets", bytes.NewBuffer(jsonValue))
	w = httptest.NewRecorder()
	router.ServeHTTP(w, bad_req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
func TestPostBid(t *testing.T) {
	router := SetUpRouter()
	router.POST("/auctions/:id/bids", database.PostBid)
	bid := database.Bid{
		Price: 5,
		Type:  2,
	}

	auctionId := "62c3b3350ff66c5091bf94bc"

	jsonValue, _ := json.Marshal(bid)
	req, _ := http.NewRequest("POST", "/auctions/"+auctionId+"/bids", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	bad_req, _ := http.NewRequest("POST", "/auctions/1/bids", bytes.NewBuffer(jsonValue))
	w = httptest.NewRecorder()
	router.ServeHTTP(w, bad_req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
func TestPostProduct(t *testing.T) {
	router := SetUpRouter()
	router.POST("/products", database.PostProduct)
	product := database.Product{
		Name: "Test_product2",
	}

	jsonValue, _ := json.Marshal(product)
	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}
func TestGetAuctions(t *testing.T) {
	router := SetUpRouter()
	router.GET("/auctions", database.GetAuctions)
	req, _ := http.NewRequest("GET", "/auctions", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var auctions []database.Auction
	json.Unmarshal(w.Body.Bytes(), &auctions)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, auctions)

}
func TestGetAuction(t *testing.T) {
	router := SetUpRouter()
	auctionId := "62c27e494d3dccdad8798f62"
	router.GET("/auctions/:id", database.GetAuction)
	req, _ := http.NewRequest("GET", "/auctions/"+auctionId, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var auction database.Auction
	json.Unmarshal(w.Body.Bytes(), &auction)
	fmt.Println(auction)
	assert.Equal(t, http.StatusOK, w.Code)

	//convert object Id to string
	IDstring := auction.ID.Hex()
	assert.Equal(t, auctionId, IDstring)

	reqNotFound, _ := http.NewRequest("GET", "/auctions/1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, reqNotFound)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetWinningBid(t *testing.T) {
	router := SetUpRouter()
	auctionId := "62c3b3350ff66c5091bf94bc"
	router.GET("/auctions/:id/winning-bid", database.GetWinningBid)
	req, _ := http.NewRequest("GET", "/auctions/"+auctionId+"/winning-bid", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var bid database.Bid
	json.Unmarshal(w.Body.Bytes(), &bid)
	assert.Equal(t, http.StatusOK, w.Code)

	reqNotFound, _ := http.NewRequest("GET", "/auctions/1/winning-bid", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, reqNotFound)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

/*
func TestPostDeleteAuction(t *testing.T) {
	router := SetUpRouter()
	auctionId := "62c836c634f9f5c766cd653c"
	router.DELETE("/auctions/:id", database.DeleteAuction)
	req, _ := http.NewRequest("DELETE", "/auctions/"+auctionId, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	reqNotFound, _ := http.NewRequest("DELETE", "/auctions/1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, reqNotFound)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
*/
