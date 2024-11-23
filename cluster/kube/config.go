package kube

const (
	Issuer        = "issuer"
	ClusterIssuer = "cluster-issuer"
)

type ClientConfig struct {
	Ssl Ssl
}

type Ssl struct {
	IssuerType string
	IssuerName string
}
