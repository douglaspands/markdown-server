package http

import (
	"errors"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

// dummyConn simula uma conexão net.Conn com endereço IP customizado para testes.
type dummyConn struct {
	net.Conn
	addr net.Addr
}

func (d *dummyConn) LocalAddr() net.Addr {
	return d.addr
}

func (d *dummyConn) Close() error {
	return nil
}

func TestDetectLANIP(t *testing.T) {
	t.Run("Given sistema operacional padrão When DetectLANIP é executado Then retorna um endereço IPv4 válido", func(t *testing.T) {
		ipStr := DetectLANIP()
		assert.NotEmpty(t, ipStr)

		parsedIP := net.ParseIP(ipStr)
		assert.NotNil(t, parsedIP, "O retorno deve ser um IP válido")
		assert.NotNil(t, parsedIP.To4(), "O retorno deve ser um endereço IPv4")
	})

	t.Run("Given UDP dialer com IP local válido When detectLANIPWithDialer executa Then retorna o IP obtido", func(t *testing.T) {
		fakeUDPAddr := &net.UDPAddr{
			IP:   net.ParseIP("192.168.1.105"),
			Port: 54321,
		}
		mockDial := func(network, address string) (net.Conn, error) {
			return &dummyConn{addr: fakeUDPAddr}, nil
		}

		ip := detectLANIPWithDialer(mockDial, nil)
		assert.Equal(t, "192.168.1.105", ip)
	})

	t.Run("Given UDP dialer com erro e interfaces vazias When detectLANIPWithDialer executa Then retorna fallback 127.0.0.1", func(t *testing.T) {
		mockDial := func(network, address string) (net.Conn, error) {
			return nil, errors.New("rede inacessível")
		}
		mockIfaces := func() ([]net.Interface, error) {
			return nil, errors.New("sem interfaces")
		}

		ip := detectLANIPWithDialer(mockDial, mockIfaces)
		assert.Equal(t, "127.0.0.1", ip)
	})

	t.Run("Given UDP dialer retornando IP loopback When detectLANIPWithDialer executa Then faz fallback para 127.0.0.1", func(t *testing.T) {
		fakeUDPAddr := &net.UDPAddr{
			IP:   net.ParseIP("127.0.0.1"),
			Port: 54321,
		}
		mockDial := func(network, address string) (net.Conn, error) {
			return &dummyConn{addr: fakeUDPAddr}, nil
		}
		mockIfaces := func() ([]net.Interface, error) {
			return []net.Interface{}, nil
		}

		ip := detectLANIPWithDialer(mockDial, mockIfaces)
		assert.Equal(t, "127.0.0.1", ip)
	})
}
