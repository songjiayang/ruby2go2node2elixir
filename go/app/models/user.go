package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

var (
	User      *_User
	tableName = "users"
)

type UserModel struct {
	Id        bson.ObjectId `bson:"_id" json:"id"`
	Username  uint32        `bson:"username" json:"username"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at" json:"updated_at"`
}

type _User struct {
}

func (_ *_User) Find(id string) (user *UserModel) {
	db.C(tableName).FindId(bson.ObjectIdHex(id)).One(&user)
	return
}

func (_ *_User) All() (users []*UserModel) {
	db.C(tableName).Find(bson.M{}).All(&users)
	return
}
