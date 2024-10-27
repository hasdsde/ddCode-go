create table authority
(
    id          int auto_increment
        primary key,
    name        varchar(255) null comment '名称',
    description varchar(40)  null comment '描述',
    type        int          null comment '类型  0-前端 1-后端',
    created_at  timestamp    null,
    updated_at  timestamp    null,
    deleted_at  timestamp    null
)
    comment '权限';

INSERT INTO authority (id, name, description, type, created_at, updated_at, deleted_at)
VALUES (4, 'server_menu_update', '新增/修改菜单', 2, '2024-01-20 21:25:46', '2024-01-20 21:28:38', null);
INSERT INTO authority (id, name, description, type, created_at, updated_at, deleted_at)
VALUES (5, 'server_menu_page', '分页查询菜单', 2, '2024-01-21 00:28:53', null, null);
INSERT INTO authority (id, name, description, type, created_at, updated_at, deleted_at)
VALUES (6, 'server_authority_update', '新增/修改权限', 2, '2024-01-21 00:43:12', null, null);
INSERT INTO authority (id, name, description, type, created_at, updated_at, deleted_at)
VALUES (7, 'server_authority_page', '分页查询全部权限', 2, '2024-01-21 00:43:40', null, null);
INSERT INTO authority (id, name, description, type, created_at, updated_at, deleted_at)
VALUES (8, 'server_role_update', '新增/修改角色', 2, '2024-01-21 00:59:36', null, null);
INSERT INTO authority (id, name, description, type, created_at, updated_at, deleted_at)
VALUES (9, 'server_role_page', '分页查询角色', 2, '2024-01-21 00:59:53', null, null);
INSERT INTO authority (id, name, description, type, created_at, updated_at, deleted_at)
VALUES (10, 'server_role_authority_update', '为角色分配权限', 2, '2024-01-21 01:01:07', '2024-01-21 01:21:52', null);
INSERT INTO authority (id, name, description, type, created_at, updated_at, deleted_at)
VALUES (11, 'server_role_menu_update', '为角色分配菜单', 2, '2024-01-21 01:01:26', '2024-01-21 01:21:41', null);
INSERT INTO authority (id, name, description, type, created_at, updated_at, deleted_at)
VALUES (12, 'server_user_role_update', '为用户分配角色', 2, '2024-01-21 01:01:49', '2024-01-21 01:23:51', null);
INSERT INTO authority (id, name, description, type, created_at, updated_at, deleted_at)
VALUES (13, 'server_user_page', '分页查询全部用户', 2, '2024-01-21 01:02:15', null, null);
INSERT INTO authority (id, name, description, type, created_at, updated_at, deleted_at)
VALUES (14, 'server_user_update', '新增/修改用户', 2, '2024-01-21 01:02:47', null, null);
INSERT INTO authority (id, name, description, type, created_at, updated_at, deleted_at)
VALUES (15, 'server_dept_page', '分页全部查询部门', 2, '2024-01-21 01:03:25', null, null);
INSERT INTO authority (id, name, description, type, created_at, updated_at, deleted_at)
VALUES (16, 'server_dept_update', '新增/修改部门', 2, '2024-01-21 01:03:45', null, null);
INSERT INTO authority (id, name, description, type, created_at, updated_at, deleted_at)
VALUES (17, 'server_user_dept', '为用户分配部门', 2, '2024-01-21 01:04:12', null, null);
INSERT INTO authority (id, name, description, type, created_at, updated_at, deleted_at)
VALUES (18, 'server_role_authority_page', '查询角色权限', 2, '2024-01-21 01:19:58', null, null);
INSERT INTO authority (id, name, description, type, created_at, updated_at, deleted_at)
VALUES (19, 'server_role_menu_page', '查询角色菜单', 2, '2024-01-21 01:20:11', '2024-01-21 01:20:30', null);
INSERT INTO authority (id, name, description, type, created_at, updated_at, deleted_at)
VALUES (20, 'admin', '其他admin权限', 2, '2024-01-21 14:38:54', null, null);

create table dept
(
    id          int auto_increment
        primary key,
    parent_id   int          null comment '父级id',
    name        varchar(40)  null,
    description varchar(100) null,
    created_at  timestamp    null,
    deleted_at  timestamp    null,
    updated_at  timestamp    null
)
    comment '部门';

