# Application rebuild foundation

Date: 2026-09-28

This source tree is a clean rebuild from the consolidated project contract. It is not presented as recovered historical source.

## Implemented in the foundation

- dependency-free Go module;
- CLI entry point with version/status/preview/apply/rollback/unapply;
- compileable UI entry point that reports the current foundation honestly;
- capability/adapter boundary;
- a real Customizer-owned local-file adapter;
- path traversal and symlink-escape rejection;
- atomic managed-file replacement;
- durable transaction journal written before mutation;
- exact before-state snapshots;
- automatic rollback after partial apply failure;
- explicit unsupported adapter reporting before mutation;
- idempotent Set/Unset and rollback behavior;
- user-local installer/uninstaller with no sudo;
- isolated install/uninstall smoke test;
- least-privilege CI with immutable Action pins.

## Deliberately not implemented yet

- KDE/GNOME/XFCE adapters;
- graphical desktop UI;
- imported community profiles;
- network/download/update behavior;
- privileged/system-wide changes;
- release publishing.

Those features must be layered on top of the tested transaction boundary rather than bypassing it.
