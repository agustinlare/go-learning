# OCP - Golang soy vago

## opc

Login basico. Da a elegir en que cluster y solicita password.  

```=bash
λ ocp
Ingrese opcion:
N | NAME        | URL
---------------------------------------------
0 | ocp         | https://api.ocp.com:6443
1 | ocpqa1      | https://api.ocpqa1.com:6443
2 | ocpqa2      | https://api.ocpqa2.com:6443
3 | ocpp1       | https://api.ocpp1.com:6443
4 | ocpp2       | https://api.ocpp2.com:6443
5 | ocpd1       | https://api.ocpd1.com:6443
6 | ocpd2       | https://api.ocpd2.com:6443
---------------------------------------------
2
Password:
Login successful.

You have access to 241 projects, the list has been suppressed. You can list all projects with 'oc projects'

Using project "default".
λ
```

## Build in windows

```=powershell
rsrc -manifest login.manifest -ico ocp.ico -o rsrc.syso
go get github.com/akavel/rsrc
go build
Move-Item .\ocp.exe C:\tools\ -Force
```

## Build in linux

```=bash
mv rsrc.syso asd
go get gopkg.in/yaml.v2
go build
mv ocp /usr/local/bin
mv asd rsrc.syso
```

## Copy occonfig file for linux

```=bash
cp /mnt/c/Users/A309788/.kube/occonfig ~/.kube/
```
