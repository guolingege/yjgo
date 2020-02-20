/*
==========================================================================
云捷GO自动生成菜单SQL,只生成一次,按需修改.
生成日期：2020-02-17 14:03:51
生成路径: document/sql/module/online_menu.sql
生成人：yunjie
==========================================================================
*/

-- 菜单 SQL
insert into sys_menu (menu_name, parent_id, order_num, url, menu_type, visible, perms, icon, create_by, create_time, update_by, update_time, remark)
values('在线用户记录', '3', '1', '/module/online', 'C', '0', 'online:view', '#', 'admin', '2020-01-01', 'admin', '2020-01-01', '在线用户记录菜单');

-- 按钮父菜单ID
SELECT @parentId := LAST_INSERT_ID();

-- 按钮 SQL
insert into sys_menu  (menu_name, parent_id, order_num, url, menu_type, visible, perms, icon, create_by, create_time, update_by, update_time, remark)
values('在线用户记录查询', @parentId, '1',  '#',  'F', '0', 'online:list',         '#', 'admin', '2020-01-01', 'admin', '2020-01-01', '');

insert into sys_menu  (menu_name, parent_id, order_num, url, menu_type, visible, perms, icon, create_by, create_time, update_by, update_time, remark)
values('在线用户记录新增', @parentId, '2',  '#',  'F', '0', 'online:add',          '#', 'admin', '2020-01-01', 'admin', '2020-01-01', '');

insert into sys_menu  (menu_name, parent_id, order_num, url, menu_type, visible, perms, icon, create_by, create_time, update_by, update_time, remark)
values('在线用户记录修改', @parentId, '3',  '#',  'F', '0', 'online:edit',         '#', 'admin', '2020-01-01', 'admin', '2020-01-01', '');

insert into sys_menu  (menu_name, parent_id, order_num, url, menu_type, visible, perms, icon, create_by, create_time, update_by, update_time, remark)
values('在线用户记录删除', @parentId, '4',  '#',  'F', '0', 'online:remove',       '#', 'admin', '2020-01-01', 'admin', '2020-01-01', '');

insert into sys_menu  (menu_name, parent_id, order_num, url, menu_type, visible, perms, icon, create_by, create_time, update_by, update_time, remark)
values('在线用户记录导出', @parentId, '5',  '#',  'F', '0', 'online:export',       '#', 'admin', '2020-01-01', 'admin', '2020-01-01', '');
