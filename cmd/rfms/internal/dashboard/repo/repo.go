package repo

import "github.com/way-platform/rfms-go/v4"

// Repo is a repo for rFMS dashboard data.
type Repo struct {
	client *rfms.Client
}

// New creates a new [Repo].
func New(client *rfms.Client) *Repo {
	return &Repo{
		client: client,
	}
}
