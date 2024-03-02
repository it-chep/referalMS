package internal

func (app *App) InitControllers() *App {
	app.controller = controllers{}
	return app
}

func (app *App) InitPgxConn() *App {
	return app
}
