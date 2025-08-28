### Script
go run *.go
go mod init example.com/basic-practice (ini seperti package.json di JS)
go get ... untuk dapetin module eksternal
go get     untuk download semua list module yang ada di mod

=============================================

### Packages
import (
	"fmt"
	"math"    (for calculation, like pow)
)

fmt.Scan()

fmt.Printf("Future Value: %f\nFuture Value (adjusted for inflation): %f\n", futureValue, futureRealValue)

formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)

%v default value
%T type data
%f decimal
%.2f 2 angka dibelakang koma

===============================================


### Variable
var investmentAmount float64 = 1000
expectedReturnRate := 5.5  (short syntax)

