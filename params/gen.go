package params

//go:generate go run $GOPATH/src/github.com/mauricelam/genny/main.go -ast -pkg=params -in=value_list.gogen -out=gen-value_list.go gen "PlaceholderType=bool,string,int,int64,uint,uint64,float64,Duration:time.Duration"
