

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
--
-- Table structure for table `countries`
--
CREATE DATABASE IF NOT EXISTS session;
USE session;

DROP TABLE IF EXISTS `countries`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `countries` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `Name` varchar(50) COLLATE utf8_bin NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=197 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `countries`
--

LOCK TABLES `countries` WRITE;
/*!40000 ALTER TABLE `countries` DISABLE KEYS */;
INSERT INTO `countries` VALUES (1,'Afghanistan'),(2,'Albania'),(3,'Algeria'),(4,'Andorra'),(5,'Angola'),(6,'Antigua & Deps'),(7,'Argentina'),(8,'Armenia'),(9,'Australia'),(10,'Austria'),(11,'Azerbaijan'),(12,'Bahamas'),(13,'Bahrain'),(14,'Bangladesh'),(15,'Barbados'),(16,'Belarus'),(17,'Belgium'),(18,'Belize'),(19,'Benin'),(20,'Bhutan'),(21,'Bolivia'),(22,'Bosnia Herzegovina'),(23,'Botswana'),(24,'Brazil'),(25,'Brunei'),(26,'Bulgaria'),(27,'Burkina'),(28,'Burundi'),(29,'Cambodia'),(30,'Cameroon'),(31,'Canada'),(32,'Cape Verde'),(33,'Central African Rep'),(34,'Chad'),(35,'Chile'),(36,'China'),(37,'Colombia'),(38,'Comoros'),(39,'Congo'),(40,'Congo {Democratic Rep}'),(41,'Costa Rica'),(42,'Croatia'),(43,'Cuba'),(44,'Cyprus'),(45,'Czech Republic'),(46,'Denmark'),(47,'Djibouti'),(48,'Dominica'),(49,'Dominican Republic'),(50,'East Timor'),(51,'Ecuador'),(52,'Egypt'),(53,'El Salvador'),(54,'Equatorial Guinea'),(55,'Eritrea'),(56,'Estonia'),(57,'Ethiopia'),(58,'Fiji'),(59,'Finland'),(60,'France'),(61,'Gabon'),(62,'Gambia'),(63,'Georgia'),(64,'Germany'),(65,'Ghana'),(66,'Greece'),(67,'Grenada'),(68,'Guatemala'),(69,'Guinea'),(70,'Guinea-Bissau'),(71,'Guyana'),(72,'Haiti'),(73,'Honduras'),(74,'Hungary'),(75,'Iceland'),(76,'India'),(77,'Indonesia'),(78,'Iran'),(79,'Iraq'),(80,'Ireland {Republic}'),(81,'Israel'),(82,'Italy'),(83,'Ivory Coast'),(84,'Jamaica'),(85,'Japan'),(86,'Jordan'),(87,'Kazakhstan'),(88,'Kenya'),(89,'Kiribati'),(90,'Korea North'),(91,'Korea South'),(92,'Kosovo'),(93,'Kuwait'),(94,'Kyrgyzstan'),(95,'Laos'),(96,'Latvia'),(97,'Lebanon'),(98,'Lesotho'),(99,'Liberia'),(100,'Libya'),(101,'Liechtenstein'),(102,'Lithuania'),(103,'Luxembourg'),(104,'Macedonia'),(105,'Madagascar'),(106,'Malawi'),(107,'Malaysia'),(108,'Maldives'),(109,'Mali'),(110,'Malta'),(111,'Marshall Islands'),(112,'Mauritania'),(113,'Mauritius'),(114,'Mexico'),(115,'Micronesia'),(116,'Moldova'),(117,'Monaco'),(118,'Mongolia'),(119,'Montenegro'),(120,'Morocco'),(121,'Mozambique'),(122,'Myanmar, {Burma}'),(123,'Namibia'),(124,'Nauru'),(125,'Nepal'),(126,'Netherlands'),(127,'New Zealand'),(128,'Nicaragua'),(129,'Niger'),(130,'Nigeria'),(131,'Norway'),(132,'Oman'),(133,'Pakistan'),(134,'Palau'),(135,'Panama'),(136,'Papua New Guinea'),(137,'Paraguay'),(138,'Peru'),(139,'Philippines'),(140,'Poland'),(141,'Portugal'),(142,'Qatar'),(143,'Romania'),(144,'Russian Federation'),(145,'Rwanda'),(146,'St Kitts & Nevis'),(147,'St Lucia'),(148,'Saint Vincent & the Grenadines'),(149,'Samoa'),(150,'San Marino'),(151,'Sao Tome & Principe'),(152,'Saudi Arabia'),(153,'Senegal'),(154,'Serbia'),(155,'Seychelles'),(156,'Sierra Leone'),(157,'Singapore'),(158,'Slovakia'),(159,'Slovenia'),(160,'Solomon Islands'),(161,'Somalia'),(162,'South Africa'),(163,'South Sudan'),(164,'Spain'),(165,'Sri Lanka'),(166,'Sudan'),(167,'Suriname'),(168,'Swaziland'),(169,'Sweden'),(170,'Switzerland'),(171,'Syria'),(172,'Taiwan'),(173,'Tajikistan'),(174,'Tanzania'),(175,'Thailand'),(176,'Togo'),(177,'Tonga'),(178,'Trinidad & Tobago'),(179,'Tunisia'),(180,'Turkey'),(181,'Turkmenistan'),(182,'Tuvalu'),(183,'Uganda'),(184,'Ukraine'),(185,'United Arab Emirates'),(186,'United Kingdom'),(187,'United States'),(188,'Uruguay'),(189,'Uzbekistan'),(190,'Vanuatu'),(191,'Vatican City'),(192,'Venezuela'),(193,'Vietnam'),(194,'Yemen'),(195,'Zambia'),(196,'Zimbabwe');
/*!40000 ALTER TABLE `countries` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `offices`
--

DROP TABLE IF EXISTS `offices`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `offices` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `CountryID` int(11) NOT NULL,
  `Title` varchar(50) COLLATE utf8_bin NOT NULL,
  `Phone` varchar(50) COLLATE utf8_bin NOT NULL,
  `Contact` varchar(250) COLLATE utf8_bin NOT NULL,
  PRIMARY KEY (`ID`),
  KEY `FK_Office_Country` (`CountryID`),
  CONSTRAINT `FK_Office_Country` FOREIGN KEY (`CountryID`) REFERENCES `countries` (`ID`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `offices`
--

LOCK TABLES `offices` WRITE;
/*!40000 ALTER TABLE `offices` DISABLE KEYS */;
INSERT INTO `offices` VALUES (1,185,'Abu dhabi','638-757-8582\r\n','MIchael Malki'),(3,52,'Cairo','252-224-8525','David Johns'),(4,13,'Bahrain','542-227-5825','Katie Ballmer'),(5,142,'Doha','758-278-9597','Ariel Levy'),(6,152,'Riyadh','285-285-1474','Andrew Hobart');
/*!40000 ALTER TABLE `offices` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `roles`
--

DROP TABLE IF EXISTS `roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `roles` (
  `ID` int(11) NOT NULL,
  `Title` varchar(50) COLLATE utf8_bin NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `roles`
--

LOCK TABLES `roles` WRITE;
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
INSERT INTO `roles` VALUES (1,'Administrator'),(2,'User');
/*!40000 ALTER TABLE `roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `RoleID` int(11) NOT NULL,
  `Email` varchar(150) COLLATE utf8_bin NOT NULL,
  `Password` varchar(50) COLLATE utf8_bin NOT NULL,
  `FirstName` varchar(50) COLLATE utf8_bin DEFAULT NULL,
  `LastName` varchar(50) COLLATE utf8_bin NOT NULL,
  `OfficeID` int(11) DEFAULT NULL,
  `Birthdate` date DEFAULT NULL,
  `Active` int(2) DEFAULT NULL,
  `Banned` BOOLEAN DEFAULT 0,
  `BanReason` TEXT DEFAULT NULL,
  PRIMARY KEY (`ID`),
  FOREIGN KEY (`Active`) REFERENCES `user_status` (`ID`),
  KEY `FK_Users_Offices` (`OfficeID`),
  KEY `FK_Users_Roles` (`RoleID`),
  CONSTRAINT `FK_Users_Offices` FOREIGN KEY (`OfficeID`) REFERENCES `offices` (`ID`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `FK_Users_Roles` FOREIGN KEY (`RoleID`) REFERENCES `roles` (`ID`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Create trigger for calc Age column
--
CREATE TRIGGER after_users_insert
AFTER INSERT ON `users`
FOR EACH ROW
    INSERT INTO `admin_panel` (`UserID`, `Age`)
    VALUES (NEW.ID, TIMESTAMPDIFF(YEAR, NEW.Birthdate, CURDATE()));

--
-- Table structure for table `user_status`
-- 

DROP TABLE IF EXISTS `user_status`;
CREATE TABLE `user_status` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `Title` varchar(250) COLLATE utf8_bin NOT NULL,
  PRIMARY KEY (`ID`)
);

ALTER TABLE `user_status` AUTO_INCREMENT=0;
LOCK TABLES `user_status` WRITE;
INSERT INTO `user_status` VALUES (1, 'Enable'),(0, 'Disable');
UNLOCK TABLES;

--
-- table structure for table `admin_panel`
-- 

drop table if exists `admin_panel`;
create table `admin_panel` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `UserID` int(11) NOT NULL,
  `Age` int(3) NOT NULL,
  PRIMARY KEY (`ID`),
  FOREIGN KEY (`UserID`) REFERENCES `users` (`ID`)
);

--
-- Table structure for table `user_panel`
-- 

DROP TABLE IF EXISTS `user_panel`;
CREATE TABLE `user_panel` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `UserID` int(11) NOT NULL,
  `Date` date NOT NULL,
  `LoginTime` date NOT NULL,
  `LogoutTime` date DEFAULT NULL,
  `LogoutReason` text DEFAULT NULL,
  PRIMARY KEY (`ID`),
  FOREIGN KEY (`UserID`) REFERENCES `users` (`ID`)
);

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
-- SET GLOBAL local_infile = 1;
INSERT INTO users (ID, RoleID, Email, Password, FirstName, LastName, OfficeID, Birthdate, Active, Banned, BanReason) VALUES
    (1, 1, 'j.doe@amonic.com', '123', 'John', 'Doe', 1, '1983-01-13', 1, 0, NULL),
    (2, 2, 'k.omar@amonic.com', '4258', 'Karim', 'Omar', 1, '1980-03-19', 1, 0, NULL),
    (3, 2, 'h.saeed@amonic.com', '2020', 'Hannan', 'Saeed', 3, '1989-12-20', 1, 0, NULL),
    (4, 2, 'a.hobart@amonic.com', '6996', 'Andrew', 'Hobart', 6, '1990-01-30', 1, 0, NULL),
    (5, 2, 'k.anderson@amonic.com', '4570', 'Katrin', 'Anderson', 5, '1992-11-10', 1, 0, NULL),
    (6, 2, 'h.wyrick@amonic.com', '1199', 'Hava', 'Wyrick', 1, '1988-08-08', 1, 0, NULL),
    (7, 2, 'marie.horn@amonic.com', '55555', 'Marie', 'Horn', 4, '1981-04-06', 1, 0, NULL),
    (8, 2, 'm.osteen@amonic.com', '9800', 'Milagros', 'Osteen', 1, '1991-02-03', 0, NULL);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;


/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

