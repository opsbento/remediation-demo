# Remediation Demo

Demo repository for validating the centralized dependency remediation workflow.

The repository intentionally pins mixed dependency sets so the workflows can show vulnerable and non-vulnerable packages across ecosystems.

Current direct dependencies include intentionally old packages such as `lodash`, `axios`, `express`, and `minimist`, plus ordinary packages such as `dayjs` and `nanoid`.

Fixtures:

- npm root project for scan and remediation.
- `go-service` for Go dependency scan coverage.
- `python-service` for Python dependency scan coverage.

## Pipeline Test

The `Dependency Remediation` workflow runs daily and can also be started manually from GitHub Actions. Expected behavior:

1. the workflow calls `opsbento/security-workflows`;
2. `security-workflows` downloads and runs the pinned `opsbento/remediation-core` CLI release;
3. dependency files are updated on a deterministic branch;
4. one Pull Request is created or updated;
5. remediation and scan results are printed to workflow logs and job summaries.

Do not run the end-to-end remediation test locally; this repository exists for pipeline verification.

## Current Workflow Inputs

The demo uses:

```yaml
remediation-core-version: v0.2.6
security-workflows-ref: v1
upload-artifacts: false
maximum-updates: 25
pr-labels: ""
```

Set the repository secret `REMEDIATION_TOKEN` when the default `GITHUB_TOKEN` is not allowed to create Pull Requests.

Pull requests from `remediation/*` branches skip the generic dependency scan gate. Those PRs rely on the remediation workflow verification, while normal pull requests and manual scans still run the blocking dependency scan.

The generic `Dependency Scan` workflow runs on normal dependency pull requests, package file changes pushed to `main`, and manual dispatch.

The scan workflow runs as a matrix over npm, Go, and Python fixtures. Automated remediation currently targets the npm root project; Go and Python remediation adapters are future work.
