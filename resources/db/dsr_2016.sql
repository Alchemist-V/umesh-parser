CREATE TABLE `umesh_app`.`dsr_2016` (
  `Code_Number` VARCHAR(20) NOT NULL,
  `Subhead_Type` VARCHAR(45) NULL,
  `Description` MEDIUMTEXT NULL,
  `Parent_Code_Number` VARCHAR(45) NULL,
  `Parent_Description` MEDIUMTEXT NULL,
  `Unit` VARCHAR(45) NULL,
  `Amount` FLOAT NULL,
  `Created_Date` DATETIME NULL,
  `Modified_Date` DATETIME NULL,
  PRIMARY KEY (`Code_Number`));
