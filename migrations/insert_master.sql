-- +migrate Up
INSERT INTO `authorities` (`authority_name`, `authority_code`) VALUES ("なし", 404);
INSERT INTO `authorities` (`authority_name`, `authority_code`) VALUES ("ルート", 777);
INSERT INTO `authorities` (`authority_name`, `authority_code`) VALUES ("管理者登録", 400);
INSERT INTO `authorities` (`authority_name`, `authority_code`) VALUES ("メルマガ送信", 123);
INSERT INTO `authorities` (`authority_name`, `authority_code`) VALUES ("スケジュール管理", 345);
INSERT INTO `authorities` (`authority_name`, `authority_code`) VALUES ("チケット取り消し", 500);
INSERT INTO `authorities` (`authority_name`, `authority_code`) VALUES ("売上管理", 100);
