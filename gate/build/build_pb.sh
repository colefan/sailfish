protoc -I ../pb/ gate_client_msg.proto --go_out=../gatemsg
protoc -I ../pb/ gate_inner_msg.proto --go_out=../gatemsg
