package apigateway

import (
	"clzrt.io/autolabel/struct/logstruct"
	"log"
	"regexp"
	"strings"
)

func Gateway(logAudit *logstruct.GatewayLog) error {
	resourceName := logAudit.ProtoPayload.ResourceName
	log.Println("Get into Getgateway function")
	gateway, err := Getgateway(resourceName)
	if err != nil {
		return err
	}
	labelSanitizer := regexp.MustCompile("[^a-zA-Z0-9-]+")
	creatorString := labelSanitizer.ReplaceAllString(strings.ToLower(logAudit.ProtoPayload.AuthenticationInfo.PrincipalEmail), "-")
	labels := map[string]string{
		"created-by":   creatorString,
		"gateway-name": gateway.GetDisplayName(),
	}
	log.Println("Get into SetGetgateway function")
	err = Setgateway(labels, gateway)
	if err != nil {
		return err
	}
	log.Println("set gateway label successfully")
	return nil
}

func Api(logAudit *logstruct.ApigatewayLog) error {
	resourceName := logAudit.ProtoPayload.ResourceName
	log.Println("Get into GetAPI function")
	api, err := Getapi(resourceName)
	if err != nil {
		return err
	}
	labelSanitizer := regexp.MustCompile("[^a-zA-Z0-9-]+")
	creatorString := labelSanitizer.ReplaceAllString(strings.ToLower(logAudit.ProtoPayload.AuthenticationInfo.PrincipalEmail), "-")
	labels := map[string]string{
		"created-by":   creatorString,
		"gateway-name": api.GetDisplayName(),
	}
	log.Println("Get into SetAPI function")
	err = Setapi(labels, api)
	if err != nil {
		return err
	}
	log.Println("set api label successfully")
	return nil

}