INSERT INTO dept (id, parent_id, name, description, created_at, deleted_at, updated_at)
VALUES (1, null, '顶级部门', '顶级部门', '2023-12-18 15:22:48', null, '2024-01-19 13:49:02');
INSERT INTO dept (id, parent_id, name, description, created_at, deleted_at, updated_at)
VALUES (6, 1, '部门1', '部门1', '2024-01-19 13:46:47', null, null);
INSERT INTO dept (id, parent_id, name, description, created_at, deleted_at, updated_at)
VALUES (7, 1, '部门2', '部门2介绍', '2024-01-19 13:54:30', null, '2024-01-19 13:58:00');
INSERT INTO dept (id, parent_id, name, description, created_at, deleted_at, updated_at)
VALUES (8, 1, '部门3', '部门3', '2024-01-19 13:59:41', '2024-01-19 14:05:19', null);
INSERT INTO dept (id, parent_id, name, description, created_at, deleted_at, updated_at)
VALUES (46, 6, '部门1的子类', null, '2024-01-21 21:20:38', null, null);
INSERT INTO dept (id, parent_id, name, description, created_at, deleted_at, updated_at)
VALUES (47, 46, '部门1的子类的子类', '修改了描述1', null, null, '2024-01-22 16:07:25');
INSERT INTO dept (id, parent_id, name, description, created_at, deleted_at, updated_at)
VALUES (48, 6, '部门1的第二个子类', '部门1的第二个子类', '2024-01-21 23:21:12', null, null);

create table log
(
    id             int auto_increment comment 'id'
        primary key,
    user_name      varchar(100)  null comment '用户名',
    user_id        int           null comment '用户id',
    authority_name varchar(100)  null comment '权限名',
    info           varchar(1000) null comment '详细信息',
    created_at     timestamp     null comment '创建时间',
    method_name    varchar(100)  not null
)
    comment 'Log';


create table menu
(
    id           int auto_increment
        primary key,
    url          varchar(40)  null comment 'URL',
    name         varchar(40)  null,
    parent_id    int          null,
    authority_id int          null,
    orders       int          null comment '排序，越高越靠前',
    icon         varchar(100) null,
    created_at   timestamp    null,
    updated_at   timestamp    null,
    deleted_at   timestamp    null
);

create index menu_authority_id_fk
    on menu (authority_id);

INSERT INTO menu (id, url, name, parent_id, authority_id, orders, icon, created_at, updated_at, deleted_at)
VALUES (31, '/dashboard', '仪表盘', null, null, 2, 'dashboard', '2023-12-22 22:01:06', '2023-12-22 22:23:55', null);
INSERT INTO menu (id, url, name, parent_id, authority_id, orders, icon, created_at, updated_at, deleted_at)
VALUES (33, '/sys', '系统管理', null, null, 2, 'settings', '2023-12-22 22:09:45', '2023-12-22 22:25:46', null);
INSERT INTO menu (id, url, name, parent_id, authority_id, orders, icon, created_at, updated_at, deleted_at)
VALUES (34, '/front', '前端工具', null, null, 2, 'auto_fix_high', '2023-12-22 22:20:43', '2023-12-22 22:25:51', null);
INSERT INTO menu (id, url, name, parent_id, authority_id, orders, icon, created_at, updated_at, deleted_at)
VALUES (35, '/server', '后端工具', null, null, 2, 'handyman', '2023-12-22 22:21:50', '2023-12-22 22:25:54', null);
INSERT INTO menu (id, url, name, parent_id, authority_id, orders, icon, created_at, updated_at, deleted_at)
VALUES (36, '/dashboard/home', '首页', 31, null, 1, 'grid_view', '2023-12-22 22:27:22', null, null);
INSERT INTO menu (id, url, name, parent_id, authority_id, orders, icon, created_at, updated_at, deleted_at)
VALUES (37, '/sys/user', '用户管理', 33, null, 1, 'account_circle', '2023-12-22 22:28:19', null, null);
INSERT INTO menu (id, url, name, parent_id, authority_id, orders, icon, created_at, updated_at, deleted_at)
VALUES (38, '/sys/authority', '权限管理', 33, null, 1, 'verified_user', '2023-12-22 22:30:32', null, null);
INSERT INTO menu (id, url, name, parent_id, authority_id, orders, icon, created_at, updated_at, deleted_at)
VALUES (39, '/sys/dept', '部门管理', 33, null, 1, 'apartment', '2023-12-22 22:31:27', null, null);
INSERT INTO menu (id, url, name, parent_id, authority_id, orders, icon, created_at, updated_at, deleted_at)
VALUES (40, '/sys/menu', '菜单管理', 33, null, 1, 'list', '2023-12-22 22:32:34', null, null);
INSERT INTO menu (id, url, name, parent_id, authority_id, orders, icon, created_at, updated_at, deleted_at)
VALUES (41, '/front/tailwind', 'Tailwind生成器', 34, null, 2, 'cyclone', '2023-12-22 22:34:09', '2023-12-24 14:09:29',
        null);
