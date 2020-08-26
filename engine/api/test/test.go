package test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ovh/cds/engine/api/authentication"
	"github.com/ovh/cds/engine/api/database/gorpmapping"
	"github.com/ovh/cds/engine/cache"
	"github.com/ovh/cds/engine/test"
	"github.com/ovh/cds/sdk"
)

// SetupPG setup PG DB for test and use gorpmapping singleton's mapper.
func SetupPG(t *testing.T, bootstrapFunc ...test.Bootstrapf) (*test.FakeTransaction, cache.Store) {
	db, cache, cancel := test.SetupPGToCancel(t, gorpmapping.Mapper, sdk.TypeAPI, bootstrapFunc...)
	t.Cleanup(cancel)

	err := authentication.Init("cds-api-test", test.SigningKey)
	require.NoError(t, err, "unable to init authentication layer")

	return db, cache
}
