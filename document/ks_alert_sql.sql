-- 现在开放了有哪些告警的资源? cluster node workspace namespace workload pod container...
create table resource_level (
	resource_level_id varchar(50) not null primary key,
	resource_type varchar(50) not null,  -- "cluster" "node" "workspace"
	description text default '' not null,
	enable boolean default true,
	create_by varchar(50),
	create_time timestamp default current_timestamp not null,
	update_time timestamp default current_timestamp not null
);

-- 资源实体，如node: "i-knzmfikc", "i-0jk3bsyk"
create table resource (
	resource_id varchar(50) not null primary key,
	resource_name varchar(50) not null,
	resource_level_id varchar(50) not null,   -- 资源级别，node or cluster or namespace...
	resource_group_id varchar(50) not null,   -- 所属哪个资源组
	create_time timestamp default current_timestamp not null,
	update_time timestamp default current_timestamp not null
);

-- resoure_group 同一个resource_group 下面的 resource_id 相同
-- 单个或者全部resource 分别视为一种resource_group
create table resource_group (
  resource_group_id varchar(50) not null primary key,
	resource_group_name varchar(50) not null, 
	enable boolean default true,
	create_time timestamp default current_timestamp not null,
	update_time timestamp default current_timestamp not null,
	create_by varchar(50) not null
);

-- metric 规则记录
create table metric (
  metric_id varchar(50) not null primary key,
  metric_name varchar(100) not null,
  rule_promql text not null,
  resource_level_id varchar(50) not null,
  resource_type varchar(50) not null,
  update_time timestamp default current_timestamp not null
);

-- alert rule，
create table alert_rule (
	alert_rule_id varchar(50) not null primary key,
	alert_name text default '' not null,
	alert_group_id varchar(50) not null, -- 所属的告警规则组
  -- 	description text default '' not null,
	metric_id varchar(50) not null,
	metric_name varchar(50) not null, -- 该字段冗余，更有有意义的指标名称，例如：node_pod_count
  condition_type varchar(50),  -- lt, le, gt, ge
  thresholds varchar(255),  -- thresholds: v1|v2...
  unit varchar(50) not null default '',  -- 监控指标的单位 raw, percent, Byte
  periods int not null, -- 告警执行周期
	consecutive_count int not null, -- 连续几次触发规则才会视为告警
	inhibit_rule boolean default false, -- 是否是 inhibit 规则，inhibit rule 会抑制其他的规则
  enable int default 0 not null, -- 0 for enabled, 1 for disabled
  create_time timestamp default current_timestamp not null,
  update_time timestamp default current_timestamp not null,
  create_by varchar(50) not null -- 系统预置的告警规则 也可以使用户自定义的告警规则
);

-- 一种类型的资源下面有哪些开放的 alert rule
create table resource_alert_rule (
	resource_alert_id varchar(50) not null primary key,
	resource_level_id varchar(50) not null,
	resource_type varchar(50) not null,
	-- 同一个resource_level_id对应多个alert template，
	-- 例如node的resource_level_id=1，对应的 alert_rule_id 具体为 cpu 内存等指标
	alert_rule_id varchar(50) not null,
	alert_name varchar(50) not null,
	enable boolean default true  -- 是否开放某种资源的某个alert rule
);

-- 一组 alert rule的集合
create table alert_rule_group (
	alert_group_id varchar(50) not null primary key,
	alert_group_name varchar(50) not null,
	enable boolean default true not null, -- 该告警组是否启动，该告警组可能属于系统或用户
	create_time timestamp default current_timestamp not null,
	update_time timestamp default current_timestamp not null,
	create_by varchar(255) not null,  -- 系统创建的规则还是用户创建的规则模板
);

-- 告警的严重程度
create table alert_severity (
	severity_id varchar(50) not null primary key,
	severity varchar(50) not null default 'critical', -- critical major minor warn
	severity_zh varchar(50) not null, -- "严重"，"较紧急"，"一般紧急" ,"轻微紧急"
	create_time timestamp default current_timestamp not null,
  update_time timestamp default current_timestamp not null,
  create_by varchar(50) not null
);

