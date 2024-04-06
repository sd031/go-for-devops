Step 1: Create Your Go Project
Start by setting up a new Go project:

mkdir project -name //e.g. kubectl-myplugin and follow naming convention
cd kubectl-myplugin
go mod init myplugin

if existing project, then install dependencies from go mod file:
go mod download


Step 2: Write Your Plugin Code (refer to the code right side)
Create a file named main.go and add the following code, which defines a simple CLI application using the Cobra library to interact with the Kubernetes API =>

Step 3: Build Your Plugin:
Compile your plugin into an executable. The executable name must start with kubectl- for kubectl to recognize it as a plugin:
go build -o kubectl-myplugin

Step 4: Make Your Plugin Discoverable
Move your plugin executable to a location in your PATH:
mv kubectl-myplugin /usr/local/bin


Step 5: Use Your Plugin
Now, you can run your plugin as if it were a native kubectl command:
kubectl myplugin --namespace my-namespace



