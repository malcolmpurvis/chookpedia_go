package models

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"
)

type Chook struct {
	Id       int    `orm:"auto;primary_key"`
	Name     string `json:"name"`
	Colour   string `json:"colour"`
	PhotoURL string `json:"photo_url"`
	Breed    *Breed `orm:"null;rel(fk);on_delete(set_null)"`
}

type Breed struct {
	Id     int    `orm:"auto;primary_key"`
	Name   string `json:"name"`
	Origin string `json:"origin"`
}

func InitModel() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=malcolmp host=127.0.0.1 port=5432 dbname=chookpedia sslmode=disable")
	orm.RegisterModel(new(Breed))
	orm.RegisterModel(new(Chook))
	orm.RunSyncdb("default", false, true)
}
