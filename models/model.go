package models

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"
)

type Chook struct {
	Id       int    `orm:"auto"`
	Name     string `json:"name"`
	Colour   string `json:"colour"`
	PhotoURL string `json:"photo_url"`
	Breed    *Breed `orm:"null;rel(fk);on_delete(set_null)"`
}

type Breed struct {
	Id     int    `orm:"auto"`
	Name   string `json:"name"`
	Origin string `json:"origin"`
}

func InitModel() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=malcolmp password=example host=db port=5432 dbname=chookpedia sslmode=disable")
	orm.RegisterModel(new(Breed))
	orm.RegisterModel(new(Chook))
	orm.RunSyncdb("default", false, true)
}

func AllChooks() ([]*Chook, error) {
	var chooks []*Chook

	o := orm.NewOrm()
	_, err := o.QueryTable("chook").All(&chooks)

	return chooks, err
}

func FindChook(id int) (Chook, error) {
	o := orm.NewOrm()
	chook := Chook{Id: id}

	err := o.Read(&chook)

	return chook, err
}

func FindBreed(id int) (Breed, error) {
	o := orm.NewOrm()
	breed := Breed{Id: id}

	err := o.Read(&breed)

	return breed, err
}
