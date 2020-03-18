package main

type List map[string]ListValue

type ListValue struct {
	Name     string `json:"name"`
	Key      string `json:"key"`
	FileName string `json:"fileName"`
	Contents string `json:"contents"`
}
