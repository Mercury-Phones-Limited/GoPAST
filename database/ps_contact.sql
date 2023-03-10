CREATE TABLE `ps_contact` (

`id`                    VARCHAR(200),
`uri`                   VARCHAR(100),
`expiration_time`       VARCHAR(20),
`qualify_frequency`     SMALLINT(5) UNSIGNED,
`qualify_timeout`       DECIMAL(3,2),
`authenticate_qualify`  ENUM('no','yes'),
`outbound_proxy`        VARCHAR(100),
`path`                  VARCHAR(100),
`user_agent`            VARCHAR(100),
`endpoint`              VARCHAR(20),
`reg_server`            VARCHAR(100),
`via_addr`              VARCHAR(100),
`via_port`              SMALLINT(5) UNSIGNED,
`call_id`               VARCHAR(200),
`prune_on_boot`         ENUM('no','yes'),
PRIMARY KEY(`id`)
)
ENGINE=InnoDB;
