# feature
* calico初始化
* k8sclient
* 网络策略初始化、打label
* 开启/关闭网络策略
* 添加网段黑名单
* 删除网络黑名单
* 添加特权ns
* 删除特权ns（内部ns不允许删除）


以上接口实现幂等，允许重复操作
以上接口暂未实现回滚

# TODO
* 添加ns网络策略
* 按名字和命名空间删除网络策略