INSERT INTO menu (id, url, name, parent_id, authority_id, orders, icon, created_at, updated_at, deleted_at)
VALUES (42, '/front/icon', 'Material Icon 列表', 34, null, 1, 'emoji_events', '2023-12-22 22:36:10', null, null);
INSERT INTO menu (id, url, name, parent_id, authority_id, orders, icon, created_at, updated_at, deleted_at)
VALUES (43, '/front/func', '函数调用指南', 34, null, 1, 'functions', '2023-12-22 22:37:04', null, null);
INSERT INTO menu (id, url, name, parent_id, authority_id, orders, icon, created_at, updated_at, deleted_at)
VALUES (44, '/front/table', '表格配置工具', 34, null, 1, 'insert_chart_outlined', '2023-12-22 22:38:22',
        '2023-12-22 23:09:22', null);
INSERT INTO menu (id, url, name, parent_id, authority_id, orders, icon, created_at, updated_at, deleted_at)
VALUES (45, '/server/data', 'SQL监控', 35, null, 1, 'timeline', '2023-12-22 22:40:25', '2024-01-22 16:56:49', null);
INSERT INTO menu (id, url, name, parent_id, authority_id, orders, icon, created_at, updated_at, deleted_at)
VALUES (46, '/server/log', '日志监控', 35, null, 1, 'post_add', '2023-12-22 22:47:13', null, null);
INSERT INTO menu (id, url, name, parent_id, authority_id, orders, icon, created_at, updated_at, deleted_at)
VALUES (47, '/server/schedule', '定时任务', 35, null, 1, 'alarm', '2023-12-22 22:47:54', '2023-12-23 15:59:24',
        '2024-01-23 00:47:17');
INSERT INTO menu (id, url, name, parent_id, authority_id, orders, icon, created_at, updated_at, deleted_at)
VALUES (48, '/front/echart', 'Echarts生成器', 34, null, 1, 'donut_large', '2023-12-23 22:40:05', null, null);
INSERT INTO menu (id, url, name, parent_id, authority_id, orders, icon, created_at, updated_at, deleted_at)
VALUES (49, '/sys/role', '角色管理', 33, null, 1, 'face', '2024-01-19 14:59:07', null, null);
INSERT INTO menu (id, url, name, parent_id, authority_id, orders, icon, created_at, updated_at, deleted_at)
VALUES (50, '/server/system', '系统监控', 35, null, 1, 'memory', '2024-01-22 23:36:26', '2024-01-22 23:42:14', null);

create table role
(
    id          int auto_increment
        primary key,
    name        varchar(40)  null comment '名称',
    description varchar(100) null,
    created_at  timestamp    null,
    updated_at  timestamp    null,
    deleted_at  timestamp    null
);

INSERT INTO role (id, name, description, created_at, updated_at, deleted_at)
VALUES (1, '角色1', '角色1', '2024-01-19 17:25:19', null, null);
INSERT INTO role (id, name, description, created_at, updated_at, deleted_at)
VALUES (3, 'Admin', '管理员', '2024-01-20 21:26:26', null, null);

