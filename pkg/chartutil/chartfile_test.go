package chartutil

import (
	"testing"

	"k8s.io/helm/pkg/proto/hapi/chart"
)

const testfile = "testdata/chartfiletest.yaml"

func TestLoadChartfile(t *testing.T) {
	f, err := LoadChartfile(testfile)
	if err != nil {
		t.Errorf("Failed to open %s: %s", testfile, err)
		return
	}
	verifyChartfile(t, f)
}

func verifyChartfile(t *testing.T, f *chart.Metadata) {

	if f == nil {
		t.Fatal("Failed verifyChartfile because f is nil")
	}

	if f.Name != "frobnitz" {
		t.Errorf("Expected frobnitz, got %s", f.Name)
	}

	if f.Description != "This is a frobnitz." {
		t.Errorf("Unexpected description %q", f.Description)
	}

	if f.Version != "1.2.3" {
		t.Errorf("Unexpected version %q", f.Version)
	}

	if len(f.Maintainers) != 2 {
		t.Errorf("Expected 2 maintainers, got %d", len(f.Maintainers))
	}

	if f.Maintainers[0].Name != "The Helm Team" {
		t.Errorf("Unexpected maintainer name.")
	}

	if f.Maintainers[1].Email != "nobody@example.com" {
		t.Errorf("Unexpected maintainer email.")
	}

	if len(f.Sources) != 1 {
		t.Fatalf("Unexpected number of sources")
	}

	if f.Sources[0] != "https://example.com/foo/bar" {
		t.Errorf("Expected https://example.com/foo/bar, got %s", f.Sources)
	}

	if f.Home != "http://example.com" {
		t.Error("Unexpected home.")
	}

	if len(f.Keywords) != 3 {
		t.Error("Unexpected keywords")
	}

	kk := []string{"frobnitz", "sprocket", "dodad"}
	for i, k := range f.Keywords {
		if kk[i] != k {
			t.Errorf("Expected %q, got %q", kk[i], k)
		}
	}

}
