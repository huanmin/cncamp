cd /sys/fs/cgroup

# 如果没有cpu目录
mkdir cpu

# cpu.shares 可出让的能获得cpu使用时间的相对值
# cpu.cfs_period_us 用来配置时间周期长度（微秒）
# CPU.cfs_quota_us  用来配置Cgroup在cfs_period_us时间最多使用的cpu时间数（微秒）
# CPU.stat  cgroup内的进程使用的cpu时间统计
# nr_periods 经过cpu.cfs_period_us的时间周期数量
# nr_throttled 在经过的周期内，有多少次因为进程在指定时间周期内用光配额而收到限制
# throttled_time cgroup中的进程被限制使用cpu的总用时，单位是ns（纳秒）

mkdir cpudemo

cd cpudemo

top
# 将进程与cgroup关联
echo 30939 > cgroup.procs

cat cpu.shares
cat cpu.cfs_period_us
cat cpu.cfs_quota_us

echo 100000 > cpu.cfs_quota_us # 1个cpu
echo 50000 > cpu.cfs_quota_us # 50% cpu
echo 10000 > cpu.cfs_quota_us # 1/10 cpu


# memory.usage_in_bytes cgroup下进程使用的内存，包含cgroup和其子group下的进程内存
# memory.max_usage_in_bytes cgroup下进程使用内存的最大值，包含子cgroup的内存使用量
# memory.limit_in_bytes 设置cgroup下进程最多能使用的内存，如果为-1，则表示不做限制
# memory.oom_control 设置是否在cgroup中使用oom，默认使用