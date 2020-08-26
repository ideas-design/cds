package cdn

import (
	"context"
	"testing"

	"github.com/gorilla/mux"

	"github.com/ovh/cds/engine/api"
	"github.com/ovh/cds/engine/cdn/index"
	"github.com/ovh/cds/engine/cdn/storage"
	"github.com/ovh/cds/engine/gorpmapper"
	"github.com/ovh/cds/engine/test"
	"github.com/ovh/cds/sdk"
)

func newRouter(m *mux.Router, p string) *api.Router {
	r := &api.Router{
		Mux:        m,
		Prefix:     p,
		URL:        "",
		Background: context.Background(),
	}
	return r
}

func newTestService(t *testing.T) (*Service, *test.FakeTransaction) {
	m := gorpmapper.New()
	index.InitDBMapping(m)
	storage.InitDBMapping(m)
	db, cache := test.SetupPGWithMapper(t, m, sdk.TypeCDN)

	router := newRouter(mux.NewRouter(), "/"+test.GetTestName(t))
	var cancel context.CancelFunc
	router.Background, cancel = context.WithCancel(context.Background())
	s := &Service{
		Router:              router,
		DBConnectionFactory: test.DBConnectionFactory,
		Cache:               cache,
		Mapper:              m,
	}
	s.initRouter(context.TODO())

	t.Cleanup(func() { cancel() })
	return s, db
}
