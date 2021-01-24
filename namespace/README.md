# namespace

## 什么是namespace

容器的核心功能，就是通过约束和修改进程的动态表现，从而为其创造一个边界。而namespace技术则是用来修改进程视图的主要方法，其本质其实只是Linux创建新进程的一个可选参数：当在创建容器进程时，指定这个进程所需要的一组namespace参数，这样该容器就只能看到当前namespace所限定的资源/文件/设备/状态，而对于宿主机其他不相关的程序则完全看不到。

## 相关namespace
| namespace类型 | 系统调用参数   | 内核版本 | 
| mount         | CLONE_NEWNS   | 2.4.19  |
| uts           | CLONE_NEWUTS  | 2.6.19  |
| ipc           | CLONE_NEWIPC  | 2.6.19  |
| pid           | CLONE_NEWPID  | 2.6.24  |
| network       | CLONE_NEWNET  | 2.6.29  |
| user          | CLONE_NEWUSER | 3.8     |

## 相关系统调用
* clone() 创建新进程
* unshare() 将进程在namespace中删除
* setns() 将进程加入到namespace中
