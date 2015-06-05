package models

type Organization struct {
	Id           string   `json:"id"`           //
	Name         string   `json:"name"`         //
	Username     string   `json:"username"`     //
	Description  string   `json:"description"`  //
	Repositories []string `json:"repositories"` //
	Teams        []string `json:"teams"`        //
	Created      int64    `json:"created"`      //
	Updated      int64    `json:"updated"`      //
	Memo         []string `json:"memo"`         //
}
