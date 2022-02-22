USE easy_testing;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for student
-- ----------------------------
DROP TABLE IF EXISTS `student`;
CREATE TABLE `student`
(
    `id`         int(10)            NOT NULL AUTO_INCREMENT COMMENT '用作主键',
    `student_id` varchar(10) UNIQUE NOT NULL COMMENT '学号',
    `name`       varchar(50)        NOT NULL COMMENT '姓名',
    `class_id`   varchar(10)        NOT NULL COMMENT '班号',
    PRIMARY KEY (`id`),
    INDEX (`student_id`),
    INDEX (`name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Table structure for teacher
-- ----------------------------
DROP TABLE IF EXISTS `teacher`;
CREATE TABLE `teacher`
(
    `id`         int(10)            NOT NULL AUTO_INCREMENT COMMENT '用作主键',
    `teacher_id` varchar(10) UNIQUE NOT NULL COMMENT '工号',
    `name`       varchar(50)        NOT NULL COMMENT '姓名',
    `password`   varchar(100)       NOT NULL COMMENT '加盐后的密码',
    `salt`       varchar(50)        NOT NULL COMMENT '盐',
    `is_admin`   bool               NOT NULL DEFAULT FALSE COMMENT '是否为超级管理员',
    PRIMARY KEY (`id`),
    INDEX (`teacher_id`),
    INDEX (`name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Table structure for exam
-- ----------------------------
DROP TABLE IF EXISTS `exam`;
CREATE TABLE `exam`
(
    `id`                   int(10)             NOT NULL AUTO_INCREMENT COMMENT '用作主键',
    `publisher_teacher_id` varchar(10)         NOT NULL COMMENT '发布考试的教师的工号',
    `scores_calculated`    bool DEFAULT FALSE COMMENT '参加这场考试的人的成绩是否已被计算过',
    `start_time`           datetime            NOT NULL COMMENT '考试开始时刻',
    `end_time`             datetime            NOT NULL COMMENT '考试结束时刻',
    `time_allowed`         tinyint(3)          NOT NULL COMMENT '考生答题时间，单位：分钟',
    `mcq_score`            tinyint(2) unsigned NOT NULL COMMENT '单选题每题分数',
    `mcq_num`              tinyint(2) unsigned NOT NULL COMMENT '单选题题数',
    `maq_score`            tinyint(2) unsigned NOT NULL COMMENT '多选题每题分数',
    `maq_num`              tinyint(2) unsigned NOT NULL COMMENT '多选题题数',
    `bfq_score`            tinyint(2) unsigned NOT NULL COMMENT '填空题每题分数',
    `bfq_num`              tinyint(2) unsigned NOT NULL COMMENT '填空题题数',
    `tfq_score`            tinyint(2) unsigned NOT NULL COMMENT '判断题每题分数',
    `tfq_num`              tinyint(2) unsigned NOT NULL COMMENT '判断题题数',
    `crq_score`            tinyint(2) unsigned NOT NULL COMMENT '代码阅读题每题分数',
    `crq_num`              tinyint(2) unsigned NOT NULL COMMENT '代码阅读题题数',
    `cq_score`             tinyint(2) unsigned NOT NULL COMMENT '写代码题每题分数',
    `cq_num`               tinyint(2) unsigned NOT NULL COMMENT '写代码题题数',
    FOREIGN KEY (`publisher_teacher_id`) REFERENCES teacher (`teacher_id`) ON DELETE CASCADE ON UPDATE CASCADE,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Table structure for exam session
-- ----------------------------
DROP TABLE IF EXISTS `exam_session`;
CREATE TABLE `exam_session`
(
    `id`           int(10)     NOT NULL AUTO_INCREMENT COMMENT '用作主键',
    `exam_id`      int(10)     NOT NULL COMMENT '连接 exam',
    `student_id`   varchar(10) NOT NULL COMMENT '连接 student',
    `student_name` varchar(50) NOT NULL COMMENT '考生的姓名',
    `start_time`   datetime    NOT NULL COMMENT '作答开始时刻',
    `time_allowed` tinyint(3)  NOT NULL COMMENT '考生答题时间，单位：分钟',
    `end_time`     datetime DEFAULT NULL COMMENT '交卷时刻',
    `score`        smallint    NOT NULL COMMENT '最终成绩*10，即保存到小数点后一位',
    FOREIGN KEY (`exam_id`) REFERENCES exam (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (`student_id`) REFERENCES student (`student_id`) ON DELETE CASCADE ON UPDATE CASCADE,
    PRIMARY KEY (`id`),
    UNIQUE INDEX (`exam_id`, `student_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Table structure for multiple choice question (mcq)
-- ----------------------------
DROP TABLE IF EXISTS `mcq`;
CREATE TABLE `mcq`
(
    `id`                    int(10)     NOT NULL AUTO_INCREMENT COMMENT '用作主键',
    `publisher_teacher_id`  varchar(10) NOT NULL COMMENT '创建本题的教师的工号',
    `stem`                  text(200)   NOT NULL COMMENT '题干',
    `choice_1`              text(200)   NOT NULL COMMENT '选项的内容',
    `choice_2`              text(200)   NOT NULL COMMENT '选项的内容',
    `choice_3`              text(200)   NOT NULL COMMENT '选项的内容',
    `choice_4`              text(200)   NOT NULL COMMENT '选项的内容',
    `right_answer`          char(1)     NOT NULL COMMENT '答案，正确选项的索引，如 "4"、"1"',
    `overall_correct_score` int         NOT NULL DEFAULT 0 COMMENT '此题在所有出现中的总得分数*10，即保存到小数点后一位',
    `overall_score`         int         NOT NULL DEFAULT 0 COMMENT '此题在所有出现中的总分数*10，即保存到小数点后一位',
    FOREIGN KEY (`publisher_teacher_id`) REFERENCES teacher (`teacher_id`) ON DELETE CASCADE ON UPDATE CASCADE,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Table structure for students' answer of multiple choice question (mcq)
-- ----------------------------
DROP TABLE IF EXISTS `mcq_answer`;
CREATE TABLE `mcq_answer`
(
    `id`              int(10) NOT NULL AUTO_INCREMENT COMMENT '用作主键',
    `mcq_id`          int(10) NOT NULL COMMENT '连接 mcq',
    `exam_session_id` int(10) NOT NULL COMMENT '连接 exam_session',
    `right_answer`    char(1) NOT NULL COMMENT '正确答案，与 mcq 中同名字段保持一致',
    `student_answer`  char(1) DEFAULT NULL COMMENT '学生的答案',
    FOREIGN KEY (`mcq_id`) REFERENCES mcq (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (`exam_session_id`) REFERENCES exam_session (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    PRIMARY KEY (`id`),
    -- unique composite index to speed up execution and avoid deadlocks on sql
    -- like `UPDATE mcq_answer SET student_answer='3' WHERE mcq_id = 4 AND exam_session_id = 22`
    UNIQUE INDEX (`mcq_id`, `exam_session_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Table structure for multiple answer question (maq)
-- ----------------------------
DROP TABLE IF EXISTS `maq`;
CREATE TABLE `maq`
(
    `id`                    int(10)     NOT NULL AUTO_INCREMENT COMMENT '用作主键',
    `publisher_teacher_id`  varchar(10) NOT NULL COMMENT '创建本题的教师的工号',
    `stem`                  text(200)   NOT NULL COMMENT '题干',
    `choice_1`              text(200)   NOT NULL COMMENT '选项的内容',
    `choice_2`              text(200)   NOT NULL COMMENT '选项的内容',
    `choice_3`              text(200)   NOT NULL COMMENT '选项的内容',
    `choice_4`              text(200)   NOT NULL COMMENT '选项的内容',
    `choice_5`              text(200)            DEFAULT NULL COMMENT '选项的内容',
    `choice_6`              text(200)            DEFAULT NULL COMMENT '选项的内容',
    `choice_7`              text(200)            DEFAULT NULL COMMENT '选项的内容',
    `right_answer`          char(7)     NOT NULL COMMENT '答案，按升序包含所有正确选项的索引，如 "5"、"124"、"67"',
    `overall_correct_score` int         NOT NULL DEFAULT 0 COMMENT '此题在所有出现中的总得分数*10，即保存到小数点后一位',
    `overall_score`         int         NOT NULL DEFAULT 0 COMMENT '此题在所有出现中的总分数*10，即保存到小数点后一位',
    FOREIGN KEY (`publisher_teacher_id`) REFERENCES teacher (`teacher_id`) ON DELETE CASCADE ON UPDATE CASCADE,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Table structure for students' answer of multiple answer question (maq)
-- ----------------------------
DROP TABLE IF EXISTS `maq_answer`;
CREATE TABLE `maq_answer`
(
    `id`              int(10) NOT NULL AUTO_INCREMENT COMMENT '用作主键',
    `maq_id`          int(10) NOT NULL COMMENT '连接 maq',
    `exam_session_id` int(10) NOT NULL COMMENT '连接 exam_session',
    `right_answer`    char(7) NOT NULL COMMENT '正确答案，与 maq 中同名字段保持一致',
    `student_answer`  char(7) DEFAULT NULL COMMENT '学生的答案',
    FOREIGN KEY (`maq_id`) REFERENCES maq (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (`exam_session_id`) REFERENCES exam_session (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    PRIMARY KEY (`id`),
    UNIQUE INDEX (`maq_id`, `exam_session_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Table structure for blank filling question (bfq)
-- ----------------------------
DROP TABLE IF EXISTS `bfq`;
CREATE TABLE `bfq`
(
    `id`                    int(10)     NOT NULL AUTO_INCREMENT COMMENT '用作主键',
    `publisher_teacher_id`  varchar(10) NOT NULL COMMENT '创建本题的教师的工号',
    `stem`                  text(200)   NOT NULL COMMENT '题干',
    `blank_num`             tinyint(2)  NOT NULL COMMENT '要填的空的数目',
    `answer_1`              text(50)    NOT NULL COMMENT '填空的答案',
    `answer_2`              text(50)             DEFAULT NULL COMMENT '填空的答案',
    `answer_3`              text(50)             DEFAULT NULL COMMENT '填空的答案',
    `overall_correct_score` int         NOT NULL DEFAULT 0 COMMENT '此题在所有出现中的总得分数*10，即保存到小数点后一位',
    `overall_score`         int         NOT NULL DEFAULT 0 COMMENT '此题在所有出现中的总分数*10，即保存到小数点后一位',
    FOREIGN KEY (`publisher_teacher_id`) REFERENCES teacher (`teacher_id`) ON DELETE CASCADE ON UPDATE CASCADE,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Table structure for students' answer of blank filling question (bfq)
-- ----------------------------
DROP TABLE IF EXISTS `bfq_answer`;
CREATE TABLE `bfq_answer`
(
    `id`               int(10) NOT NULL AUTO_INCREMENT COMMENT '用作主键',
    `bfq_id`           int(10) NOT NULL COMMENT '连接 bfq',
    `exam_session_id`  int(10) NOT NULL COMMENT '连接 exam_session',
    `student_answer_1` text(50) DEFAULT NULL COMMENT '学生的答案',
    `student_answer_2` text(50) DEFAULT NULL COMMENT '学生的答案',
    `student_answer_3` text(50) DEFAULT NULL COMMENT '学生的答案',
    FOREIGN KEY (`bfq_id`) REFERENCES bfq (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (`exam_session_id`) REFERENCES exam_session (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    PRIMARY KEY (`id`),
    UNIQUE INDEX (`bfq_id`, `exam_session_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Table structure for true false question (tfq)
-- ----------------------------
DROP TABLE IF EXISTS `tfq`;
CREATE TABLE `tfq`
(
    `id`                    int(10)     NOT NULL AUTO_INCREMENT COMMENT '用作主键',
    `publisher_teacher_id`  varchar(10) NOT NULL COMMENT '创建本题的教师的工号',
    `stem`                  text(200)   NOT NULL COMMENT '题干',
    `answer`                bool        NOT NULL COMMENT '正确答案',
    `overall_correct_score` int         NOT NULL DEFAULT 0 COMMENT '此题在所有出现中的总得分数*10，即保存到小数点后一位',
    `overall_score`         int         NOT NULL DEFAULT 0 COMMENT '此题在所有出现中的总分数*10，即保存到小数点后一位',
    FOREIGN KEY (`publisher_teacher_id`) REFERENCES teacher (`teacher_id`) ON DELETE CASCADE ON UPDATE CASCADE,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Table structure for students' answer of true false question (tfq)
-- ----------------------------
DROP TABLE IF EXISTS `tfq_answer`;
CREATE TABLE `tfq_answer`
(
    `id`              int(10) NOT NULL AUTO_INCREMENT COMMENT '用作主键',
    `tfq_id`          int(10) NOT NULL COMMENT '连接 tfq',
    `exam_session_id` int(10) NOT NULL COMMENT '连接 exam_session',
    `right_answer`    bool    NOT NULL COMMENT '正确答案，与 tfq 中同名字段保持一致',
    `student_answer`  bool DEFAULT NULL COMMENT '学生的答案',
    FOREIGN KEY (`tfq_id`) REFERENCES tfq (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (`exam_session_id`) REFERENCES exam_session (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    PRIMARY KEY (`id`),
    UNIQUE INDEX (`tfq_id`, `exam_session_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Table structure for code reading question (crq)
-- ----------------------------
DROP TABLE IF EXISTS `crq`;
CREATE TABLE `crq`
(
    `id`                    int(10)     NOT NULL AUTO_INCREMENT COMMENT '用作主键',
    `publisher_teacher_id`  varchar(10) NOT NULL COMMENT '创建本题的教师的工号',
    `stem`                  text(200)   NOT NULL COMMENT '题干',
    `blank_num`             tinyint(2)  NOT NULL COMMENT '要填的空的数目',
    `answer_1`              text(50)    NOT NULL COMMENT '填空的答案',
    `answer_2`              text(50)    NOT NULL COMMENT '填空的答案',
    `answer_3`              text(50)             DEFAULT NULL COMMENT '填空的答案',
    `answer_4`              text(50)             DEFAULT NULL COMMENT '填空的答案',
    `answer_5`              text(50)             DEFAULT NULL COMMENT '填空的答案',
    `answer_6`              text(50)             DEFAULT NULL COMMENT '填空的答案',
    `overall_correct_score` int         NOT NULL DEFAULT 0 COMMENT '此题在所有出现中的总得分数*10，即保存到小数点后一位',
    `overall_score`         int         NOT NULL DEFAULT 0 COMMENT '此题在所有出现中的总分数*10，即保存到小数点后一位',
    FOREIGN KEY (`publisher_teacher_id`) REFERENCES teacher (`teacher_id`) ON DELETE CASCADE ON UPDATE CASCADE,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Table structure for students' answer of code reading question (crq)
-- ----------------------------
DROP TABLE IF EXISTS `crq_answer`;
CREATE TABLE `crq_answer`
(
    `id`               int(10) NOT NULL AUTO_INCREMENT COMMENT '用作主键',
    `crq_id`           int(10) NOT NULL COMMENT '连接 crq',
    `exam_session_id`  int(10) NOT NULL COMMENT '连接 exam_session',
    `student_answer_1` text(50) DEFAULT NULL COMMENT '学生的答案',
    `student_answer_2` text(50) DEFAULT NULL COMMENT '学生的答案',
    `student_answer_3` text(50) DEFAULT NULL COMMENT '学生的答案',
    `student_answer_4` text(50) DEFAULT NULL COMMENT '学生的答案',
    `student_answer_5` text(50) DEFAULT NULL COMMENT '学生的答案',
    `student_answer_6` text(50) DEFAULT NULL COMMENT '学生的答案',
    FOREIGN KEY (`crq_id`) REFERENCES crq (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (`exam_session_id`) REFERENCES exam_session (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    PRIMARY KEY (`id`),
    UNIQUE INDEX (`crq_id`, `exam_session_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Table structure for coding question (cq)
-- ----------------------------
DROP TABLE IF EXISTS `cq`;
CREATE TABLE `cq`
(
    `id`                    int(10)     NOT NULL AUTO_INCREMENT COMMENT '用作主键',
    `publisher_teacher_id`  varchar(10) NOT NULL COMMENT '创建本题的教师的工号',
    `stem`                  text(200)   NOT NULL COMMENT '题干',
    `is_input_from_file`    bool        NOT NULL COMMENT '程序输入是否从文件读取，若为否，从命令行读取',
    `is_output_to_file`     bool        NOT NULL COMMENT '程序输出是否写入文件，若为否，输出到命令行',
    `input`                 text(200)   NOT NULL COMMENT '程序的输入',
    `output`                text(200)   NOT NULL COMMENT '程序的输出',
    `template`              text(200)   NOT NULL COMMENT '题目的初始模板',
    `overall_correct_score` int         NOT NULL DEFAULT 0 COMMENT '此题在所有出现中的总得分数*10，即保存到小数点后一位',
    `overall_score`         int         NOT NULL DEFAULT 0 COMMENT '此题在所有出现中的总分数*10，即保存到小数点后一位',
    FOREIGN KEY (`publisher_teacher_id`) REFERENCES teacher (`teacher_id`) ON DELETE CASCADE ON UPDATE CASCADE,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

DROP TABLE IF EXISTS `cq_answer`;
CREATE TABLE `cq_answer`
(
    `id`              int(10) NOT NULL AUTO_INCREMENT COMMENT '用作主键',
    `cq_id`           int(10) NOT NULL COMMENT '连接 cq',
    `exam_session_id` int(10) NOT NULL COMMENT '连接 exam_session',
    `student_answer`  text(200) DEFAULT NULL COMMENT '学生的答案，即代码',
    `is_answer_right` bool      DEFAULT FALSE COMMENT '学生的代码是否正确',
    FOREIGN KEY (`cq_id`) REFERENCES cq (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (`exam_session_id`) REFERENCES exam_session (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    PRIMARY KEY (`id`),
    UNIQUE INDEX (`cq_id`, `exam_session_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;