# serverMonitor

## 仓库介绍

其实最开始来说，本仓库的目的就是写一个服务器监控系统，用来监控我自己跑的一些web服务，检测一下这些服务是否在线。

后来，由于一些开源库的影响，比如gin-vue-admin，就试图通过这个demo练习下全栈开发，但是前端代码好难，进度几乎为0...

现在，我把这个项目当作一个练习demo，把我目前可能感兴趣的内容都做到这个项目里，比如一些开源库，一些微服务机制，grpc等等


## 项目进展

以下是一些已经实现的功能：
1. 日志封装，配置统一读取
2. 日志服务通过grpc实现（可能效率不如同一个进程，主要是为了练习）
3. 对web服务的增删改查，以及状态的监控

下面是一些想实现（练习）的功能（不定时更新）：
1. 微服务的注册，发现，以及管理模块
2. 增加用户管理
3. 把所有的模块都做成微服务（后续考虑）
4. Makefile 不断完善
5. 增加微服务监控模块
6. 增加容器资产管理模块