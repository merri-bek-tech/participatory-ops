module paropd

go 1.21.0

replace parops.libs/msg => ../libs/msg

require (
	github.com/BurntSushi/toml v1.4.0
	github.com/Ullaakut/nmap/v3 v3.0.3
	github.com/eclipse/paho.mqtt.golang v1.5.0
	github.com/google/uuid v1.6.0
	parops.libs/msg v0.0.0-00010101000000-000000000000
)

require (
	github.com/gorilla/websocket v1.5.3 // indirect
	golang.org/x/net v0.27.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
)
