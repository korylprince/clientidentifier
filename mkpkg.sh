#!/bin/bash
ver="1.2"
./build.sh
sudo cp build/clientidentifier /usr/local/bin/
sudo chown root:wheel /usr/local/bin/clientidentifier
sudo pkggen -i tk.unstac.clientidentifier -v "$ver" files out.pkg
sudo chown administrator:staff out.pkg
mv out.pkg "clientidentifier-$ver.pkg"
