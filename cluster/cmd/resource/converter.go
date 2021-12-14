package resource

import "github.com/rostislavgit/cloudformation-shipa/cluster/internal/shipa"

func convertModel(currentModel *Model) *shipa.Cluster {
    return &shipa.Cluster{
        Name:      optionalString(currentModel.Name),
        Endpoint:  convertEndpoint(currentModel.Endpoint),
        Resources: convertResources(currentModel.Resources),
    }
}

func convertResources(resources *ClusterResources) *shipa.ClusterResources {
    if resources == nil {
        return nil
    }
    return &shipa.ClusterResources{
        Frameworks:         convertFrameworks(resources.Frameworks),
        IngressControllers: convertIngressControllers(resources.IngressControllers),
    }
}

func convertIngressControllers(controllers []IngressController) (out []*shipa.IngressController) {
    for _, c := range controllers {
        out = append(out, &shipa.IngressController{
            IngressIP:     optionalString(c.IngressIP),
            ServiceType:   optionalString(c.ServiceType),
            Type:          optionalString(c.Type),
            HTTPPort:      int64(optionalInt(c.HTTPPort)),
            HTTPSPort:     int64(optionalInt(c.HTTPSPort)),
            ProtectedPort: int64(optionalInt(c.ProtectedPort)),
            Debug:         optionalBool(c.Debug),
            AcmeEmail:     optionalString(c.AcmeEmail),
            AcmeServer:    optionalString(c.AcmeServer),
        })
    }
    return out
}

func convertFrameworks(frameworks []Framework) (out []*shipa.Framework) {
    for _, f := range frameworks {
        out = append(out, &shipa.Framework{
            Name: optionalString(f.Name),
        })
    }
    return out
}

func convertEndpoint(endpoint *ClusterEndpoint) *shipa.ClusterEndpoint {
    if endpoint == nil {
        return nil
    }
    return &shipa.ClusterEndpoint{
        Addresses:         endpoint.Addresses,
        Certificate:       optionalString(endpoint.Certificate),
        ClientCertificate: optionalString(endpoint.ClientCertificate),
        ClientKey:         optionalString(endpoint.ClientKey),
        Token:             optionalString(endpoint.Token),
        Username:          optionalString(endpoint.Username),
        Password:          optionalString(endpoint.Password),
    }
}

func optionalString(s *string) string {
    if s == nil {
        return ""
    }
    return *s
}

func optionalBool(b *bool) bool {
    if b == nil {
        return false
    }
    return *b
}

func optionalInt(val *int) int {
    if val == nil {
        return 0
    }
    return *val
}
