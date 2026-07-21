package testutil

import (
	"context"
	"testing"
)

func TestInstallRLPCorpusRequiresStore(t *testing.T) {
	t.Parallel()
	if _, err := InstallRLPCorpus(context.Background(), nil, nil); err == nil {
		t.Fatal("nil Store accepted")
	}
}
