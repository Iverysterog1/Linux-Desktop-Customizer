# Source recovery evidence

GitHub is the only authoritative storage still available for this project. No PC/local/external source tree exists.

## Route 1 — preserved 0.11 stream

The ten preserved `source-package` fragments reconstruct 120,000 bytes with SHA-256 `321faa7210e54178d33da1ea9f194df67e64d1ea9ff1e10a0e066010b779b1fd`. Strict XZ validation fails with `Unexpected end of input`.

**Result: FAIL — truncated.**

## Route 2 — historical 56-object archive

The pinned archive manifest cannot be reconstructed because Git blob `cfd027fdf53fcd403a020f7a934b74d0febdc98c` is not reachable from GitHub (HTTP 404). Reconstruction stops before complete SHA-256/XZ/source validation.

**Result: BLOCKED/FAIL — authoritative object missing.**

## Route 3 — older three-part source archive

The surviving `bootstrap/source.part.00..02.b64` blobs were re-tested using only GitHub-hosted bytes in Actions run `36413607567`.

The reconstructed archive does **not** match its historical pinned SHA-256:

`444f8803b764e2e7f709185428d1b519a44cd902291c40aef891cee8992413bd`

The job stopped before extraction. Accessible history contains no `source.part.03+`.

**Result: FAIL — incomplete/mismatched stream.**

## Readable patch route

Accessible history contains only `import-parts/part-000.patch`; no additional readable patch parts were found. It is preserved as evidence but is not a complete application tree.

## Conclusion

The complete original application source is not reconstructible from the authoritative GitHub bytes that remain. Do not splice unrelated streams, guess missing bytes or relabel partial archives as recovered source.

The project now proceeds by **clean rebuild from the consolidated requirements and surviving evidence**. Historical fragments remain provenance/reference material only.
