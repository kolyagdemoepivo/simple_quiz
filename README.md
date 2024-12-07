# simple_quiz


To run the simple_quiz client and server, you need to install protoc, protoc-gen-go, and protoc-gen-go-grpc.
protoc for ubuntu: "apt install -y protobuf-compiler"
protoc for macos: "brew install protobuf"
protoc-gen-go use command "go install google.golang.org/protobuf/cmd/protoc-gen-go@latest"
protoc-gen-go-grpc use command "go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest"

After that, execute the build.sh script to build the client, server and generate the proto structures.

If the build is successful, you should see two binary files: server and client, located in their respective server and client directories.

Start the server by running ./server, and then start the client by running ./client.

Once the client is running, you will need to enter your username and follow the further instructions in the client interface.
