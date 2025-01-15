# 使用已有的基础镜像
FROM docker-cf.registry.cyou/yedf/dtm

# 设置工作目录（假设基础镜像中已经存在 /app/dtm 目录）
WORKDIR /app/dtm

# 设置时区环境变量
ARG TZ=UTC
ENV TZ=${TZ}

# 复制配置文件到镜像中
COPY ./config.yml /app/dtm/configs/config.yaml

# 配置 entrypoint
ENTRYPOINT ["/app/dtm/dtm", "-c=/app/dtm/configs/config.yaml"]

# 暴露端口
EXPOSE 36789
EXPOSE 36790
