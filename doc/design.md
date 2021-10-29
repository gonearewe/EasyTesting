# 软件设计说明书

## 概述

EasyTesting 是一个自动化的在线考试平台，主要针对 Python 课程测试。 其允许教师出题，发布考试，考试结束后查看学生成绩统计信息； 允许学生在考试发布后登录作答，限时提交答案。系统主要由三部分组成：
考生客户端、教师客户端、服务端软件。两个客户端软件分别为两种角色服务， 各自通过网络连接服务端软件，而服务端软件主要提供数据库相关功能（角色管理、题目管理、成绩管理等）。

![system_composition](./img/system_composition.svg)

## 技术选型

- 数据库：MySQL
- 服务端软件：Go、[Gin](https://gin-gonic.com/)
- 学生客户端：Python、PyQt5、[qtmodern](https://github.com/gmarull/qtmodern)
- 教师客户端：Python、PyQt5、[qtmodern](https://github.com/gmarull/qtmodern)

## 教师客户端设计

教师客户端为教师服务。教师角色中又有一种特殊的超级管理员角色，他除了具有基本的教师权限外， 还可以增删改查教师用户本身。

客户端启动后会进入登录界面，需要输入服务器 IP 及端口号和用户名（工号）及其密码。 如果信息正确，即可进入客户端软件。之后，用户可以在主界面中看到一些功能菜单项：“试题管理”、 “考试管理”、”个人信息“与”用户管理“（超级管理员专有）。

### 试题管理

用户可以在这里操作数据库中的试题。他可以新建试题，在软件支持的 5 种题目类型 （选择题、填空题、判断对错题、读程序题、写程序题）中选择一种进入相应的试题创建向导， 并按提示完成题干。题干编辑是在一个内嵌的富文本编辑器中进行的，支持
Markdown 格式。 试题完成后会被提交到数据库成为题库的一部分，学生参加考试时得到的试题就是从中随机抽取的。 他也可以查看已有的试题，并对其进行修改或删除。

### 考试管理

教师在“考试管理”中创建一个新的考试是学生能参加考试的前提。 教师用户需要根据提示选择考试开始时间、结束时间、答题时间、试卷题量等。

> 答题时间不等于结束时间减开始时间，而是应当小于它。
>
> 因为学生参加考试的时间是弹性的，他可以选择在开始时间与结束时间间的任意时刻登录。
> 每个学生的答题时间都是从各自的登录作答时刻开始计算的。但是结束时间之后服务器将不再接受答卷的提交。
> 教师应当引导学生在考试开始后尽快登录作答。

![exam_timeline](./img/exam_timeline.svg)

考试发布后系统会生成一个考试 ID，用户应将其告知考生，使用不同的 ID 将进入同时进行的不同考试， 使用无效的 ID 将无法进入考试。 用户还能查看处于不同状态（还未开始、正在进行、已经结束）的所有考试，
删除还未开始与已经结束的考试，修改还未开始的考试信息。对于已经结束的考试， 用户可以查看某次考试中所有人的成绩以及对应的统计信息（平均分、不及格率等）， 数据以可视化方式给出，且支持以 Excel 等格式导出。

### 个人信息

用户可以查看和修改自己的密码、称呼等，但不能修改用户名，也不能删除自己的账户。

### 用户管理

这是超级管理员专有的权限，他可以在这里增加新的超级管理员与新的教师用户，重置其他普通教师的密码。 可以增加与删除考生，通过 Excel 等格式批量导入、导出所有考生信息（学号、姓名、班级等）。 只有预先记录进系统的考生才能参加后续的考试。

## 学生客户端设计

学生用户是参与考试的主体。他需要在学生客户端登录界面输入正确的服务器 IP 及端口号、考试 ID、 用户名（学号），并确认个人信息、阅读注意事项、勾选诚信保证，才能进入考试答题主界面。
主界面包括两个部分：工具栏与试题。通过切换工具栏中的标签页，用户可以进入要作答的不同题型。 同一题型包含多道小题，用户可以自由决定答题顺序。工具栏还会显示已答题数、未答题数、剩余答题时间等。
客户端会每隔一段时间就自动保存当前作答情况至服务器，用户也可通过“保存”按钮手动保存。 而通过“提交”按钮，用户可以保存作答情况并结束本次考试。即使用户不手动提交，
剩余答题时间归零时软件也会自动提交。提交成功后用户即可看到系统自动计算的考试成绩。

系统支持 5 种题目类型：选择题、填空题、判断对错题、读程序题、写程序题。 选择题又分为单项选择与多项选择，普通的填空题每题一个空，读程序题是给出一段代码和多个空的特殊的填空题， 写程序题要求根据题意完成整段代码，支持类似 OJ
的“给定输入，测试输出”功能。

![student_client_demo](./img/student_client_demo.png)

## 服务端与数据库设计

为了支持所需的服务，EasyTesting 系统设计了如下的数据库表结构。

> Working In Progress

```mysql
USE easy_testing;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for student
-- ----------------------------
DROP TABLE IF EXISTS `student`;
CREATE TABLE `student`
(
    `id`         int(10)             NOT NULL AUTO_INCREMENT COMMENT '用作主键',
    `student_id` smallint(10) UNIQUE NOT NULL COMMENT '学号',
    `name`       varchar(50)         NOT NULL COMMENT '姓名',
    `class_id`   smallint(10)        NOT NULL COMMENT '班号',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Table structure for teacher
-- ----------------------------
DROP TABLE IF EXISTS `teacher`;
CREATE TABLE `teacher`
(
    `id`         int(10)             NOT NULL AUTO_INCREMENT COMMENT '用作主键',
    `teacher_id` smallint(10) UNIQUE NOT NULL COMMENT '工号',
    `name`       varchar(50)         NOT NULL COMMENT '姓名',
    `password`   varchar(200)        NOT NULL COMMENT '加盐后的密码',
    `is_admin`   bit(1)              NOT NULL DEFAULT 0 COMMENT '是否为超级管理员，0：否，1：是',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Table structure for exam
-- ----------------------------
DROP TABLE IF EXISTS `exam`;
CREATE TABLE `exam`
(
    `id`                   int(10)      NOT NULL AUTO_INCREMENT COMMENT '用作主键',
    `publisher_teacher_id` smallint(10) NOT NULL COMMENT '发布考试的教师的工号',
    `start_time`           datetime     NOT NULL COMMENT '考试开始时间',
    `end_time`             datetime     NOT NULL COMMENT '考试结束时间',
    `time_allowed`         varchar(200) NOT NULL COMMENT '考生答题时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Table structure for exam session
-- ----------------------------
DROP TABLE IF EXISTS `exam_session`;
CREATE TABLE `exam_session`
(
    `id`         int(10)      NOT NULL AUTO_INCREMENT COMMENT '用作主键',
    `exam_id`    int(10)      NOT NULL COMMENT '连接 exam',
    `student_id` smallint(10) NOT NULL COMMENT '连接 student',
    `start_time` datetime     NOT NULL COMMENT '作答开始时间',
    `end_time`   datetime         DEFAULT NULL COMMENT '交卷时间',
    `score`      tinyint unsigned DEFAULT NULL COMMENT '最终成绩',
    FOREIGN KEY (`exam_id`) REFERENCES exam (`id`),
    FOREIGN KEY (`student_id`) REFERENCES student (`student_id`),
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Table structure for multiple choice question (mcq), including multiple answer question (maq)
-- ----------------------------

DROP TABLE IF EXISTS `mcq`;
CREATE TABLE `mcq`
(
    `id`                   int(10)      NOT NULL AUTO_INCREMENT COMMENT '用作主键',
    `publisher_teacher_id` smallint(10) NOT NULL COMMENT '创建本题的教师的工号',
    `stem`                 text(20)     NOT NULL COMMENT '题干',
    `choice_1`             text(20)     NOT NULL COMMENT '选项的内容',
    `choice_2`             text(20)     NOT NULL COMMENT '选项的内容',
    `choice_3`             text(20)     NOT NULL COMMENT '选项的内容',
    `choice_4`             text(20)     NOT NULL COMMENT '选项的内容',
    `choice_5`             text(20)              DEFAULT NULL COMMENT '选项的内容',
    `choice_6`             text(20)              DEFAULT NULL COMMENT '选项的内容',
    `choice_7`             text(20)              DEFAULT NULL COMMENT '选项的内容',
    `is_maq`               bit(1)       NOT NULL DEFAULT 0 COMMENT '是不是多选题，0：否，1：是',
    `right_answer`         char(7)               DEFAULT NULL COMMENT '答案，按升序包含所有正确选项的索引，如 "5"、"124"、"67"',
    FOREIGN KEY (`publisher_teacher_id`) REFERENCES teacher (`teacher_id`),
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Table structure for students' answer of multiple choice question (mcq)
-- ----------------------------

DROP TABLE IF EXISTS `mcq_answer`;
CREATE TABLE `mcq_answer`
(
    `id`              int(10)          NOT NULL AUTO_INCREMENT COMMENT '用作主键',
    `mcq_id`          int(10)          NOT NULL COMMENT '连接 mcq',
    `exam_session_id` int(10)          NOT NULL COMMENT '连接 exam_session',
    `right_answer`    char(7) DEFAULT NULL COMMENT '正确答案，与 mcq 中同名字段保持一致',
    `student_answer`  char(7) DEFAULT NULL COMMENT '学生的答案',
    `score`           tinyint unsigned NOT NULL COMMENT '本题分值',
    FOREIGN KEY (`mcq_id`) REFERENCES mcq (`id`),
    FOREIGN KEY (`exam_session_id`) REFERENCES exam_session (`id`),
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Table structure for blank filling question (bfq), including code reading question (crq)
-- ----------------------------

DROP TABLE IF EXISTS `bfq`;
CREATE TABLE `bfq`
(
    `id`                   int(10)      NOT NULL AUTO_INCREMENT COMMENT '用作主键',
    `publisher_teacher_id` smallint(10) NOT NULL COMMENT '创建本题的教师的工号',
    `stem`                 text(20)     NOT NULL COMMENT '题干',
    `blank_num`            tinyint(1)   NOT NULL COMMENT '要填的空的数目，若大于 1，则说明是 crq',
    `answer_1`             text(20)     NOT NULL COMMENT '填空的答案',
    `answer_2`             text(20) DEFAULT NULL COMMENT '填空的答案',
    `answer_3`             text(20) DEFAULT NULL COMMENT '填空的答案',
    `answer_4`             text(20) DEFAULT NULL COMMENT '填空的答案',
    `answer_5`             text(20) DEFAULT NULL COMMENT '填空的答案',
    FOREIGN KEY (`publisher_teacher_id`) REFERENCES teacher (`teacher_id`),
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Table structure for students' answer of blank filling question (bfq)
-- ----------------------------

DROP TABLE IF EXISTS `bfq_answer`;
CREATE TABLE `bfq_answer`
(
    `id`               int(10)          NOT NULL AUTO_INCREMENT COMMENT '用作主键',
    `bfq_id`           int(10)          NOT NULL COMMENT '连接 bfq',
    `exam_session_id`  int(10)          NOT NULL COMMENT '连接 exam_session',
    `student_answer_1` text(20)         NOT NULL COMMENT '学生的答案',
    `student_answer_2` text(20) DEFAULT NULL COMMENT '学生的答案',
    `student_answer_3` text(20) DEFAULT NULL COMMENT '学生的答案',
    `student_answer_4` text(20) DEFAULT NULL COMMENT '学生的答案',
    `student_answer_5` text(20) DEFAULT NULL COMMENT '学生的答案',
    `score`            tinyint unsigned NOT NULL COMMENT '本题分值',
    FOREIGN KEY (`bfq_id`) REFERENCES bfq (`id`),
    FOREIGN KEY (`exam_session_id`) REFERENCES exam_session (`id`),
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- 写代码题待定

SET FOREIGN_KEY_CHECKS = 1;
```

![database_er](./img/database_er.png)

服务端通过网络响应两种客户端的请求。 服务器会对登录的用户进行鉴权以决定其能访问的接口。 试题管理、用户管理等增删改查操作不作详解，这里仅展示核心的考试流程。

![server_exam_sequence](./img/server_exam_sequence.svg)
