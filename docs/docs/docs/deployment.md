

如果服务器出现

> Error 1292: Incorrect datetime value: '0000-00-00' for column 'end_time'

说明是 MySQL 禁用了 datetime 类型的零值。
可以通过在 MySQL console 中输入 `SELECT @@GLOBAL.sql_mode;` 查看
全局的 sql_mode，它可能包括这些内容：

```
ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION
```

我们只需要删除其中的 NO_ZERO_IN_DATE 与 NO_ZERO_DATE，通过重新全局设置 sql_mode：

> SET GLOBAL sql_mode = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

参考 Stackoverflow 上的问题 [MySQL Incorrect datetime value: '0000-00-00 00:00:00'](https://stackoverflow.com/questions/35565128/mysql-incorrect-datetime-value-0000-00-00-000000)