package swift

import (
	"io/ioutil"
	"testing"

	"github.com/ovh/cds/engine/cdn/index"
	"github.com/ovh/cds/engine/cdn/storage"
	"github.com/ovh/cds/sdk/log"
	"github.com/ovh/symmecrypt/ciphers/aesgcm"
	"github.com/ovh/symmecrypt/convergent"
	"github.com/stretchr/testify/require"
)

func TestSwift(t *testing.T) {
	log.SetLogger(t)
	var driver = new(Swift)
	err := driver.Init(&storage.SwiftStorageConfiguration{
		Encryption: []convergent.ConvergentEncryptionConfig{
			{
				Cipher:      aesgcm.CipherName,
				LocatorSalt: "secret_locator_salt",
				SecretValue: "secret_value",
			},
		},
	})
	require.NoError(t, err, "unable to initialiaze webdav driver")

	err = driver.client.ApplyEnvironment()
	if err != nil {
		t.Logf("skipping this test: %v", err)
		t.SkipNow()
	}

	err = driver.client.Authenticate()
	if err != nil {
		t.Logf("skipping this test: %v", err)
		t.SkipNow()
	}

	itemUnit := storage.ItemUnit{
		Locator: "a_locator",
		Item: &index.Item{
			Type: index.TypeItemStepLog,
		},
	}
	w, err := driver.NewWriter(itemUnit)
	require.NoError(t, err)
	require.NotNil(t, w)

	_, err = w.Write([]byte("something"))
	require.NoError(t, err)

	err = w.Close()
	require.NoError(t, err)

	r, err := driver.NewReader(itemUnit)
	require.NoError(t, err)
	require.NotNil(t, r)

	btes, err := ioutil.ReadAll(r)
	require.NoError(t, err)
	err = r.Close()
	require.NoError(t, err)

	require.Equal(t, "something", string(btes))
}
