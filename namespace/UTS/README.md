# UTS
1. 运行go程序进入交互程序
```shell
go run uts.go
```
2. 打印进程树，可以看到sh进程的父进程是uts，并且有个子进程pstree
```shell
pstree -pl
```
3. 查看当前的PID
```shell
echo $$
```
4. 查看当前进程的父子进程
```shell
pstree -ps $$
```
5. 验证父子进程是否在同一个uts namespace中
```shell
readlink /proc/父进程PID/ns/uts
readlink /proc/子进程PID/ns/uts
```
6. 在子进程shell中修改hostname后去父进程shell中查看是否生效

