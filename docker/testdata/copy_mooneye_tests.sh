#!/bin/bash

set -x

mkdir -p /roms/mooneye

while read -r test; do
  mkdir -p "/roms/mooneye/$(dirname "$test")"
  cp -f "/opt/mooneye/build/$test.gb" "/roms/mooneye/$test.gb"
  cp -f "/opt/mooneye/$test-expected.png" "/roms/mooneye/$test-expected.png" 2>/dev/null
  cp -f "/opt/mooneye/$test.s" "/roms/mooneye/$test.s" 2>/dev/null
done < /mooneye_tests.txt
