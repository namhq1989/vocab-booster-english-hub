package grpcclient

// func NewStaffClient(_ *appcontext.AppContext, addr string) (staffpb.StaffServiceClient, error) {
// 	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return staffpb.NewStaffServiceClient(conn), nil
// }

// func NewAuthClient(ctx *appcontext.AppContext, addr string) (authpb.AuthServiceClient, error) {
// 	conn, err := grpc.DialContext(ctx.Context(), addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return authpb.NewAuthServiceClient(conn), nil
// }
