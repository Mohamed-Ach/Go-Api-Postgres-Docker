package main

import (
	controller "github.com/mohamed-ach/go-repo/blog-rest-api/controller"
	model "github.com/mohamed-ach/go-repo/blog-rest-api/model"
)

func main() {
	model.Init()
	controller.Start()
}
