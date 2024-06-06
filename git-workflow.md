# git-workflow

## 基本概念

### 1.Branch

### 2.Remote

### 3.Local

### 4.Disk

### 5.Commit

## 流程

### 1.Remote2Local

```shell
git clone https://github.com/clzrt/autolabel

## 复制了当前branch到新branch上
git checkout -b login

## 改变了什么
git diff
##  告知git 想要修改文件
git add file
#git commit(local git多了个commit)
git commit
# 同步main-branch并更新login-branch到main-branch
git checkout main

git pull origin main

git checkout login
## 把main-branch上的修改作为基础，添加login-franch上的修改，可能出现（rebase-conflict)
git rebase main
# 测试是否可跑通
go run .
## 与merge比较

## 同步login-brench到remote(由于使用rebase导致localgit上的login的commit与remote的commit不同，需要加上 -force)
git push -f origin login(push之前 一定要将远程的main分支同步下来)

## 将remote上的login-branch合并到main-branch New Pull Request
Squash and merge(将此分支上的所有commit合并为一个commit)

# 删除本地分支
git branch -D login
#同步远程主分支 到 本地
git pull origin main


# github 远程新建分支 同步到本地
git fetch origin <branch_name>
git checkout -b <branch_name> origin/<branch_name>
```

### 2.Local2Disk

3.Remote2Disk

### 3.Disk2Local

### 4.Local2Remote

### New Branch

```shell
git checkout -b login
```

# cluster(true),batch(true),job(true),session(true)