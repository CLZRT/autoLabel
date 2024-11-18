# Just  Label some gcp resources with audit log.

用Terraform 创建 GCP 资源 用于 在创建资源时打上资源标签

### 1.克隆本项目到本地 并进入工作目录

```shell
git clonehttps://github.com/CLZRT/autoLabel.git && cd autoLabel
```

### 2. 初始化 Terraform环境 

```shell
terraform init
```

### 3. 创建 告警政策

```shell
terraform apply
```

### 4. 输入关键参数

> var.region "部署autoLabel的区域"
>
> var.project_id "用于部署autoLabel的项目id"
>

### 5. 删除相关资源

```shell
terraform destroy
```
