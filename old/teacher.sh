#!/bin/bash

if [ -d mystery ]; then
  cd mystery || exit 1
fi

export INTERVIEW=$(sed -n '179p' streets/Buckingham_Place | grep -oE '[0-9]+')
printf '%s\n' "$INTERVIEW"
cat "interviews/interview-$INTERVIEW"
printf '%s\n' "$MAIN_SUSPECT"
