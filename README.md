# Arch CLI

Arch is a lightweight CLI tool to scaffold **feature-based Clean Architecture** folder structures.
It can be used with any project type — .NET, Blazor, Web, Flutter, Node.js, or even plain folders.

---

## 📦 Installation

1. Download the installer (`arch-installer.exe`) from the [Releases](./dist/installer).
2. Run the installer — this will place `arch.exe` into `C:\Program Files\Arch\bin`.

After installation, you can run `arch` globally from any terminal.

---

## 🚀 Usage

Arch provides two commands:

### 1. Create a feature in the current directory

```sh
arch Todo
```

➡ Creates a `Todo` feature folder with the standard Clean Architecture structure in your **current working directory**.

### 2. Create a feature in a specific path

```sh
arch Todo ../MyApp/Features
```

➡ Creates the same `Todo` feature folder inside the **path you provide**.

---

## 🗂 Folder Structure

Every feature generated will follow this structure:

```
<FeatureName>/
 ├── Data/
 │   ├── Datasources/
 │   └── Repositories/
     └── DTOs/
     └── Mappings/
 ├── Domain/
 │   ├── Entities/
 │   ├── Repositories/
 │   └── UseCases/
 └── Presentation/
     └── Pages/
```

---

## 🔧 Examples

Create an **Auth** feature in the current directory:

```sh
arch Auth
```

Create a **Product** feature inside a custom path:

```sh
arch Product ./src/Features
```

---

## 🧑 Author

Developed by **Muhammad Hamza**
