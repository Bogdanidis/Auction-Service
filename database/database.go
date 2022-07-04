package database

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Trade struct {
	TradeID   primitive.ObjectID `bson:"_tradeid,omitempty"`
	AuctionID primitive.ObjectID `bson:"_auctionid,omitempty"`
	BidderID  primitive.ObjectID `bson:"_bidderid,omitempty"`
	Volume    float32            `bson:"volume,omitempty"`
	Price     float32            `bson:"price,omitempty"`
	Fees      float32            `bson:"fees,omitempty"`
	Total     float32            `bson:"total,omitempty"`
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
	BidID     primitive.ObjectID `bson:"_bidid,omitempty"`
	UserID    primitive.ObjectID `bson:"_userid,omitempty"`
	AuctionID primitive.ObjectID `bson:"_auctionid,omitempty"`
	Timestamp time.Time          `bson:"timestamp,omitempty"`
	Type      Type               `bson:"type,omitempty"`
}

type Auction struct {
	AuctionID       primitive.ObjectID   `bson:"_auctionid,omitempty"`
	AssetIDs        []primitive.ObjectID `bson:"_assetids,omitempty"`
	SellerID        primitive.ObjectID   `bson:"_sellerid,omitempty"`
	Charge          float32              `bson:"charge,omitempty"`
	StartingTime    time.Time            `bson:"startingtime,omitempty"`
	EndingTimeStart time.Time            `bson:"endingtimestart,omitempty"`
	EndingTimeEnd   time.Time            `bson:"endingtimeend,omitempty"`
	Type            Type                 `bson:"type,omitempty"`
}

type Product struct {
	ProductID   primitive.ObjectID `bson:"_productid,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Description string             `bson:"description,omitempty"`
	Images      []string           `bson:"images,omitempty"`
}

type Asset struct {
	AssetID       primitive.ObjectID `bson:"_assetid,omitempty"`
	ProductID     primitive.ObjectID `bson:"_productid,omitempty"`
	Volume        float32            `bson:"volume,omitempty"`
	ShippingFrom  string             `bson:"shippingfrom,omitempty"`
	ShippingCosts [][]string         `bson:"shippingcosts,omitempty"`
}

type Blacklist struct {
	BlacklistID primitive.ObjectID   `bson:"_blacklistid,omitempty"`
	UserIDs     []primitive.ObjectID `bson:"_userids,omitempty"`
	AuctionID   primitive.ObjectID   `bson:"_auctionid,omitempty"`
}

type Seller struct {
	SellerID    primitive.ObjectID `bson:"_sellerid,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Logo        string             `bson:"logo,omitempty"`
	Credentials string             `bson:"credentials,omitempty"`
}

type SellerUser struct {
	SellerUserID primitive.ObjectID `bson:"_selleruserid,omitempty"`
	SellerID     primitive.ObjectID `bson:"_sellerid,omitempty"`
	Name         string             `bson:"name,omitempty"`
}

type Bidder struct {
	BidderID    primitive.ObjectID `bson:"_bidderid,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Email       string             `bson:"email,omitempty"`
	Location    string             `bson:"location,omitempty"`
	TotalBudget float32            `bson:"totalbudget,omitempty"`
	Budget      float32            `bson:"budget,omitempty"`
	Reserved    float32            `bson:"reserved,omitempty"`
	Credentials string             `bson:"credentials,omitempty"`
}
