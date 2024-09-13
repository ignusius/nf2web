#!/bin/sh
nfcapd  -E  -p 2025 -w /var/cache/nfdump/l1 -t 5   > /dev/null 2>&1 &
nfcapd  -E  -p 2026 -w /var/cache/nfdump/g1 -t 5    > /dev/null 2>&1 &
nfcapd  -E  -p 2027 -w /var/cache/nfdump/m1 -t 5    > /dev/null 2>&1 &
cd /opt/nf2web && ./nf2web
