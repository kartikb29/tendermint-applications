package nameservice

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Initial Starting Price for a name that was never previously owned
//var MinNamePrice = sdk.Coins{sdk.NewInt64Coin("nametoken", 1)}

// Record is a struct that contains all the metadata of a DB entry you want to add
type Record struct {
	//_id          string         `json:"_id"`
	Owner        sdk.AccAddress `json:"owner"`
	creationTime int64          `json:"timeStamp"`
	data         string         `json:"data"`
}

// NewRecord returns a record with a randomly generated _id and its creationTime
func NewRecord() Record {
	//val, _ := randomHex(10)
	currentTime := time.Now()
	emptyString := ""
	return Record{
		//_id:          val,
		creationTime: currentTime,
		data:         emptyString,
	}
}

// func randomHex(n int) (string, error) {
// 	bytes := make([]byte, n)
// 	if _, err := rand.Read(bytes); err != nil {
// 		return "", err
// 	}
// 	return hex.EncodeToString(bytes), nil
// }
