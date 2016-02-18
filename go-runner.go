package main

import (
	"github.com/codegangsta/cli"

	"github.com/maps90/go-runner/actions"

	"os"
)

var (
	Bucket         = "mm-storage-dev"
	Prefix         = "images"
	LocalDirectory = "web/images"
)

type Image struct {
	ParentSku     string `json:"parent_or_vendor_sku"`
	ImageSequence string `json:"image_sequence"`
	ImagePath     string `json:"image_path"`
}

func main() {
	app := cli.NewApp()
	app.Name = "imageProcessor"
	app.Commands = []cli.Command{
		{
			Name:  "load",
			Usage: "load images from json file.",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file, f",
					Usage: "file path",
				},
			},
			Action: func(c *cli.Context) {
				if len(c.String("file")) < 1 {
					return
				}
				action.LoadImages(c.String("file"))
			},
		},
	}
	app.Run(os.Args)
}
