module github.com/partitio/protoc-gen-gorm

require (
	github.com/dgrijalva/jwt-go v0.0.0-20180921172315-3af4c746e1c2 // indirect
	github.com/envoyproxy/protoc-gen-validate v0.1.0
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/gogo/protobuf v1.0.0
	github.com/golang/protobuf v1.3.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.4.1 // indirect
	github.com/jinzhu/gorm v1.9.1
	github.com/jinzhu/inflection v0.0.0-20180308033659-04140366298a
	github.com/lib/pq v1.0.0
	github.com/partitio/atlas-app-toolkit v0.0.0-20190328190718-c1fb12986e25
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.3.0 // indirect
	golang.org/x/crypto v0.0.0-20190211182817-74369b46fc67 // indirect
	golang.org/x/net v0.0.0-20190213061140-3a22650c66bd // indirect
	golang.org/x/sys v0.0.0-20190213121743-983097b1a8a3 // indirect
	google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8
	google.golang.org/grpc v1.18.0
	gopkg.in/yaml.v2 v2.2.1 // indirect
)

go 1.13

replace github.com/gogo/protobuf => github.com/partitio/gogo-protobuf v1.0.1-0.20210521094738-1cd42496b840
