-- -----------------------------------------------------
-- Schema article
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `article` DEFAULT CHARACTER SET utf8mb4 ;
USE `article` ;

-- -----------------------------------------------------
-- Table `article`.`article`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `article`.`article` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(128) NOT NULL COMMENT '文章标题',
  `preview_img` VARCHAR(256) NOT NULL COMMENT '预览图片',
  `desc` VARCHAR(500) NOT NULL COMMENT '文章描述，用于预览视图，如果不填，需要取内容的前100字符',
  `content` LONGTEXT NOT NULL COMMENT '文章内容',
  `status` TINYINT UNSIGNED NOT NULL COMMENT '文章状态：0-未启用，1-启用，2-禁用。启用时，前台将能够看到该文章',
  `gmt_publish` DATETIME NULL COMMENT '启用时间，在每次启用时更新',
  `gmt_create` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `gmt_modified` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新事件',
  PRIMARY KEY (`id`),
  INDEX `idx_gmt_publish` (`gmt_publish` DESC));
