module github.com/go-saas/kit

go 1.19

require (
	github.com/alexedwards/argon2id v0.0.0-20211130144151-3585854a6387
	github.com/casbin/casbin/v2 v2.55.1
	github.com/casbin/gorm-adapter/v3 v3.12.1
	github.com/envoyproxy/protoc-gen-validate v0.6.8
	github.com/go-chi/chi/v5 v5.0.7
	github.com/go-kratos/kratos/v2 v2.5.0
	github.com/go-saas/saas v0.5.1
	github.com/go-saas/uow v0.0.5-0.20220710170855-20a530670db0
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/google/uuid v1.3.0
	github.com/gorilla/csrf v1.7.1
	github.com/gorilla/handlers v1.5.1
	github.com/goxiaoy/go-eventbus v0.0.5
	github.com/goxiaoy/gorm-concurrency v1.1.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.11.3
	github.com/mennanov/fmutils v0.2.0
	github.com/mitchellh/mapstructure v1.5.0
	github.com/nbutton23/zxcvbn-go v0.0.0-20210217022336-fa2cb2858354
	github.com/stretchr/testify v1.8.0
	golang.org/x/net v0.0.0-20221002022538-bcab6841153b // indirect
	golang.org/x/sys v0.0.0-20220928140112-f11e5e49a4ec // indirect
	google.golang.org/genproto v0.0.0-20220930163606-c98284e70a91
	google.golang.org/grpc v1.49.0
	google.golang.org/protobuf v1.28.1
	gorm.io/driver/mysql v1.4.1
	gorm.io/driver/sqlite v1.4.3
	gorm.io/gorm v1.24.0
)

