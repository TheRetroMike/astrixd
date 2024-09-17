#!/bin/bash

APPDIR=/tmp/astrixd-temp
ASTRIXD_RPC_PORT=29587

rm -rf "${APPDIR}"

astrixd --simnet --appdir="${APPDIR}" --rpclisten=0.0.0.0:"${ASTRIXD_RPC_PORT}" --profile=6061 &
ASTRIXD_PID=$!

sleep 1

RUN_STABILITY_TESTS=true go test ../ -v -timeout 86400s -- --rpc-address=127.0.0.1:"${ASTRIXD_RPC_PORT}" --profile=7000
TEST_EXIT_CODE=$?

kill $ASTRIXD_PID

wait $ASTRIXD_PID
ASTRIXD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Astrixd exit code: $ASTRIXD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $ASTRIXD_EXIT_CODE -eq 0 ]; then
  echo "mempool-limits test: PASSED"
  exit 0
fi
echo "mempool-limits test: FAILED"
exit 1
