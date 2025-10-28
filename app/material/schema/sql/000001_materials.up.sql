CREATE TABLE `materials` (
                             `id` BIGINT UNSIGNED AUTO_INCREMENT ,
                             `code` VARCHAR(100) NOT NULL UNIQUE COMMENT '物料编码，业务唯一键',
                             `name` VARCHAR(255) NOT NULL COMMENT '物料名称',
                             `material_type` VARCHAR(50) NULL COMMENT '物料类型 (例如: 成品, 半成品, 原料)',
                             `spec` VARCHAR(255) NULL COMMENT '规格型号',
                             `unit` VARCHAR(20) NULL COMMENT '单位 (例如: 个, kg, 米)',
                             `price` BIGINT NULL DEFAULT 0 COMMENT '单价, 存储为分，避免浮点数精度问题',
                             `stock_quantity` INT NOT NULL DEFAULT 0 COMMENT '库存数量',
                             `supplier_name` VARCHAR(255) NULL COMMENT '供应商, 简化处理',
                             `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间 (微秒精度)',
                             `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新时间 (微秒精度)',
                             PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='物料信息表';