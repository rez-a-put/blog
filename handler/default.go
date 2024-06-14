package handler

import (
	c "blog/controller"
	cd "blog/database/connection"
)

var (
	db = cd.ConnectDatabase()
	bh = c.NewBaseHandler(db)
)
