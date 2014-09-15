#!/bin/bash
./build.sh
sudo cp build/clientidentifier /usr/bin/
sudo chown root:wheel /usr/bin/clientidentifier
sudo pkggen -i tk.unstac.clientidentifier -v "1.1" files out.pkg
sudo chown administrator:staff out.pkg
mv out.pkg "clientidentifier.pkg"
