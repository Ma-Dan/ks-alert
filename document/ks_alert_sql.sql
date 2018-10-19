-- resource 现在开放了有哪些告警的资源? cluster node workspace namespace workload pod container...
create table resource_level (
	resource_level_id varchar(50) not null primary key,
	resource_type varchar(50) not null,
	description text default '' not null,
	enable boolean default true,
	create_by varchar(50),
	create_time timestamp default current_timestamp not null,
	update_time timestamp default current_timestamp not null
);

-- 一种类型的资源下面有哪些开放的 alert template
create table resource_alert_tmpls (
	resource_tmpl_id varchar(50) not null primary key,
	resource_level_id varchar(50) not null,
	alert_tmpl_id varchar(50) not null,
	enable boolean default true,
	description text default '' not null
)

create table resource (
	resource_id varchar(50) not null primary key,
	resource_name varchar(50) not null,
	resource_group_id varchar(50) not null,   -- 所属哪个资源组
	resource_level_id varchar(50) not null
);

-- resoure_group 同一个resource_group 下面的 resource_id 相同
-- 单个或者全部resource 分别视为一种resource_group
create table resource_group (
    resource_group_id varchar(50) not null primary key,	
	resource_group_name varchar(50) not null, 
	create_by varchar(50),
	enable boolean default true,
	create_time timestamp default current_timestamp not null,
	update_time timestamp default current_timestamp not null
);

create table alert_templ (
	alert_tmpl_id varchar(50) not null primary key,
	templ_group_id varchar(50) not null,
	templ_name text default '' not null,
	description text default '' not null,
    usage varchar(50) default 'alert' not null,  -- alert, autoscaling
	-- status_time timestamp default current_timestamp not null,
    create_time timestamp default current_timestamp not null,
	metric varchar(50) not null,
    consecutive_periods int not null, -- 该时间维度的频次
	frequency int not null, -- 频次
    condition_type varchar(50),  -- lt, le, gt, ge
    thresholds varchar(255),  -- thresholds: v1|v2...
    unit varchar(50) not null,  -- 单位 raw, percent
	inhibit_rule boolean default false -- 是否是 inhibit 规则，inhibit rule 会抑制其他的规则
    enable int default 0 not null, -- 0 for enabled, 1 for disabled
);


-- 一组 alert 模板的集合
create table alert_templ_group (
	templ_group_id varchar(50) not null primary key,
	enable int default 0 not null, -- 0 for enabled, 1 for disabled
	create_time timestamp default current_timestamp not null,
	update_time timestamp default current_timestamp not null
	create_by varchar(255) not null,  -- 系统创建的规则还是用户创建的规则模板
)


-- 告警的严重程度
create table alert_severity (
	severity_id varchar(50) not null primary key,
	severity varchar(50) not null default 'critical',
	create_time timestamp default current_timestamp not null,
    update_time timestamp default current_timestamp not null
);


create table receiver (
	receiver_id varchar(50) not null primary key,
	receiver_group_id varchar(50) not null,
	use_system_user boolean default true,  -- 使用KS系统已有的用户信息，如果不用，用户也可以指定email...
	receiver_name varchar(50) not null,
	-- mail_validated boolean default false,
	mail varchar(50)  default '',
    -- sms_validated boolean default true,
    phone varchar(50) default '',
	-- wechat_validated boolean default true,
	wechat varchar(50) default '',
    -- webhook_validated boolean default true,
	webhook varchar(50) not null,
    create_time timestamp default current_timestamp not null,
    update_time timestamp default current_timestamp not null
)


create table receiver_group {
	receiver_group_id varchar(50) not null primary key,
	receiver_group_name varchar(50) not null,
	severity_level varchar(50) not null,  -- 绑定严重程度
	default_webhook varchar(50) not null,
	default_email varchar(50) not null,
	create_time timestamp default current_timestamp not null,
    update_time timestamp default current_timestamp not null,
	create_by varchar(50) not null
}


-- alert 
CREATE TABLE alert (
    alert_id varchar(50) not null primary key,
    alert_name  varchar(50) not null,
    enable_time timestamp default current_timestamp not null,
	disable_time timestamp default current_timestamp not null,
	group_wait int -- 初次告警产生后，累积多长时间发送
	group_interval int -- 两次累积发送的时间间隔
	repeat_interval int -- 重发告警的时间间隔  
	create_time timestamp default current_timestamp not null,
    update_time timestamp default current_timestamp not null,
);

-- alert 绑定资源和receivers和报警规则
create table alert_binding {
	id varchar(50) not null primary key, 
	user_id varchar(50) not null,
	status varchar(50) not null,  -- ok, alert, insufficient
    reason varchar(255) default '' not null,  -- meters which status is alert  触发告警的原因
	alert_id varchar(50) not null,   -- 该用户的某个告警
	resource_group_id varchar(50),  -- 告警资源绑定
	tmpl_group_id varchar(50) not null,  -- 告警规则绑定
	receiver_group_id varchar(50),  -- 绑定receiver group，若有多个receiver_group,则插入多条记录
	receiver_id varchar(50) -- 绑定receiver，同上
}


-- alert history 只记录和alert状态变更相关的 记录，用来推算alert的开始时间-结束时间
-- 如：alert被触发， alert被silenced， alert 起作用时间已到， alert 被删除。
create table alert_history (
    -- id int unsigned primary key auto_increment,
	id varchar(50) primary key,
	user_id varchar(50) not null,
    alert_id varchar(50) not null,
	alert_name varchar(50) not null,
	resource_level varchar(50) not null, -- 资源类型
	resource_group_name varchar(50) not null -- 用户自定义的资源组
	resource_names text not null -- 该时刻所有的资源
	severity varchar(50) not null  -- 严重程度 与 receiver_group 绑定
	receiver_group varchar(300) -- 多个接受者信息的拼接成一个大的字符串
	receiver varchar(50) -- 单个接受者
	update_time timestamp default current_timestamp not null -- alert变更的时间，
	operation varchar(50) not null -- alert被触发， alert被silenced， alert 起作用时间已到， alert 被删除。
	trigger_alert_rule varchar(200) not null -- 触发告警的 alert rule 
    description text default '' not null,
    create_time timestamp default current_timestamp not null
);
