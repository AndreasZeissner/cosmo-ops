// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Country struct {
	Key      *CountryKey `json:"key"`
	Language *string     `json:"language,omitempty"`
}

func (Country) IsEntity() {}

type CountryKey struct {
	Name string `json:"name"`
}

type Query struct {
}
