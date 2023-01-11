## 分布式锁redis版本实现

### 实现思路

* 利用`set nx ex`获取锁，并设置过期时间，保存协程标识
* 释放锁时先判断线程标识是否与自己一致，一致则删除锁
    * 特性：
        * 利用`set nx`满足互斥性
        * 利用`set ex`保证故障时锁依然能释放，避免死锁，提高安全性
        * 利用Redis集群保证高可用和高并发特性

### 后续想法

本实现仅仅只是实现了一个redis的分布式锁，但是对于重入问题、不可重试问题、主从一致性问题等。
后续精力和水平达到之后再做考虑