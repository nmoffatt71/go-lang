package main

import (
	"html/template"
	"net/http"
	"patient-manager/graphql/generated"
	"patient-manager/graphql/resolver"
	"patient-manager/internal/db"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func main() {
	db.ConnectMongo()
	r := gin.Default()

	r.SetHTMLTemplate(template.Must(template.ParseFiles("templates/patients.html")))

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "patients.html", nil)
	})

	r.GET("/patients", func(c *gin.Context) {
		search := c.Query("search")
		resolver := resolver.Resolver{}
		patients, err := resolver.SearchPatientsResolver(c.Request.Context(), search)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error loading patients")
			return
		}
		html := ""
		for _, p := range patients {
			html += "<div>" + p.Firstname + " " + p.Lastname + " (" + p.Gender + ", DOB: " + p.Dateofbirth + ")</div>"
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
	})

	r.POST("/query", gin.WrapH(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))))
	r.GET("/playground", gin.WrapH(playground.Handler("GraphQL", "/query")))

	r.Run(":8080")
}