// Package mind implementa los invariantes cross-runtime del blueprint Anima
// (https://github.com/joshuamoreno1/anima — spec doc 03, matriz C.1).
package mind

import "math"

// Plasticidad decreciente del SelfModel (spec §B.4).
// Invariante cross-runtime: misma fórmula y valores canónicos que AnimaKit (Swift).
//
//	p(n) = pMin + (1 − pMin) · e^(−n/τ)
//
// donde n = ciclos de consolidación exitosos (edad en experiencia, no wall-time).
const (
	PMin      = 0.05
	TauServer = 100.0 // perfil server (spec §B.4); el edge usa τ=30
	TauEdge   = 30.0
)

// Plasticity devuelve p(n) para el τ dado. Panics con n negativo.
func Plasticity(n int, tau float64) float64 {
	if n < 0 {
		panic("mind: cycles must be non-negative")
	}
	return PMin + (1-PMin)*math.Exp(-float64(n)/tau)
}

// Regime es el régimen del período crítico (plan doc 04 §5.5).
type Regime int

const (
	Bootstrap   Regime = iota // p ≥ 0.7 — la mente se forma
	Adolescence               // 0.3 ≤ p < 0.7
	Maturity                  // p < 0.3 — identity/values requieren aprobación del Otro
)

// RegimeFor clasifica el régimen para n ciclos con el τ dado.
func RegimeFor(n int, tau float64) Regime {
	p := Plasticity(n, tau)
	switch {
	case p >= 0.7:
		return Bootstrap
	case p >= 0.3:
		return Adolescence
	default:
		return Maturity
	}
}