-- 单个 receiver
create table receiver (
	receiver_id varchar(50) not null primary key,
	receiver_group_id varchar(50) not null,
	-- 使用KS系统已有的用户信息，如果不用，用户也可以指定email...
	system_user boolean default true,
	receiver_name varchar(50) not null,
	-- mail_validated boolean default false,
	mail varchar(50)  default '',
  -- sms_validated boolean default true,
  phone varchar(50) default '',
	-- wechat_validated boolean default true,
	wechat varchar(50) default '',
  -- webhook_validated boolean default true,
	webhook varchar(50),
  create_time timestamp default current_timestamp not null,
  update_time timestamp default current_timestamp not null
);

-- 一组 receivers
create table receiver_group (
	receiver_group_id varchar(50) not null primary key,
	receiver_group_name varchar(50) not null,
	severity_level varchar(50) not null,  -- 绑定严重程度
  -- 	default_webhook varchar(50) not null,
  -- 	default_email varchar(50) not null,
	create_time timestamp default current_timestamp not null,
  update_time timestamp default current_timestamp not null,
	create_by varchar(50) not null
);

-- alert 绑定资源和receivers和报警规则
create table alert_binding (
	id varchar(50) not null primary key, 
	user_id varchar(50) not null,
	user_alert_id varchar(50) not null,   -- 该用户的某个告警
	resource_group_id varchar(50),  -- 告警资源绑定
	alert_group_id varchar(50) not null,  -- 告警规则绑定
	receiver_group_id varchar(50) not null,  -- 绑定receiver group，若有多个receiver_group,则插入多条记录
  reason varchar(255) default '' not null,  -- meters which status is alert  触发告警的原因
  silence_enable boolean default false,
  silence_start timestamp default current_timestamp not null,
  silence_end timestamp default current_timestamp not null,
  repeat_type varchar(20) default exponential,  -- 多久重复一次告警通知，重发周期保持不变 stationary 还是服从指数 exponential
  repeat_interval int default 8,  -- 单位是小时, 如果重发周期是固定的 stationary，该字段才有效
  init_repeat_period int default 1,  --  如果重发周期是exponential， 初始重发周期是 1小时
  max_repeat_period int default 168,  --  如果重发周期是exponential， 最大重发周期是 168 小时。
  enable_time timestamp default current_timestamp not null,  -- 告警的开始时间， 若不设置，则就是当前时间
	disable_time timestamp default current_timestamp not null,  -- 告警的结束时间， 若不设置，则告警一直有效
	enable boolean default true not null, -- 该告警组是否启用
	description text default '' not null  -- 告警描述
);

-- alert history 只记录和alert状态变更相关的 记录，用来推算alert的开始时间-结束时间
-- 如：alert被触发， alert被silenced， alert 起作用时间已到， alert 被删除。
create table alert_history (
	id varchar(50) primary key,
	user_id varchar(50) not null,
  user_alert_id varchar(50) not null,
	alert_name varchar(50) not null,
	resource_level varchar(50) not null, -- 资源类型
	resource_type varchar(50) not null, -- "node" , "pod"  ...
	resource_group_id varchar(50) not null, -- 用户自定义的资源组
	resource_group_name varchar(50) not null,
	resource_names text not null, -- 该时刻所有的资源,以数组的形式展示
	alerted_resource text not null, -- 正在触发告警的资源

	severity varchar(20) not null, -- 严重程度 与 receiver_group 绑定
	receiver_group_name varchar(50) not null,
	receiver_group varchar(300),  -- 多个接受者信息的拼接成一个大的字符串

	alert_group_id varchar(200) not null,
	alert_group_name varchar(200) not null,
	trigger_alert_rule varchar(200) not null, -- 触发告警的 alert rule
	event_time timestamp default current_timestamp not null, -- alert变更的时间，
	operation varchar(50) not null, -- alert被触发， alert被silenced， alert 起作用时间已到， alert 被删除。
  description text default '' not null
);
