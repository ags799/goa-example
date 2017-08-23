package main

import (
	"github.com/ags799/goa-example/app"
	"github.com/goadesign/goa"
	"fmt"
)

// BottleController implements the bottle resource.
type BottleController struct {
	*goa.Controller
}

// NewBottleController creates a bottle controller.
func NewBottleController(service *goa.Service) *BottleController {
	return &BottleController{Controller: service.NewController("BottleController")}
}

// Show runs the show action.
func (c *BottleController) Show(ctx *app.ShowBottleContext) error {
	if ctx.BottleID == 0 {
		return ctx.NotFound()
	}
	bottle := app.GoaExampleBottle{
		ID: ctx.BottleID,
		Name: fmt.Sprintf("Bottle #%d", ctx.BottleID),
		Href: app.BottleHref(ctx.BottleID),
	}
	return ctx.OK(&bottle)
}
