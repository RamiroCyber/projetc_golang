package handler

import (
	"errors"
	"github.com/RamiroCyber/projetc_golang/config/database"
	"github.com/RamiroCyber/projetc_golang/util"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func DeleteUser(c *fiber.Ctx) error {
	objID, err := getObjectIDFromParams(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("invalid user id")
	}

	err = deleteUser(c, objID)
	if err != nil {
		return handleError(err)
	}

	return c.SendStatus(fiber.StatusOK)
}

func getObjectIDFromParams(c *fiber.Ctx) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(c.Params("id"))
}

func deleteUser(c *fiber.Ctx, objID primitive.ObjectID) error {
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{"deleted_at": time.Now()}}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	res := database.UserCollection.FindOneAndUpdate(c.Context(), filter, update, opts)
	if err := res.Err(); err != nil {
		return err
	}

	return nil
}

func handleError(err error) error {
	if errors.Is(err, mongo.ErrNoDocuments) {
		return fiber.NewError(fiber.StatusNotFound, "user not found")
	}

	util.Logger("ERROR", err.Error())
	return fiber.NewError(fiber.StatusInternalServerError, err.Error())
}
