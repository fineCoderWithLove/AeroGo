# AeroGo
## 目前处于个人学习阶段
Aero 和 GO
一个快速、轻量级和高性能的Http框架。
## 框架开发目标
### 1.多缓存
* **FIFO/LFU/LRU** 多缓存保证各方面的需求
### 2.快速开发
* 代码生成
* 参数绑定
### 3.自动集成
* 一套框架即可完成单体项目的开发，内置Aorm，松耦合可拆卸
### 4.可观测
* 缓存命中率
* 平均接口耗时
* 慢sql排查日志
## 框架开发步骤
1. 封装了路由信息
2. 封装了Context的内容
3. 使用Trie树来实现路由分组和路由匹配
4. 控制Group,将路由分组重新封装
   * 为了进行统一的鉴权
   * 为了更容易管理路由
5. 中间件的扩展
6. 静态资源服务器
7. recovery中间件，保证程序不强行panic导致宕机
