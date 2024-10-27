## DDCode-基于TailwindCss的代码生成系统-golang

配套前端：https://github.com/hasdsde/ddCode_vue

### 主要功能

- 基于树状图、符合程序员逻辑
- 使用TailwindCss，不再写CSS
- 基于Node服务器对本地文件读取，直接渲染页面
- 读取代码，自动识别分类变量、函数、引入
- 使用Quasar框架，支持开箱即用图标、组件更丰富

### 系统功能

- 基本权限管理功能
- Material Icon速览
- 读取Swagger结构体，快速生成表格和示例数据
- Echart基本功能生成、在线预览

## 计划表

- [x] 用户管理
	- [x] Jwt验证
	- [ ] 登录验证码
	- [x] 用户密码加密
	- [ ] 多次登录限制

- [x] 部门管理
- [x] 岗位管理
- [x] 角色管理
- [x] 权限管理
	- [x] 自定义注解

- [x] 菜单管理
	- [x] 前后端同步

- [x] 全局日志
	- [x] 自定义注解实现
	- [x] 操作日志
	- [ ] 登录日志
- [ ] 定时任务
- [x] SWAGGER接口
- [x] 系统监控
- [x] 缓存监控
	- [x] Druid监控
- [x] 文件下载上传
- [ ] 通知公告
- [x] 系统配置
- [ ] XSS
- [x] 代码生成器
	- [x] 后端sql代码
	- [x] MBS代码生成器
	- [x] 基于前端TailWindCSS低代码生成
	- [x] Echarts可视化
- [ ] NPM CLi

## 预览

登录页

![https://raw.githubusercontent.com/hasdsde/img_bed/main/QQ%E6%88%AA%E5%9B%BE20240326203612.png](https://raw.githubusercontent.com/hasdsde/img_bed/main/QQ%E6%88%AA%E5%9B%BE20240326203612.png)

管理页面

![https://raw.githubusercontent.com/hasdsde/img_bed/main/QQ%E6%88%AA%E5%9B%BE20240326203637.png](https://raw.githubusercontent.com/hasdsde/img_bed/main/QQ%E6%88%AA%E5%9B%BE20240326203637.png)

![https://raw.githubusercontent.com/hasdsde/img_bed/main/QQ%E6%88%AA%E5%9B%BE20240326203657.png](https://raw.githubusercontent.com/hasdsde/img_bed/main/QQ%E6%88%AA%E5%9B%BE20240326203657.png)

![https://raw.githubusercontent.com/hasdsde/img_bed/main/QQ%E6%88%AA%E5%9B%BE20240326203857.png](https://raw.githubusercontent.com/hasdsde/img_bed/main/QQ%E6%88%AA%E5%9B%BE20240326203857.png)

代码生成器

![https://raw.githubusercontent.com/hasdsde/img_bed/main/QQ%E6%88%AA%E5%9B%BE20240326203718.png](https://raw.githubusercontent.com/hasdsde/img_bed/main/QQ%E6%88%AA%E5%9B%BE20240326203718.png)

![https://raw.githubusercontent.com/hasdsde/img_bed/main/QQ%E6%88%AA%E5%9B%BE20240326203735.png](https://raw.githubusercontent.com/hasdsde/img_bed/main/QQ%E6%88%AA%E5%9B%BE20240326203735.png)

![https://raw.githubusercontent.com/hasdsde/img_bed/main/QQ%E6%88%AA%E5%9B%BE20240326203752.png](https://raw.githubusercontent.com/hasdsde/img_bed/main/QQ%E6%88%AA%E5%9B%BE20240326203752.png)

![https://raw.githubusercontent.com/hasdsde/img_bed/main/QQ%E6%88%AA%E5%9B%BE20240326203811.png](https://raw.githubusercontent.com/hasdsde/img_bed/main/QQ%E6%88%AA%E5%9B%BE20240326203811.png)

![https://raw.githubusercontent.com/hasdsde/img_bed/main/QQ%E6%88%AA%E5%9B%BE20240326203829.png](https://raw.githubusercontent.com/hasdsde/img_bed/main/QQ%E6%88%AA%E5%9B%BE20240326203829.png)

## 配置文件

`config/config.yml`
