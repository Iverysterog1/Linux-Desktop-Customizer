# Workflow security model

Validation observes; it does not publish. Recovery reconstructs; it does not publish. Integration changes arrive through reviewed commits/PRs. Publication is a distinct explicit operation after release gates. Keeping these responsibilities separate prevents a validation or recovery trigger from becoming a repository takeover path.
