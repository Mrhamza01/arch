---

### ✅ `RELEASE.md`

```markdown
# Arch CLI Release Guide

This document explains how to build and package the Arch CLI tool.

---

## 1. Build the CLI

Run the following command to build the Windows executable:

```sh
go build -o dist/build/arch.exe main.go
This will produce:

bash
Copy code
dist/build/arch.exe
2. Build Installer (Inno Setup)
Install Inno Setup → Download here.

Run the compiler with:

sh
Copy code
"C:\Program Files (x86)\Inno Setup 6\ISCC.exe" arch-installer.iss
This will generate:

bash
Copy code
dist/installer/arch-installer.exe
3. Test Installation
Run dist/installer/arch-installer.exe.

Open a new terminal and type:

sh
Copy code
arch --help
✅ You should see the CLI help globally without moving the .exe.

4. Release
Upload dist/installer/arch-installer.exe as the release artifact.

Optionally, keep dist/build/arch.exe as a raw binary for advanced users.

Notes
arch.exe is installed into C:\Program Files\Arch\bin.

Installer automatically updates PATH environment variable.

yaml
Copy code

---

