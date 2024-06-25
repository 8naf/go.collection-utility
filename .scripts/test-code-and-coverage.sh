source ./.scripts/_common.sh

printf "${YELLOW}---------- TESTING ----------${RESET}\n"

TEST_COVERAGE_FILE_NAME=coverage.out
go test -race -vet=off -v ./... -covermode=atomic -coverpkg=./... -coverprofile=$TEST_COVERAGE_FILE_NAME

TEST_COVERAGE_THRESHOLD=100
TEST_COVERAGE_ACTUAL=$(go tool cover -func=$TEST_COVERAGE_FILE_NAME | grep total | grep -Eo '[0-9]+\.[0-9]+')

is_less_than_threshold=$(echo "$TEST_COVERAGE_ACTUAL < $TEST_COVERAGE_THRESHOLD" | bc -l)

actual_coverage_color=$GREEN
if [ $is_less_than_threshold -eq 1 ]; then
	actual_coverage_color=$RED
fi

printf "${YELLOW}---------- TEST COVERAGE CHECKING ----------${RESET}\n"
printf "Test Coverage (Threshold): ${PURPLE}%10s${RESET}\n" "$TEST_COVERAGE_THRESHOLD%"
printf "Test Coverage (Actual)   : ${actual_coverage_color}%10s${RESET}\n" "$TEST_COVERAGE_ACTUAL%"
echo "----------------------------------------"

if [ $is_less_than_threshold -eq 1 ]; then
	printf "${CYAN}Please add more unit tests to increase the test coverage.${RESET}\n";
	printf "${RED}FAILED${RESET}\n";
	exit 1;
else
	printf "${GREEN}PASS${RESET}\n";
fi
