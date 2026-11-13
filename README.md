# gosec-demo

`openlibing-test` 组织的 **Gosec GitHub Actions 验证仓**，与 openLiBing 文档 [gosec 指南 · GitHub 接入](https://gitcode.com/openlibing/docs/blob/main/static-code-analysis/solutions/full-analysis/tools/gosec/README.md) 及 **`demos/github/`** 样例对齐。

## 两种方案（本仓都有）

| 方案 | workflow 文件 | 要不要 pre-commit |
|------|-----------------|-------------------|
| **pre-commit** | `.github/workflows/gosec-precommit.yml` | 要，读 `.pre-commit-config.yaml` |
| **workflow 全量** | `.github/workflows/gosec-scan.yml` | 不要，直接 `gosec ./...` |

共用 **`gosec.json`**（与文档 `demos/gosec.json` 一致）与 **`main.go`**（故意保留可扫问题：弱哈希 G401、硬编码 token G101）。

## 仓库里有什么

| 路径 | 作用 |
|------|------|
| `.pre-commit-config.yaml` | 与 `demos/github/.pre-commit-config.yaml` 一致 |
| `.github/workflows/gosec-precommit.yml` | 与 `demos/github/gosec-precommit.yml` 一致（本仓 upload 固定 `env_type: beta`） |
| `.github/workflows/gosec-scan.yml` | 与 `demos/github/gosec-scan.yml` 一致（同上） |
| `gosec.json` | 两场景共用规则 |
| `main.go` | 演示用可扫代码 |

## 前置条件

1. 组织内 Action：`openlibing-test/pre-commit-action@main`、`openlibing-test/upload-sarif-action@source-platform-detection` 可解析（本组织公开仓即可）；
2. **上传认证**：workflow 已声明 `permissions: id-token: write`，使用 **OIDC 免密换证**，**无需**配置 `OBS_AK` / `OBS_SK`。

## 怎么测

1. push 到 `main`，或在 Actions 页手动 **Run workflow**；
2. **场景一**（`gosec-precommit`）：看 job `pre-commit-report`；pre-commit 扫到问题显示 Failed 属正常，SARIF 仍会通过 `if: always()` 上传；
3. **场景二**（`gosec-upload-sarif`）：看 job `Gosec 扫描 - SARIF - 上传`；扫描步使用 `-no-fail`，job 一般全绿；
4. openLiBing **测试环境**（`env_type: beta`）「开发管理 - 代码检查」：仓库 = **gosec-demo**、扫描类别 = **business**、分支 = 触发分支。

## 与文档 demo 的唯一刻意差异

| 项 | 文档 `demos/github/*.yml` | 本验证仓 |
|----|---------------------------|----------|
| `env_type` | 默认 prod（注释掉） | 固定 `beta`，便于在测试环境验收 |

业务仓落地时复制 `demos/github/` 即可，一般删掉 `env_type: beta` 行。

## 可选：本地 pre-commit（Linux / WSL / macOS）

```bash
mkdir -p .github/bin /tmp/pre-commit-reports
GOBIN="$(pwd)/.github/bin" go install github.com/securego/gosec/v2/cmd/gosec@v2.22.4
pip install pre-commit
pre-commit run gosec --all-files
```

Windows 原生 shell 需改 hook 里 `-out=` 路径（见文档 FAQ）。
