package handlers

import (
	"go-crud-api-NoSQL/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func HandleAllCollections(c *fiber.Ctx) error {

	collection := c.Params("collection")
	// fetch all Collections
	coll := database.GetCollection(collection)

	// return all Collections
	filter := bson.M{}
	opts := options.Find().SetSkip(0).SetLimit(100)

	// find all Collections
	cursor, err := coll.Find(c.Context(), filter, opts)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"internal server error": err.Error()})
	}

	// convert cursor to slice
	Collections := make([]bson.M, 0)
	if err = cursor.All(c.Context(), &Collections); err != nil {
		return c.Status(500).JSON(fiber.Map{"internal server error": err.Error()})
	}

	// return Collections
	return c.Status(200).JSON(Collections)
}

func HandleCreateCollection(c *fiber.Ctx) error {
	collection := c.Params("collection")
	// get the Collection from the request body
	nCollection := new(bson.M)

	// validate the request body
	if err := c.BodyParser(nCollection); err != nil {
		return c.Status(400).JSON(fiber.Map{"bad input": err.Error()})
	}

	// insert the Collection into the database
	coll := database.GetCollection(collection)
	_, err := coll.InsertOne(c.Context(), nCollection)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"internal server error": err.Error()})
	}

	// return the inserted Collection
	return HandleAllCollections(c)

}

func HandleUpdateCollection(c *fiber.Ctx) error {

	collection := c.Params("collection")

	// get the id from the request params
	id := c.Params("id")
	dbId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"invalid id": err.Error()})
	}

	// get the Collection from the request body
	uCollection := new(bson.M)

	// validate the request body
	if err := c.BodyParser(uCollection); err != nil {
		return c.Status(400).JSON(fiber.Map{"bad input": err.Error()})
	}

	// update the Collection in the database
	coll := database.GetCollection(collection)
	filter := bson.M{"_id": dbId}
	update := bson.M{"$set": uCollection}
	res, err := coll.UpdateOne(c.Context(), filter, update)

	if err != nil || res == nil {
		return c.Status(500).JSON(fiber.Map{"internal server error": err.Error()})
	}

	// return the updated Collection
	return HandleAllCollections(c)

}

func HandleAllUpdateCollection(c *fiber.Ctx) error {

	collection := c.Params("collection")

	// get the Collection from the request body
	uCollection := new(bson.M)

	// validate the request body
	if err := c.BodyParser(uCollection); err != nil {
		return c.Status(400).JSON(fiber.Map{"bad input": err.Error()})
	}

	// update the Collection in the database
	coll := database.GetCollection(collection)
	update := bson.M{"$set": uCollection}
	res, err := coll.UpdateMany(c.Context(), bson.M{}, update)

	if err != nil || res == nil {
		return c.Status(500).JSON(fiber.Map{"internal server error": err.Error()})
	}

	// return the updated Collection
	return HandleAllCollections(c)
}

func HandleGetOneCollection(c *fiber.Ctx) error {

	collection := c.Params("collection")

	// get the id from the request params
	id := c.Params("id")
	dbId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"invalid id": err.Error()})
	}

	// fetch the Collection from the database
	coll := database.GetCollection(collection)
	filter := bson.M{"_id": dbId}
	var Collection bson.M
	err = coll.FindOne(c.Context(), filter).Decode(&Collection)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	// return the Collection
	return c.Status(200).JSON(Collection)
}

func HandleDeleteCollection(c *fiber.Ctx) error {

	collection := c.Params("collection")

	// get the id from the request params
	id := c.Params("id")
	dbId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"invalid id": err.Error()})
	}

	// delete the Collection from the database
	coll := database.GetCollection(collection)
	filter := bson.M{"_id": dbId}
	res, err := coll.DeleteOne(c.Context(), filter)

	if err != nil || res == nil {
		return c.Status(500).JSON(fiber.Map{"internal server error": err.Error()})
	}

	// return the deleted Collection
	return HandleAllCollections(c)
}

func HandleAllDeleteCollection(c *fiber.Ctx) error {
	collection := c.Params("collection")

	// delete the Collection from the database
	coll := database.GetCollection(collection)

	res, err := coll.DeleteMany(c.Context(), bson.D{})

	if err != nil || res == nil {
		return c.Status(500).JSON(fiber.Map{"internal server error": err.Error()})
	}

	// return the deleted Collection
	return HandleAllCollections(c)
}
