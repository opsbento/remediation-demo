# Remediation Demo

Demo repository for validating the centralized dependency remediation workflow.

The repository intentionally pins a mixed npm dependency set so the workflows can show vulnerable and non-vulnerable packages in one run.

Current direct dependencies include intentionally old packages such as `lodash`, `axios`, `express`, and `minimist`, plus ordinary packages such as `dayjs` and `nanoid`.

## Pipeline Test

Run the `Dependency Remediation` workflow manually from GitHub Actions. Expected behavior:

1. the workflow calls `opsbento/security-workflows`;
2. `security-workflows` downloads and runs the pinned `opsbento/remediation-core` CLI release;
3. dependency files are updated on a deterministic branch;
4. one Pull Request is created or updated;
5. remediation and scan results are printed to workflow logs and job summaries.

Do not run the end-to-end remediation test locally; this repository exists for pipeline verification.

## Current Workflow Inputs

The demo uses:

```yaml
remediation-core-version: v0.1.1
security-workflows-ref: main
upload-artifacts: false
pr-labels: ""
```

Set the repository secret `REMEDIATION_TOKEN` when the default `GITHUB_TOKEN` is not allowed to create Pull Requests.
