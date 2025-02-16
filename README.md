# redrock-ecommerce - Go + Gin + MySQL 实现
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

## 项目概述

本项目是基于 Go 语言、Gin 框架和 MySQL 数据库实现的电商平台后端服务，旨在完成 RedRock 电商平台的基础功能。项目实现了用户注册、登录、商品浏览、购物车管理、订单创建等核心功能，并使用了 JWT 进行用户身份验证。

## 项目结构

```
├── config
│   └── database.go          # 数据库配置
├── controllers
│   ├── user.go              # 用户相关控制器
│   ├── product.go           # 商品相关控制器
│   ├── cart.go              # 购物车相关控制器
│   ├── order.go             # 订单相关控制器
│   └── comment.go           # 评论相关控制器
├── models
│   ├── user.go              # 用户模型
│   ├── product.go           # 商品模型
│   ├── cart.go              # 购物车模型
│   ├── order.go             # 订单模型
│   └── comment.go           # 评论模型
├── routers
│   ├── user.go              # 用户路由
│   ├── product.go           # 商品路由
│   ├── cart.go              # 购物车路由
│   ├── order.go             # 订单路由
│   └── comment.go           # 评论路由
├── middleware
│   └── auth.go              # 认证中间件
├── go.mod                   # Go 模块文件
├── main.go                  # 项目入口文件
└── README.md                
```

## 功能列表

### 基础功能
- **用户注册与登录**
  - 用户注册
  - 用户登录
- **用户信息管理**
  - 获取用户信息
  - 修改用户信息
  - 修改用户密码
- **商品管理**
  - 获取商品列表
  - 搜索商品
  - 获取商品详情
  - 获取分类下的商品列表
- **购物车管理**
  - 添加商品到购物车
  - 获取购物车商品列表
- **订单管理**
  - 下单
- **评论管理**
  - 获取商品评论
  - 发布评论
  - 删除评论
  - 更新评论

## 技术栈

- **编程语言**: Go
- **Web 框架**: Gin
- **数据库**: MySQL
- **身份验证**: JWT (JSON Web Token)
- **密码加密**: bcrypt
- **ORM**: GORM

## 快速开始

### 1. 环境准备
- 安装 Go (版本 1.16+)
- 安装 MySQL (版本 5.7+)

### 2. 克隆项目
```bash
git clone https://github.com/YEsoup91/redrock-ecommerce.git
cd redrock-ecommerce
```

### 3. 配置数据库
添加环境变量
```bash
set DB_DSN="your_username:password@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
```
本项目使用AutoMigrate自动调整数据库的结构

### 4. 启动服务
```bash
go run main.go
```

服务默认运行在 `http://localhost:8080`。

---

# 接口文档

## 用户模块
### 请求体/响应体结构
```json
// 注册/登录请求体
{
  "username": "string",
  "password": "string"
}

// 用户信息响应体
{
  "user_id": "uint",
  "nickname": "string",
  "email": "string",
  "gender": "string",
  "birthday": "2025/2/16"
}
```

| 接口名称       | 路径                | 方法 | 请求参数                  | 鉴权 | 说明               |
|----------------|---------------------|------|---------------------------|------|--------------------|
| 用户注册       | /user/register      | POST | 见请求体                  | 否   | 新用户注册         |
| 用户登录       | /user/login         | POST | 见请求体                  | 否   | 返回 JWT Token     |
| 获取用户信息   | /user/info/:user_id | GET  | URL参数：user_id          | 是   | 返回用户详细信息   |
| 修改密码       | /user/password      | PUT  | { "old_password": "string", "new_password": "string" } | 是 | 需原密码验证       |
| 修改用户信息   | /user/info          | PUT  | { "nickname": "string", "email": "string"... } | 是 | 可更新非敏感信息   |

---

## 产品模块
### 响应体示例
```json
{
  "products": [
    {
      "id": 1,
      "name": "商品名称",
      "price": 99.9,
      "cover": "https://example.com/cover.jpg",
      "comment_num": 5,
      "publish_time": "2025/2/16"
    }
  ]
}
```

