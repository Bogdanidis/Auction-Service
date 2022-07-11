package database

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var databaseURI string = "mongodb://localhost:27017"

type Trade struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	AuctionID primitive.ObjectID `bson:"_auctionid,omitempty"`
	BidderID  primitive.ObjectID `bson:"_bidderid,omitempty"`
	Volume    float64            `bson:"volume,omitempty"`
	Price     float64            `bson:"price,omitempty"`
	Fees      float64            `bson:"fees,omitempty"`
	Total     float64            `bson:"total,omitempty"`
}

//Bid-Auction Type Enum
type Type int

const (
	AscendingPrice Type = iota
	AscendingClock
	UniformPrice
	PayAsBid
)

type Bid struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"_userid,omitempty"`
	AuctionID primitive.ObjectID `bson:"_auctionid,omitempty"`
	Price     float64            `bson:"price,omitempty"`
	Accepted  bool               `bson:"accepted,omitempty"`
	Timestamp time.Time          `bson:"timestamp,omitempty"`
	Type      Type               `bson:"type,omitempty"`
}

type Auction struct {
	ID              primitive.ObjectID   `bson:"_id,omitempty"`
	AssetIDs        []primitive.ObjectID `bson:"_assetids,omitempty"`
	SellerID        primitive.ObjectID   `bson:"_sellerid,omitempty"`
	Charge          float64              `bson:"charge,omitempty"`
	StartingTime    time.Time            `bson:"startingtime,omitempty"`
	EndingTimeStart time.Time            `bson:"endingtimestart,omitempty"`
	EndingTimeEnd   time.Time            `bson:"endingtimeend,omitempty"`
	Type            Type                 `bson:"type,omitempty"`
	Active          bool                 `bson:"active,omitempty"`
}

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Description string             `bson:"description,omitempty"`
	Images      []string           `bson:"images,omitempty"`
}
type ShippingCost struct {
	ShippingTo string
	Cost       float64
}
type Asset struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	ProductID     primitive.ObjectID `bson:"_productid,omitempty"`
	Volume        float64            `bson:"volume,omitempty"`
	ShippingFrom  string             `bson:"shippingfrom,omitempty"`
	ShippingCosts []ShippingCost     `bson:"shippingcosts,omitempty"`
}

type Blacklist struct {
	ID        primitive.ObjectID   `bson:"_tid,omitempty"`
	UserIDs   []primitive.ObjectID `bson:"_userids,omitempty"`
	AuctionID primitive.ObjectID   `bson:"_auctionid,omitempty"`
}

type Seller struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name,omitempty"`
	///Logo        string             `bson:"logo,omitempty"`
	//Credentials string             `bson:"credentials,omitempty"`
}

type SellerUser struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	SellerID primitive.ObjectID `bson:"_sellerid,omitempty"`
	Name     string             `bson:"name,omitempty"`
}

type Bidder struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Email       string             `bson:"email,omitempty"`
	Location    string             `bson:"location,omitempty"`
	TotalBudget float64            `bson:"totalbudget,omitempty"`
	Budget      float64            `bson:"budget,omitempty"`
	Reserved    float64            `bson:"reserved,omitempty"`
	//Credentials string             `bson:"credentials,omitempty"`
}
