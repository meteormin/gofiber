package monitor

import "github.com/gofiber/fiber/v2"

type fiberInfo struct {
	HandlersCount uint32
}

type routerInfo struct {
	Name   string
	Method string
	Path   string
	Params []string
}

func newFiberInfo(f *fiber.App) fiberInfo {
	return fiberInfo{
		HandlersCount: f.HandlersCount(),
	}
}

func newRouterInfo(f *fiber.App) []routerInfo {
	ri := make([]routerInfo, 0)
	for _, r := range f.GetRoutes() {
		ri = append(ri, routerInfo{
			Name:   r.Name,
			Method: r.Method,
			Path:   r.Path,
			Params: r.Params,
		})
	}

	return ri
}
