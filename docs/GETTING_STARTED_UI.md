# Graphical UI onboarding / Início da interface gráfica

The current graphical runtime is an **early, read-only local UI**. It is not a release build and it does not yet apply desktop changes.

A interface gráfica atual é uma **interface local inicial e apenas de leitura**. Ainda não é uma versão final e ainda não aplica alterações ao ambiente de trabalho.

## Build / Compilar

From the repository root / A partir da raiz do repositório:

```bash
go build -o ./bin/ltc-ui ./cmd/ltc-ui
```

## Run safely / Executar em segurança

English:

```bash
./bin/ltc-ui --web --listen 127.0.0.1:7788 --locale en
```

Português (Portugal):

```bash
./bin/ltc-ui --web --listen 127.0.0.1:7788 --locale pt-PT
```

Then open `http://127.0.0.1:7788/` in a browser on the same machine.

Depois abra `http://127.0.0.1:7788/` num navegador no mesmo computador.

> **Safety / Segurança:** keep the listener on an explicit loopback IP (`127.0.0.1` or `::1`). The runtime intentionally rejects wildcard and hostname binds. Do not expose this development UI to the network.
>
> Mantenha o serviço num IP de loopback explícito (`127.0.0.1` ou `::1`). O programa rejeita intencionalmente endereços wildcard e nomes de host. Não exponha esta interface de desenvolvimento à rede.

## What to expect / O que esperar

- The UI presents the current navigation and desktop-capability state.
- English and Portuguese (pt-PT) are supported for the shipped UI model, with deterministic fallback to English.
- Reduced-motion preferences are respected by the graphical presentation.
- Apply/undo controls that require mutation remain gated until the KDE transaction path is fully validated.

- A interface apresenta a navegação atual e o estado das capacidades do ambiente de trabalho.
- Inglês e português (pt-PT) são suportados no modelo atual, com fallback determinístico para inglês.
- A preferência por movimento reduzido é respeitada pela apresentação gráfica.
- Controlos de aplicar/desfazer que exigem alterações permanecem bloqueados até o caminho transacional KDE estar totalmente validado.

## Validation before screenshots / Validação antes de screenshots

Before capturing project screenshots, run the required repository checks and follow [REAL_LINUX_SCREENSHOTS.md](REAL_LINUX_SCREENSHOTS.md). Screenshots are evidence only when captured from the actual running application on a real Linux graphical session; mockups and generated images do not satisfy the release gate.

Antes de capturar screenshots do projeto, execute as verificações exigidas do repositório e siga [REAL_LINUX_SCREENSHOTS.md](REAL_LINUX_SCREENSHOTS.md). As screenshots só contam como evidência quando forem capturadas da aplicação realmente em execução numa sessão gráfica Linux real; mockups e imagens geradas não cumprem o requisito de release.
