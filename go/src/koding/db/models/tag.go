package models

import "labix.org/v2/mgo/bson"

type Tag struct {
  Id     bson.ObjectId `bson:"_id" json:"-"`
  Title  string        `bson:"title"`
  Slug   string        `bson:"slug"`
  Group  string        `bson:"group"`
  Type   string        `bson:"type,omitempty"`
  Counts TagCount      `bson:"counts"`
  Meta   Meta          `bson:"meta"`
}

type TagCount struct {
  Followers int `bson:"followers"`
  Following int `bson:"following"`
  Post      int `bson:"post,omitempty"`
  Tagged    int `bson:"tagged"`
}
