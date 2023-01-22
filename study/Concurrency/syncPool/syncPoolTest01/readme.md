# sync.Pool的工作原理
sync.Pool有两个containers来存储对象，分别是：local pool和victim cache

根据sync.pool.go的源码中可以看出，这个package的init函数会注册一个PoolCleanUp函数，而这个函数就是通过GC触发的

```go
func init() {
   runtime_registerPoolCleanup(poolCleanup)
}
```
当GC的触发的时候，在victim中的对象就会被收集回收，而在local pool中的对象会被移动victim cache当中；下面是poolCleanUp的代码：

```
func poolCleanup() {
   // Drop victim caches from all pools.
   for _, p := range oldPools {
      p.victim = nil
      p.victimSize = 0
   }

   // Move primary cache to victim cache.
   for _, p := range allPools {
      p.victim = p.local
      p.victimSize = p.localSize
      p.local = nil
      p.localSize = 0
   }

   oldPools, allPools = allPools, nil
}
```
新对象是放在local pool当中的，调用pool.Put也是将对象放在local pool当中的。

调用pool.Get时，会先从victim cache中获取，如果没有找到，则就会从local pool中获取，如果local pool中也没有，就会执行初始化时的New Functionle，否则就返回nil了

sync-pool.gif
从Go1.12之后添加了mutex来保证线程安全；从Go1.13在sync.pool中引入了双向链表，移除了mutex，改善了共享操作

# 结论
总之，如果有一个对象需要频繁的创建，并且有很大的开销，那么你就可以使用sync.pool
