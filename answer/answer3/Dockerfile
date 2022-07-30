
#设置镜像
FROM golang AS golang

#配置模块代理
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /workspace

#安装git和项目
RUN apt-get install -y git
RUN git clone https://github.com/huanmin/cncamp.git


#编译打包
WORKDIR /workspace/cncamp
RUN go mod tidy
WORKDIR /workspace/cncamp/answer/answer2
RUN go build -o httpServer

#设置空的镜像
FROM scratch

#copy编译好的可执行文件
COPY --from=golang /workspace/cncamp/answer/answer2/httpServer /cncamp

ENTRYPOINT ["/cncamp/httpServer"]