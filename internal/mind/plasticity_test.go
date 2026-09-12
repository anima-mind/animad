package mind

import (
	"math"
	"testing"
)

func TestNewbornIsFullyPlastic(t *testing.T) {
	if got := Plasticity(0, TauEdge); got != 1.0 {
		t.Fatalf("p(0) = %v, want 1.0", got)
	}
}

func TestDecreasesMonotonically(t *testing.T) {
	prev := Plasticity(0, TauEdge)
	for n := 1; n <= 100; n++ {
		p := Plasticity(n, TauEdge)
		if p >= prev {
			t.Fatalf("p(%d) = %v not < p(%d) = %v", n, p, n-1, prev)
		}
		prev = p
	}
}

func TestFloorsAtPMin(t *testing.T) {
	if got := Plasticity(10_000, TauEdge); got < PMin {
		t.Fatalf("p(10000) = %v below floor %v", got, PMin)
	}
}

// Valores canónicos cross-runtime: AnimaKit (Swift) asserta EXACTAMENTE estos
// (Tests/AnimaKitTests/PlasticityTests.swift). Si este test y el de Swift
// divergen, la matriz de portabilidad (spec C.1) está rota.
func TestCanonicalCrossRuntimeValues(t *testing.T) {
	cases := []struct {
		n    int
		want float64
	}{
		{10, 0.7307047450},
		{30, 0.3994854691},
		{90, 0.0972977149},
	}
	for _, c := range cases {
		if got := Plasticity(c.n, TauEdge); math.Abs(got-c.want) > 1e-8 {
			t.Errorf("p(%d) = %.10f, want %.10f", c.n, got, c.want)
		}
	}
}

func TestRegimes(t *testing.T) {
	if RegimeFor(0, TauEdge) != Bootstrap {
		t.Error("n=0 should be Bootstrap")
	}
	if RegimeFor(20, TauEdge) != Adolescence {
		t.Error("n=20 should be Adolescence")
	}
	if RegimeFor(90, TauEdge) != Maturity {
		t.Error("n=90 should be Maturity")
	}
}
