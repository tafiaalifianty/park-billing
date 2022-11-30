# Simple Parking Bill
### How to start
1. Clone this repo to your local machine
2. Open terminal and `cd` to this project
3. Install dependencies with `go install`
4. You can use `main.go` to test the output, build and run main fn with command `go run .`
5. After all functions completed run `go test ./...`, make sure all the test is passed!

### Parking Lot Billing
Parking Price List

| Duration   | Motorcycle | Car       |
|------------|------------|-----------|
| First 1 hr | 3000       | 7000      |
| After 1 hr | 2000/hour  | 5000/hour |

If vehicle parked more than 24hrs then they will be an extra charged, as below:

| Vehicle Type | Extra Charge |
|--------------|--------------|
| Motorcycle   | +20000       |
| Car          | +50000       |
