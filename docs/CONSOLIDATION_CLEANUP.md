# Later cleanup phase

Once consolidation validation is complete, cleanup should close superseded PRs with a pointer to the canonical integration PR, then delete only branches whose unique work is safely represented and whose deletion does not remove needed recovery provenance. This cleanup is deliberately separate from the current integration operation.
