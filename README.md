# animad

**Perfil server del harness [Anima](https://github.com/anima-mind/anima)** — el runtime Go de la misma mente: daemon y/o servicio REST. Brain compartido, consolidación continua, capa intersubjetiva (spec §B.9). El Otro = la organización/el operador.

> Una mente = LLM (dotación) + harness (desarrollo) + historia (experiencia).
> Un blueprint, N cuerpos: [`anima-ios`](https://github.com/anima-mind/anima-ios) es el cuerpo edge; `animad` es el cuerpo server.

## Estado

**Pre-implementación.** El repo arranca por el primer invariante cross-runtime del blueprint: la [fórmula de plasticidad](internal/mind/plasticity.go) del SelfModel (spec §B.4), con los **mismos valores canónicos** que asserta `AnimaKit` en Swift — la matriz de portabilidad (spec §C.1) probada desde el commit 1. El plan de implementación del perfil server (hermano del doc 04) se escribirá en el blueprint antes de la Fase 0 de este runtime.

## Dev

```bash
go build ./... && go test -race ./...
go run ./cmd/animad -version
```

- Módulo: `github.com/anima-mind/animad` · Go 1.26.
- CI: build + vet + test -race + gofmt en cada push/PR.
- CD (futuro): goreleaser en tags para binarios multi-plataforma del daemon.

## Licencia

MIT
