package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("cellar", func() {
	Title("The virtual wine cellar")
	Description("A simple goa service")
})

var _ = Resource("bottle", func() {
	BasePath("/bottles")
	DefaultMedia(BottleMedia)

	Action("show", func() {
		Description("Get bottle by id")
		Routing(GET("/:bottleID"))
		Params(func() {
			Param("bottleID", Integer, "Bottle ID")  // TODO(ags799): use UUID instead of Integer
		})
		Response(OK)
		Response(NotFound)
	})
})

var BottleMedia = MediaType("application/goa-example.bottle+json", func() {
	Description("A bottle of wine")
	Attributes(func() {
		Attribute("id", Integer, "Unique bottle ID")
		Attribute("href", String, "API href for making requests on the bottle")
		Attribute("name", String, "Name of wine")
		Required("id", "href", "name")
	})
	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("name")
	})
})
