package main

type Task struct {
	UUID  string `json:"uuid"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}
