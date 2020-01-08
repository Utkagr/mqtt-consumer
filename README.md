# mqtt-consumer

go build -buildmode=plugin -o stdout-sink/connector.so stdout-sink/connector.go
go build -o mqtt-consumer .
./mqtt-consumer -t topic1
