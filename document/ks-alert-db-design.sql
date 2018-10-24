
-- SET SESSION FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS alert_binding;
DROP TABLE IF EXISTS alert_history;
DROP TABLE IF EXISTS alert_rule;
DROP TABLE IF EXISTS alert_rule_group;
DROP TABLE IF EXISTS alert_severity;
DROP TABLE IF EXISTS metric;
DROP TABLE IF EXISTS receiver;
DROP TABLE IF EXISTS receiver_group;
DROP TABLE IF EXISTS resource;
DROP TABLE IF EXISTS resource_group;
DROP TABLE IF EXISTS resource_level;

CREATE TABLE alert_binding
(
	id varchar(50) NOT NULL,
	user_id varchar(50) NOT NULL,
	user_alert_id varchar(50) NOT NULL,
	user_alert_name varchar(100) NOT NULL,
	alert_group_id varchar(50) NOT NULL,
	resource_group_id varchar(50) NOT NULL,
	receiver_group_id varchar(50) NOT NULL,
	severity_id varchar(50) NOT NULL,
	severity varchar(20),
	repeat_type varchar(20) DEFAULT 'exponential' NOT NULL,
	repeat_interval int DEFAULT 8 NOT NULL,
	init_repeat_period int DEFAULT 1,
	max_repeat_period int DEFAULT 168 NOT NULL,
	enable_time timestamp DEFAULT CURRENT_TIMESTAMP,
	disable_time timestamp DEFAULT CURRENT_TIMESTAMP,
	enable boolean DEFAULT '49' NOT NULL,
	description text,
	create_time timestamp,
	update_time timestamp NOT NULL,
	update_type varchar(20),
	PRIMARY KEY (id),
	UNIQUE (user_id, user_alert_id)
) DEFAULT CHARACTER SET latin1 COLLATE latin1_swedish_ci;


CREATE TABLE alert_history
(
	id bigint NOT NULL AUTO_INCREMENT,
	user_id varchar(50) NOT NULL,
	user_alert_id varchar(50) NOT NULL,
	alert_name varchar(50) NOT NULL,
	resource_level_id varchar(50) NOT NULL,
	resource_type varchar(20) NOT NULL,
	resource_group_id varchar(50) NOT NULL,
	resource_group_name varchar(50) NOT NULL,
	resource_names text NOT NULL,
	alerted_resource text NOT NULL,
	severity_id varchar(50) NOT NULL,
	severity varchar(20) NOT NULL,
	receiver_group_id varchar(50) NOT NULL,
	receiver_group_name varchar(50) NOT NULL,
	receiver_group text NOT NULL,
	alert_group_id varchar(50) NOT NULL,
	alert_group_name varchar(50) NOT NULL,
	trigger_alert_rule varchar(100) NOT NULL,
	event_time timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	operation varchar(50),
	cause text,
	silence_enable boolean DEFAULT '48' NOT NULL,
	silence_start timestamp DEFAULT CURRENT_TIMESTAMP,
	silence_end timestamp DEFAULT CURRENT_TIMESTAMP,
	repeat_type varchar(20) DEFAULT 'exponential' NOT NULL,
	repeat_interval int DEFAULT 8 NOT NULL,
	init_repeat_period int DEFAULT 1,
	max_repeat_period int DEFAULT 168 NOT NULL,
	send_status text,
	-- fired, resolved, deleted, silenced
	alert_status varchar(20) NOT NULL COMMENT 'fired, resolved, deleted, silenced',
	PRIMARY KEY (id)
) DEFAULT CHARACTER SET latin1 COLLATE latin1_swedish_ci;


CREATE TABLE alert_rule
(
	alert_rule_id varchar(50) NOT NULL,
	alert_name varchar(50) NOT NULL,
	alert_group_id varchar(50) NOT NULL,
	metric_id varchar(50) NOT NULL,
	metric_name varchar(50) NOT NULL,
	resource_level_id varchar(50) NOT NULL,
	resource_type varchar(20) NOT NULL,
	condition_type varchar(10) NOT NULL,
	thresholds int NOT NULL,
	unit varchar(10) NOT NULL,
	periods int NOT NULL,
	consecutive_count int NOT NULL,
	inhibit_rule boolean DEFAULT '48' NOT NULL,
	enable int DEFAULT 0 NOT NULL,
	create_time timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	update_time timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	create_by varchar(50),
	ref_alert_rule_id varchar(50) NOT NULL,
	PRIMARY KEY (alert_rule_id)
) DEFAULT CHARACTER SET latin1 COLLATE latin1_swedish_ci;


