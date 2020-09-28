module service-1

go 1.14

require (
	github.com/golang/protobuf v1.4.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.1
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/jinzhu/gorm v1.9.16
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/kr/pretty v0.1.0 // indirect
	github.com/labstack/echo-contrib v0.9.0
	github.com/labstack/echo/v4 v4.1.16
	github.com/opentracing-contrib/go-grpc v0.0.0-20200813121455-4a6760c71486
	github.com/opentracing/opentracing-go v1.1.0
	github.com/smacker/opentracing-gorm v0.0.0-20181207094635-cd4974441042
	github.com/soheilhy/cmux v0.1.4
	github.com/uber/jaeger-client-go v2.19.1-0.20191002155754-0be28c34dabf+incompatible
	google.golang.org/grpc v1.31.0
	google.golang.org/grpc/examples v0.0.0-20200814200710-a3740e5ed326 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
)

replace github.com/labstack/echo-contrib v0.9.0 => github.com/phanletrunghieu/echo-contrib v0.9.2

replace github.com/smacker/opentracing-gorm v0.0.0-20181207094635-cd4974441042 => github.com/phanletrunghieu/opentracing-gorm v0.0.1
