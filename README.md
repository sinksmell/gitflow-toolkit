# GitFlow ToolKit

> 这是一个使用 go 编写的简单的 GitFlow 提交工具，用于快速生成 Angular 规范 git commit 提交格式。
>
> 主要有以下用途：
>
> 	- 快速生成规范化 commit
> 	- 快速创建指定前缀的分支：例如 feature/xxx , hotfix/xxx
> 	- 简化常用操作，例如推送至远程同名分支

## 支持命令

| 命令 | 描述 |
| --- | --- |
| git ci | 交互式输入 `commit message`，可选择 feat，fix，style 等类型 |
| git cm | 接受一个 `commit message` 字符串，并校验其格式(用作 commit hook) |
| git feat xxx [branch] | 接受一个分支名和可选的追溯分支，创建一个 feature 新分支，分支名称格式为 `feature/xxx`。如果不指定追溯分支，则从当前分支创建。 |
| git fix xxx [branch] | 创建 fix 分支，类似 `git feat xxx [branch]`，但分支格式为 `fix/xxx` |
| git hotfix xxx [branch] | 创建 hotfix 分支(通常用于对 master 紧急修复)，类似 `git feat xxx [branch]，`但分支格式为 `hotfix/xxx` |
| git docs xxx [branch] | 创建 docs 分支类似 `git feat xxx [branch]`,但分支格式为 `docs/xxx` |
| git style xxx [branch] | 创建 style 分支，类似 `git feat xxx [branch]，`但分支格式为 `style/xxx` |
| git refactor xxx [branch] | 创建 refactor 分支，类似 `git feat xxx [branch]，`但分支格式为 `refactor/xxx` |
| git chore xxx [branch] | 创建 chore 分支，类似 `git feat xxx [branch]，`但分支格式为 chore/xxx` |
| git perf xxx [branch] | 创建 perf 分支，类似 `git feat xxx [branch]，`但分支格式为 `perf/xxx` |
| git style xxx [branch] | 创建 style 分支，类似 `git feat xxx [branch]，`但分支格式为 `style/xxx` |
| git ps | 推送当前分支到远程同名分支(git push origin xxxx) |

## 示例

- 创建分支

  ![image.png](https://i.loli.net/2020/07/08/Q8PkBGTZLW2cdDh.png)

- 提交信息

  ![image.png](https://i.loli.net/2020/07/08/mYvJAbfxZH1qSa6.png)



## 安装与卸载
- 手动安装

```shell
# 1. 克隆代码
git clone https://github.com/sinksmell/gitflow-toolkit.git

# 2. 安装
cd gitflow-toolkit
make install

# 卸载
make uninstall
```

- 下载 release 安装

``` shell
# 1. 到 release 标签下载对应的二进制文件压缩包，解压

# 2. 找到解压后的二进制文件，在命令行执行以下命令
# 安装
sudo gitflow-toolkit install

# 卸载
sudo gitflow-toolkit uninstall

```