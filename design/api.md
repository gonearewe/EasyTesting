接下来介绍系统的前后端接口。
接口相当于一道切面，可以明确地划分开前后端的职责，
本系统的接口通过 [OpenAPI](https://www.openapis.org/) v3.0.0 规范描述，详细的描述文件见[这里](https://github.com/gonearewe/EasyTesting/tree/gh-pages/design/easy_testing.yaml)，可以通过 [Swagger Editor](https://editor.swagger.io/) 查看。这里仅仅对接口作一个概要的介绍。

> #### info::OpenAPI 描述文件编辑器
> 
> OpenAPI 的描述文件可以用 yaml 或 json 格式，
> 但无论如何，直接对着文本文件进行编辑效率都很低下。
> 我选用的带 UI 的编辑器是来自 [Stoplight](https://stoplight.io/) 的。

## 特殊 API

两个特殊的 API 是 `GET ping` 与 `GET hello`，主要用于测试。
`ping` 始终返回文本 pong，可以测试服务端是否在正常工作。
而 `hello` 仅支持已认证的用户访问，会返回用户 id（学生的学号或教师的工号），可以测试登录状态。

## 用户登录 API

`GET teacher_auth` 接口用于教师的登录，提供参数 `teacher_id` 与 `password`。如果教师工号存在且密码正确，服务器将完成认证并返回一个 JSON Web Token。关于 JWT 的详细介绍可参考[这里](https://jwt.io/introduction)，简单来说，它包含三部分：Header、Payload 与 Signature。
Payload 可携带明文信息供客户端使用，教师的 Token Payload 中包含：

- id: 教师的系统唯一 ID
- teacher_id: 工号
- name：教师的姓名
- is_admin：教师是否为管理员

为了保持登录状态（即让服务器相信“我就是XX”），
客户端必须把 TOKEN 放进之后的接口请求中，
具体为首部字段 Authorization 中，以 Bearer 方式，
即 `Authorization: Bearer <token_here>`。
输入 Header、Payload 与一个密钥，利用指定的算法即可计算出 Signature，
因为密钥仅供服务器访问，所以计算仅有服务器能完成，
服务器就根据请求携带的 TOKEN 是否能通过验算进行认证。

`GET student_auth` 则是学生的登录接口，接受参数 `student_id`、
`name` 与 `exam_id`。学生登录特殊的地方在于不仅要学号与姓名正确，
还要求指定的考试正在进行（即当前时间处于该考试的*开始时刻*与*结束时刻*间），
并且如果学生还未参加这场考试，则自动参加（即答题倒计时自接口访问成功开始计算），
当然如果学生已经参加了，就只会返回一个 TOKEN（这种情况一般不会出现，因为 TOKEN 的
有效期都设置得很长，基本不会在考试期间过期）。
学生登录成功获得的 Token 的 Payload 中包含：

- student_id：学号
- name：学生的姓名
- class_id：学生的班号
- exam_session_id：由当前学生 ID 和考试 ID 共同决定的系统唯一 ID 
- exam_deadline：学生在这个考试中的作答截止时刻

除开 `GET ping` 与这两个登录认证 API，
其他的接口都要求 JWT 认证。并且教师与学生的可访问接口也是完全分开的
（唯一一个两用户均可访问的认证接口是 `GET hello`）。

## 教师客户端的其他 API

`mcq`、`maq`、`bfq`、`tfq`、`crq` 与 `cq` 这六个接口都支持
GET、POST、PUT、DELETE 四大方法，分别用于单选题（Multiple Choice Question）、多选题（Multiple Answer Question）、填空题（Blank-Filling Question）、
判断题（True or False Question）、代码阅读题（Code Reading Question）与编程题（Coding Question）的管理（增删改查）。

`teachers` 与 `students` 接口同样支持四大方法，它们分别用于
教师用户和学生用户的增删改查。但是前者仅有管理员可以访问，
服务器通过 TOKEN Payload 中的 is_admin 字段认证管理员；
因为服务器读取 is_admin 字段之前仍会验证 Signature，
所以不必担心恶意请求篡改 is_admin 字段的问题。
`PUT profile` 同样是用来修改教师的接口，但它仅限于修改自己的信息，
并且不支持修改*是否为管理员*属性。

`exams` 接口同样支持四大方法，用于支持对考试的增删改查。
`GET exams/ended` 获得所有已结束考试的列表，因为它不返回考试的详细信息且列表一般不会很长，
所以不支持分页。`GET exams/examinees` 接受参数 `exam_id`，
返回指定考试的所有 exam session，同样不支持分页。

## 学生客户端的其他 API

教师客户端本质上是管理平台，所以相关的接口较多，而学生客户端的 API 就要少很多。

`GET exams/my_questions` 用于获取学生的试题。`PUT exams/my_answers` 则用于提交作答。

> WIP

## API 的生命周期

我们希望在有考试进行时禁用对题目的修改，也希望学生仅能在自己的作答时间内提交答案。
显然，这要求 API 有一定的生命周期（lifecycle），各种 API 的生命周期如下所示：

![lifecycle](../img/api-lifecycle.drawio.svg)

当然，这是由生命周期决定的可访问性，如果接口参数有误或不满足其他条件，
即使在生命周期内也会访问失败。
