CREATE TABLE `message_task` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `media_id` int(11) NOT NULL COMMENT '媒体ID',
  `template_id` char(64) NOT NULL DEFAULT '' COMMENT '模板ID',
  `template_title` char(32) NOT NULL DEFAULT '' COMMENT '模板名称',
  `name` varchar(32) NOT NULL DEFAULT '' COMMENT '任务名称',
  `user_ids` text NOT NULL COMMENT 'Json 待发送user_id',
  `sql` varchar(64) NOT NULL DEFAULT '' COMMENT '待发送获取用户id sql',
  `message` varchar(255) NOT NULL DEFAULT '' COMMENT 'Json 模板消息内容',
  `action_path` char(64) NOT NULL DEFAULT '' COMMENT '跳转小程序路径',
  `start_at` int(11) NOT NULL DEFAULT '0' COMMENT '开始时间',
  `send_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '发送状态,1:待发送 2:发送中 3:完成 4:失败',
  `succ_count` int(11) NOT NULL DEFAULT '0' COMMENT '成功数',
  `fail_count` int(11) NOT NULL DEFAULT '0' COMMENT '失败数',
  `status` tinyint(2) NOT NULL DEFAULT '1' COMMENT '状态 0:正常 1:删除',
  `created_at` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `updated_at` int(11) NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='模板消息任务表';
