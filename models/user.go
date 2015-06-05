package models

type User struct {
	Id                string   `json:"id"`                //
	Username          string   `json:"username"`          //
	Password          string   `json:"password"`          //
	Email             string   `json:"email"`             //
	Fullname          string   `json:"fullname"`          //
	Company           string   `json:"company"`           //
	Location          string   `json:"location"`          //
	Mobile            string   `json:"mobile"`            //
	URL               string   `json:"url"`               //
	Gravatar          string   `json:"gravatar"`          //
	Created           int64    `json:"created"`           //
	Updated           int64    `json:"updated"`           //
	Repositories      []string `json:"repositories"`      //
	Organizations     []string `json:"organizations"`     //
	Teams             []string `json:"teams"`             //
	JoinOrganizations []string `json:"joinorganizations"` //
	JoinTeams         []string `json:"jointeams"`         //
	Starts            []string `json:"starts"`            //
	Comments          []string `json:"comments"`          //
	Memo              []string `json:"memo"`              //
}
