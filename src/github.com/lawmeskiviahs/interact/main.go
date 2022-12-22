package main 

import (
    // "context"
    "fmt"

    "google.golang.org/grpc"

    "github.com/cosmos/cosmos-sdk/codec"
    sdk "github.com/cosmos/cosmos-sdk/types"
    // banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	bellchain "github.com/lawmeskiviahs/wasmd/x/bellchain/types"
)

// func main() {
// 	queryState()
// }
func main() {
    myAddress, err := sdk.AccAddressFromBech32("cosmos1vjhxfjakfhua5urxzqmkxncl8v6z40pp4s89dg")
    if err != nil {
        // return err
		fmt.Println(err)
    }

    // Create a connection to the gRPC server.
    grpcConn, err := grpc.Dial(
        "127.0.0.1:9090", // your gRPC server address.
        grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism. 
        // This instantiates a general gRPC codec which handles proto bytes. We pass in a nil interface registry
        // if the request/response types contain interface instead of 'nil' you should pass the application specific codec.
        grpc.WithDefaultCallOptions(grpc.ForceCodec(codec.NewProtoCodec(nil).GRPCCodec())),
    )
    if err != nil {
        // return err
		fmt.Println(err)
    }
    defer grpcConn.Close()

    // This creates a gRPC client to query the x/bank service.
    // bankClient := banktypes.NewQueryClient(grpcConn)
    // bankRes, err := bankClient.Balance(
    //     context.Background(),
    //     &banktypes.QueryBalanceRequest{Address: myAddress.String(), Denom: "stake"},
    // )
    // if err != nil {
    //     // return err
	// 	fmt.Println(err)
    // }

    // fmt.Println(bankRes.GetBalance()) // Prints the account balance

	bellClient := bellchain.NewQueryClient(grpcConn)
	fmt.Println(bellClient)

}