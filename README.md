# gRPC sample  
  
This repository is published for sharing my knowledge.  
  
# Usage  
  
1. Clone to your local.  
`git clone git@github.com:y-harashima/grpc-sample.git`  
  
2. Change current directory to project root.  
`cd grpc-sample`  
  
3. Launch container with docker-compose command.  
`docker-compose up -d --build`  

# Testing gRPC

You can test with gRPC-cli tool.  
[grpcurl](https://github.com/fullstorydev/grpcurl), [evans](https://github.com/ktr0731/evans)  
  
e.g.) (Testcase - Success call)
`grpcurl -plaintext -d '{"id": 1}' localhost:50051 sample.Student.Get`

(Testcase - Failure call)
`grpcurl -plaintext -d '{"id": 2}' localhost:50051 sample.Student.Get`

You can also check "grpc-zap" logging in docker log.
`docker logs -f 

# Author  

y-harashima