package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
	"path"

	"github.com/ovh/cds/sdk/log"
)

// LoadTestingConf loads test configuration tests.cfg.json
func LoadTestingConf(t log.Logger, serviceType string) map[string]string {
	var f string
	u, _ := user.Current()
	if u != nil {
		f = path.Join(u.HomeDir, ".cds", serviceType+".tests.cfg.json")
	}

	if _, err := os.Stat(f); err == nil {
		t.Logf("Tests configuration read from %s", f)
		btes, err := ioutil.ReadFile(f)
		if err != nil {
			t.Fatalf("Error reading %s: %v", f, err)
		}
		if len(btes) != 0 {
			cfg := map[string]string{}
			if err := json.Unmarshal(btes, &cfg); err != nil {
				t.Fatalf("Error reading %s: %v", f, err)
			}
			return cfg
		}
	} else {
		t.Fatalf("Error reading %s: %v", f, err)
	}
	return nil
}
