module gonotes

go 1.13

require (
	github.com/robfig/cron/v3 v3.0.1
	github.com/sirupsen/logrus v1.4.2
	github.com/smallnest/rpcx v0.0.0-20200310110228-122cece1047a
	golang.org/x/crypto v0.0.0-20200302210943-78000ba7a073 // indirect
)

replace google.golang.org/grpc v1.27.0 => google.golang.org/grpc v1.26.0
