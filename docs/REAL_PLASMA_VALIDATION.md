# Real KDE Plasma validation

This runbook is the acceptance path for the first real KDE/Plasma vertical slice. Automated and in-memory tests remain useful, but they do **not** count as real-session evidence.

## Scope

The current mutation scope is intentionally narrow:

`detect -> preview -> snapshot/journal -> color-scheme apply -> config read-back -> visible confirmation -> rollback/unapply`

No panel, widget, wallpaper, window-layout or monitor-specific state is changed in this validation.

## Why this test is required

KDE documents color schemes as system-wide color definitions used by KDE/Qt applications and compatible Plasma surfaces. Applying a scheme copies values into `~/.config/kdeglobals`. KDE also ships the native `plasma-apply-colorscheme` command-line utility.

The current reviewed adapter is deliberately smaller: it writes only the allowlisted `General/ColorScheme` setting through `kwriteconfig6`. Therefore a successful config read-back is **not enough** to claim that the visible Plasma state changed. The real session must measure that distinction.

References:
- https://develop.kde.org/docs/plasma/
- https://docs.kde.org/stable_kf6/en/plasma-workspace/kcontrol/colors/index.html
- https://invent.kde.org/plasma/plasma-workspace

## Safety rules

- Run only in a real KDE Plasma session.
- Use a DEMO/test desktop session or a state you are comfortable changing temporarily.
- The harness is read-only unless both `--apply` and `LDC_REAL_PLASMA_MUTATION=YES` are supplied.
- The target must be an installed, simple color-scheme identifier.
- The harness always attempts rollback after apply.
- Do not upload `ltc status` or local evidence automatically; review it locally first because it may contain local paths.
- A missing tool/session is **BLOCKED**, not FAIL.
- A config change without visible effect is **FAIL** for the real vertical-slice requirement.
- If visible confirmation is not performed, the result is **NOT RUN** even when config apply/rollback succeeds.

## Read-only preflight

From the repository root:

```bash
scripts/validate-real-plasma.sh
```

Expected evidence includes:
- Plasma detection;
- Wayland/X11 session type;
- current `ColorScheme`;
- `kreadconfig6`, `kwriteconfig6`, and `plasma-apply-colorscheme` availability;
- local `ltc status`;
- native list of installed color schemes.

## Mutation test

Choose an installed scheme that is visibly different from the current one. For example, when the current scheme is light and `BreezeDark` is installed:

```bash
LDC_REAL_PLASMA_MUTATION=YES \
  scripts/validate-real-plasma.sh --apply --target BreezeDark
```

The harness:
1. previews the exact declarative change;
2. applies through the transaction engine;
3. records the transaction ID;
4. reads the config value back;
5. asks for a genuine visible-effect confirmation when run interactively;
6. rolls back the transaction;
7. verifies the original config value returned.

For non-interactive evidence collection, a human may explicitly supply the visible observation:

```bash
LDC_REAL_PLASMA_MUTATION=YES \
LDC_REAL_PLASMA_VISIBLE_RESULT=PASS \
  scripts/validate-real-plasma.sh --apply --target BreezeDark
```

Only use `PASS` after actually observing the visible change.

## Acceptance result

**PASS** requires all three:
- config read-back equals the requested target after apply;
- visible Plasma/application change is genuinely observed;
- rollback restores the exact original `ColorScheme` value.

**FAIL** includes:
- apply succeeds in config but no visible change occurs;
- rollback does not restore the original value;
- transaction/read-back errors.

**BLOCKED** includes:
- not running inside KDE Plasma;
- required commands missing;
- no safe alternate installed scheme available.

**NOT RUN** includes:
- only the read-only preflight was executed;
- config apply/rollback ran but visible confirmation was not performed.

## Evidence to preserve in Issue #47

Record:
- exact canonical SHA;
- Plasma version/session type;
- Wayland or X11;
- current and target scheme identifiers;
- transaction ID;
- CONFIG_APPLY / VISIBLE_APPLY / ROLLBACK results;
- final PASS / FAIL / BLOCKED / NOT RUN;
- no private screenshots or unrelated environment values.

Real screenshots for the README are a later release gate and must follow `docs/REAL_LINUX_SCREENSHOTS.md`.
