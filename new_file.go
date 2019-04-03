package main

import (
	"context"
	"google.golang.org/grpc"
	"kbtg.tech/template/protosvr"
	"log"
	"time"
)

func aa() {

	cc, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()
	c := protosvr.NewYottaServiceClient(cc)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	b := &protosvr.CardHolder{BaseInfoSeg: &protosvr.CardHolder_BASEINFOSEG{
		CardNo: "1222233",
	},
		AnnualFeeSeg: &protosvr.CardHolder_ANNUALFEESEG{
			NextDueDate: "",
		},
		EMVSeg: &protosvr.CardHolder_EMVSEG{ChipATC: 1, ChipFlag: true}}
	c.Add(ctx, b)
}
