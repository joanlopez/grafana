package inbox

import (
	"net/http"
	"strings"

	"github.com/grafana/grafana/pkg/api"
	"github.com/grafana/grafana/pkg/infra/log"
	"github.com/grafana/grafana/pkg/web"
)

type Service struct {
	handlers   map[string]map[string]func()
	httpServer *api.HTTPServer
	logger     log.Logger
}

func ProvideService(httpServer *api.HTTPServer) *Service {

	logger := log.New("inbox")
	hc := newHandlerCollection(logger)

	s := &Service{
		handlers: map[string]map[string]func(){
			"/api/datasources": {
				http.MethodPost: hc.AddDataSource,
			},
			"/api/datasources/:id": {
				http.MethodPut:    hc.UpdateDataSource,
				http.MethodDelete: hc.DeleteDataSource,
			},
		},
		httpServer: httpServer,
		logger:     logger,
	}

	s.logger.Error("Service being initialized...")
	s.httpServer.AddMiddleware(s.middleware())
	s.logger.Error("Middleware registered!")

	return s
}

func (s *Service) middleware() web.Handler {
	return func(c *web.Context) {
		s.logger.Error("Request received!")

		handler := s.getHandler(c)
		if handler == nil {
			s.logger.Error("No handler found!")
			c.Next()
			return
		}

		go handler()
		c.Next()
	}
}

func (s *Service) getHandler(c *web.Context) func() {
	path := s.getReqPath(c)
	method := c.Req.Method

	s.logger.Error("Finding handler for...", "path", path, "method", method)

	if pathHandlers, ok := s.handlers[path]; ok {
		if handler, ok := pathHandlers[method]; ok {
			return handler
		}
	}

	return nil
}

func (s *Service) getReqPath(c *web.Context) string {
	params := web.Params(c.Req)
	path := c.Req.URL.Path
	for key, value := range params {
		path = strings.Replace(path, value, key, 1)
	}
	// Remove trailing '/'
	if last := len(path) - 1; last >= 0 && path[last] == '/' {
		path = path[:last]
	}

	return path
}
