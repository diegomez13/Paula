
CREATE TABLE `categorias` (
  `id_cat` int(4) NOT NULL AUTO_INCREMENT,
  `nombre` varchar(100) NOT NULL,
  `fecha` datetime NOT NULL,
  `id_sup` int(4) NOT NULL,
  PRIMARY KEY (`id_cat`),
  KEY `id_sup` (`id_sup`),
  CONSTRAINT `categorias_ibfk_1` FOREIGN KEY (`id_sup`) REFERENCES `supermercados` (`id_sup`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf32 COLLATE=utf32_spanish2_ci;

CREATE TABLE `permisos` (
  `id_usr` int(4) NOT NULL,
  `id_sup` int(4) NOT NULL,
  `p_addcat` tinyint(1) NOT NULL,
  `p_addpro` tinyint(1) NOT NULL,
  PRIMARY KEY (`id_usr`,`id_sup`),
  KEY `id_sup` (`id_sup`),
  CONSTRAINT `permisos_ibfk_1` FOREIGN KEY (`id_sup`) REFERENCES `supermercados` (`id_sup`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `permisos_ibfk_2` FOREIGN KEY (`id_usr`) REFERENCES `usuarios` (`id_usr`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf32 COLLATE=utf32_spanish2_ci;

CREATE TABLE `productos` (
  `id_pro` int(4) NOT NULL AUTO_INCREMENT,
  `nombre` varchar(100) NOT NULL,
  `precio` int(4) NOT NULL,
  `fecha` datetime NOT NULL,
  `id_cat` int(4) NOT NULL,
  PRIMARY KEY (`id_pro`),
  KEY `id_cat` (`id_cat`),
  CONSTRAINT `productos_ibfk_1` FOREIGN KEY (`id_cat`) REFERENCES `categorias` (`id_cat`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf32 COLLATE=utf32_spanish2_ci;

CREATE TABLE `supermercados` (
  `id_sup` int(4) NOT NULL AUTO_INCREMENT,
  `nombre` varchar(100) NOT NULL,
  `direccion` varchar(100) NOT NULL,
  `fecha` datetime NOT NULL,
  PRIMARY KEY (`id_sup`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf32 COLLATE=utf32_spanish2_ci;

CREATE TABLE `usuarios` (
  `id_usr` int(4) NOT NULL AUTO_INCREMENT,
  `nombre` varchar(100) DEFAULT NULL,
  `password` varchar(100) NOT NULL,
  `correo` varchar(100) NOT NULL,
  `cookie` varchar(100) DEFAULT NULL,
  `p_addsup` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id_usr`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf32 COLLATE=utf32_spanish2_ci;

