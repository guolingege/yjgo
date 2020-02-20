## 框架简介

以前是做Java企业应用开发，想转Go,看了很多优秀的开源项目但是发现没有合适的，一直没找到类似于若依开发思路的快速开发框架。于是狠下心自己写了一套后台系统。这个框架可以用于所有的Web应用程序，如网站管理后台，网站会员中心，CMS，CRM，OA。所有前端后台代码封装过后十分精简易上手，出错概率低。同时支持移动客户端访问。系统会陆续更新一些实用功能。

受之开源，回馈社区，本框架在GoFrame为基础库，继续沿用MIT开源协议，架构思路沿袭着若依的以辅助生成重复代码为主，不过度封装，生成的代码可以快速修改适应不同的需求，适应每个开发者自己的习惯和风格。所以命名为"云捷GO"（寓意：GO云原生业务快捷开发）

> 如需前后端分离版本，正在酝酿，打算以ant design pro为前端。

> 阿里云通用云产品1888优惠券 ：[点我领取](https://www.aliyun.com/minisite/goods?userCode=vxqr4qh3)    

> 阿里云Hi拼购 限量爆款 低至199元/年 [点我进入](https://www.aliyun.com/acts/hi-group-buying?userCode=vxqr4qh3)  `(仅限新用户)`

## 内置功能

1.  用户管理：用户是系统操作者，该功能主要完成系统用户配置。
2.  部门管理：配置系统组织机构（公司、部门、小组），树结构展现支持数据权限。
3.  岗位管理：配置系统用户所属担任职务。
4.  菜单管理：配置系统菜单，操作权限，按钮权限标识等。
5.  角色管理：角色菜单权限分配、设置角色按机构进行数据范围权限划分。
6.  字典管理：对系统中经常使用的一些较为固定的数据进行维护。
7.  参数管理：对系统动态配置常用参数。
8.  通知公告：系统通知公告信息发布维护。
9.  操作日志：系统正常操作日志记录和查询；系统异常信息日志记录和查询。
10.  登录日志：系统登录日志记录查询包含登录异常。
11.  在线用户：当前系统中活跃用户状态监控。
12.  定时任务：在线（添加、修改、删除)任务调度包含执行结果日志。
13.  代码生成：前后端代码的生成（Go、html、xml、sql） 。
14.  系统接口：根据业务代码自动生成相关的api接口文档。
15.  服务监控：监视当前系统CPU、内存、磁盘、堆栈等相关信息。
16.  在线构建器：拖动表单元素生成相应的HTML代码。
17.  案例演示：常用的前端组件整合演示。

## 在线体验

> admin/yjg123456

演示地址：[http://www.yunjie.info:8199/login](http://www.yunjie.info:8199/login)

文档地址：正在编写

## 演示图

| ![](https://oscimg.oschina.net/oscnet/up-80920b74c9c2dcbe02650f7c127b445373a.png) | ![](https://oscimg.oschina.net/oscnet/up-8fcd14be9bb673dc3d80e2acc23fd962e20.png) |
| ![](https://oscimg.oschina.net/oscnet/up-90c964378cafc0aac0340b2a0d9951de820.png) | ![](https://oscimg.oschina.net/oscnet/up-bad8bfe2ae30a939dc4b45c8575791d4952.png) |
| ![](https://oscimg.oschina.net/oscnet/up-3278734c643a09b630bb8e4020c04584ab3.png) | ![](https://oscimg.oschina.net/oscnet/up-295a93f887b221acd2ccb66f40c40c9c202.png) |
| ![](https://oscimg.oschina.net/oscnet/up-9127f2c19ef7125bdad5f6f55b0a040371f.png) | ![](https://oscimg.oschina.net/oscnet/up-8cac96fe235b930b32eda064b5a707e3227.png) |
| ![](https://oscimg.oschina.net/oscnet/up-d9b2704523edcb32cdd4e5d2ba32238ae43.png) | ![](https://oscimg.oschina.net/oscnet/up-6c31e3d5d85136ddc2cc0e6143bad8ef39a.png) |
| ![](https://oscimg.oschina.net/oscnet/up-1d534c91e26375ceb5a9af20c5a0ebaa484.png) | ![](https://oscimg.oschina.net/oscnet/up-17c94cc7f2710158fd3495f03caeb17b33e.png) |
| ![](https://oscimg.oschina.net/oscnet/up-c59a5cec71dfe5b57b78c15229fd800faa3.png) | ![](https://oscimg.oschina.net/oscnet/up-5f69aa236d6cff05ed626479e636565ee9c.png) |
| ![](https://oscimg.oschina.net/oscnet/up-b272e09917b75d9e0d14372049a291aecfa.png) | ![](https://oscimg.oschina.net/oscnet/up-002f836244ca3b1971e8c67eaa972cf1d7b.png) |
| ![](https://oscimg.oschina.net/oscnet/up-ae81782cebfac71945f5301eabb96631089.png) | ![](https://oscimg.oschina.net/oscnet/up-454c053ee1d75c0d67b0d87a47925a2c6e4.png) |

## 微信交流群

如有兴趣一起完善，可以加我微信，进群

![](https://oscimg.oschina.net/oscnet/up-ba0417e087dc0869036610774da45a32e45.JPEG)