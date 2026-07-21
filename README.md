# Remediation Demo

Demo repository for validating the centralized dependency remediation workflow.

The repository intentionally pins `lodash` to `4.17.19` so the remediation workflow can detect vulnerabilities, update dependency files, and create a Pull Request.

## Pipeline Test

Run the `Dependency Remediation` workflow manually from GitHub Actions. Expected behavior:

1. the workflow calls `opsbento/security-workflows`;
2. `security-workflows` builds and runs `opsbento/remediation-core`;
3. dependency files are updated on a deterministic branch;
4. one Pull Request is created or updated;
5. SBOM and scan reports are uploaded as workflow artifacts.

Do not run the end-to-end remediation test locally; this repository exists for pipeline verification.
