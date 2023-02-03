CREATE TABLE `ps_aor` (

`id`                    VARCHAR(20) NOT NULL,
`customer_id`           BIGINT UNSIGNED NOT NULL,
`contact`               VARCHAR(20),
`default_expiration`    SMALLINT(5) UNSIGNED,
`mailboxes`             VARCHAR(20),
`voicemail_extension`   VARCHAR(20),
`maximum_expiration`    SMALLINT(5) UNSIGNED,
`max_contacts`          TINYINT(3) UNSIGNED,
`minimum_expiration`    SMALLINT(5) UNSIGNED,
`remove_existing`       ENUM('yes','no') NOT NULL,
`remove_unavailable`    ENUM('no','yes') NOT NULL,
`qualify_frequency`     SMALLINT(5) UNSIGNED,
`qualify_timeout`       DECIMAL(3,2),
`authenticate_qualify`  ENUM('no','yes'),
`outbound_proxy`        VARCHAR(100),
`support_path`          ENUM('no','yes'),
PRIMARY KEY(`id`)
)
ENGINE=InnoDB;
