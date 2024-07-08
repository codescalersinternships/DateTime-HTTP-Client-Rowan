# DateTime-HTTP-Client-Rowan
 An HTTP client in Go that consumes the datetime server APIs implemented in the previous project: DateTime-Server-Rowan 

## 1. Client Functionality:

- Endpoint: `/datetime`
- Request is send at end point using `Get`

## 2. Implementations:

The client is implemented using:
1. Standard library `(net/http)`

## 3. Directories Walkthrough
        root  
        |____cmd
        |       |___datetimecli
        |                      |___main.go
        |
        |____pkg
                |___client.go (pkg code)
                |___client_test.go (test file)
## 4. EndPoint 
   - Client takes base URL and concatinates it with the `/datetime` endpoint

## 5. Package Testing
### How Testing Is Done (Integration/Mock):
- Using `httptest` package, we can mock a server, so that our client sends actual requests but to a mocked server.
- By accessing mockServer.URL --> we can send its path to our client smoothly.
- In tests we're assuming _a tolerance of 1 second maximum._
- Tests cover 2 main points here:
  - `Correctness` by testing against Time.Now()
  - `Error Detections` by testing wrong path sending
- 2 main Testing Functions:
  - `Test_GetDateTime`
  - `Test_AllPossibilities` runs the previous function numerous times to ensure correctness

### Coverage
- Test coverage: 94.7% of statements

## 6. Package Importing
```
import
"github.com/codescalersinternships/DateTime-HTTP-Client-Rowan/pkg"
```
If not available
```
go get github.com/codescalersinternships/DateTime-HTTP-Client-Rowan/pkg
```
## 7. UseCases

<!-- ### Commands:  -->
 To use package, user must intialize it using a server URL
```
client := pkg.NewDateTimeClient("http://localhost:8080")
```
 User can then invoke the GetDateTime method which returns [ ]byte and err
```
data, err := client.GetDateTime()
```
User can use environment variables as a path as well:
```
URL=http://localhost:8080 go run cmd/datetimecli/main.go
```