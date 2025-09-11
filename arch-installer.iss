; arch-installer.iss
; Inno Setup script for Arch CLI tool

[Setup]
AppName=Arch CLI
AppVersion=1.0.0
DefaultDirName={commonpf}\Arch
DefaultGroupName=Arch
UninstallDisplayIcon={app}\bin\arch.exe
UninstallDisplayName=Arch CLI
OutputDir=dist\installer
OutputBaseFilename=arch-installer
Compression=lzma
SolidCompression=yes

[Files]
Source: "dist\build\arch.exe"; DestDir: "{app}\bin"; Flags: ignoreversion

[Icons]
Name: "{group}\Arch CLI"; Filename: "{app}\bin\arch.exe"

[Run]
; Add Arch/bin to PATH permanently
Filename: "{cmd}"; Parameters: "/C setx PATH ""%PATH%;{app}\bin"""; Flags: runhidden
