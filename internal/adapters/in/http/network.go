package http

import (
	"net"
)

// DetectLANIP descobre o endereço IPv4 não-loopback primário da máquina na rede local (LAN).
// Caso nenhuma interface de rede externa esteja ativa, retorna "127.0.0.1" como fallback seguro.
func DetectLANIP() string {
	return detectLANIPWithDialer(func(network, address string) (net.Conn, error) {
		return net.Dial(network, address)
	}, net.Interfaces)
}

// dialerFunc abstrai net.Dial para viabilizar testes unitários determinísticos.
type dialerFunc func(network, address string) (net.Conn, error)

// ifacesFunc abstrai net.Interfaces para viabilizar testes unitários determinísticos.
type ifacesFunc func() ([]net.Interface, error)

func detectLANIPWithDialer(dial dialerFunc, getIfaces ifacesFunc) string {
	// 1. Tenta identificar o IP de saída ativo via socket UDP (sem trafegar pacotes reais)
	if dial != nil {
		conn, err := dial("udp", "8.8.8.8:80")
		if err == nil && conn != nil {
			defer conn.Close()
			if localAddr, ok := conn.LocalAddr().(*net.UDPAddr); ok {
				ip := localAddr.IP.To4()
				if ip != nil && !ip.IsLoopback() && !ip.IsUnspecified() {
					return ip.String()
				}
			}
		}
	}

	// 2. Fallback: itera sobre todas as interfaces de rede do sistema operacional
	if getIfaces != nil {
		ifaces, err := getIfaces()
		if err == nil {
			for _, iface := range ifaces {
				// Ignora interfaces inativas ou de loopback
				if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
					continue
				}

				addrs, err := iface.Addrs()
				if err != nil {
					continue
				}

				for _, addr := range addrs {
					var ip net.IP
					switch v := addr.(type) {
					case *net.IPNet:
						ip = v.IP
					case *net.IPAddr:
						ip = v.IP
					}

					ip4 := ip.To4()
					if ip4 != nil && !ip4.IsLoopback() && !ip4.IsUnspecified() {
						return ip4.String()
					}
				}
			}
		}
	}

	return "127.0.0.1"
}
