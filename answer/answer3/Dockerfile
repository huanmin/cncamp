
#设置镜像
FROM golang AS golang

#配置模块代理
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

#创建代码目录
RUN mkdir /root/workspace
#安装git和项目
RUN apt-get install -y git
RUN go install github.com/huanmin/learn_go@v1.0.0

WORKDIR /root/workspace/pkg/mod/github.com/huanmin/learn_go@v1.0.0/answer/answer2

#编译打包
RUN GOOS=linux GOARCH=amd64 go build -o httpServer

#设置空的镜像
FROM scratch

#暴露服务端口
EXPOSE 8080

#copy编译好的可执行文件
COPY --from=golang /root/workspace/pkg/mod/github.com/huanmin/learn_go@v1.0.0/answer/answer2/httpServer /bin/project

ENTRYPOINT ["/bin/project/httpServer"]