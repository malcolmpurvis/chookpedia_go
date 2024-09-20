package main

import (
	"chookpedia_go/models"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"io/ioutil"
	"log"
)

type ChookFile struct {
	Id       int    `orm:"auto;primary_key"`
	Name     string `json:"name"`
	Colour   string `json:"colour"`
	PhotoURL string `json:"photo_url"`
	Breed    string `json:"breed"`
}

func populate() {
	o := orm.NewOrm()

	content, err := ioutil.ReadFile("static/db/breeds.json")
	if err != nil {
		log.Fatal(err)
	}

	var breeds []models.Breed

	err = json.Unmarshal(content, &breeds)

	if err != nil {
		log.Fatal(err)
	}

	res, err := o.Raw("DELETE from breed").Exec()

	if err != nil {
		log.Fatal(err)
	}

	rows, _ := res.RowsAffected()
	fmt.Printf("Deleted %v rows from breeds table.\n", rows)

	for _, breed := range breeds {
		o.Insert(&breed)
	}

	content, err = ioutil.ReadFile("static/db/chooks.json")

	if err != nil {
		log.Fatal(err)
	}

	var chooks []ChookFile

	err = json.Unmarshal(content, &chooks)

	if err != nil {
		log.Fatal(err)
	}

	res, err = o.Raw("DELETE from chook").Exec()

	if err != nil {
		log.Fatal(err)
	}

	rows, _ = res.RowsAffected()
	fmt.Printf("Deleted %v rows from chooks table.\n", rows)

	fmt.Println(chooks)

	for _, chook := range chooks {
		dbChook := new(models.Chook)
		dbChook.Name = chook.Name
		dbChook.Colour = chook.Colour
		dbChook.PhotoURL = chook.PhotoURL

		var thisBreed models.Breed

		fmt.Println(chook)
		err := o.QueryTable("breed").Filter("name", chook.Breed).One(&thisBreed)
		if err == orm.ErrMultiRows {
			// Have multiple records
			log.Fatal("Returned Multi Rows Not One")
		}
		if err == orm.ErrNoRows {
			// No result
			log.Fatal("No row found")
		}

		dbChook.Breed = &thisBreed

		fmt.Println(dbChook)

		fmt.Println(o.Insert(dbChook))
	}
}

func main() {
	models.InitModel()
	orm.RunCommand()

	populate()
}
