package handler

import (
	"net/http"
    "github.com/gin-gonic/gin"  
	"github.com/bmonikraj/goql/model"
	"github.com/bmonikraj/goql/gql/schema"
	"github.com/graphql-go/graphql"
)

func GraphQLHandler(c *gin.Context) {

	qgl_schema := schema.InitSchema()
	
	q := model.RequestBody{}
	c.BindJSON(&q)

	params := graphql.Params{Schema: qgl_schema, RequestString: q.Query}
	res := graphql.Do(params)
	if len(res.Errors) > 0 {
		c.JSON(500, gin.H{"error" : res.Errors})
	}
	c.JSON(http.StatusOK, res)
}