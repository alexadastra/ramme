module github.com/alexadastra/ramme

go 1.17

replace (
	github.com/alexadastra/auth-service/pkg/api => ../auth-service/pkg/api
	github.com/alexadastra/ramme => ../ramme
)

require (
	github.com/gorilla/mux v1.8.0
	github.com/pkg/errors v0.9.1
	github.com/rs/xlog v0.0.0-20171227185259-131980fab91b
	github.com/sirupsen/logrus v1.8.1
	gopkg.in/fsnotify.v1 v1.4.7
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/fsnotify/fsnotify v1.5.2 // indirect
	github.com/justinas/alice v1.2.0 // indirect
	github.com/rs/cors v1.8.2 // indirect
	github.com/rs/xhandler v0.0.0-20170707052532-1eb70cf1520d // indirect
	github.com/rs/xid v1.4.0 // indirect
	golang.org/x/net v0.0.0-20220412020605-290c469a71a5 // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
)
