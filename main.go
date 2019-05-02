package main

import (
	"bytes"
	"image/jpeg"
	"log"
	"os"

	"github.com/aofei/air"
	"github.com/aofei/cameron"
)

var a = air.Default

func main() {
	a.Address = ":" + os.Getenv("PORT")
	a.GET("/:Name", identicon)
	log.Fatal(a.Serve(), nil)
}



func identicon(req *air.Request, res *air.Response) error {

	res.Header.Set("Access-Control-Allow-Origin", "*")

	buf := bytes.Buffer{}
	err:= jpeg.Encode(
		&buf,
		cameron.Identicon(
			[]byte(req.Param("Name").Value().String()),
			540,
			60,
		),
		&jpeg.Options{
			Quality: 100,
		},
	)

	if err!=nil{
		log.Fatal("error", nil)
	}

	res.Header.Set("Content-Type", "image/jpeg")

	return res.Write(bytes.NewReader(buf.Bytes()))
}