require (
	github.com/BurntSushi/toml v1.2.0
	github.com/Shopify/sarama v1.37.0
	github.com/apache/pulsar-client-go v0.8.1
	github.com/aws/aws-sdk-go v1.44.109
	github.com/centrifugal/centrifuge v0.26.0
	github.com/dtm-labs/dtmcli v1.14.2
	github.com/dtm-labs/dtmdriver-kratos v0.0.8
	github.com/dtm-labs/dtmgrpc v1.14.2
	github.com/eko/gocache/v3 v3.1.1
	github.com/fclairamb/afero-s3 v0.3.1
	github.com/flowchartsman/swaggerui v0.0.0-20210303154956-0e71c297862e
	github.com/go-kratos/kratos/contrib/log/zap/v2 v2.0.0-20220927151429-0ecc2b422f28
	github.com/go-kratos/kratos/contrib/registry/consul/v2 v2.0.0-20220927151429-0ecc2b422f28
	github.com/go-kratos/kratos/contrib/registry/etcd/v2 v2.0.0-20220927151429-0ecc2b422f28
	github.com/go-redis/redis/extra/redisotel/v8 v8.11.5
	github.com/go-redis/redis/v8 v8.11.5
	github.com/go-saas/go-i18n/v2 v2.2.1-0.20220719043234-2e8ea1df2661
	github.com/go-saas/lazy v1.0.0
	github.com/go-saas/sessions v1.2.2-0.20220626044315-a7a87c87f120
	github.com/go-sql-driver/mysql v1.6.0
	github.com/goava/di v1.11.1
	github.com/gorilla/mux v1.8.0
	github.com/hashicorp/consul/api v1.15.2
	github.com/hibiken/asynq v0.23.0
	github.com/hibiken/asynqmon v0.7.1
	github.com/lithammer/shortuuid/v3 v3.0.7
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/nyaruka/phonenumbers v1.1.1
	github.com/ory/hydra-client-go v1.11.8
	github.com/pilagod/gorm-cursor-paginator/v2 v2.3.0
	github.com/pkg/errors v0.9.1
	github.com/samber/lo v1.28.2
	github.com/segmentio/ksuid v1.0.4
	github.com/spf13/afero v1.9.2
	github.com/uptrace/opentelemetry-go-extra/otelgorm v0.1.16
	github.com/xhit/go-simple-mail/v2 v2.12.0
	go.etcd.io/etcd/client/v3 v3.5.5
	go.opentelemetry.io/otel v1.10.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.10.0
	go.opentelemetry.io/otel/sdk v1.10.0
	go.opentelemetry.io/otel/trace v1.10.0
	go.uber.org/zap v1.23.0
	golang.org/x/sync v0.0.0-20220929204114-8fcdb60fdcc0
	golang.org/x/text v0.3.7
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4 // indirect
	github.com/99designs/keyring v1.2.1 // indirect
	github.com/AthenZ/athenz v1.11.10 // indirect
	github.com/DataDog/zstd v1.5.2 // indirect
	github.com/FZambia/eagle v0.0.2 // indirect
	github.com/FZambia/sentinel v1.1.1 // indirect
	github.com/Knetic/govaluate v3.0.1-0.20171022003610-9aa49832a739+incompatible // indirect
	github.com/XiaoMi/pegasus-go-client v0.0.0-20210427083443-f3b6b08bc4c2 // indirect
	github.com/apache/pulsar-client-go/oauth2 v0.0.0-20220810085107-c2476193159f // indirect
	github.com/ardielle/ardielle-go v1.5.2 // indirect
	github.com/armon/go-metrics v0.4.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bradfitz/gomemcache v0.0.0-20220106215444-fb4bf637b56d // indirect
	github.com/cenkalti/backoff/v4 v4.1.3 // indirect
	github.com/centrifugal/protocol v0.8.11 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.4.0 // indirect
	github.com/danieljoos/wincred v1.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dtm-labs/dtmdriver v0.0.3 // indirect
	github.com/dvsekhvalnov/jose2go v1.5.0 // indirect
	github.com/eapache/go-resiliency v1.3.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20180814174437-776d5712da21 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/elliotchance/orderedmap/v2 v2.2.0 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/glebarez/go-sqlite v1.19.1 // indirect
	github.com/glebarez/sqlite v1.5.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-playground/form/v4 v4.2.0 // indirect
	github.com/go-redis/redis/extra/rediscmd/v8 v8.11.5 // indirect
	github.com/go-resty/resty/v2 v2.7.0 // indirect
	github.com/go-test/deep v1.0.8 // indirect
	github.com/godbus/dbus v0.0.0-20190726142602-4481cbc300e2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/golang-sql/sqlexp v0.1.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/gomodule/redigo v1.8.9 // indirect
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/gsterjov/go-libsecret v0.0.0-20161001094733-a6f4afe4910c // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.3.1 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/serf v0.10.0 // indirect
	github.com/iancoleman/strcase v0.2.0 // indirect
	github.com/igm/sockjs-go/v3 v3.0.2 // indirect
	github.com/imdario/mergo v0.3.13 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.13.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.1 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.12.0 // indirect
	github.com/jackc/pgx/v4 v4.17.2 // indirect
	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
	github.com/jcmturner/gofork v1.7.6 // indirect
	github.com/jcmturner/gokrb5/v8 v8.4.3 // indirect
	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/klauspost/compress v1.15.11 // indirect
	github.com/lib/pq v1.10.4 // indirect
	github.com/linkedin/goavro/v2 v2.12.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mattn/go-sqlite3 v1.14.15 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2 // indirect
	github.com/microsoft/go-mssqldb v0.17.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mna/redisc v1.3.2 // indirect
	github.com/montanaflynn/stats v0.6.6 // indirect
	github.com/mtibben/percent v0.2.1 // indirect
	github.com/pegasus-kv/thrift v0.13.0 // indirect
	github.com/pierrec/lz4 v2.6.1+incompatible // indirect
	github.com/pierrec/lz4/v4 v4.1.17 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.13.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20220927061507-ef77025ab5aa // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/segmentio/asm v1.2.0 // indirect
	github.com/segmentio/encoding v0.3.5 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/toorop/go-dkim v0.0.0-20201103131630-e1cd1a0a5208 // indirect
	github.com/uptrace/opentelemetry-go-extra/otelsql v0.1.16 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.1 // indirect
	github.com/xdg-go/stringprep v1.0.3 // indirect
	github.com/youmark/pkcs8 v0.0.0-20201027041543-1326539a0a0a // indirect
	go.etcd.io/etcd/api/v3 v3.5.5 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.5 // indirect
	go.mongodb.org/mongo-driver v1.10.2 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.10.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.10.0 // indirect
	go.opentelemetry.io/otel/metric v0.32.1 // indirect
	go.opentelemetry.io/proto/otlp v0.19.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	golang.org/x/crypto v0.0.0-20221005025214-4161e89ecf1b // indirect
	golang.org/x/exp v0.0.0-20221002003631-540bb7301a08 // indirect
	golang.org/x/oauth2 v0.0.0-20220909003341-f21342109be1 // indirect
	golang.org/x/term v0.0.0-20220919170432-7a66f970e087 // indirect
	golang.org/x/time v0.0.0-20220922220347-f3bd1da661af // indirect
	google.golang.org/appengine v1.6.7 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/tomb.v2 v2.0.0-20161208151619-d5d1b5820637 // indirect
	gorm.io/driver/postgres v1.4.4 // indirect
	gorm.io/driver/sqlserver v1.4.1 // indirect
	gorm.io/plugin/dbresolver v1.3.0 // indirect
	k8s.io/apimachinery v0.25.2 // indirect
	modernc.org/libc v1.20.0 // indirect
	modernc.org/mathutil v1.5.0 // indirect
	modernc.org/memory v1.4.0 // indirect
	modernc.org/sqlite v1.19.1 // indirect
)

replace (
	github.com/99designs/keyring => github.com/99designs/keyring v1.2.1
	github.com/hibiken/asynqmon => github.com/Goxiaoy/asynqmon v0.7.2-0.20220620161318-80231130441b
)
