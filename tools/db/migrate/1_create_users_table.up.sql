CREATE TABLE IF NOT EXISTS `users` (
    id INTEGER NOT NULL AUTO_INCREMENT,
    uuid varchar(255) NOT NULL ,
    name varchar(255) NOT NULL CHECK (name <> ''),
    email varchar(255) NOT NULL,CHECK (email <> ''),
    admin ENUM('0', '1') NOT NULL DEFAULT '0',
    hashed_password varchar(255) NOT NULL CHECK( octet_length(hashed_password) <> 0),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB COLLATE utf8mb4_bin;