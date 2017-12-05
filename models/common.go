package models

import (
	"github.com/goweb4/database"
)

type Model interface {
	GetRelationship() map[string]interface{}
}

func GetRelatedData(model Model, relationshipName string) {
	data := model.GetRelationship()
	database.DBCon.Model(model).Association(relationshipName).Find(data[relationshipName])
}
