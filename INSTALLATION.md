# Installation

## MacOS

- Download Docker for mac [here](https://docs.docker.com/docker-for-mac/install/)
- Download Go progamming language [here](https://golang.org/doc/install)
- Install kind

```bash
brew install kind
```

## Windows

- Check for build number, need to be higher than Version 2004
- Install WSL2

```bash
dism.exe /online /enable-feature /featurename:Microsoft-Windows-Subsystem-Linux /all /norestart
dism.exe /online /enable-feature /featurename:VirtualMachinePlatform /all /norestart
```

- Restart windows
- Set default version

```bash
wsl --set-default-version 2
```

- Download linux update [here](https://docs.microsoft.com/en-gb/windows/wsl/install-win10#step-4---download-the-linux-kernel-update-package)
- Download Docker desktop [here](https://hub.docker.com/editions/community/docker-ce-desktop-windows/)
- Download Go progamming language [here](https://golang.org/doc/install)
- Download Windows Terminal from Microsoft Store [here](https://aka.ms/terminal)
- Install kind

```bash
curl.exe -Lo kind-windows-amd64.exe https://kind.sigs.k8s.io/dl/v0.11.0/kind-windows-amd64
Move-Item .\kind-windows-amd64.exe c:\some-dir-in-your-PATH\kind.exe

# OR via Chocolatey (https://chocolatey.org/packages/kind)
choco install kind
```
