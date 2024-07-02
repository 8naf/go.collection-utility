source ./.scripts/_common.sh

COVERAGE_PROFILE="$1"
CODE_COVERAGE_THRESHOLD=$2
CODE_COVERAGE_ACTUAL=$(go tool cover -func=$COVERAGE_PROFILE | grep total | grep -Eo '[0-9]+\.[0-9]+')

is_less_than_threshold=$(echo "$CODE_COVERAGE_ACTUAL < $CODE_COVERAGE_THRESHOLD" | bc -l)

actual_coverage_color=$GREEN
if [[ $is_less_than_threshold -eq 1 ]]; then
	actual_coverage_color=$RED
fi

printf "${BLUE}CODE COVERAGE REPORT${RESET}\n"
hline

printf "Code Coverage (Threshold): ${PURPLE}%10s${RESET}\n" "$CODE_COVERAGE_THRESHOLD%"
printf "Code Coverage (Actual)   : ${actual_coverage_color}%10s${RESET}\n" "$CODE_COVERAGE_ACTUAL%"
hline

if [[ $is_less_than_threshold -eq 1 ]]; then
	printf "${CYAN}Please add more unit tests to increase the code coverage.${RESET}\n";
	printf "${RED}FAILED${RESET}\n";
	exit 1;
else
	printf "${CYAN}Passed the code coverage threshold.${RESET}\n";
	printf "${GREEN}PASS${RESET}\n";
fi
