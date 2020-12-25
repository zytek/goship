package resources

// DefaultShortOutputTemplate defines default template for the list output
var DefaultShortOutputTemplate = "{{range .NativeObject.Tags}}{{if eq (.Key | String) \"Name\"}}{{.Value | String | printf \"%-33s\"}}{{end}}{{end}}   {{range .NativeObject.Tags}}{{if eq (.Key | String) \"environment\"}}{{.Value | String | printf \"%-10s\"}}{{end}}{{end}} {{.NativeObject.Placement.AvailabilityZone}}  {{.NativeObject.PrivateIpAddress| String | printf \"%-16s\"}} {{.NativeObject.PublicIpAddress| String |printf \"%-16s\"}}{{.ProfileName | String |printf \"%-20s\"}}\n"

// DefaultLongOutputTemplate defines default template for the detailed output
var DefaultLongOutputTemplate = `
{{range .NativeObject.Tags}}{{if eq (.Key | String) "Name"}}{{.Value}}{{end}}{{end}}
  profile             {.ProfileName}
  ami_id              {.NativeObject.ImageId}}
  az                  {.NativeObject.Placement.AvailabilityZone}}
  dns_name            {.NativeObject.PublicDnsName}}
  id                  {.NativeObject.InstanceId}}
  instance_type       {.NativeObject.InstanceType}}
  key_name            {.NativeObject.KeyName}}
  private_dns_name    {.NativeObject.PrivateDnsName}}
  private_ip          {.NativeObject.PrivateIpAddress}}
  public_ip           {.NativeObject.PublicIpAddress}}
  tags
    {{range .NativeObject.Tags}}{{.Key | String | printf "%-10s"}}        {{.Value}}
    {{end}}
`
