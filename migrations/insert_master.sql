-- +migrate Up
INSERT INTO `authorities` (`authority_name`, `authority_code`) VALUES ("なし", 404);
INSERT INTO `authorities` (`authority_name`, `authority_code`) VALUES ("ルート", 777);
INSERT INTO `authorities` (`authority_name`, `authority_code`) VALUES ("管理者登録", 400);
INSERT INTO `authorities` (`authority_name`, `authority_code`) VALUES ("メルマガ送信", 123);
INSERT INTO `authorities` (`authority_name`, `authority_code`) VALUES ("スケジュール管理", 345);
INSERT INTO `authorities` (`authority_name`, `authority_code`) VALUES ("チケット取り消し", 500);
INSERT INTO `authorities` (`authority_name`, `authority_code`) VALUES ("売上管理", 100);

INSERT INTO `prices` (`customer_type`, `price`, `last_updated_administrator_id`) VALUES ("一般", 1800, "0");
INSERT INTO `prices` (`customer_type`, `price`, `last_updated_administrator_id`) VALUES ("学生（大高生）", 1500, "0");
INSERT INTO `prices` (`customer_type`, `price`, `last_updated_administrator_id`) VALUES ("小中学生", 800, "0");
INSERT INTO `prices` (`customer_type`, `price`, `last_updated_administrator_id`) VALUES ("幼児", 500, "0");
INSERT INTO `prices` (`customer_type`, `price`, `last_updated_administrator_id`) VALUES ("3D専用メガネ代", 300, "0");

-- +migrate Down
