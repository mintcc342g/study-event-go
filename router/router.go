package router

import (
	"study_event_go/cmd/container"

	"github.com/labstack/echo/v4"
)

// AssaultLilyRoute ...
func AssaultLilyRoute(sys *echo.Group, ctrlContainer *container.ControllerContainer) {

	// TODO
	// sys.POST("/garden", ctrlContainer.Garden.New)
	// sys.GET("/garden", ctrlContainer.Garden.List)
	// sys.GET("/garden/:id", ctrlContainer.Garden.Get)
	// sys.PUT("/garden/:id", ctrlContainer.Garden.Update)
	// sys.DELETE("/garden/:id", ctrlContainer.Garden.Delete)

	// TODO
	// sys.POST("/lily", ctrlContainer.Lily.New)
	// sys.GET("/lily", ctrlContainer.Lily.List)
	// sys.GET("/lily/:id", ctrlContainer.Lily.Get)
	// sys.PUT("/lily/:id", ctrlContainer.Lily.Update)
	// sys.DELETE("/lily/:id", ctrlContainer.Lily.Delete)

	sys.POST("/garden/:id/alarm", ctrlContainer.Protection.Alarm)
}
