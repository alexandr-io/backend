package database

import "github.com/gofiber/fiber"

// FindOneWithFilter fill the given object with a mongodb single result filtered by the given filters.
func FindOneWithFilter(ctx *fiber.Ctx, object interface{}, filters interface{}) error {
	collection := Instance.Db.Collection(CollectionUser)
	filteredSingleResult := collection.FindOne(ctx.Fasthttp, filters)
	return filteredSingleResult.Decode(object)
}
