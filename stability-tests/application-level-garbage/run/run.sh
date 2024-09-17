#!/bin/bash
rm -rf /tmp/astrixd-temp

astrixd --devnet --appdir=/tmp/astrixd-temp --profile=6061 --loglevel=debug &
ASTRIXD_PID=$!
ASTRIXD_KILLED=0
function killAstrixdIfNotKilled() {
    if [ $ASTRIXD_KILLED -eq 0 ]; then
      kill $ASTRIXD_PID
    fi
}
trap "killAstrixdIfNotKilled" EXIT

sleep 1

application-level-garbage --devnet -alocalhost:16611 -b blocks.dat --profile=7000
TEST_EXIT_CODE=$?

kill $ASTRIXD_PID

wait $ASTRIXD_PID
ASTRIXD_KILLED=1
ASTRIXD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Astrixd exit code: $ASTRIXD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $ASTRIXD_EXIT_CODE -eq 0 ]; then
  echo "application-level-garbage test: PASSED"
  exit 0
fi
echo "application-level-garbage test: FAILED"
exit 1
