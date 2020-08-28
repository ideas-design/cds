package index

import (
	"time"

	"github.com/ovh/cds/engine/gorpmapper"
	"github.com/ovh/cds/sdk"
)

var (
	TypeItemStepLog    = "StepLog"
	TypeItemServiceLog = "ServiceLog"

	StatusItemIncoming  = "Incoming"
	StatusItemCompleted = "Completed"
)

type Item struct {
	gorpmapper.SignedEntity
	ID           string           `json:"id" db:"id"`
	Created      time.Time        `json:"created" db:"created"`
	LastModified time.Time        `json:"last_modified" db:"last_modified"`
	Hash         string           `json:"-" db:"cipher_hash" gorpmapping:"encrypted,ID,ApiRefHash,Type"`
	ApiRef       sdk.CDNLogAPIRef `json:"api_ref" db:"api_ref"`
	ApiRefHash   string           `json:"api_ref_hash" db:"api_ref_hash"`
	Status       string           `json:"status" db:"status"`
	Type         string           `json:"type" db:"type"`
}
