CREATE TABLE USER
(
    id           BIGINT AUTO_INCREMENT COMMENT '用户id'
        PRIMARY KEY,
    username     VARCHAR(256)                       NULL COMMENT '用户昵称',
    userAccount  VARCHAR(256)                       NULL COMMENT '账号',
    avatarUrl    VARCHAR(1024)                      NULL COMMENT '用户头像',
    gender       TINYINT                            NULL COMMENT '性别',
    userPassword VARCHAR(512)                       NOT NULL COMMENT '密码',
    phone        VARCHAR(128)                       NULL COMMENT '电话号码',
    email        VARCHAR(512)                       NULL COMMENT '邮箱',
    userStatus   INT      DEFAULT 0                 NULL COMMENT '用户状态',
    userRole     INT      DEFAULT 0                 NULL COMMENT '用户权限',
    createTime   DATETIME DEFAULT CURRENT_TIMESTAMP NULL COMMENT '创建时间',
    updateTime   DATETIME DEFAULT CURRENT_TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    isDelete     TINYINT  DEFAULT 0                 NULL COMMENT '是否删除'
);
