module controller.go

go 1.13

replace (
	dbconn => /home/aniket/Desktop/GO/GO-Boiler-Plate/src/dbconn
	structs => /home/aniket/Desktop/GO/GO-Boiler-Plate/src/structs
)

require (
	dbconn v0.0.0-00010101000000-000000000000
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fatih/color v1.7.0
	github.com/gookit/color v1.2.0
	github.com/gorilla/mux v1.6.2
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/mattn/go-isatty v0.0.10 // indirect
	github.com/sirupsen/logrus v1.2.0
	github.com/withmandala/go-log v0.1.0
	structs v0.0.0-00010101000000-000000000000
)
