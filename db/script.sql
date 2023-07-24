CREATE  TABLE `todo_items` (
    `id` int NOT NULL AUTO_INCREMENT,
    `title` varchar(150) CHARACTER SET utf8 NOT NULL,
    `status` enum('Doing','Finished') DEFAULT 'Doing',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO golang.todo_items
	( title) VALUES ("study golang"), ("study c#");