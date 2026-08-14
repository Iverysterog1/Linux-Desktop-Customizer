# Security testing policy

Public CI keeps ordinary unit, integration, fuzz-regression and security-regression tests visible so contributors can reproduce fixes and verify changes.

Private security work is reserved for material that would disclose real secrets, private infrastructure, unpublished exploit chains, weaponized proof-of-concept details, or internal red-team methodology that is not required to reproduce a public regression.

## Public

- Unit and integration tests
- Race-detector tests
- Fuzz regression tests and non-weaponized fuzz seeds
- Path traversal, symlink, archive, SSRF, rollback and input-validation regressions
- Secret-redaction tests that use synthetic credentials
- Build, packaging and reproducibility checks

## Private

- Real credentials, tokens, infrastructure details or private endpoints
- Unpublished exploit chains and weaponized proof-of-concept material
- Internal red-team playbooks and attack automation
- Sensitive incident material before coordinated disclosure

A vulnerability fix should gain a public regression test whenever the test can be safely expressed without exposing sensitive material.
