package go_saas_basic_tls

import "sync"

type Tls struct {
	CertPath string
	KeyPath  string
	*sync.RWMutex
}

func (tls *Tls) GetCertPath() string {
	tls.RLock()
	defer tls.RUnlock()

	return tls.CertPath
}

func (tls *Tls) GetKeyPath() string {
	tls.RLock()
	defer tls.RUnlock()

	return tls.CertPath
}
