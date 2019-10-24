package main

import (
	_ "github.com/mehrdadnekopour/go-tools/helpers"
	_ "github.com/mehrdadnekopour/go-tools/iis/jwt"
	_ "github.com/mehrdadnekopour/go-tools/iis/models"
	_ "github.com/mehrdadnekopour/go-tools/menv"

	_ "github.com/mehrdadnekopour/go-tools/mexcel"
	_ "github.com/mehrdadnekopour/go-tools/monfig/env"
	_ "github.com/mehrdadnekopour/go-tools/monfig/middleware"

	_ "github.com/mehrdadnekopour/go-tools/morm"
	_ "github.com/mehrdadnekopour/go-tools/mypes"
	_ "github.com/mehrdadnekopour/go-tools/rest"
	_ "github.com/mehrdadnekopour/go-tools/router"
	_ "github.com/mehrdadnekopour/go-tools/templates"
)

func main() {}
