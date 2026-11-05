# gosec-demo

**Gosec · pre-commit + workflow 全量扫描** 的 GitHub Actions 演示仓，属于 `openlibing-test` 组织。扫描结果通过 `openlibing-test/pre-commit-action` 与 `openlibing-test/upload-sarif-action`（AK/SK 模式）上传 OBS 并回调 openLiBing 平台。

与 openLiBing 文档对照：[gosec 接入指南（GitHub 章节）](https://gitcode.com/openlibing/docs/blob/main/static-code-analysis/solutions/full-analysis/tools/gosec/README.md)

## 两种方案（本仓都有）

| 方案 | workflow 文件 | 要不要 pre-commit |
|------|-----------------|-------------------|
| **pre-commit** | `.github/workflows/gosec-precommit.yml` | 要，读 `.pre-commit-config.yaml` |
| **workflow 全量** | `.github/workflows/gosec-scan.yml` | 不要，直接 `gosec ./...` |

共用 **`gosec.json`**（规则配置）与 **`main.go`**（故意保留可扫问题：弱哈希 G401、硬编码 token G101）。

## 仓库里有什么

| 路径 | 作用 |
|------|------|
| `.pre-commit-config.yaml` | 仅 **pre-commit 方案** 使用；hook `entry` 固定 `.github/bin/gosec` |
| `.github/workflows/gosec-precommit.yml` | 场景一：setup → install gosec 到 `.github/bin` → pre-commit-action → upload-sarif-action |
| `.github/workflows/gosec-scan.yml` | 场景二：checkout → setup-go → gosec 全量 → upload-sarif-action |
| `gosec.json` | 两场景共用规则 |
| `main.go` | 故意保留可扫问题 |

## 前置条件

1. **组织内 action 仓**：`openlibing-test/pre-commit-action`、`openlibing-test/upload-sarif-action` 需存在于本组织（公开，或私有 + 本仓同组织即可解析）；
2. **Secrets**：仓库 `Settings → Secrets and variables → Actions` 配置 `OBS_AK`、`OBS_SK`。

## 怎么测

1. push 到 `main`（或 Actions 页手动 Run workflow）；
2. **场景一**：看 job `pre-commit-report`；pre-commit 步骤扫到问题显示 Failed 属正常（hook 未加 `-no-fail`，用于拦截语义），`if: always()` 保证 SARIF 仍上传；
3. **场景二**：看 job `gosec-scan-upload`；已加 `-no-fail`，一般全绿，SARIF 里仍有 `results`；
4. openLiBing 平台「开发管理 - 代码检查」按 仓库 = gosec-demo、扫描类别 = business 查看。

## 可选：本地 pre-commit（Linux / WSL / macOS）

```bash
mkdir -p .github/bin /tmp/pre-commit-reports
GOBIN="$(pwd)/.github/bin" go install github.com/securego/gosec/v2/cmd/gosec@v2.22.4
pip install pre-commit
pre-commit run gosec --all-files
```

Windows 原生 shell 需改 hook 里 `-out=` 路径（见 `.pre-commit-config.yaml` 顶部注释）。
