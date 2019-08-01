FROM python:2.7

WORKDIR /usr/src/app

# RUN sed -i 's=deb.debian.org=mirrors.ustc.edu.cn=g' /etc/apt/sources.list && \
#     sed -i 's=security.debian.org=mirrors.ustc.edu.cn=g' /etc/apt/sources.list

RUN apt update && apt install -y vim &&  apt install -y python-pip && apt clean && apt auto-clean                                      
RUN python -m pip install --upgrade pip && \
    python -m pip install --no-cache-dir grpcio  &&  \
    python -m pip install --no-cache-dir grpcio-tools  

COPY . .
EXPOSE 50052
CMD [ "python", "./server.py" ]