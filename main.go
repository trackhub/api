package main

import (
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/trackhub/api/graph"
	"github.com/trackhub/api/graph/generated"
)

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{}},
		),
	)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	godotenv.Load()

	r := gin.Default()

	hf := func(ctx *gin.Context) {
		auth := ctx.Request.Header.Get("auth")
		if auth != "test123" {
			// ctx.JSON(500, map[string]string{"test123": "da"})
			// do ctx.Abort() to cancel the request
			return
		}

		ctx.Next()
	}

	g := r.Group("/")
	g.Use(hf)

	g.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())

	r.Run(":" + os.Getenv("PORT"))
}
