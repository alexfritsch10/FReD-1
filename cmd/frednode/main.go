package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"gitlab.tu-berlin.de/mcc-fred/fred/pkg/app"
)

// Fred is the main app logic.
type Fred interface {
	CreateKeygroup(kg string) error
	DeleteKeygroup(kg string) error
	Read(kg string, id string) (string, error)
	Update(kg string, id string, data string) error
	Delete(kg string, id string) error
}

var addr = flag.String("addr", ":9001", "http service address")
var a Fred

func postKeygroup(context *gin.Context) {

	kgname := context.Params.ByName("kgname")

	err := a.CreateKeygroup(kgname)

	if err != nil {
		context.Status(http.StatusConflict)
		return
	}

	context.Status(http.StatusOK)
	return
}

func deleteKeygroup(context *gin.Context) {
	kgname := context.Params.ByName("kgname")

	err := a.DeleteKeygroup(kgname)

	if err != nil {
		context.Status(http.StatusConflict)
		return
	}

	context.Status(http.StatusOK)
	return
}

func getItem(context *gin.Context) {
	kgname := context.Params.ByName("kgname")

	id := context.Params.ByName("id")

	data, err := a.Read(kgname, id)

	if err != nil {
		context.Status(http.StatusNotFound)
		return
	}

	context.String(http.StatusOK, data)
	return
}

func putItem(context *gin.Context) {
	kgname := context.Params.ByName("kgname")

	id := context.Params.ByName("id")

	var json struct {
		Data string `json:"data" binding:"required"`
	}

	if err := context.ShouldBindJSON(&json); err != nil {
		log.Print(err)
		context.Status(http.StatusBadRequest)
		return
	}

	data := json.Data
	err := a.Update(kgname, id, data)

	if err != nil {
		context.Status(http.StatusConflict)
		return
	}

	context.Status(http.StatusOK)
	return
}

func deleteItem(context *gin.Context) {
	kgname := context.Params.ByName("kgname")

	id := context.Params.ByName("id")

	err := a.Delete(kgname, id)

	if err != nil {
		context.Status(http.StatusNotFound)
		return
	}

	context.Status(http.StatusOK)
	return
}

func setupRouter() (r *gin.Engine) {
	r = gin.Default()

	r.POST("/keygroup/:kgname", postKeygroup)
	r.DELETE("/keygroup/:kgname", deleteKeygroup)

	r.GET("/keygroup/:kgname/items/:id", getItem)
	r.PUT("/keygroup/:kgname/items/:id", putItem)
	r.DELETE("/keygroup/:kgname/items/:id", deleteItem)

	return
}

func main() {
	a = app.New()

	r := setupRouter()

	log.Fatal(r.Run(*addr))
}
