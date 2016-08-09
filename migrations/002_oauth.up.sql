CREATE TABLE `access` (
  `client` varchar(255) NOT NULL,
  `authorize` varchar(255) NOT NULL,
  `previous` varchar(255) NOT NULL,
  `access_token` varchar(255) NOT NULL,
  `refresh_token` varchar(255) NOT NULL,
  `expires_in` int(10) NOT NULL,
  `scope` varchar(255) NOT NULL,
  `redirect_uri` varchar(255) NOT NULL,
  `extra` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`access_token`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `authorize` (
  `client` varchar(255) NOT NULL,
  `code` varchar(255) NOT NULL,
  `expires_in` int(10) NOT NULL,
  `scope` varchar(255) NOT NULL,
  `redirect_uri` varchar(255) NOT NULL,
  `state` varchar(255) NOT NULL,
  `extra` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `client` (
  `id` varchar(255) NOT NULL,
  `secret` varchar(255) NOT NULL,
  `extra` varchar(255) NOT NULL,
  `redirect_uri` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `refresh` (
  `token` varchar(255) NOT NULL,
  `access` varchar(255) NOT NULL,
  PRIMARY KEY (`token`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
