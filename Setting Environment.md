# To set environment variables 
in each platform and run your CLI tool using those variables, you can follow these platform-specific instructions. Here's how you can do it:

## 1. macOS/Linux
Step 1: Set Environment Variables Temporarily
You can set environment variables temporarily for a single command by prepending the command with the environment variables.

```bash
MY_VAR="some_value" OTHER_VAR="another_value" ./build/gohexa-mac
```

Step 2: Set Environment Variables Persistently
To set environment variables persistently, you can add them to your shell profile (e.g., `~/.bashrc`, `~/.bash_profile`, `~/.zshrc` for Zsh users):

```bash
alias gohexa="~/path/to/build/gohexa-mac"
```

After adding the variables, source the profile to apply the changes:
```bash
source ~/.bashrc
```

Now, you can run your CLI tool:
```bash
gohexa -name new_project
```

## 2. Windows
Step 1: Set Environment Variables Temporarily
To set environment variables temporarily for a single command in Command Prompt:
```bash
set MY_VAR=some_value && set OTHER_VAR=another_value && build\gohexa.exe
```
In PowerShell:
```bash
$env:MY_VAR="some_value"; $env:OTHER_VAR="another_value"; .\build\gohexa.exe
```

Step 2: Set Environment Variables Persistently
To set environment variables persistently in Windows:

1. Open the Start menu, search for "Environment Variables," and select "Edit the system environment variables."
2. In the System Properties window, click "Environment Variables."
Under "User variables" or "System variables," click "New" and add your variable name and value.
3. Once set, these environment variables will be available to all command-line sessions, and you can run your CLI tool:

```bash
build\gohexa.exe
```

## 3. Using Environment Variables in Go
In your Go code, you can access these environment variables using the os.Getenv function:
```
package main

import (
	"fmt"
	"os"
)

func main() {
	myVar := os.Getenv("MY_VAR")
	otherVar := os.Getenv("OTHER_VAR")

	fmt.Println("MY_VAR:", myVar)
	fmt.Println("OTHER_VAR:", otherVar)

	// Your CLI logic here
}
```
## 4. Running the CLI with Environment Variables
After setting the environment variables, you can run the CLI on each platform, and it will pick up those variables.

macOS/Linux:
```bash
./build/gohexa-mac
```

Windows:
```bash
build\gohexa.exe
```

## Conclusion
By setting environment variables either temporarily or persistently on each platform, you can control the runtime environment of your CLI tool. The CLI can then access these variables using Go’s `os.Getenv` function, allowing you to configure the tool's behavior based on the environment.








