package controller

type AppController struct {
	Product interface{ ProductController }
}
