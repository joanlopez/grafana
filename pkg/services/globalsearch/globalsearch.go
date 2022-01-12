package globalsearch

import (
	"context"
	"net/http"

	"github.com/grafana/grafana/pkg/services/sqlstore"

	"github.com/blevesearch/bleve"
	"github.com/grafana/grafana/pkg/api/response"
	"github.com/grafana/grafana/pkg/api/routing"
	"github.com/grafana/grafana/pkg/infra/log"
	"github.com/grafana/grafana/pkg/models"
)

type Service struct {
	routerRegister routing.RouteRegister
	sqlStore       *sqlstore.SQLStore

	index  bleve.Index
	logger log.MultiLoggers
}

func ProvideService(
	routerRegister routing.RouteRegister,
	sqlStore *sqlstore.SQLStore,
) (*Service, error) {
	mapping := bleve.NewIndexMapping()
	index, err := bleve.NewMemOnly(mapping)
	if err != nil {
		return nil, err
	}

	message := struct {
		Id   string
		From string
		Body string
	}{
		Id:   "example",
		From: "marty.schoch@gmail.com",
		Body: "bleve indexing is easy",
	}

	err = index.Index(message.Id, message)
	if err != nil {
		return nil, err
	}

	srv := &Service{
		routerRegister: routerRegister,
		sqlStore:       sqlStore,

		index:  index,
		logger: log.New("globalsearch"),
	}

	srv.registerRoutes()

	return srv, nil
}

func (s *Service) registerRoutes() {
	s.routerRegister.Get("globalsearch", routing.Wrap(s.GetResults))
}

func (s *Service) GetResults(reqCtx *models.ReqContext) response.Response {
	query := bleve.NewQueryStringQuery(reqCtx.Query("q"))
	searchRequest := bleve.NewSearchRequest(query)

	searchResult, err := s.index.Search(searchRequest)
	if err != nil {
		return response.JSON(http.StatusInternalServerError, err)
	}

	return response.JSON(http.StatusOK, searchResult)
}

func (s *Service) Run(ctx context.Context) error {
	dsResults, err := s.sqlStore.GetAllDataSources(ctx)
	if err != nil {
		s.logger.Error("sqlStore.GetAllDataSources failed", "error", err)
	}

	for _, ds := range dsResults {
		if err := s.index.Index(ds.Uid, *ds); err != nil {
			s.logger.Error("error while indexing ds", "uid", ds.Uid, "error", err)
		}
	}

	dbResults, err := sqlstore.GetAllDashboards(ctx)
	if err != nil {
		s.logger.Error("sqlstore.GetAllDashboards failed", "error", err)
	}

	for _, db := range dbResults {
		bytes, err := db.Data.ToDB()
		if err != nil {
			s.logger.Error("error while indexing db", "uid", db.Uid, "error", err)
			return err
		}

		if err := s.index.Index(db.Uid, string(bytes)); err != nil {
			s.logger.Error("error while indexing db", "uid", db.Uid, "error", err)
		}
	}

	s.logger.Info("Index built successfully")

	return nil
}
