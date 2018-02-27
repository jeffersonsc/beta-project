package handler

import (
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/jeffersonsc/beta-project/conf"
	"github.com/jeffersonsc/beta-project/lib/context"
)

type User struct {
	ID   bson.ObjectId `bson:"_id"`
	Name string        `bson:"name"`
}

func TestConnMongo(ctx *context.Context) {
	col, err := conf.GetMongoCollection("users")
	if err != nil {
		log.Println("Error ", err.Error())
		return
	}

	user := User{
		ID:   bson.NewObjectId(),
		Name: "Jefferson",
	}

	if err := col.Insert(&user); err != nil {
		log.Println("Erro ", err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "<html><body>OK</body></html>")
}
