# Preserved 0.11 source-package evidence

This directory preserves the ten exact `part00.b64` through `part09.b64`
objects from historical commit
`80270fe4e7d9f01f2e4438e9d7c2bc1658088b8e` on
`publish-0.11-effects-composer`.

These files are recovery evidence, not canonical application source. Their
concatenated Base64 stream decodes to 120,000 bytes with SHA-256
`321faa7210e54178d33da1ea9f194df67e64d1ea9ff1e10a0e066010b779b1fd`,
but XZ validation fails with `Unexpected end of input`. Do not extract,
execute, splice, pad, publish, or represent this incomplete stream as a
release.

Run `bash scripts/validate-preserved-source-package.sh` from the repository
root. A non-zero result at XZ validation is the expected truthful state until
complete authoritative bytes are recovered.
