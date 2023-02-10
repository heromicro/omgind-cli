# omgind-cli - [omgind](https://github.com/heromicro/omgind)

> omgind 辅助工具，提供创建项目、快速生成功能模块的功能

## build
```
go build -o omgind-cli main.go
```

## 下载并使用

```bash
go get -u github.com/heromicro/omgind-cli
```

### 创建项目

```bash
USAGE:
   omgind-cli new [command options] [arguments...]

OPTIONS:
   --dir value, -d value     项目生成目录(默认GOPATH)
   --pkg value, -p value     项目包名
   --branch value, -b value  指定分支(默认main)
  #  --mirror, -m              使用国内镜像(gitee.com)
   --web, -w                 包含web项目
```

> 使用示例

```bash
omgind-cli new -p test-omgind -m
```

### 生成业务模块

#### 指定模块名称和说明生成模块

```bash
USAGE:
   omgind-cli generate [command options] [arguments...]

OPTIONS:
   --dir value, -d value      项目生成目录(默认GOPATH)
   --pkg value, -p value      项目包名
   --name value, -n value     业务模块名称(结构体名称)
   --comment value, -c value  业务模块注释(结构体注释)
   --file value, -f value     指定模板文件(.yaml，模板配置可参考说明)
   --module value, -m value   指定生成模块（默认生成全部模块，以逗号分隔，支持：schema, repo, service, api, mock, router）
   --storage value, -s value  指定model的实现存储（默认gorm，支持：mongo/gorm）
```

> 使用示例

```bash
omgind-cli g -p test-omgind -n Task -c '任务管理'
```

#### 指定配置文件生成模块

```bash
omgind-cli g -p 包名 -f 配置文件(yaml)
```

> 配置文件说明

```yaml
---
name: 结构体名称
comment: 结构体注释说明
mixin: 自定义entgo mixin,目前有,sort, time, status, active, memo
fields:
  - name: 结构体字段名称
    type: 结构体字段类型
    comment: 结构体字段注释
    binding_options: binding配置项（不包含required，required由required字段控制）
    index: 数据库索引
    max: string类型最长, 数字类型为最大
    min:  string类型最少, 数字类型为最小
    
```

##### 使用示例

> 创建`task.yaml`文件

```yaml
name: Job
comment: 任务管理
mixin: "sort, time, status, active, memo"
fields:
  - name: Code
    storage: code
    type: string
    comment: 任务编号
    binding_options: ""
    index: unique
    max: 32
    min: 6

  - name: Name
    storage: name
    type: string
    comment: 任务名称
    binding_options: ""
    index: true

  - name: viewed
    storage: viewed
    type: int
    comment: 查看次数
    max: 200
    min: 88
    default: 122
    binding_options: ""

```

```bash
omgind-cli g -p test-omgind -f task.yaml
```

## MIT License

  Copyright (c) 2021 heromicro/omgind-cli
