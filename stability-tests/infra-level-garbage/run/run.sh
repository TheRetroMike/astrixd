#!/bin/bash
rm -rf /tmp/astrixd-temp

astrixd --devnet --appdir=/tmp/astrixd-temp --profile=6061 &
ASTRIXD_PID=$!

sleep 1

infra-level-garbage --devnet -alocalhost:16611 -m messages.dat --profile=7000
TEST_EXIT_CODE=$?

kill $ASTRIXD_PID

wait $ASTRIXD_PID
ASTRIXD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Astrixd exit code: $ASTRIXD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $ASTRIXD_EXIT_CODE -eq 0 ]; then
  echo "infra-level-garbage test: PASSED"
  exit 0
fi
echo "infra-level-garbage test: FAILED"
exit 1
