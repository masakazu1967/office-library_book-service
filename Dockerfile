FROM alpine:3.12
LABEL maintainer="Masakazu Kobayashi"

ARG OS=linux
ARG ARCH=amd64

RUN adduser user
WORKDIR /home/user
COPY book_service-${OS}-${ARCH} /home/user
RUN chmod +x book_service-${OS}-${ARCH}
USER user

CMD [ "./book_service-${OS}-${ARCH}" ]
