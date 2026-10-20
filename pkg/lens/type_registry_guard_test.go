package lens_test

import (
	"testing"

	"github.com/stellar-optics/stellar-xdr-lens/pkg/lens"
)

func TestEvaluateTypeRegistryAccess(t *testing.T) {
	t.Parallel()

	res, err := lens.EvaluateTypeRegistryAccess("xdr-node-01", 20, 10, 100)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.StatusTag != "nominal" || res.ScoreRate != 3000 {
		t.Fatalf("got status=%q score=%d, want nominal/3000", res.StatusTag, res.ScoreRate)
	}

	elev, err := lens.EvaluateTypeRegistryAccess("xdr-node-02", 40, 20, 0)
	if err != nil || elev.StatusTag != "elevated" {
		t.Fatalf("got status=%q err=%v, want elevated", elev.StatusTag, err)
	}

	crit, err := lens.EvaluateTypeRegistryAccess("xdr-node-03", 90, 30, 100)
	if err != nil || crit.StatusTag != "critical" || crit.ScoreRate != 10000 {
		t.Fatalf("got status=%q score=%d err=%v, want critical/10000", crit.StatusTag, crit.ScoreRate, err)
	}

	if _, err := lens.EvaluateTypeRegistryAccess("   ", -1, -1, 100); err == nil {
		t.Fatal("expected error for blank targetKey, got nil")
	}
}
