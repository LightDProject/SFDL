# SFDL Language Specification

## 1. Overview

SFDL (Server File Definition Language) is a configuration language for service packaging, inspired by Terraform.

## 2. Syntax

### 2.1 Identifiers

标识符用于命名资源、块、函数等。命名规则：

- 以字母或下划线开头
- 后续字符可以是字母、数字或下划线
- 区分大小写
- 不能使用保留关键字

```SFDL
my_server    # 有效
server_1     # 有效
_private     # 有效
1server      # 无效
my-server    # 无效（连字符）
```

### 2.2 Comments

支持单行注释，以 `#` 开头到行尾。

```SFDL
# 这是单行注释
name = "server"  # 行内注释
```

### 2.3 Strings

字符串用双引号包裹，支持转义字符。

```SFDL
name = "my-server"
path = "C:\\Program Files\\app"
unicode = "你好世界"
```

转义序列：

| 序列   | 说明  |
|------|-----|
| `\n` | 换行  |
| `\t` | 制表符 |
| `\\` | 反斜杠 |
| `\"` | 双引号 |

### 2.4 Numbers

支持整数和浮点数。

```SFDL
count = 42
port = 8080
rate = 3.14
negative = -10
```

### 2.5 Lists

用方括号定义列表，元素类型可混合。

```SFDL
ports = [80, 443, 8080]
names = ["web", "api", "db"]
mixed = [1, "two", 3.0]
empty = []
```

### 2.6 Blocks

块是 SFDL 的核心结构，用于组织配置。

```SFDL
block_type "label" {
  key1 = value1
  key2 = value2

  nested {
    # 嵌套块
  }
}
```

- `block_type`：块类型（如 `provider`、`registry`、`function`）
- `label`：可选标签，用于唯一标识
- 大括号内为块属性和嵌套块

## 3. Providers

Provider 用于声明和配置外部提供者，支持源码编译和二进制分发两种模式。

### 3.1 Provider 声明

```SFDL
SFDL {
  required_providers {
    example_for_source_code = {
      source   = "git@exmaple.com:example.org/example"
      version  = "1.0.0"
      fileHash = "sha256:abc123"

      # 如果没有发布文件，Provider 将在开发时自动编译
      # version 必须为 VCS 标签
      # 目前仅支持 Go 语言
      needCompilation = true
    }
    example_for_binary = {
      source   = "https://example.com/example.elf"
      fileHash = "sha256:def456"
    }
  }
}
```

### 3.2 Provider 配置

```SFDL
provider "example_for_source_code" {
  # Provider 特定配置
}
```

### 3.3 Provider 属性

| 属性              | 类型  | 必填      | 说明                     |
|-----------------|-----|---------|------------------------|
| source          | 字符串 | 是       | 来源地址，支持 Git 仓库或 URL    |
| version         | 字符串 | 源码模式必填  | 版本号，必须为 VCS 标签         |
| fileHash        | 字符串 | 二进制模式必填 | 文件校验和，格式为 `sha256:xxx` |
| needCompilation | 布尔值 | 否       | 是否需要源码编译，默认为 false     |

## 4. Registry

Registry 用于配置远程文件注册中心，支持多种认证方式。

### 4.1 Registry 声明

```SFDL
registry "example" {
  API_root          = "https://example.com/registry"
  hashSign          = "/registry/sign/gpg"
  defaultPermission = "read-only"
  private {
    authType = "RFC5054"
  }
}
```

### 4.2 Registry 属性

| 属性                | 类型  | 必填 | 说明                                      |
|-------------------|-----|----|-----------------------------------------|
| API_root          | 字符串 | 是  | Registry API 根地址                        |
| hashSign          | 字符串 | 否  | GPG 签名验证路径                              |
| defaultPermission | 字符串 | 否  | 默认权限，可选 `read-only`、`read-write`        |
| private           | 对象  | 否  | 认证配置                                    |
| private.authType  | 字符串 | 是  | 认证类型：JWT、Basic、OAuth2、RFC5054、publicKey |

## 5. Functions

Function 用于定义可执行的处理函数，支持输入输出规范和命令执行。

### 5.1 Function 定义

```SFDL
function "serve" {
  input {
    type = "string,default 'http'"
    port = "int,default = 80"
  }
  output {
    type = "string"
    port = "int"
  }
  program {
    exec = "start"
    args = [
      "serve",
      "--type=${input.type}",
      "--port=${input.port}",
      "serve",
      input.type,
      input.port
    ]
  }
}
```

### 5.2 Input 属性

输入参数定义，支持类型标注和默认值语法。

| 语法                            | 说明        |
|-------------------------------|-----------|
| `name = "type"`               | 必填参数      |
| `name = "type,default value"` | 可选参数，带默认值 |

### 5.3 Output 属性

输出参数定义，仅包含类型声明。

### 5.4 Program 配置

| 属性   | 类型  | 必填 | 说明      |
|------|-----|----|---------|
| exec | 字符串 | 是  | 可执行命令名称 |
| args | 列表  | 是  | 命令参数列表  |

### 5.5 变量引用

在 args 中可使用 `${input.name}` 或 `input.name` 语法引用输入参数。

## 6. Examples

```SFDL
SFDL {
  required_providers {
    webserver = {
      source   = "git@example.com:org/webserver"
      version  = "1.0.0"
      needCompilation = true
    }
  }
}

registry "hub" {
  API_root = "https://registry.example.com"
  private {
    authType = "OAuth2"
  }
}

function "deploy" {
  input {
    env = "string,default 'production'"
    replicas = "int,default = 3"
  }
  output {
    status = "string"
  }
  program {
    exec = "deploy.sh"
    args = ["--env=${input.env}", "--replicas=${input.replicas}"]
  }
}

provider "webserver" {
  config = "value"
}

name = "my-server"
port = 8080

server "web" {
  host = "localhost"
}
```