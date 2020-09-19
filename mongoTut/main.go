package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// Product describes an electronic product i.e phone
type Product struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	Name string `json:"product_name" bson:"product_name"`
	Price int `json:"price" bson:"price"`
	Currency string `json:"currency" bson:"currency"`
	Quantity string `json:"quantity" bson:"quantity"`
	Discount int `json:"discount,omitempty" bson:"discount,omitempty"`
	Vendor string `json:"vendor" bson:"vendor"`
	Accessories []string `json:"accessories,omitempty" bson:"accessories,omitempty"`
	SkuID string `json:"sku_id" bson:"sku_id"`
}

var iphone = Product{
	ID:          primitive.NewObjectID(),
	Name:        "iphone10",
	Price:       900,
	Currency:    "USD",
	Quantity:    "40",
	Discount:    0,
	Vendor:      "apple",
	Accessories: []string{"charger", "headset"},
	SkuID:       "1234",
}

var toy = Product{
	ID:          primitive.NewObjectID(),
	Name:        "toy",
	Price:       5,
	Currency:    "USD",
	Quantity:    "10",
	Discount:    0,
	Vendor:      "Tokyo",
	Accessories: []string{"charger", "headset"},
	SkuID:       "1234",
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(err)
	}

	db := client.Database("tronics")
	collection := db.Collection("products")

	// **** WRITE ****

	// Insert one
	res, err := collection.InsertOne(context.Background(), iphone)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.InsertedID.(primitive.ObjectID).Timestamp())

	// Using bson.D
	res, err = collection.InsertOne(context.Background(), bson.D{
		{"name", "test"},
		{"age", 1},
		{"hobbies", bson.A{"mcdonalds"}},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.InsertedID)

	resMany, err := collection.InsertMany(context.Background(), []interface{}{iphone, toy})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resMany.InsertedIDs)

	// **** READ ****

	// Equality operator using FindOne
	var findOne Product
	err = collection.FindOne(context.Background(), bson.M{"price": 800}).Decode(&findOne)
	fmt.Println(findOne)

	// Comparison operator using Find
	var find Product
	findCursor, err := collection.Find(context.Background(), bson.M{"price": bson.M{"$gt": 100}})
	if err != nil {
		fmt.Println(err)
	}
	for findCursor.Next(context.Background()) {
		if err := findCursor.Decode(&find); err != nil {
			fmt.Println(err)
		}
		fmt.Println(find.Name)
	}

	// Logical operator using Find
	var findLogic Product
	logicFilter := bson.M{
		"$and": bson.A{
			bson.M{"price": bson.M{"$gt": 100}},
			bson.M{"quantity": bson.M{"$gt": 30}},
		},
	}

	findLogicRes, err := collection.Find(context.Background(), logicFilter)
	if err != nil {
		fmt.Println(err)
	}
	for findLogicRes.Next(context.Background()) {
		if err := findLogicRes.Decode(&findLogic); err != nil {
			fmt.Println(err)
		}
		fmt.Println(findLogic.Name)
	}

	// **** UPDATE ****

	// Update operator for Field
	updateFieldCon := bson.M{"$set": bson.M{"IsEssential": "false"}}
	updateFieldRes, err := collection.UpdateMany(context.Background(), bson.M{}, updateFieldCon)
	fmt.Println(updateFieldRes.ModifiedCount)

	// Delete operation
	delRes, err := collection.DeleteMany(context.Background(), updateFieldCon)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(delRes.DeletedCount)
}
