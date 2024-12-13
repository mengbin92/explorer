-- 用户表（users）
CREATE TABLE `users` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `username` VARCHAR(255) NOT NULL UNIQUE,
  `email` VARCHAR(255) NOT NULL UNIQUE,
  `password_hash` VARCHAR(255) NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `status` ENUM('active', 'inactive', 'suspended', 'deleted') NOT NULL DEFAULT 'active',
  `last_login_at` DATETIME DEFAULT NULL,
  `api_key` VARCHAR(255) DEFAULT NULL,
  `two_factor_enabled` BOOLEAN DEFAULT FALSE,
  `role` ENUM('user', 'admin', 'super_admin') NOT NULL DEFAULT 'user'
);

-- 用户身份验证表（auth_logs）
CREATE TABLE `auth_logs` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `user_id` INT NOT NULL,
  `ip_address` VARCHAR(45) NOT NULL,
  `device_info` TEXT,
  `login_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `logout_at` DATETIME DEFAULT NULL,
  `success` BOOLEAN NOT NULL,
  `reason` VARCHAR(255) DEFAULT NULL,
  FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
);

-- API密钥表（api_keys）
CREATE TABLE `api_keys` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `user_id` INT NOT NULL,
  `api_key` VARCHAR(255) NOT NULL UNIQUE,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `expires_at` DATETIME DEFAULT NULL,
  `permissions` JSON DEFAULT NULL,
  `status` ENUM('active', 'revoked') NOT NULL DEFAULT 'active',
  FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
);

-- 权限表（permissions）
CREATE TABLE `permissions` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `api_key_id` INT NOT NULL,
  `permission` ENUM('read', 'write', 'delete', 'admin') NOT NULL,
  `resource` VARCHAR(255) NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (`api_key_id`) REFERENCES `api_keys`(`id`)
);

-- 用户活动表（user_activity）
CREATE TABLE `user_activity` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `user_id` INT NOT NULL,
  `activity_type` ENUM('login', 'api_call', 'update_profile', 'create_project') NOT NULL,
  `details` TEXT,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
);

-- 角色权限表（role_permissions）
CREATE TABLE `role_permissions` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `role` ENUM('user', 'admin', 'super_admin') NOT NULL,
  `permission` ENUM('read', 'write', 'delete', 'admin') NOT NULL,
  `resource` VARCHAR(255) NOT NULL
);

-- 用户支付表（billing_info） - 可选
CREATE TABLE `billing_info` (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `user_id` INT NOT NULL,
  `subscription_plan` VARCHAR(50) NOT NULL,
  `billing_cycle` ENUM('monthly', 'yearly') NOT NULL,
  `amount` DECIMAL(10, 2) NOT NULL,
  `payment_status` ENUM('paid', 'pending', 'failed') NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
);

-- 创建索引
CREATE INDEX idx_users_username ON `users` (`username`);
CREATE INDEX idx_users_email ON `users` (`email`);
CREATE INDEX idx_api_keys_api_key ON `api_keys` (`api_key`);
CREATE INDEX idx_auth_logs_user_id ON `auth_logs` (`user_id`);

