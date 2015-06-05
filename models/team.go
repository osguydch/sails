package models

type Team struct {
	Id           string   `json:"id"`           //
	Name         string   `json:"name"`         //
	Organization string   `json:"organization"` //
	Username     string   `json:"username"`     //
	Description  string   `json:"description"`  //
	Access       string   `json:"access"`       //
	Users        []string `json:"users"`        //
	Repositories []string `json:"repositories"` //
	Created      int64    `json:"created"`      //
	Updated      int64    `json:"updated"`      //
	Memo         []string `json:"memo"`         //
}
