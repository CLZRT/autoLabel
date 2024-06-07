package ipaddress

import (
	"clzrt.io/autolabel/struct/logstruct"
	"log"
	"regexp"
	"strings"
)

func StaticIp(logAudit *logstruct.IpaddressLog) error {

	resourceNameArray := strings.Split(logAudit.ProtoPayload.ResourceName, "/")
	ipLocation := map[string]string{
		"project_id": resourceNameArray[1],
		"region":     resourceNameArray[3],
		"name":       resourceNameArray[5],
	}
	for i, v := range resourceNameArray {
		log.Println("Get IP Address", i, v)
	}
	ipaddress, err := getIpaddress(ipLocation)
	if err != nil || ipaddress == nil {
		return err
	}
	labelSanitizer := regexp.MustCompile("[^a-zA-Z0-9-]+")
	creatorString := labelSanitizer.ReplaceAllString(strings.ToLower(logAudit.ProtoPayload.AuthenticationInfo.PrincipalEmail), "-")
	labels := map[string]string{
		"created-by":     creatorString,
		"ipaddress-name": ipaddress.GetName(),
		"ipaddress-type": strings.ToLower(*ipaddress.AddressType),
	}
	for k, v := range labels {
		log.Println("Set Label", k, v)
	}
	log.Println("Set IP Address", ipaddress.GetName())
	err = setIpaddress(ipLocation, labels, ipaddress)
	if err != nil {
		return err
	}
	return nil
}

func GlobalStaticIp(logAudit *logstruct.GlobalAddressLog) error {
	resourceNameArray := strings.Split(logAudit.ProtoPayload.ResourceName, "/")
	ipLocation := map[string]string{
		"project_id": resourceNameArray[1],
		"name":       resourceNameArray[4],
	}
	for i, v := range resourceNameArray {
		log.Println("Get IP Address", i, v)
	}
	ipaddress, err := getGlobalIP(ipLocation)
	if err != nil || ipaddress == nil {
		return err
	}
	labelSanitizer := regexp.MustCompile("[^a-zA-Z0-9-]+")
	creatorString := labelSanitizer.ReplaceAllString(strings.ToLower(logAudit.ProtoPayload.AuthenticationInfo.PrincipalEmail), "-")
	labels := map[string]string{
		"created-by":     creatorString,
		"ipaddress-name": ipaddress.GetName(),
		"ipaddress-type": strings.ToLower(*ipaddress.AddressType),
	}
	for k, v := range labels {
		log.Println("Set Label", k, v)
	}
	log.Println("Set IP Address", ipaddress.GetName())
	err = setGlobalIp(ipLocation, labels, ipaddress)
	if err != nil {
		return err
	}
	return nil
}
