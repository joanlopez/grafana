package inbox

import "github.com/grafana/grafana/pkg/infra/log"

type handlerCollection struct {
	logger log.Logger
}

func newHandlerCollection(logger log.Logger) handlerCollection {
	return handlerCollection{
		logger: logger,
	}
}

func (h handlerCollection) AddDataSource() {
	h.logger.Error("Executing handler...", "handler", "AddDataSource")
}

func (h handlerCollection) UpdateDataSource() {
	h.logger.Error("Executing handler...", "handler", "UpdateDataSource")
}

func (h handlerCollection) DeleteDataSource() {
	h.logger.Error("Executing handler...", "handler", "DeleteDataSource")
}