CREATE TABLE alert_rule_group
(
	alert_group_id varchar(50) NOT NULL,
	alert_group_name varchar(50) NOT NULL,
	enable boolean DEFAULT '49' NOT NULL,
	create_time timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	update_time timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	create_by varchar(255) NOT NULL,
	PRIMARY KEY (alert_group_id)
) DEFAULT CHARACTER SET latin1 COLLATE latin1_swedish_ci;


CREATE TABLE alert_severity
(
	severity_id varchar(50) NOT NULL,
	severity varchar(20) DEFAULT 'critical' NOT NULL,
	severity_zh varchar(20) NOT NULL,
	create_time timestamp DEFAULT CURRENT_TIMESTAMP,
	update_time timestamp DEFAULT CURRENT_TIMESTAMP,
	create_by varchar(50),
	PRIMARY KEY (severity_id)
) DEFAULT CHARACTER SET latin1 COLLATE latin1_swedish_ci;


CREATE TABLE metric
(
	metric_id varchar(50) NOT NULL,
	metric_name varchar(100) NOT NULL,
	rule_promql text,
	resource_level_id varchar(50) NOT NULL,
	resource_type varchar(20) NOT NULL,
	update_time timestamp DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (metric_id)
) DEFAULT CHARACTER SET latin1 COLLATE latin1_swedish_ci;


CREATE TABLE receiver
(
	receiver_id varchar(50) NOT NULL,
	receiver_group_id varchar(50) NOT NULL,
	system_user boolean DEFAULT '49',
	receiver_name varchar(50) NOT NULL,
	mail varchar(50),
	phone varchar(50),
	wechat varchar(100),
	webhook varchar(200),
	create_time timestamp DEFAULT CURRENT_TIMESTAMP,
	update_time timestamp DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (receiver_id)
) DEFAULT CHARACTER SET latin1 COLLATE latin1_swedish_ci;


CREATE TABLE receiver_group
(
	receiver_group_id varchar(50) NOT NULL,
	receiver_group_name varchar(50) NOT NULL,
	create_time timestamp DEFAULT CURRENT_TIMESTAMP,
	update_time timestamp DEFAULT CURRENT_TIMESTAMP,
	webhook varchar(300),
	webhook_enable boolean,
	create_by varchar(50) NOT NULL,
	PRIMARY KEY (receiver_group_id)
) DEFAULT CHARACTER SET latin1 COLLATE latin1_swedish_ci;


CREATE TABLE resource
(
	resource_id varchar(50) NOT NULL,
	resource_name varchar(50) NOT NULL,
	resource_group_id varchar(50) NOT NULL,
	resource_type varchar(20) NOT NULL,
	resource_level_id varchar(50) NOT NULL,
	owner varchar(50),
	create_time timestamp DEFAULT CURRENT_TIMESTAMP,
	update_time timestamp DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (resource_id)
) DEFAULT CHARACTER SET latin1 COLLATE latin1_swedish_ci;


CREATE TABLE resource_group
(
	resource_group_id varchar(50) NOT NULL,
	resource_group_name varchar(50) NOT NULL,
	create_time timestamp DEFAULT CURRENT_TIMESTAMP,
	update_time timestamp DEFAULT CURRENT_TIMESTAMP,
	create_by varchar(50) NOT NULL,
	PRIMARY KEY (resource_group_id)
) DEFAULT CHARACTER SET latin1 COLLATE latin1_swedish_ci;


CREATE TABLE resource_level
(
	resource_level_id varchar(50) NOT NULL,
	resource_type varchar(20) NOT NULL,
	description text,
	enable boolean DEFAULT '49',
	create_by varchar(50),
	create_time timestamp DEFAULT CURRENT_TIMESTAMP,
	update_time timestamp DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (resource_level_id)
) DEFAULT CHARACTER SET latin1 COLLATE latin1_swedish_ci;


