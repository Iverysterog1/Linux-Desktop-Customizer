# Evidence policy

Use PASS only for a gate actually executed successfully against the stated commit/input. Use FAIL for executed failures, BLOCKED when required authoritative input is unavailable/invalid, and NOT RUN when a gate was not executed. Consolidation documentation must not upgrade BLOCKED/NOT RUN into PASS.