create table role_authority
(
    role_id      int null,
    authority_id int null,
    id           int auto_increment
        primary key,
    constraint role_authority_authority_id_fk
        foreign key (authority_id) references authority (id),
    constraint role_authority_role_id_fk
        foreign key (role_id) references role (id)
);

INSERT INTO role_authority (role_id, authority_id, id)
VALUES (3, 4, 11);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (3, 5, 12);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (3, 6, 13);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (3, 7, 14);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (3, 8, 15);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (3, 9, 16);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (3, 10, 17);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (3, 11, 18);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (3, 12, 19);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (3, 13, 20);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (3, 14, 21);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (3, 15, 22);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (3, 16, 23);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (3, 17, 24);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (3, 18, 25);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (3, 19, 26);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (3, 20, 27);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 4, 28);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 5, 29);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 6, 30);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 18, 33);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 19, 34);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 14, 35);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 4, 223);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 5, 224);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 6, 225);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 18, 226);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 19, 227);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 14, 228);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 7, 229);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 13, 230);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 15, 231);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 4, 232);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 5, 233);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 6, 234);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 18, 235);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 19, 236);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 14, 237);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 4, 238);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 5, 239);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 6, 240);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 18, 241);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 19, 242);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 14, 243);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 7, 244);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 13, 245);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 15, 246);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 9, 247);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 4, 248);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 5, 249);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 6, 250);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 18, 251);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 19, 252);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 14, 253);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 4, 254);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 5, 255);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 6, 256);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 18, 257);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 19, 258);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 14, 259);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 7, 260);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 13, 261);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 15, 262);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 4, 263);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 5, 264);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 6, 265);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 18, 266);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 19, 267);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 14, 268);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 4, 269);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 5, 270);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 6, 271);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 18, 272);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 19, 273);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 14, 274);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 7, 275);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 13, 276);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 15, 277);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 9, 278);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 16, 279);

create table role_authority
(
    role_id      int null,
    authority_id int null,
    id           int auto_increment
        primary key,
    constraint role_authority_authority_id_fk
        foreign key (authority_id) references authority (id),
    constraint role_authority_role_id_fk
        foreign key (role_id) references role (id)
);

INSERT INTO role_authority (role_id, authority_id, id)
VALUES (3, 4, 11);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (3, 5, 12);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (3, 6, 13);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (3, 7, 14);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (3, 8, 15);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (3, 9, 16);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (3, 10, 17);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (3, 11, 18);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (3, 12, 19);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (3, 13, 20);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (3, 14, 21);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (3, 15, 22);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (3, 16, 23);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (3, 17, 24);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (3, 18, 25);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (3, 19, 26);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (3, 20, 27);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 4, 28);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 5, 29);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 6, 30);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 18, 33);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 19, 34);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 14, 35);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 4, 223);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 5, 224);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 6, 225);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 18, 226);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 19, 227);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 14, 228);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 7, 229);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 13, 230);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 15, 231);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 4, 232);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 5, 233);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 6, 234);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 18, 235);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 19, 236);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 14, 237);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 4, 238);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 5, 239);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 6, 240);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 18, 241);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 19, 242);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 14, 243);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 7, 244);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 13, 245);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 15, 246);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 9, 247);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 4, 248);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 5, 249);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 6, 250);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 18, 251);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 19, 252);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 14, 253);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 4, 254);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 5, 255);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 6, 256);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 18, 257);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 19, 258);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 14, 259);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 7, 260);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 13, 261);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 15, 262);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 4, 263);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 5, 264);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 6, 265);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 18, 266);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 19, 267);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 14, 268);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 4, 269);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 5, 270);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 6, 271);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 18, 272);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 19, 273);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 14, 274);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 7, 275);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 13, 276);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 15, 277);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 9, 278);
INSERT INTO role_authority (role_id, authority_id, id)
VALUES (1, 16, 279);


create table role_menu
(
    id      int auto_increment
        primary key,
    role_id int null comment '角色id',
    menu_id int null comment '菜单id'
);

