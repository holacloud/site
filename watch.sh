#!/bin/bash

inotifywait -r -m src/ -e close_write --format "%T" --timefmt "%S" | while read events; do
  ./bin/sitegen --src src/ --www statics/www/ --versions "v1,v2" --languages "en,es,zh"
done