--
-- /* Create Foreign Keys */
--
-- ALTER TABLE alert_rule
-- 	ADD FOREIGN KEY (ref_alert_rule_id)
-- 	REFERENCES alert_rule (alert_rule_id)
-- 	ON UPDATE RESTRICT
-- 	ON DELETE RESTRICT
-- ;
--
--
-- ALTER TABLE alert_binding
-- 	ADD FOREIGN KEY (alert_group_id)
-- 	REFERENCES alert_rule_group (alert_group_id)
-- 	ON UPDATE RESTRICT
-- 	ON DELETE RESTRICT
-- ;
--
--
-- ALTER TABLE alert_history
-- 	ADD FOREIGN KEY (alert_group_id)
-- 	REFERENCES alert_rule_group (alert_group_id)
-- 	ON UPDATE RESTRICT
-- 	ON DELETE RESTRICT
-- ;
--
--
-- ALTER TABLE alert_rule
-- 	ADD FOREIGN KEY (alert_group_id)
-- 	REFERENCES alert_rule_group (alert_group_id)
-- 	ON UPDATE RESTRICT
-- 	ON DELETE RESTRICT
-- ;
--
--
-- ALTER TABLE alert_binding
-- 	ADD FOREIGN KEY (severity_id)
-- 	REFERENCES alert_severity (severity_id)
-- 	ON UPDATE RESTRICT
-- 	ON DELETE RESTRICT
-- ;
--
--
-- ALTER TABLE alert_history
-- 	ADD FOREIGN KEY (severity_id)
-- 	REFERENCES alert_severity (severity_id)
-- 	ON UPDATE RESTRICT
-- 	ON DELETE RESTRICT
-- ;
--
--
-- ALTER TABLE alert_rule
-- 	ADD FOREIGN KEY (metric_id)
-- 	REFERENCES metric (metric_id)
-- 	ON UPDATE RESTRICT
-- 	ON DELETE RESTRICT
-- ;
--
--
-- ALTER TABLE alert_binding
-- 	ADD FOREIGN KEY (receiver_group_id)
-- 	REFERENCES receiver_group (receiver_group_id)
-- 	ON UPDATE RESTRICT
-- 	ON DELETE RESTRICT
-- ;
--
--
-- ALTER TABLE alert_history
-- 	ADD FOREIGN KEY (receiver_group_id)
-- 	REFERENCES receiver_group (receiver_group_id)
-- 	ON UPDATE RESTRICT
-- 	ON DELETE RESTRICT
-- ;
--
--
-- ALTER TABLE receiver
-- 	ADD FOREIGN KEY (receiver_group_id)
-- 	REFERENCES receiver_group (receiver_group_id)
-- 	ON UPDATE RESTRICT
-- 	ON DELETE RESTRICT
-- ;
--
--
-- ALTER TABLE alert_binding
-- 	ADD FOREIGN KEY (resource_group_id)
-- 	REFERENCES resource_group (resource_group_id)
-- 	ON UPDATE RESTRICT
-- 	ON DELETE RESTRICT
-- ;
--
--
-- ALTER TABLE alert_history
-- 	ADD FOREIGN KEY (resource_group_id)
-- 	REFERENCES resource_group (resource_group_id)
-- 	ON UPDATE RESTRICT
-- 	ON DELETE RESTRICT
-- ;
--
--
-- ALTER TABLE resource
-- 	ADD FOREIGN KEY (resource_group_id)
-- 	REFERENCES resource_group (resource_group_id)
-- 	ON UPDATE RESTRICT
-- 	ON DELETE RESTRICT
-- ;
--
--
-- ALTER TABLE alert_history
-- 	ADD FOREIGN KEY (resource_level_id)
-- 	REFERENCES resource_level (resource_level_id)
-- 	ON UPDATE RESTRICT
-- 	ON DELETE RESTRICT
-- ;
--
--
-- ALTER TABLE metric
-- 	ADD FOREIGN KEY (resource_level_id)
-- 	REFERENCES resource_level (resource_level_id)
-- 	ON UPDATE RESTRICT
-- 	ON DELETE RESTRICT
-- ;
--
--
-- ALTER TABLE resource
-- 	ADD FOREIGN KEY (resource_level_id)
-- 	REFERENCES resource_level (resource_level_id)
-- 	ON UPDATE RESTRICT
-- 	ON DELETE RESTRICT
-- ;
--
--
--
