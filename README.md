# grpc-GO


In phase-01 created two service for shop and products. (Unary state)


 ____________________
|                    |
|  User->Shop        |
|  User->Product     |
|____________________|

1) Client can view shop products from product service
2) Client can view shops and all shops from shop service

## Generate proto by running: 
 - `protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/user_service.proto`
 - `protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/shop_service.proto`
 - `protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/product_service.proto` 

### User-Service Env
 - USER_SERVICE_PORT=8090
 - DB_HOST=host.docker.internal 

### Shop-Service Env
 - SHOP_SERVICE_PORT=8083
 - DB_HOST=host.docker.internal 


### Product-Service Env
 - PORDUCT_SERVICE_PORT=8082
 - PRODUCT_PRICE=1200
 - DB_HOST=host.docker.internal 