### 概述

个人小工具。包括：

+ 时间戳转换；
+ 常用格式转化；

**不包括公司相关的工具**

### API

### time

`one time [-f <format>] <val1> <val2> <val3>`：生成时间对应的时间戳
+ `-f <format>`：指定时间格式。

+ 支持批量数据处理。
+ 默认支持多种格式。
+ 如果 val 为空/now，那么返回当前时间戳；
+ 如果 val 为 today，那么返回今天零点的时间戳

`one timestamp last <days>`：返回最近几天的时间戳


### timestamp

`one timestamp [-ns|-s] <timestamp>`：生成时间戳对应的时间

+ 支持批量数据处理。

`one base36 <number>`
`one base36 -d/--decode <number>`

### 路由
