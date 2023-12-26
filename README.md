# 背景

> 最近在完成青训营项目的简易版抖音项目考虑到使用了二级缓存，看到Github上也存在许多的库，但是并不很符合项目的需求，需要对其进行二次开发。但是由于考虑到二次开发还不如自己自研一个。至此machine
> cache就出来了。

# 技术架构

## 技术设想

- 支持设置过期时间，删除是一种惰性删除的策略。（参考redis的key过期策略，后续也会支持定时删除定期删除）
- 支持缓存回写，遇到不存在的缓存时进行缓存回调，之后回写到本地缓存。
- singleflight机制防止缓存击穿。

