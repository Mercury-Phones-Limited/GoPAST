CREATE TABLE `ps_auth` (

`id`              VARCHAR(20) NOT NULL,
`customer_id`     BIGINT UNSIGNED NOT NULL,
`auth_type`       ENUM('userpass','md5') NOT NULL,
`nonce_lifetime`  SMALLINT(5) UNSIGNED,
`md5_cred`        VARCHAR(20),
`password`        VARCHAR(20),
`refresh_token`   VARCHAR(20),
`oauth_clientid`  VARCHAR(20),
`oauth_secret`    VARCHAR(20),
`realm`           VARCHAR(20),
`username`        VARCHAR(20),
PRIMARY KEY(`id`)
)
ENGINE=InnoDB;
