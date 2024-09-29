# Random Password Generator

A simple golang script to generate random passwords with specified lengths and quantities. It utilizes uppercase letters, lowercase letters, digits, and symbols to create secure passwords.

## Prerequisites

- **Golang**

##Installation

### For Linux or macOS:

1. **Open your terminal.**

2. **Download the Go tarball:**
   Use `curl` to download Go. Adjust the version as needed.

   ```bash
   curl -O https://dl.google.com/go/go1.20.5.linux-amd64.tar.gz
   ```

3. **Extract the tarball:**

   ```bash
   sudo tar -C /usr/local -xzf go1.20.5.linux-amd64.tar.gz
   ```

4. **Set up your environment variables:**
   Add Go to your `PATH` by adding the following line to your `~/.bash_profile`, `~/.bashrc`, or `~/.zshrc`:

   ```bash
   export PATH=$PATH:/usr/local/go/bin
   ```

   Then, reload your profile:

   ```bash
   source ~/.bash_profile
   ```

   or

   ```bash
   source ~/.bashrc
   ```

5. **Verify the installation:**

   ```bash
   go version
   ```

### For Windows

1. **Open Command Prompt or PowerShell.**

2. **Download the Go installer:**
   Use the following command (adjust the version if necessary):

   ```powershell
   Invoke-WebRequest -Uri https://dl.google.com/go/go1.20.5.windows-amd64.msi -OutFile goinstaller.msi
   ```

3. **Run the installer:**

   ```powershell
   Start-Process msiexec.exe -ArgumentList '/i goinstaller.msi' -NoNewWindow -Wait
   ```

4. **Add Go to your `PATH`:**
   This may be done automatically by the installer, but if not, add `C:\Go\bin` to your `PATH` manually through System Properties.

5. **Verify the installation:**

   ```powershell
   go version
   ```

## Usage

1. Clone the repository:
   ```bash
   git clone https://github.com/d1v45/Password-Generator.git
   ```

2. Navigate to the project directory:
   ```bash
   cd Password-Generator
   ```

3. Run the script:
   ```golang
   go run pass_gen.go
   ```

4. Follow the prompts to specify the length of passwords and the number of passwords you need.

## Example

Generating a password of length 12 and getting 3 passwords:
```
Enter the length of the password (minimum 6): 12
Enter the number of passwords you need: 5
Include uppercase letters? (y/n): y
Include lowercase letters? (y/n): y
Include numbers? (y/n): y
Include symbols? (y/n): y
Would you like to enter a custom character set? (y/n): n

;*Jsfr3^U_E5
lo;XI6MpG^AL
g[%w(T<cSy=~
F%-${LO2>3bg
W>1v&wK1;HZ{
