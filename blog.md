### 1. 用户认证与管理

这是所有需要用户登录的网站的基础。

| 接口                        | HTTP 方法 | 功能描述                 |
| --------------------------- | --------- | ------------------------ |
| `/api/auth/register`        | POST      | 用户注册                 |
| `/api/auth/login`           | POST      | 用户登录                 |
| `/api/auth/logout`          | POST      | 用户登出                 |
| `/api/auth/refresh-token`   | POST      | 刷新访问令牌             |
| `/api/auth/forgot-password` | POST      | 忘记密码（发送重置链接） |
| `/api/auth/reset-password`  | POST      | 重置密码                 |

### 2. 用户信息

| 接口                 | HTTP 方法 | 功能描述                            |
| -------------------- | --------- | ----------------------------------- |
| `/api/users/profile` | GET       | 获取当前登录用户信息                |
| `/api/users/profile` | PUT       | 更新当前登录用户信息                |
| `/api/users/avatar`  | POST      | 上传 / 更新用户头像                 |
| `/api/users/{id}`    | GET       | 获取指定用户信息 (管理员或用户本人) |
| `/api/users`         | GET       | 获取用户列表 (管理员)               |
| `/api/users/{id}`    | PUT       | 更新指定用户信息 (管理员)           |
| `/api/users/{id}`    | DELETE    | 删除用户 (管理员)                   |

### 3. 文章管理

这是博客的核心功能。

| 接口                         | HTTP 方法 | 功能描述                             |
| ---------------------------- | --------- | ------------------------------------ |
| `/api/articles`              | GET       | 获取文章列表（支持分页、排序、筛选） |
| `/api/articles`              | POST      | 创建新文章                           |
| `/api/articles/{id}`         | GET       | 获取单篇文章详情                     |
| `/api/articles/{id}`         | PUT       | 更新文章                             |
| `/api/articles/{id}`         | DELETE    | 删除文章                             |
| `/api/articles/{id}/publish` | POST      | 发布文章                             |
| `/api/articles/{id}/draft`   | POST      | 保存为草稿                           |
| `/api/articles/{id}/view`    | POST      | 增加文章阅读量                       |

### 4. 分类管理

| 接口                            | HTTP 方法 | 功能描述             |
| ------------------------------- | --------- | -------------------- |
| `/api/categories`               | GET       | 获取所有分类         |
| `/api/categories`               | POST      | 创建新分类           |
| `/api/categories/{id}`          | GET       | 获取指定分类详情     |
| `/api/categories/{id}`          | PUT       | 更新分类             |
| `/api/categories/{id}`          | DELETE    | 删除分类             |
| `/api/categories/{id}/articles` | GET       | 获取指定分类下的文章 |

### 5. 标签管理

| 接口                      | HTTP 方法 | 功能描述             |
| ------------------------- | --------- | -------------------- |
| `/api/tags`               | GET       | 获取所有标签         |
| `/api/tags`               | POST      | 创建新标签           |
| `/api/tags/{id}`          | GET       | 获取指定标签详情     |
| `/api/tags/{id}`          | PUT       | 更新标签             |
| `/api/tags/{id}`          | DELETE    | 删除标签             |
| `/api/tags/{id}/articles` | GET       | 获取指定标签下的文章 |

### 6. 评论管理

| 接口                          | HTTP 方法 | 功能描述             |
| ----------------------------- | --------- | -------------------- |
| `/api/articles/{id}/comments` | GET       | 获取文章下的所有评论 |
| `/api/articles/{id}/comments` | POST      | 为文章添加评论       |
| `/api/comments/{id}`          | PUT       | 更新评论             |
| `/api/comments/{id}`          | DELETE    | 删除评论             |
| `/api/comments/{id}/reply`    | POST      | 回复评论             |

### 7. 文件上传

| 接口                | HTTP 方法 | 功能描述                        |
| ------------------- | --------- | ------------------------------- |
| `/api/upload/image` | POST      | 上传图片 (文章封面、内容图片等) |
| `/api/upload/file`  | POST      | 上传文件 (附件)                 |

### 8. 网站设置

| 接口            | HTTP 方法 | 功能描述         |
| --------------- | --------- | ---------------- |
| `/api/settings` | GET       | 获取网站设置信息 |
| `/api/settings` | PUT       | 更新网站设置信息 |

### 9. 统计信息

| 接口                  | HTTP 方法 | 功能描述                            |
| --------------------- | --------- | ----------------------------------- |
| `/api/stats/articles` | GET       | 获取文章统计信息 (总数、月发布量等) |
| `/api/stats/comments` | GET       | 获取评论统计信息                    |
| `/api/stats/users`    | GET       | 获取用户统计信息                    |
| `/api/stats/views`    | GET       | 获取网站访问量统计                  |

### 通用设计原则

- **认证与授权**：使用 JWT (JSON Web Token) 进行身份验证，在请求头中携带 `Authorization: Bearer <token>`。

- **数据校验**：对所有用户输入进行严格校验，返回清晰的错误信息。

- **分页、排序与筛选**：列表接口支持 `page`, `limit`, `sort`, `order`, `search` 等参数。

- 响应格式

  ：统一的 JSON 响应格式，包含状态码、消息和数据。

  json

  

  

  

  

  

  ```json
  {
    "code": 200,
    "message": "Success",
    "data": { ... }
  }
  ```

  

- **错误处理**：统一的错误处理机制，返回合适的 HTTP 状态码和错误信息。

- **日志记录**：记录关键操作和错误日志，便于排查问题。

- 安全性

  ：

  - 使用 HTTPS
  - 防止 SQL 注入、XSS、CSRF 等攻击
  - 密码加密存储 (使用 bcrypt 等算法)
  - 合理设置接口访问频率限制 (Rate Limiting)