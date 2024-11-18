package apigateway

import (
	apigateway "cloud.google.com/go/apigateway/apiv1"
	"cloud.google.com/go/apigateway/apiv1/apigatewaypb"
	"context"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"log"
)

func Getgateway(resourceName string) (*apigatewaypb.Gateway, error) {
	ctx := context.Background()
	client, err := apigateway.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	log.Println("Get API gateway:", resourceName)
	gateway, err := client.GetGateway(ctx, &apigatewaypb.GetGatewayRequest{
		// Required. Resource name of the form:
		// `projects/*/locations/*/gateways/*`
		Name: resourceName,
	})
	if err != nil {
		return nil, err
	}
	return gateway, nil
}

func Setgateway(labels map[string]string, gateway *apigatewaypb.Gateway) error {
	ctx := context.Background()
	client, err := apigateway.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()
	log.Println("Set API gateway:", gateway.Name)
	gateway.Labels = labels
	_, err = client.UpdateGateway(ctx, &apigatewaypb.UpdateGatewayRequest{
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: []string{"labels"},
		},
		Gateway: gateway,
	})
	if err != nil {
		return err
	}
	log.Println("Label API gateway:", gateway.Name+" Successfully")
	return nil
}

func Getapi(resourceName string) (*apigatewaypb.Api, error) {
	ctx := context.Background()
	client, err := apigateway.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	log.Println("Get API :", resourceName)
	api, err := client.GetApi(ctx, &apigatewaypb.GetApiRequest{
		// Required. Resource name of the form:
		// `projects/*/locations/global/apis/*`
		Name: resourceName,
	})
	if err != nil {
		return nil, err
	}
	return api, nil
}

func Setapi(labels map[string]string, api *apigatewaypb.Api) error {
	ctx := context.Background()
	client, err := apigateway.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()
	api.Labels = labels
	_, err = client.UpdateApi(ctx, &apigatewaypb.UpdateApiRequest{
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: []string{"labels"},
		},
		Api: api,
	})
	if err != nil {
		return err
	}
	log.Println("Set API :", labels["name"], "Successfully")
	return nil
}
