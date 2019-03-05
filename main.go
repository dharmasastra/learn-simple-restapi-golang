package main

import "github.com/dharmasastra/lerning-restapi/app"

func main()  {
	app := &app.App{}
	app.Initialize()
	app.Run(":3000")
}
