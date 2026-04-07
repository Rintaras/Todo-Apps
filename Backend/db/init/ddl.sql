-- Todo-Apps: MySQL 8 初期スキーマ（OpenAPI `components/schemas/Todo` に対応）
-- 既定 DB 名は `TodoApp`（`env-sample` / docker-compose の MYSQL_DATABASE と一致）

SET NAMES utf8mb4;
SET time_zone = '+00:00';

CREATE DATABASE IF NOT EXISTS `TodoApp`
  CHARACTER SET utf8mb4
  COLLATE utf8mb4_unicode_ci;

USE `TodoApp`;

CREATE TABLE IF NOT EXISTS `todos` (
  `id` CHAR(36) NOT NULL COMMENT 'UUID（API の id と同一）',
  `title` VARCHAR(500) NOT NULL COMMENT 'タイトル（最大500文字）',
  `completed` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '1=完了, 0=未完了',
  `created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '作成日時(UTC)',
  `updated_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '更新日時(UTC)',
  PRIMARY KEY (`id`),
  KEY `idx_todos_completed` (`completed`),
  KEY `idx_todos_created_at` (`created_at`)
) ENGINE=InnoDB
  DEFAULT CHARSET=utf8mb4
  COLLATE=utf8mb4_unicode_ci
  COMMENT='Todo タスク';
