package db

func initMySQLTable() {
	initMysqlDB()
	sqlTable := "CREATE TABLE IF NOT EXISTS `user` ( " +
		"\n  `uid` varchar(64) NOT NULL," +
		"\n  `nick_name` varchar(64) DEFAULT NULL," +
		"\n  `head_img` varchar(1024) DEFAULT NULL," +
		"\n  `gender` int(1) unsigned zerofill DEFAULT NULL," +
		"\n  `phone` varchar(32) DEFAULT NULL," +
		"\n  `birth` varchar(16) DEFAULT NULL," +
		"\n  `email` varchar(64) DEFAULT NULL," +
		"\n  `pass_word` varchar(64) DEFAULT NULL," +
		"\n  `ex` varchar(1024) DEFAULT NULL," +
		"\n  `create_time` datetime DEFAULT NULL," +
		"\n  `update_time` datetime DEFAULT NULL," +
		"\n  PRIMARY KEY (`uid`)," +
		"\n  UNIQUE KEY `uk_uid` (`uid`)" +
		"\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;"
	_, err := Db.Exec(sqlTable)
	if err != nil {
		panic(err)
	}

	sqlTable = "CREATE TABLE IF NOT EXISTS `friend` (" +
		"\n  `owner_id` varchar(64) NOT NULL," +
		"\n  `friend_id` varchar(64) NOT NULL," +
		"\n  `comment` varchar(255) DEFAULT NULL," +
		"\n  `friend_flag` int(11) NOT NULL," +
		"\n  `create_time` datetime NOT NULL," +
		"\n  PRIMARY KEY (`owner_id`,`friend_id`) USING BTREE" +
		"\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;"
	_, err = Db.Exec(sqlTable)
	if err != nil {
		panic(err)
	}

	sqlTable = "CREATE TABLE IF NOT EXISTS  `friend_request` (" +
		"\n  `req_id` varchar(64) NOT NULL," +
		"\n  `user_id` varchar(64) NOT NULL," +
		"\n  `flag` int(11) NOT NULL DEFAULT '0'," +
		"\n  `req_message` varchar(255) DEFAULT NULL," +
		"\n  `create_time` datetime NOT NULL," +
		"\n  PRIMARY KEY (`user_id`,`req_id`) USING BTREE" +
		"\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;"
	_, err = Db.Exec(sqlTable)
	if err != nil {
		panic(err)
	}

	sqlTable = "CREATE TABLE IF NOT EXISTS `black_list` (" +
		"\n  `uid` varchar(64) NOT NULL COMMENT 'uid'," +
		"\n  `begin_disable_time` datetime DEFAULT NULL," +
		"\n  `end_disable_time` datetime DEFAULT NULL," +
		"\n  `ex` varchar(1024) DEFAULT NULL," +
		"\n  PRIMARY KEY (`uid`) USING BTREE" +
		"\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;"
	_, err = Db.Exec(sqlTable)
	if err != nil {
		panic(err)
	}

	sqlTable = "CREATE TABLE IF NOT EXISTS `user_black_list` (" +
		"\n  `owner_id` varchar(64) NOT NULL," +
		"\n  `black_id` varchar(64) NOT NULL," +
		"\n  `create_time` datetime NOT NULL," +
		"\n  PRIMARY KEY (`owner_id`,`black_id`) USING BTREE" +
		"\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;"
	_, err = Db.Exec(sqlTable)
	if err != nil {
		panic(err)
	}

	sqlTable = "CREATE TABLE IF NOT EXISTS `group` (" +
		"\n  `group_id` varchar(64) NOT NULL," +
		"\n  `group_name` varchar(255) DEFAULT NULL," +
		"\n  `introduction` varchar(255) DEFAULT NULL," +
		"\n  `notification` varchar(255) DEFAULT NULL," +
		"\n  `face_img` varchar(255) DEFAULT NULL," +
		"\n  `create_time` datetime DEFAULT NULL," +
		"\n  `update_time` datetime DEFAULT NULL," +
		"\n  `ex` varchar(255) DEFAULT NULL," +
		"\n  PRIMARY KEY (`group_id`)" +
		"\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;"
	_, err = Db.Exec(sqlTable)
	if err != nil {
		panic(err)
	}

	sqlTable = "CREATE TABLE IF NOT EXISTS `group_member` (" +
		"\n  `group_id` varchar(64) NOT NULL," +
		"\n  `user_id` varchar(64) NOT NULL," +
		"\n  `nick_name` varchar(255) DEFAULT NULL," +
		"\n  `user_group_face_url` varchar(255) DEFAULT NULL," +
		"\n  `administrator_level` int(11) NOT NULL," +
		"\n  `join_time` datetime NOT NULL," +
		"\n  PRIMARY KEY (`group_id`,`user_id`) USING BTREE" +
		"\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;"
	_, err = Db.Exec(sqlTable)
	if err != nil {
		panic(err)
	}

	sqlTable = "CREATE TABLE IF NOT EXISTS `group_request` (" +
		"\n  `id` int(11) NOT NULL AUTO_INCREMENT," +
		"\n  `group_id` varchar(64) NOT NULL," +
		"\n  `from_user_id` varchar(64) NOT NULL," +
		"\n  `to_user_id` varchar(64) NOT NULL," +
		"\n  `flag` int(10) NOT NULL DEFAULT '0'," +
		"\n  `req_message` varchar(255) DEFAULT ''," +
		"\n  `handled_msg` varchar(255) DEFAULT ''," +
		"\n  `create_time` datetime NOT NULL," +
		"\n  `from_user_nick_name` varchar(255) DEFAULT ''," +
		"\n  `to_user_nick_name` varchar(255) DEFAULT NULL," +
		"\n  `from_user_head_img` varchar(255) DEFAULT ''," +
		"\n  `to_user_head_img` varchar(255) DEFAULT ''," +
		"\n  `handled_user` varchar(255) DEFAULT ''," +
		"\n  PRIMARY KEY (`id`)\n) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8mb4;"
	_, err = Db.Exec(sqlTable)
	if err != nil {
		panic(err)
	}

	sqlTable = "CREATE TABLE IF NOT EXISTS  `chat_msg` (" +
		"\n  `msg_id` varchar(64) NOT NULL," +
		"\n  `send_id` varchar(64) NOT NULL," +
		"\n  `del_flag` int(11) NOT NULL," +
		"\n  `msg_type` int(11) NOT NULL," +
		"\n  `receive_id` varchar(255) NOT NULL," +
		"\n  `content_type` int(11) NOT NULL," +
		"\n  `sign_type` int(11) NOT NULL," +
		"\n  `content` varchar(1000) NOT NULL," +
		"\n  `send_time` datetime NOT NULL," +
		"\n  PRIMARY KEY (`msg_id`) USING BTREE" +
		"\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;"
	_, err = Db.Exec(sqlTable)
	if err != nil {
		panic(err)
	}

	defer Db.Close()
}
