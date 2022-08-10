#!/bin/bash

default_wgpm_bin=$(command -v wgpm)
default_host=0.0.0.0
default_port=32090
default_device=wg0

RED="$(tput setaf 1 2>/dev/null || echo '')"
RESET="$(tput sgr0 2>/dev/null || echo '')"

ask() {
    if [ -z "$NO_INTERACTIVE" ]; then
        >&2 printf "$@ "

        local ans

        if [ -n "$ASK_SILENT" ]; then
            read -s ans </dev/tty
        else
            read ans </dev/tty
        fi

        if [ $? -ne 0 ]; then
            fatal "Cannot read from stdin"
        fi

        if [ -n "$ans" ]; then
            echo "$ans"
        elif [ -n "$2" ]; then
            echo "$2"
        fi
    fi
}

fatal() {
    echo -e "${RED}✘" "$@${RESET}" >&2
    exit 1
}

tpl() {
    local vars=""
    local script="{"

    for var_name in ${@:2}; do
        local tpl_var_name=$(echo "$var_name" | awk -F '=' {'print $1'})
        local script_var_name=$(echo "$var_name" | awk -F '=' {'print $2'})

        if [ -z "$script_var_name" ]; then
            script_var_name="$tpl_var_name"
        fi

        script="$script gsub(\"{{ $tpl_var_name }}\",$tpl_var_name);"
        vars="$vars -v $tpl_var_name=\"${!script_var_name}\""
    done

    script="$script }1"

    eval awk "$vars" "'$script'" "$1"
}

if [ -z "$default_wgpm_bin" ]; then
    ask_wgpm_bin_default="Not found"
else
    ask_wgpm_bin_default="$default_wgpm_bin"
fi

wgpm_bin=$(ask "wgpm binary path (default: $ask_wgpm_bin_default): " "$default_wgpm_bin")

if [ -z "$wgpm_bin" ]; then
    fatal "No bianry for wgpm"
elif [ ! -e "$wgpm_bin" ]; then
    fatal "$wgpm_bin is not an executable binary"
fi

host=$(ask "Listen host (default: $default_host): " "$default_host")
port=$(ask "Listen port (default: $default_port): " "$default_port")
device=$(ask "Device (default: $default_device): " "$default_device")
bearer_token_auth=$(ASK_SILENT=true ask "Bearer token auth (default: none)")

echo '##'

tpl wgpm.service wgpm_bin host port device | sudo dd status=none of=/etc/systemd/system/wgpm.service
tpl wgpm.service.env bearer_token_auth | sudo dd status=none of=/etc/systemd/system/wgpm.service.env
sudo systemctl daemon-reload
sudo systemctl enable wgpm.service
sudo systemctl start wgpm.service
