package registry

import (
	"errors"
	"strings"

	"github.com/micro/mdns"
)

type mdnsWatcher struct {
	domain string
	wo     WatchOptions
	ch     chan *mdns.ServiceEntry
	exit   chan struct{}
}

func (m *mdnsWatcher) Next() (*Result, error) {
	for {
		select {
		case e := <-m.ch:

			txt, err := decode(e.InfoFields)
			if err != nil {
				continue
			}

			if len(e.Name) == 0 || len(txt.Version) == 0 {
				continue
			}

			// Filter watch options
			// wo.Service: Only keep services we care about
			if len(m.wo.Service) > 0 && e.Name != m.wo.Service {
				continue
			}

			var action string

			if e.TTL == 0 {
				action = "delete"
			} else {
				action = "create"
			}

			service := &Service{
				Name:      e.Name,
				Version:   txt.Version,
				Endpoints: txt.Endpoints,
			}

			// TODO: don't hardcode .local.
			if !strings.HasSuffix(e.Name, "."+service.Name+"."+m.domain+".") {
				continue
			}

			service.Nodes = append(service.Nodes, &Node{
				Id:       strings.TrimSuffix(e.Name, "."+service.Name+"."+m.domain+"."),
				Address:  e.AddrV4.String(),
				Port:     e.Port,
				Metadata: txt.Metadata,
			})

			return &Result{
				Action:  action,
				Service: service,
			}, nil
		case <-m.exit:
			return nil, errors.New("watcher stopped")
		}
	}
}

func (m *mdnsWatcher) Stop() {
	select {
	case <-m.exit:
		return
	default:
		close(m.exit)
	}
}
