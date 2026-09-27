# Workflow permission target

Repository validation workflows should normally declare `permissions: contents: read`. Any future write permission requires a narrow documented purpose and review; default-branch mutation must not be hidden inside validation/recovery jobs.
