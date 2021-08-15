#!/bin/bash

if [ ! -e gedcom.js ] || [ $(stat -c %s gedcom.js) -ne 36 ]; then
	echo "export const data = [], users = [];" > data.js;
fi;

(
	echo "package main";
	echo;
	echo "// File automatically generated with ./embed.sh";
	echo;
	echo "const (";
	echo "	htmlStart = \"<html><head><title>HCL Timeline</title><link rel=\\\"shortcut icon\\\" sizes=\\\"any\\\" href=\\\"data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 256 256'%3E%3Crect x='15' y='15' height='45' width='175' rx='5' fill='%23f00' /%3E%3Crect x='15' y='75' height='45' width='90' rx='5' fill='%230f0' /%3E%3Crect x='160' y='75' height='45' width='85' rx='5' fill='%230f0' /%3E%3Crect x='70' y='135' height='45' width='125' rx='5' fill='%2300f' /%3E%3Crect x='50' y='195' height='45' width='70' rx='5' fill='%23000' /%3E%3Crect x='170' y='195' height='45' width='50' rx='5' fill='%23000' /%3E%3C/svg%3E\\\" /><script type=\\\"module\\\">\"";
	echo -n "	htmlEnd   = \"</script><style type=\\\"text/css\\\">";
	uglifycss style.css | tr -d "\n"; 
	echo "</style></head><body></body></html>\"";

	echo -n "	jsStart   = \"";

	jslib -i /script.js -n | tail -n+2 | sed -e 's/pageLoad/(document.readyState == "complete" ? Promise.resolve() : new Promise(successFn => globalThis.addEventListener("load", successFn, {once: true})))/' | terser -m  --module --compress pure_getters,passes=3 --ecma 6 | tr -d '\n' | sed -e 's/\\/\\\\/g' -e 's/"/\\\"/g' -e 's/\(.*=\[\)\(\],[^=]*=\[\)/\1\"'"\n"'	jsMid     = \"\2\"'"\n"'	jsEnd     = \"/';

	echo "\"";

	echo ")";
) > parts.go