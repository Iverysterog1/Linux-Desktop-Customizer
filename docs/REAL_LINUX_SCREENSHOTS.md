# Real Linux screenshot evidence

Final project screenshots must come from the **actual Linux Desktop Customizer binary running in a real Linux graphical session**. Mockups, generated images, browser-devtool substitutions and screenshots of unrelated prototypes do not satisfy this gate.

## Preconditions

1. Work from the exact reviewed `integration/canonical` commit intended for validation.
2. Record the commit SHA, Linux distribution/version, desktop environment/version, display protocol (Wayland/X11), and application build command.
3. Run the repository's required unit/integration, vet, race, build and installer checks applicable to that commit. Preserve failures; do not capture screenshots as if a failed build were release-ready.
4. Build `ltc-ui` from that same commit.
5. Start the graphical runtime on explicit loopback only, for example:

```bash
./bin/ltc-ui --web --listen 127.0.0.1:7788 --locale en
```

## Required captures

Capture the real application at a readable desktop scale. At minimum collect:

- English graphical home/navigation state.
- Portuguese (pt-PT) graphical home/navigation state.
- A capability or unsupported/gated state that demonstrates honest desktop support reporting.
- The preview/review/apply/undo flow only after those controls are genuinely enabled and validated; do not stage or fake them.
- Reduced-motion/accessibility presentation when it is visually meaningful and implemented.

Avoid including terminals, usernames, home-directory paths, notifications, browser profiles, tokens, email addresses, private files or other personal information unless they are essential evidence and have been sanitized safely.

## Repository destination

Commit approved captures under:

```text
docs/screenshots/
```

Use descriptive stable names such as:

```text
home-en.png
home-pt-PT.png
capability-state.png
preview-review.png
```

Do **not** add placeholder image files merely to make the directory exist.

## Evidence note

For each screenshot set, add a short note to the coordination ledger (Issue #47) containing:

- exact application commit SHA;
- Linux distribution and version;
- KDE Plasma/desktop version;
- Wayland or X11;
- command used to build and launch;
- validation run IDs/results available for that exact head;
- screenshot filenames;
- PASS / FAIL / BLOCKED classification.

A screenshot proves only what is visibly demonstrated. It does not replace automated tests, real KDE mutation/rollback validation, clean install/uninstall testing, security review or release authorization.
