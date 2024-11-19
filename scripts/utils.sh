#!/bin/bash

function run() {
  # 使用 ANSI 转义码设置输出颜色
  echo -e "$(log_debug "\e[95m% $* \e[0m")"
  # 执行命令
  "$@"

  # 检查命令是否执行成功
  local status=$?
  if [ $status -ne 0 ]; then
    echo -e "\e[91mcommand failed with $status 。\e[0m" >&2
    exit $status
  fi
}

function prepare_dir() {
  local dir="$1"
  if [ -d "$dir" ]; then
    log_warn "directory $dir exists, deleting all the files under."
    run find "$dir" -mindepth 1 -delete
  else
    log_info "directory $dir does not exist, creating."
    run mkdir -p "$dir"
  fi

}

function log_info() {
  echo -e "[\e[92minfo\e[0m]: $*"
}

function log_error() {
  echo -e "[\e[91merror\e[0m]: $*"
}

function log_warn() {
  echo -e "[\e[93mwarn\e[0m]: $*"
}

function log_debug() {
  echo -e "[\e[97mdebug\e[0m]: $*"
}

function colors() {

  # 输出所有颜色
  for i in {0..100}; do
    echo -e "\e[${i}m @@@<color $i>@@@ \e[0m"
  done
}

