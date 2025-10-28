CREATE TABLE `bom` (
                       `id` BIGINT UNSIGNED AUTO_INCREMENT ,
                       `parent_material_code` VARCHAR(100) NOT NULL COMMENT '父物料编码, 关联 materials.code',
                       `child_material_code` VARCHAR(100) NOT NULL COMMENT '子物料编码, 关联 materials.code',
                       `quantity` DECIMAL(10, 2) NOT NULL COMMENT '子物料用量，即组成1个父物料需要多少子物料',
                       `status` TINYINT NOT NULL DEFAULT 1 COMMENT '生效状态 (1:生效, 0:失效)',
                       `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间 (微秒精度)',
                       `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新时间 (微秒精度)',
                       PRIMARY KEY (`id`),
                       UNIQUE KEY `uk_parent_child` (`parent_material_code`, `child_material_code`),
                       FOREIGN KEY (`parent_material_code`) REFERENCES `materials`(`code`) ON DELETE CASCADE,
                       FOREIGN KEY (`child_material_code`) REFERENCES `materials`(`code`) ON DELETE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='BOM (物料清单) 关系表';