 to build plugins
 go build -buildmode=plugin -o addition/addition.so addition/Addition.go
 go build -buildmode=plugin -o division/division.so Division/Division.go      
 go build -buildmode=plugin -o multiply/multiply.so Multiply/Multiply.go      
 go build -buildmode=plugin -o subtract/subtract.so Subtract/Subtract.go 

 go run main.go addition 5 2 
 go run main.go multiply 5 2

few points-

Plugin is build in separate module and using file path to get plugin 
To make it truly dynamically so that plugin can be added at run time - i have done some research i know it is possible using  - grpc , hashicorp etc but seems like it will take more effort and time 


