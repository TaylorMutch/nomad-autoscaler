set -ex

# Setup the directory to put coverage reports.
COVERAGE_DIR="cover"
mkdir -p "cover"

# Get the list of packages, and test each one.
PKG_LIST=$(go list ./... | grep -v /vendor/);
for package in ${PKG_LIST}; do
    go test -v -race -timeout=120s -failfast -vet="all"  -tags test -benchmem -bench=. -covermode=atomic -coverprofile "${COVERAGE_DIR}/${package##*/}.cov" "$package";
done;

# Initialize the global code coverage report.
# Move logs from individual reports to global report.
echo 'mode: atomic' > coverage.cov;
tail -q -n +2 cover/*.cov >> coverage.cov;

# Analyze the global report.
go tool cover -func=coverage.cov;
go tool cover -html=coverage.cov -o coverage.html;

# Remove the leftover files.
rm -rf "$COVERAGE_DIR";
rm coverage.cov;