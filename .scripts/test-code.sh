COVERAGE_PROFILE="$1"
go test -race -vet=off -v ./... -covermode=atomic -coverpkg=./... -coverprofile=$COVERAGE_PROFILE
