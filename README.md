# Linux Desktop Customizer

> **Created by one. Improved by many. Available to all.**

![Status](https://img.shields.io/badge/status-KDE%20validation-blue)
![Safety](https://img.shields.io/badge/safety-fail--closed-brightgreen)
![Languages](https://img.shields.io/badge/languages-EN%20%7C%20pt--PT-blue)
![License](https://img.shields.io/badge/license-0BSD-green)

**Linux Desktop Customizer (LDC)** is an open-source, privacy-first application for customizing Linux desktops from one place while keeping every supported change reviewable and reversible.

The core product loop is:

**detect → preview → snapshot → apply → validate → rollback / undo**

> [!IMPORTANT]
> LDC is under active development and is **not yet a supported final release**. The rebuilt foundation is complete and the project is currently validating its first KDE Plasma vertical slice on a real Linux desktop.

## Current status

The repository foundation is stable and protected. The current development source of truth is **`integration/canonical`**; **`main`** is reserved for protected final promotion.

Completed foundations include:

- buildable Go CLI and local graphical UI;
- safe declarative profiles;
- bounded KDE adapter architecture;
- transaction journal, snapshot and rollback primitives;
- explicit preview/read-back validation;
- English and Portuguese (pt-PT) user-facing support for current flows;
- user-level install/uninstall foundations;
- CI, vet and race testing;
- least-privilege workflows with immutable Action pins;
- Security A and Security B review of the current integrated boundary.

### Current P0 — real KDE Plasma proof

The immediate gate is a real KDE Plasma session. A successful end-to-end result must prove independently:

1. **REQUESTED_SCHEME** — the intended scheme is explicitly requested.
2. **CONFIG_READBACK** — the resulting configuration is read back and matches.
3. **VISIBLE_EFFECT** — the desktop visibly reflects the requested change.
4. **ROLLBACK_CONFIG** — the exact previous configuration is restored.
5. **ROLLBACK_VISIBLE** — the desktop visibly returns to its previous state.

A configuration write alone is never enough to declare success. If configuration changes but the visible desktop does not, the result is **FAIL**. If real visual confirmation has not been performed, the result remains **NOT RUN**.

The guarded procedure is documented in [Real KDE Plasma validation](docs/REAL_PLASMA_VALIDATION.md).

## Safety model

LDC is designed around **fail-closed** behavior:

- mutation is never inferred from a read-only operation;
- unsupported or ambiguous states are reported rather than guessed;
- downloaded profiles cannot introduce arbitrary commands;
- filesystem paths and identifiers are bounded and validated;
- changes are journalled so recovery remains possible after partial failure;
- graphical mutation remains locked until the corresponding real-environment evidence exists;
- CI and validation tooling do not install packages, invoke `sudo`, publish releases or silently expand permissions.

Monitor-aware restoration is being designed with the same rule: the monitor topology must be identified deterministically and rechecked immediately before mutation. A changed or ambiguous topology must block the operation rather than risk restoring state to the wrong display.

## What works today

The current source contains a rebuilt application foundation, local UI, declarative profile path, transaction/recovery engine and a deliberately narrow KDE color-scheme integration.

The automated harness exercises read-only behavior, explicit mutation opt-in, configuration read-back, visible-result control flow, rollback and emergency recovery after partial failure. These simulated tests improve confidence in the control flow, but **do not replace a real Plasma end-to-end test**.

The graphical UI therefore reports the validation state honestly and does not expose unproven graphical mutation as a finished feature.

## Road to the first release

Development proceeds in this order:

**real Plasma validation → visible-effect hardening → safe monitor-aware restoration → complete graphical product flow → clean Linux install/run/uninstall → real screenshots → EN/pt-PT documentation review → final security/supply-chain audit → protected promotion to `main`**

No release is published automatically. Final publication requires all applicable Definition-of-Done gates and explicit project-owner authorization.

## Português (pt-PT)

O **Linux Desktop Customizer** pretende permitir personalizar o ambiente Linux num único local, mantendo as alterações suportadas verificáveis e reversíveis.

A fundação reconstruída já inclui CLI, interface gráfica local, perfis declarativos seguros, transações com rollback, integração KDE limitada, EN + pt-PT, instalação por utilizador e testes automatizados.

O projeto **ainda não é uma versão final**. A prioridade atual é provar num KDE Plasma real que uma alteração pedida é gravada, aparece realmente no ambiente gráfico e pode ser totalmente desfeita. Enquanto essa prova não existir, a mutação gráfica correspondente permanece bloqueada.

## Project governance

- **`integration/canonical`** is the single integration/development source of truth.
- **`main`** is the protected stable/release line.
- Work uses short-lived branches and returns through pull requests with required checks.
- Direct/force writes to `main` are forbidden.
- Failures and unavailable evidence remain visible as **FAIL**, **BLOCKED** or **NOT RUN**.
- Historical recovery material is evidence only and is never presented as recovered complete source.

Coordination and final-readiness gates are tracked in [Issue #47](https://github.com/Iverysterog1/Linux-Desktop-Customizer/issues/47). The rebuilt foundation milestone is recorded in [Issue #40](https://github.com/Iverysterog1/Linux-Desktop-Customizer/issues/40).

## Documentation

- [Project status](docs/PROJECT_STATUS.md)
- [Getting started with the graphical UI](docs/GETTING_STARTED_UI.md)
- [Real KDE Plasma validation](docs/REAL_PLASMA_VALIDATION.md)
- [Real Linux screenshot evidence](docs/REAL_LINUX_SCREENSHOTS.md)
- [Branch governance](docs/BRANCH_GOVERNANCE.md)
- [Security baseline](docs/SECURITY_BASELINE.md)
- [Security testing policy](SECURITY_TESTING_POLICY.md)
- [Recovery evidence](docs/RECOVERY_EVIDENCE.md)

Final GitHub screenshots must come from the real running Linux application. Mockups or generated substitutes do not satisfy the release gate.

## Contributing

Start implementation work from the latest verified **`integration/canonical`** commit. Declare ownership before implementation, keep write scopes bounded, avoid duplicate work, and submit changes through a reviewed pull request with all required checks visible.

Security-sensitive reports must never include credentials, private files or secrets.

## License

Linux Desktop Customizer is licensed under the **BSD Zero Clause License (0BSD)**. You may use, copy, modify and distribute the software for any purpose, with or without fee. See [LICENSE](LICENSE).

## Author

**JP99** — original concept, product direction and project testing.