INSERT INTO role_menu (id, role_id, menu_id)
VALUES (1, 1, 31);
INSERT INTO role_menu (id, role_id, menu_id)
VALUES (4, 3, 31);
INSERT INTO role_menu (id, role_id, menu_id)
VALUES (5, 3, 33);
INSERT INTO role_menu (id, role_id, menu_id)
VALUES (6, 3, 34);
INSERT INTO role_menu (id, role_id, menu_id)
VALUES (7, 3, 36);
INSERT INTO role_menu (id, role_id, menu_id)
VALUES (8, 3, 35);
INSERT INTO role_menu (id, role_id, menu_id)
VALUES (9, 3, 37);
INSERT INTO role_menu (id, role_id, menu_id)
VALUES (10, 3, 38);
INSERT INTO role_menu (id, role_id, menu_id)
VALUES (11, 3, 39);
INSERT INTO role_menu (id, role_id, menu_id)
VALUES (12, 3, 40);
INSERT INTO role_menu (id, role_id, menu_id)
VALUES (13, 3, 41);
INSERT INTO role_menu (id, role_id, menu_id)
VALUES (14, 3, 42);
INSERT INTO role_menu (id, role_id, menu_id)
VALUES (15, 3, 43);
INSERT INTO role_menu (id, role_id, menu_id)
VALUES (16, 3, 44);
INSERT INTO role_menu (id, role_id, menu_id)
VALUES (17, 3, 45);
INSERT INTO role_menu (id, role_id, menu_id)
VALUES (18, 3, 46);
INSERT INTO role_menu (id, role_id, menu_id)
VALUES (19, 3, 47);
INSERT INTO role_menu (id, role_id, menu_id)
VALUES (20, 3, 48);
INSERT INTO role_menu (id, role_id, menu_id)
VALUES (21, 3, 49);
INSERT INTO role_menu (id, role_id, menu_id)
VALUES (22, 1, 33);
INSERT INTO role_menu (id, role_id, menu_id)
VALUES (25, 1, 36);
INSERT INTO role_menu (id, role_id, menu_id)
VALUES (26, 1, 37);
INSERT INTO role_menu (id, role_id, menu_id)
VALUES (27, 1, 38);
INSERT INTO role_menu (id, role_id, menu_id)
VALUES (28, 1, 39);
INSERT INTO role_menu (id, role_id, menu_id)
VALUES (29, 1, 40);
INSERT INTO role_menu (id, role_id, menu_id)
VALUES (30, 1, 49);
INSERT INTO role_menu (id, role_id, menu_id)
VALUES (31, 3, 50);


create table user
(
    id         int auto_increment comment '主键'
        primary key,
    user_name  varchar(40)  null comment '登录名',
    nick_name  varchar(40)  null comment '显示名称',
    email      varchar(100) null comment '邮箱',
    phone      varchar(20)  null comment '手机号',
    sex        int          null comment '性别 0-未知 1男 2女',
    role_id    int          null,
    avatar     varchar(200) null comment '头像',
    password   varchar(40)  null comment '密码',
    dept_id    int          null comment '部门id',
    created_at timestamp    null comment '创建时间',
    updated_at timestamp    null comment '更新时间',
    deleted_at timestamp    null comment '删除时间',
    comment    varchar(100) null comment '备注'
)
    comment '用户表';

INSERT INTO user (id, user_name, nick_name, email, phone, sex, role_id, avatar, password, dept_id, created_at,
                         updated_at, deleted_at, comment)
VALUES (1, 'zhangsan', '张三', 'zhangsan@qq.com', '17565286323', 1, 1, null, 'E10ADC3949BA59ABBE56E057F20F883E', 6,
        '2023-12-17 17:27:26', '2024-01-22 00:06:40', null, null);
INSERT INTO user (id, user_name, nick_name, email, phone, sex, role_id, avatar, password, dept_id, created_at,
                         updated_at, deleted_at, comment)
VALUES (3, 'admin', '管理员', '114514191810@qq.com', '1145141919810', 0, 3, '', 'E10ADC3949BA59ABBE56E057F20F883E', 1,
        '2024-01-20 15:21:47', null, null, '最高の管理员');
INSERT INTO user (id, user_name, nick_name, email, phone, sex, role_id, avatar, password, dept_id, created_at,
                         updated_at, deleted_at, comment)
VALUES (4, 'lisi', '李四', '123', '123', 1, 1, '', '', 48, '2024-01-21 17:33:47', '2024-01-22 00:08:22', null, '');