| 接口名称         | 路径                     | 方法 | 请求参数                          | 鉴权 | 说明                     |
|------------------|--------------------------|------|-----------------------------------|------|--------------------------|
| 获取产品列表     | /product/list            | GET  | -                                 | 否   | 分页参数需在代码实现     |
| 搜索产品         | /product/search          | GET  | Query参数：`keyword=手机`         | 否   | 模糊匹配名称/描述        |
| 获取产品详情     | /product/info/:product_id| GET  | URL参数：product_id               | 否   | 包含完整商品描述和链接   |
| 按类型筛选产品   | /product/:type           | GET  | URL参数：type（如`electronics`）  | 否   | 需预先定义商品类型       |

---

## 购物车模块
### 请求体/响应体
```json
// 添加购物车请求体
{
  "product_id": 123,
  "quantity": 2
}

// 查看购物车响应体
{
  "products": [
    {
      "id": 123,
      "name": "商品名称",
      "price": 199.0,
      "quantity": 2,
      "total": 398.0
    }
  ]
}
```

| 接口名称       | 路径          | 方法 | 请求参数                  | 鉴权 | 说明               |
|----------------|---------------|------|---------------------------|------|--------------------|
| 添加商品到购物车 | /cart/add     | POST | 见请求体                  | 是   | 重复添加会覆盖数量   |
| 查看购物车      | /cart         | GET  | -                         | 是   | 返回商品列表及小计   |

---

## 评论模块
### 请求体/响应体
```json
// 发布评论请求体
{
  "content": "这个商品非常好！"
}

// 评论响应体
{
  "id": 1,
  "content": "评论内容",
  "user_id": 1001,
  "publish_time": "2025/2/16",
  "praise_count": 5
}
```

| 接口名称       | 路径                     | 方法 | 请求参数                  | 鉴权 | 说明               |
|----------------|--------------------------|------|---------------------------|------|--------------------|
| 获取商品评论   | /comment/:product_id     | GET  | URL参数：product_id       | 否   | 按时间倒序排列     |
| 发布评论       | /comment/:product_id     | POST | 见请求体                  | 是   | 返回新评论ID       |
| 修改评论       | /comment/:comment_id     | PUT  | 新评论内容                | 是   | 仅限作者操作       |
| 删除评论       | /comment/:comment_id     | DELETE| -                        | 是   | 软删除记录         |

---

## 订单模块
### 请求体/响应体
```json
// 下单请求体
{
  "items": [
    { "product_id": 123, "quantity": 1 }
  ]
}

// 订单响应体
{
  "order_id": "20230901123456",
  "total_price": 299.0,
  "status": "pending",
  "created_at": "2025/2/16"
}
```

| 接口名称       | 路径          | 方法 | 请求参数                  | 鉴权 | 说明               |
|----------------|---------------|------|---------------------------|------|--------------------|
| 下单           | /order        | POST | 商品列表                  | 是   | 生成待支付订单     |

---

## 全局说明
### 1. 鉴权方式
需在请求头中添加：
```http
Authorization: Bearer <JWT_TOKEN>
```

### 2. 状态码说明
| 状态码 | 说明                  |
|--------|-----------------------|
| 401    | 未提供有效 Token      |
| 403    | 无操作权限            |
| 404    | 资源不存在            |
| 500    | 服务器内部错误        |


### 3. 时间格式
统一使用 **2025/2/16** 格式。

---

## 第三方库说明

- **Gin**: 高性能 HTTP Web 框架，用于快速构建 RESTful API。
- **GORM**: ORM 库，用于简化数据库操作。
- **JWT**: 用于用户身份验证和授权。
- **bcrypt**: 用于密码加密存储，确保用户密码安全。

## 贡献指南

欢迎提交 Issue 和 Pull Request。请在提交代码前确保通过所有测试，并遵循项目的代码风格。

## 许可证

本项目采用 MIT 许可证。详情请参阅 [LICENSE](LICENSE) 文件。

---
**这是本人的第一个项目，感谢redrock，bilibili，github，csdn以及各种AI大模型在我的学习和开发过程中的支持和帮助！**
