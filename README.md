# modemBandSelector #

## Description ##
Simple tool for Huawei HiLInk modems. It's allow manage
(disable/enable) LTE bands.

For now it support only devices, which web UI accesible
without authorization

## Build ##
First install fyne:
``go install fyne.io/fyne/v2/cmd/fyne``

Then build app:
``fyne build``

On windows before build you must specify environment variables like:

Enable CGO:
``CGO_ENABLE=1``

Define path to your GCC-binary:
``CC=C:\msys64\clangarm64\bin\gcc.exe``

## Windows on Arm64 issues ##
On arm64 devices with windows you need "OpenCL and OpenGL Compatibility pack":

* Store link: [ms-windows-store://pdp/?productid=9NQPSL29BFFF](ms-windows-store://pdp/?productid=9NQPSL29BFFF)
* Web link: [https://apps.microsoft.com/store/detail/9NQPSL29BFFF](https://apps.microsoft.com/store/detail/9NQPSL29BFFF)

