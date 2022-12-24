FROM golang as Build

WORKDIR /App/file-sharing-server

COPY . .

RUN go build

FROM busybox as Final
COPY --from=Build /App/file-sharing-server .
CMD ./file-sharing-server 19090 /file-sharing