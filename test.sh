#!/bin/sh

cat $1 | { 
  while read line ;
  do
    if echo ${line} | grep "#" > /dev/null; then
      :
    elif echo ${line} | grep interface > /dev/null; then
      interface="$( echo ${line} | awk -F ':' '{print $2}' )"
    elif echo ${line} | grep id > /dev/null; then
      id="$( echo ${line} | awk -F ':' '{print $2}' )"
    fi
  done

  if [ -n "$interface" -a -n "$id" ]; then
      echo $interface
      echo $id
    else
      echo "YAML Error"
  fi

}
