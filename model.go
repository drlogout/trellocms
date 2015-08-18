package trellocms

import(
	"net/http"
	"encoding/json"
	"io/ioutil"
	"strings"
	"fmt"
)

type Card struct {
	Name string
	Id string
	Desc string
}

type List struct {
	Name string
	Id string
	Slug string
}

type Lists struct {
	Lists []List
}



func GetLists(config Config) (Lists, error) {
	url := config.API + "boards/" + config.BoardId + "/lists"
	res, err := http.Get(url)

	defer res.Body.Close()

	if err != nil {
		panic(err)
	}

	jsonData, err := ioutil.ReadAll(res.Body)

	var lists Lists

	err = json.Unmarshal(jsonData, &lists.Lists)

	if err == nil {
		lists.Slugify()
	}

	return lists, err
}

func (this List) GetCards(config Config) ([]Card, error) {

	url := config.API + "lists/" + this.Id + "/cards"
	res, err := http.Get(url)

	defer res.Body.Close()

	if err != nil {
		panic(err)
	}

	jsonData, err := ioutil.ReadAll(res.Body)

	var cards []Card

	err = json.Unmarshal([]byte(jsonData), &cards)

	if err != nil {
		panic(err)
	}

	return cards, err
}


func (this Lists) GetByName(name string) List{
	var foundList List
	for _,list := range this.Lists {
		if list.Name == name {
			foundList = list
			break
		}
	}
	return foundList
}

func (this *Lists) Slugify() {
	for i, list := range this.Lists {
		this.Lists[i].Slug = strings.ToLower(list.Name)
	}
	fmt.Print("slug", this.Lists)
}

