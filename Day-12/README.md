# Steps to build plugin
```bash
go mod tidy  # Ensure all dependencies are fetched
go build -o terraform-provider-example
```

# Move the provider to proper path
1. For macOS (Intel):
On macOS with an Intel or AMD processor (x86 architecture), place the provider binary in the following path:
```bash
mkdir -p ~/.terraform.d/plugins/example.com/example/example/0.1.0/darwin_amd64/
mv terraform-provider-example ~/.terraform.d/plugins/example.com/example/example/0.1.0/darwin_amd64/
```

2. For macOS (Apple Silicon - M1/M2):
On macOS with an Apple Silicon (ARM64) processor, the architecture folder should be darwin_arm64.
```bash
mkdir -p ~/.terraform.d/plugins/example.com/example/example/0.1.0/darwin_arm64/
mv terraform-provider-example ~/.terraform.d/plugins/example.com/example/example/0.1.0/darwin_arm64/
```

3. For Linux (x86_64):
On Linux, if you're using the x86_64 architecture (Intel/AMD), place the provider binary in the following path:
```bash
mkdir -p ~/.terraform.d/plugins/example.com/example/example/0.1.0/linux_amd64/
mv terraform-provider-example ~/.terraform.d/plugins/example.com/example/example/0.1.0/linux_amd64/
```

4. For Linux (ARM64):
For Linux on ARM64 architecture, such as a Raspberry Pi or other ARM-based systems, use the following path:
```bash
mkdir -p ~/.terraform.d/plugins/example.com/example/example/0.1.0/linux_arm64/
mv terraform-provider-example ~/.terraform.d/plugins/example.com/example/example/0.1.0/linux_arm64/
```

5. For Windows (x86_64):
On Windows, use the following path and rename the binary to terraform-provider-example.exe:
```bash
mkdir -p %APPDATA%\terraform.d\plugins\example.com\example\example\0.1.0\windows_amd64\
move terraform-provider-example.exe %APPDATA%\terraform.d\plugins\example.com\example\example\0.1.0\windows_amd64\
```

6. For Windows (ARM64):
On ARM64-based Windows machines (e.g., Surface Pro X), use the following path and also rename the binary:
```bash
mkdir -p %APPDATA%\terraform.d\plugins\example.com\example\example\0.1.0\windows_arm64\
move terraform-provider-example.exe %APPDATA%\terraform.d\plugins\example.com\example\example\0.1.0\windows_arm64\
```


# Test the provider via terraform 
To to go-for-devops/Day-12/terraform_provider/example_use folder 
Run command :
```bash
terraform init 
terraform plan 
terraform apply 
```