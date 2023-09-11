# GitLab Jenkins Runner

GitLab Jenkins Runner 是一个用于在 GitLab CI/CD 流水线中触发 Jenkins 作业的命令行工具。它允许你轻松地将 GitLab 与 Jenkins 集成，以实现自动化构建和部署工作流程。

## 安装

在开始之前，请确保你的系统上已安装了 Go 编程语言。

使用以下命令获取和构建 `gitlab-jenkins-runner` 工具：

```bash
go install github.com/yourusername/gitlab-jenkins-runner
```

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

## 许可证

本项目基于 MIT 许可证。有关详细信息，请参阅 [LICENSE.md](LICENSE.md)。

## 支持

如果你需要帮助或有其他相关问题，请联系我们。

## 致谢

感谢使用 GitLab Jenkins Runner！
