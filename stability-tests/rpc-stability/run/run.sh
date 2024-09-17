#!/bin/bash
rm -rf /tmp/astrixd-temp

astrixd --devnet --appdir=/tmp/astrixd-temp --profile=6061 --loglevel=debug &
ASTRIXD_PID=$!

sleep 1

rpc-stability --devnet -p commands.json --profile=7000
TEST_EXIT_CODE=$?

kill $ASTRIXD_PID

wait $ASTRIXD_PID
ASTRIXD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Astrixd exit code: $ASTRIXD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $ASTRIXD_EXIT_CODE -eq 0 ]; then
  echo "rpc-stability test: PASSED"
  exit 0
fi
echo "rpc-stability test: FAILED"
exit 1
