package main

import (
	"rxPatient/graphql/generated"
	"rxPatient/graphql/resolver"
	"rxPatient/internal/db"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func main() {
	db.ConnectMongo()
	r := gin.Default()

	r.POST("/query", gin.WrapH(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))))
	r.GET("/playground", gin.WrapH(playground.Handler("GraphQL", "/query")))

	r.Run(":8080")
}
