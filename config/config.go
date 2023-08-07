package config

import (
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/template/html/v2"
	"github.com/haji-sudo/ShabehRoshan/router/routes"
	"github.com/joho/godotenv"
)

var Engine *html.Engine

func Init() {

	//Load and add ENV
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create view engine
	Engine = html.New("./views", ".html")

	// Disable this in production
	Engine.Reload(true)

	Engine.AddFunc("getCssAsset", func(name string) (res template.HTML) {
		filepath.Walk("public/assets", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.Name() == name {
				res = template.HTML("<link rel=\"stylesheet\" href=\"/" + path + "\">")
			}
			return nil
		})
		return
	})

	Engine.AddFunc("URL", func(url string) template.HTML {
		return template.HTML(routes.Geturlpath()[url])
	})

}
