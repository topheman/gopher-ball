#!/bin/bash

# from https://gist.github.com/drkibitz/8896013

list=""

otool_list() {
	local lib="$1"
	local libs=""
	if [[ "$list" == "${list/:$lib/}" ]]; then
		echo "$lib"
		list="$list:$lib"
		libs=$(otool -L "$lib" | sed -n 's/^	\(.*\) (compatibility version.*$/\1/p')
		for lib in $libs
		do
			otool_list "$lib"
		done
	fi
}

otool_list "$@"