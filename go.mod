module github.com/zjllib/go-micro

go 1.16

require (
	cloud.google.com/go/pubsub v1.24.0
	cuelang.org/go v0.4.3
	github.com/BurntSushi/toml v0.3.1
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/Shopify/sarama v1.22.0
	github.com/VictoriaMetrics/metrics v1.22.2
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/anacrolix/envpprof v1.2.1 // indirect
	github.com/anacrolix/missinggo v1.3.0 // indirect
	github.com/anacrolix/sync v0.4.0 // indirect
	github.com/anacrolix/utp v0.1.0
	github.com/apex/log v1.9.0
	github.com/asim/go-awsxray v0.0.0-20161209120537-0d8a60b6e205
	github.com/aws/aws-sdk-go v1.44.68
	github.com/bitly/go-simplejson v0.5.0
	github.com/bradfitz/gomemcache v0.0.0-20220106215444-fb4bf637b56d
	github.com/caddyserver/certmagic v0.17.2
	github.com/davecgh/go-spew v1.1.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/eclipse/paho.mqtt.golang v1.4.1
	github.com/ef-ds/deque v1.0.4
	github.com/evanphx/json-patch/v5 v5.5.0
	github.com/fsnotify/fsnotify v1.5.4
	github.com/fsouza/go-dockerclient v1.7.3
	github.com/ghodss/yaml v1.0.0
	github.com/gin-gonic/gin v1.8.1
	github.com/go-acme/lego/v4 v4.4.0
	github.com/go-git/go-git/v5 v5.4.2
	github.com/go-log/log v0.2.0
	github.com/go-redis/redis/v8 v8.11.5
	github.com/go-sql-driver/mysql v1.6.0
	github.com/go-stomp/stomp/v3 v3.0.5
	github.com/go-zookeeper/zk v1.0.3
	github.com/gobwas/httphead v0.1.0
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/gobwas/ws v1.0.4
	github.com/golang/glog v1.0.0
	github.com/golang/protobuf v1.5.2
	github.com/gomodule/redigo v1.8.9
	github.com/google/uuid v1.3.0
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/websocket v1.4.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/hashicorp/consul/api v1.13.0
	github.com/hashicorp/go-hclog v1.1.0
	github.com/hashicorp/hcl v1.0.0
	github.com/hashicorp/memberlist v0.3.1
	github.com/hashicorp/vault/api v1.8.0
	github.com/hudl/fargo v1.4.0
	github.com/imdario/mergo v0.3.12
	github.com/juju/ratelimit v1.0.2
	github.com/kr/pretty v0.3.0
	github.com/lib/pq v1.10.7
	github.com/lucas-clemente/quic-go v0.29.1
	github.com/m3o/m3o-go/client v0.0.0-20210421144725-8bfd7992ada3
	github.com/markbates/pkger v0.17.1
	github.com/miekg/dns v1.1.50
	github.com/minio/highwayhash v1.0.2
	github.com/mitchellh/hashstructure v1.1.0
	github.com/nacos-group/nacos-sdk-go/v2 v2.1.0
	github.com/nats-io/nats-streaming-server v0.25.2 // indirect
	github.com/nats-io/nats.go v1.17.0
	github.com/nats-io/stan.go v0.10.3
	github.com/nsqio/go-nsq v1.1.0
	github.com/nxadm/tail v1.4.8
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7
	github.com/opentracing/opentracing-go v1.2.0
	github.com/oxtoacart/bpool v0.0.0-20190530202638-03653db5a59c
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pborman/uuid v1.2.1
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.12.2
	github.com/prometheus/client_model v0.2.0
	github.com/rs/zerolog v1.28.0
	github.com/segmentio/kafka-go v0.4.35
	github.com/sirupsen/logrus v1.8.1
	github.com/sony/gobreaker v0.5.0
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.8.0
	github.com/uber/jaeger-client-go v2.30.0+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible
	github.com/urfave/cli/v2 v2.3.0
	go.etcd.io/bbolt v1.3.6
	go.etcd.io/etcd/api/v3 v3.5.5
	go.etcd.io/etcd/client/v3 v3.5.5
	go.opencensus.io v0.23.0
	go.uber.org/ratelimit v0.0.0-20180316092928-c15da0234277
	go.uber.org/zap v1.23.0
	gocloud.dev v0.27.0
	gocloud.dev/pubsub/rabbitpubsub v0.27.0
	golang.org/x/crypto v0.0.0-20221010152910-d6f0a8c073c2
	golang.org/x/net v0.0.0-20220805013720-a33c5aa5df48
	golang.org/x/oauth2 v0.0.0-20220722155238-128564f6959c
	golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4
	golang.org/x/sys v0.0.0-20221010170243-090e33056c14
	golang.org/x/text v0.3.7
	google.golang.org/api v0.91.0
	google.golang.org/genproto v0.0.0-20220802133213-ce4fa296bf78
	google.golang.org/grpc v1.48.0
	google.golang.org/grpc/examples v0.0.0-20221011233738-e81d0a276fb3
	google.golang.org/protobuf v1.28.1
	gopkg.in/DataDog/dd-trace-go.v1 v1.42.1
	gopkg.in/yaml.v2 v2.4.0
)
