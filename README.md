# GitLab Jenkins Runner

GitLab Jenkins Runner 是一个用于在 GitLab CI/CD 流水线中触发 Jenkins 作业的命令行工具。它允许你轻松地将 GitLab 与 Jenkins 集成，实时输出Jenkins的构建日志，以实现自动化构建和部署工作流程。

## 安装

在开始之前，请确保你的系统上已安装了 Go 编程语言。

使用以下命令获取和构建 `gitlab-jenkins-runner` 工具：

```bash
$ go install github.com/raojinlin/gitlab-jenkins-runner@latest
```

或者在[Releases](https://github.com/raojinlin/gitlab-jenkins-runner/releases)页面下载已经编译好的可执行文件。


## 使用

使用以下命令来触发 Jenkins 作业：

```bash
gitlab-jenkins-runner -base <your-jenkins-url> -job <job-name> -params PARAM1=Value1,PARAM2=Value2 -user <jenkins-user> -token <jenkins-token> -build
```

### 参数说明

- `-base <your-jenkins-url>`：Jenkins 服务器的基本 URL。
- `-job <job-name>`：要触发的 Jenkins 作业的名称。
- `-params PARAM1=Value1,PARAM2=Value2`：要传递给 Jenkins 作业的参数，以逗号分隔的键值对。
- `-user <jenkins-user>`：Jenkins 用户名。
- `-token <jenkins-token>`：Jenkins 用户令牌或密码。
- `-build`：触发 Jenkins 作业的构建。

## 示例

以下是一个示例命令，演示如何使用 `gitlab-jenkins-runner` 工具：

```bash
gitlab-jenkins-runner -base https://your-jenkins-url.com -job my-build-job -params BRANCH=main,ENV=prod -user jenkinsuser -token myapitoken -build
```

## 在 GitLab 流水线中使用

这个章节将指导你如何在 GitLab CI/CD 流水线中集成和使用 `gitlab-jenkins-runner` 命令来触发 Jenkins 构建作业。

### 步骤 1: 配置 GitLab 项目

确保你的 GitLab 项目正确配置了 CI/CD 设置以及与 Jenkins 的连接。你需要在项目设置中添加 Jenkins 服务器的 URL、认证凭据等信息。请参考 GitLab 文档以了解如何配置 GitLab 项目以与 Jenkins 集成。

### 步骤 2: 配置 GitLab CI/CD 流水线

在你的 GitLab 项目中，打开 `.gitlab-ci.yml` 文件并添加一个新的阶段，以使用 `gitlab-jenkins-runner` 工具触发 Jenkins 构建。

```yaml
# .gitlab-ci.yml

stages:
  - trigger_jenkins_build

trigger_jenkins_build:
  stage: trigger_jenkins_build
  script:
    - gitlab-jenkins-runner -base https://your-jenkins-url.com -job my-build-job -params BRANCH=$CI_MERGE_REQUEST_SOURCE_BRANCH,ENV=prod,${CI_MERGE_REQUEST_LABELS} -user jenkinsuser -token myapitoken -build

```

### 步骤 4：创建一个merge_request
创建一个merge_request

### 步骤 5：为merge_request创建标签
如果要想要Jenkins构建时使用参数，需要在merge request上设置label，label的格式如下：
```
PARAM_1=PARAM_VALUE_1
```

如：`env=PROD`

## 许可证

本项目基于 MIT 许可证。有关详细信息，请参阅 [LICENSE.md](LICENSE.md)。

## 支持

如果你需要帮助或有其他相关问题，请联系我们。

## 致谢

感谢使用 GitLab Jenkins Runner！
