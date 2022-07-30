# 查看当前系统的namespace
lsns

lsns -t net

# 查看某进程的namespace
ls -la /proc/pid/ns/

# 进入namespace运行命令
nsenter -t pid -n ip addr

unshare -fn sleep 60

ps -ef | grep sleep

lsns -t net

nsenter -t pid -n ip a