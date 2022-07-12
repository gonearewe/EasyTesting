首先介绍系统的后台数据库，
这有助于我们了解系统的数据模型（Model）。

创建数据库的 sql 文件可在源代码中找到（参见 [introduction](./introduction.md) 中介绍的文件结构）。

*student* 与 *teacher* 是为两种用户分别准备的表（Table）。
它们各有几个复合索引（Composite Index），这用于加速教师客户端的列表查询。
例如，我们可以确保下列查询全都走索引。

```
SELECT * FROM `student` WHERE student_id LIKE '20%' AND class_id LIKE '10%' AND name LIKE '%小%'
SELECT * FROM `student` WHERE student_id LIKE '20%' AND class_id LIKE '10%'
SELECT * FROM `student` WHERE student_id LIKE '20%'
SELECT * FROM `student` WHERE class_id LIKE '10%' AND name LIKE '%小%'
SELECT * FROM `student` WHERE class_id LIKE '10%'
SELECT * FROM `student` WHERE name LIKE '%小%'
```

*teacher* 的 *password* 与 *salt* 字段用于支持密码登录，
详情将于 [server 文档](./server.md) 中介绍。

*exam* 与 *mcq*、*maq*、*bfq*、*tfq*、*crq*、*cq* 分别代表考试与各种题型。
对应于每种题型都有记录考生作答的表：*mcq_answer*、*maq_answer*、*bfq_answer*、*tfq_answer*、*crq_answer* 和 *cq_answer*，
它们都通过各自的外键连接不同 id 的题目，
其中的一些（如 *tfq_answer*）还记录有对应 id 题目的正确答案以加速答案比对过程（避免了 join 表），而 *cq_answer* 还有专门的触发器来比对答案。
因为一个考生可能参加过多场考试，所以 *answer* 系列的表不能仅靠 *student_id* 区分考生。
我们引入了 *exam_session*（即考试会话） 的概念，它由 *exam_id* 与 *student_id* 联合确定，并用于连接 *answer* 系列的表，
该表中的每一条记录都表示某学生参加了某考试。*exam_session* 中的 *time_allowed* 与对应 *exam* 中的同名字段一致，为冗余字段；
但其中的 *start_time* 表示的是对应学生进入考试的时刻，*end_time* 表示最后一次提交答卷的时刻，与对应 *exam* 中的同名字段意义均不同。
另外，*exam_session* 还有 *score* 字段，用于保存考生在某考试中的最终得分。

为了加速 sql 查询，整个数据库中使用了不少冗余字段。
其中的外键（如 *publisher_teacher_id*）都创建了 DELETE 与 UPDATE 的触发器，由数据库自动维护数据一致性。
而 *answer* 系列的表的冗余字段（正确答案）与 *exam_session* 表中的冗余字段 *time_allowed* 根据业务逻辑不会被修改，无需顾及一致性：
*answer* 系列的表有存档（archive）性质，始装保存出题时刻的正确答案，事后修改的题目自然不会影响历史作答；
*exam_session* 中的一条记录被创建就意味着对应考试已开始，而开始后的考试禁用修改。
唯一需要服务器软件手动维护一致性的情景是，在修改 *student* 表时同步更新 *exam_session* 表中的冗余字段 *student_name*。

另外，所有的表中都有 *created_at* 和 *last_updated_at* 两个字段，
分别记录行的创建时刻和上次修改的时刻，而且由触发器自动维护，
不需要在后续的 sql 语句中人为关注。不过，这两个字段目前还没有被服务端实际使用。


