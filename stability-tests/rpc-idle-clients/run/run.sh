#!/bin/bash
rm -rf /tmp/astrixd-temp

NUM_CLIENTS=128
astrixd --devnet --appdir=/tmp/astrixd-temp --profile=6061 --rpcmaxwebsockets=$NUM_CLIENTS &
ASTRIXD_PID=$!
ASTRIXD_KILLED=0
function killAstrixdIfNotKilled() {
  if [ $ASTRIXD_KILLED -eq 0 ]; then
    kill $ASTRIXD_PID
  fi
}
trap "killAstrixdIfNotKilled" EXIT

sleep 1

rpc-idle-clients --devnet --profile=7000 -n=$NUM_CLIENTS
TEST_EXIT_CODE=$?

kill $ASTRIXD_PID

wait $ASTRIXD_PID
ASTRIXD_EXIT_CODE=$?
ASTRIXD_KILLED=1

echo "Exit code: $TEST_EXIT_CODE"
echo "Astrixd exit code: $ASTRIXD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $ASTRIXD_EXIT_CODE -eq 0 ]; then
  echo "rpc-idle-clients test: PASSED"
  exit 0
fi
echo "rpc-idle-clients test: FAILED"
exit 1
