# Step 1: Build for Windows
If you are on a Windows machine, you can build the CLI with:

go build -o gohexa.exe
This will generate a gohexa.exe binary for Windows.

# Step 2: Cross-Compile for Multiple Platforms
Go makes it easy to cross-compile binaries for different operating systems and architectures. You can use the GOOS and GOARCH environment variables to specify the target platform.

Here's how to build your CLI for the most common platforms:

macOS (Intel/AMD64)
```bash
GOOS=darwin GOARCH=amd64 go build -o gohexa-mac
```
macOS (Apple Silicon/ARM64)
```bash
GOOS=darwin GOARCH=arm64 go build -o gohexa-mac-arm
```

Linux (AMD64)
```bash
GOOS=linux GOARCH=amd64 go build -o gohexa-linux
```

Windows (AMD64)
```
GOOS=windows GOARCH=amd64 go build -o gohexa.exe
```
Windows (32-bit)
```bash
GOOS=windows GOARCH=386 go build -o gohexa-32.exe
```

# Step 3: Automate Cross-Platform Builds with a Script
You can create a shell script (or batch script on Windows) to automate the build process for all platforms.

Example Shell Script (build.sh)
```bash
#!/bin/bash

# Create the build directory if it doesn't exist
mkdir -p build

# Build for macOS (Intel/AMD64)
GOOS=darwin GOARCH=amd64 go build -o build/gohexa-mac

# Build for macOS (Apple Silicon/ARM64)
GOOS=darwin GOARCH=arm64 go build -o build/gohexa-mac-arm

# Build for Linux (AMD64)
GOOS=linux GOARCH=amd64 go build -o build/gohexa-linux

# Build for Windows (AMD64)
GOOS=windows GOARCH=amd64 go build -o build/gohexa.exe

# Build for Windows (32-bit)
GOOS=windows GOARCH=386 go build -o build/gohexa-32.exe

echo "Builds completed for macOS, Linux, and Windows in the build/ directory."
```

# Step 4: Run the Script
Make the script executable and run it:
```bash
chmod +x build.sh
./build.sh
```