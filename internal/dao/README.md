数据访问层（Database Access Object）

CREATE DATABASE
IF
    NOT EXISTS blog_service DEFAULT CHARACTER
    SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;


更新环境变量

export PATH=$PATH:/usr/local/mysql/bin
source /etc/profile