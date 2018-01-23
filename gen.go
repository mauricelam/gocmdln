package params

//go:generate go run $GOPATH/src/github.com/cheekybits/genny/main.go -pkg=params -in=value_list.gogen -out=gen-value_list.go gen "gen=bool,string,int,int64,uint,uint64,float64,time.Duration"
