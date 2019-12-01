module main

go 1.13

require (
	apis/controller v0.0.0-00010101000000-000000000000
	dbconn v0.0.0-00010101000000-000000000000
	github.com/go-delve/delve v1.3.2
	github.com/gorilla/mux v1.7.3
	github.com/sirupsen/logrus v1.2.0
)

replace (
	apis/controller => /home/aniket/Desktop/GO/GO-Boiler-Plate/src/apis/controller
	dbconn => /home/aniket/Desktop/GO/GO-Boiler-Plate/src/dbconn
	structs => /home/aniket/Desktop/GO/GO-Boiler-Plate/src/structs
)
