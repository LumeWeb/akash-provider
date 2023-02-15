package kube

const (
	Issuer        = "Issuer"
	ClusterIssuer = "cluster-Issuer"
)

type ClientConfig struct {
	Ssl Ssl
}

type Ssl struct {
	IssuerType string
	IssuerName string
}
