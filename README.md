# Arch CLI

Arch is a CLI tool to scaffold **feature-based Clean Architecture** folder structures for .NET, Blazor Hybrid, and Web projects.

## 📦 Installation

1. Download the installer (`arch-installer.exe`) from [Releases](./dist/installer).
2. Run the installer — this will:
   - Install `arch.exe` into `C:\Program Files\Arch\bin`
   - Add `Arch/bin` to your system PATH

After installation, you can run `arch` globally from any terminal.

## 🚀 Usage

### Create a new feature
```sh
arch feature Todo
Creates the following structure in your Ecom.Shared/Features folder:

markdown
Copy code
Features/
 └── Todo/
     ├── Data/
     │   ├── Datasources/
     │   └── Repositories/
     ├── Domain/
     │   ├── Entities/
     │   ├── Repositories/
     │   └── UseCases/
     └── Presentation/
         └── Pages/
Create a feature in a specific path
sh
Copy code
arch feature Todo --path ../MyApp.Shared/Features