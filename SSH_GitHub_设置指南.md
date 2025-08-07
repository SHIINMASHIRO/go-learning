# SSH 连接 GitHub 设置指南

## 第一步：检查是否已有SSH密钥

打开终端，运行以下命令检查是否已经有SSH密钥：

```bash
ls -la ~/.ssh
```

如果看到 `id_rsa` 和 `id_rsa.pub` 文件，说明你已经有SSH密钥了，可以跳到第三步。

## 第二步：生成新的SSH密钥

如果没有SSH密钥，需要生成一个新的：

```bash
ssh-keygen -t ed25519 -C "你的邮箱@example.com"
```

**注意**：把 `你的邮箱@example.com` 替换为你的GitHub邮箱地址

当提示输入文件位置时，直接按回车使用默认位置。
当提示输入密码时，可以直接按回车（不设密码）或设置一个安全密码。

## 第三步：启动SSH代理并添加密钥

```bash
# 启动ssh-agent
eval "$(ssh-agent -s)"

# 添加SSH密钥到ssh-agent
ssh-add ~/.ssh/id_ed25519
```

如果你使用的是旧版本的RSA密钥，使用：
```bash
ssh-add ~/.ssh/id_rsa
```

## 第四步：复制SSH公钥

```bash
# 对于ed25519密钥
cat ~/.ssh/id_ed25519.pub

# 对于RSA密钥
cat ~/.ssh/id_rsa.pub
```

复制输出的整个公钥内容。

## 第五步：在GitHub中添加SSH密钥

1. 登录GitHub网站
2. 点击右上角头像 → Settings
3. 在左侧菜单中点击 "SSH and GPG keys"
4. 点击 "New SSH key"
5. 在 "Title" 中输入一个描述性名称（如："我的Mac电脑"）
6. 在 "Key" 框中粘贴刚才复制的公钥
7. 点击 "Add SSH key"

## 第六步：测试SSH连接

```bash
ssh -T git@github.com
```

如果设置成功，你会看到类似这样的消息：
```
Hi 你的用户名! You've successfully authenticated, but GitHub does not provide shell access.
```

## 第七步：配置Git用户信息

```bash
git config --global user.name "你的用户名"
git config --global user.email "你的邮箱@example.com"
```

## 第八步：在GitHub创建新仓库

1. 登录GitHub
2. 点击右上角的 "+" → "New repository"
3. 输入仓库名称，如：`go-learning`
4. 选择 "Public" 或 "Private"
5. 勾选 "Add a README file"
6. 点击 "Create repository"

## 第九步：克隆仓库到本地

复制仓库的SSH地址（以git@github.com开头），然后运行：

```bash
git clone git@github.com:你的用户名/go-learning.git
cd go-learning
```

## 日常使用：上传代码的基本流程

```bash
# 添加文件到暂存区
git add .

# 提交更改
git commit -m "第1章：Go基础语法"

# 推送到GitHub
git push origin main
```

## 按章节组织代码的建议结构

```
go-learning/
├── chapter01-basics/
│   ├── hello.go
│   ├── variables.go
│   └── README.md
├── chapter02-control-flow/
│   ├── if-else.go
│   ├── loops.go
│   └── README.md
├── chapter03-functions/
│   ├── basic-functions.go
│   ├── advanced-functions.go
│   └── README.md
└── README.md
```

## 常用Git命令备忘

```bash
# 查看当前状态
git status

# 查看提交历史
git log --oneline

# 创建新分支
git checkout -b new-feature

# 切换分支
git checkout main

# 合并分支
git merge new-feature

# 拉取最新更改
git pull origin main
```

## 故障排除

### 如果SSH连接失败
1. 确保公钥已正确添加到GitHub
2. 检查SSH代理是否运行：`ssh-add -l`
3. 重新添加密钥：`ssh-add ~/.ssh/id_ed25519`

### 如果推送失败
1. 确保你有仓库的写权限
2. 先拉取最新更改：`git pull origin main`
3. 解决冲突后再推送

现在你就可以开始你的Go学习之旅了！🚀 