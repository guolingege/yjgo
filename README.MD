## 框架简介

以前是做Java企业应用开发，想转Go,看了很多优秀的开源项目但是发现没有合适的，一直没找到类似于若依开发思路的快速开发框架。于是狠下心自己写了一套后台系统。这个框架可以用于所有的Web应用程序，如网站管理后台，网站会员中心，CMS，CRM，OA。所有前端后台代码封装过后十分精简易上手，出错概率低。同时支持移动客户端访问。系统会陆续更新一些实用功能。

受之开源，回馈社区，本框架以GoFrame为web服务框架，继续沿用MIT开源协议，架构思路沿袭着若依的以辅助生成重复代码为主，不过度封装，生成的代码可以快速修改适应不同的需求，适应每个开发者自己的习惯和风格。所以命名为"云捷GO"（寓意：GO云原生业务快捷开发）



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

<table>
	<tbody>
		<tr>
			<td><img alt="" src="https://images.gitee.com/uploads/images/2020/0221/075216_b6e05db0_779478.png" /></td>
			<td><img alt="" src="https://images.gitee.com/uploads/images/2020/0221/075216_86e787b5_779478.png" /></td>
		</tr>
		<tr>
			<td><img alt="" src="https://images.gitee.com/uploads/images/2020/0221/075216_168d55c8_779478.png" /></td>
			<td><img alt="" src="https://images.gitee.com/uploads/images/2020/0221/075216_14f443a1_779478.png" /></td>
		</tr>
		<tr>
			<td><img alt="" src="https://images.gitee.com/uploads/images/2020/0221/075216_10acfb8a_779478.png" /></td>
			<td><img alt="" src="https://images.gitee.com/uploads/images/2020/0221/075217_8381993d_779478.png" /></td>
		</tr>
		<tr>
			<td><img alt="" src="https://images.gitee.com/uploads/images/2020/0221/075217_1bc1bb34_779478.png" /></td>
			<td><img alt="" src="https://images.gitee.com/uploads/images/2020/0221/075217_0a269a55_779478.png" /></td>
		</tr>
		<tr>
			<td><img alt="" src="https://images.gitee.com/uploads/images/2020/0221/075217_31384ef1_779478.png" /></td>
			<td><img alt="" src="https://images.gitee.com/uploads/images/2020/0221/075217_aa7f375b_779478.png" /></td>
		</tr>
		<tr>
			<td><img alt="" src="https://images.gitee.com/uploads/images/2020/0221/075217_74b541a0_779478.png" /></td>
			<td><img alt="" src="https://images.gitee.com/uploads/images/2020/0221/075217_97556567_779478.png" /></td>
		</tr>
		<tr>
			<td><img alt="" src="https://images.gitee.com/uploads/images/2020/0221/075217_e8f10ec9_779478.png" /></td>
			<td><img alt="" src="https://images.gitee.com/uploads/images/2020/0221/075217_ff915e4a_779478.png" /></td>
		</tr>
		<tr>
			<td><img alt="" src="https://images.gitee.com/uploads/images/2020/0221/075217_9527bcd4_779478.png" /></td>
			<td><img alt="" src="https://images.gitee.com/uploads/images/2020/0221/075217_41e24951_779478.png" /></td>
		</tr>
		<tr>
			<td><img alt="" src="https://images.gitee.com/uploads/images/2020/0221/075217_c0897638_779478.png" /></td>
			<td><img alt="" src="https://images.gitee.com/uploads/images/2020/0221/075218_8888e0e9_779478.png" /></td>
		</tr>
	</tbody>
</table>

## 感谢
1. gf框架 [https://github.com/gogf/gf](https://github.com/gogf/gf) 
2. RuoYi框架 [https://github.com/yangzongzhuan/RuoYi](https://github.com/yangzongzhuan/RuoYi)
3. tealeg [https://github.com/tealeg/xlsx](https://github.com/tealeg/xlsx)
4. swaggo [https://github.com/swaggo/swag](https://github.com/swaggo/swag)

## 项目支持

> 阿里云通用云产品1888优惠券 ：[点我领取](https://www.aliyun.com/minisite/goods?userCode=vxqr4qh3)    

> 阿里云Hi拼购 限量爆款 低至199元/年 [点我进入](https://www.aliyun.com/acts/hi-group-buying?userCode=vxqr4qh3)  `(仅限新用户)`

> 也可以请作者喝一杯咖啡:)



## 微信交流群

如有兴趣一起完善，可以加我微信，进群

<img alt="" height="389" src="https://images.gitee.com/uploads/images/2020/0221/075218_1a479f04_779478.jpeg" width="300" />