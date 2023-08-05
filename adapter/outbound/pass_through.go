package outbound

import (
	"context"
	"errors"

	"github.com/Dreamacro/clash/component/dialer"
	C "github.com/Dreamacro/clash/constant"
)

type PassThrough struct {
	*Base
}

// DialContext implements C.ProxyAdapter
func (p *PassThrough) DialContext(ctx context.Context, metadata *C.Metadata, opts ...dialer.Option) (C.Conn, error) {
	return nil, errors.New("no support")
}

// ListenPacketContext implements C.ProxyAdapter
func (p *PassThrough) ListenPacketContext(ctx context.Context, metadata *C.Metadata, opts ...dialer.Option) (C.PacketConn, error) {
	return nil, errors.New("no support")
}

func NewPassThrough() *PassThrough {
	return &PassThrough{
		Base: &Base{
			name: "PASS_THROUGH",
			tp:   C.PassThrough,
			udp:  true,
		},
	}
}
