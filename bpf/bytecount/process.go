package bytecount

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"time"

	"github.com/cilium/cilium/pkg/identity"
	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/perf"
	"github.com/dinoallo/sealos-networkmanager-agent/store"
	"github.com/dinoallo/sealos-networkmanager-agent/util"
)

func (bf *Factory) readTraffic(ctx context.Context, t uint32) error {
	log := bf.logger
	objs := bf.objs
	var eventArray *ebpf.Map
	switch t {
	case IPv4Ingress.TypeInt:
		eventArray = objs.IngressTrafficEvents
	case IPv4Egress.TypeInt:
		eventArray = objs.EgressTrafficEvents
	default:
		return fmt.Errorf("unknown direction")
	}
	er, err := perf.NewReader(eventArray, PERF_BUFFER_SIZE)
	if err != nil {
		return fmt.Errorf("failed to create a new reader: %v", err)
	}
	go func(ctx context.Context) {
		defer er.Close()
		<-ctx.Done()
	}(ctx)

	for {
		rec, err := er.Read()
		if err != nil {
			if errors.Is(err, perf.ErrClosed) {
				log.Infof("the perf event channel is closed")
				return nil
			} else {
				log.Errorf("unable to read from perf event reader: %v", err)
				return nil
			}
		}
		if rec.LostSamples != 0 {
			log.Infof("perf event ring buffer full, dropped %d samples", rec.LostSamples)
			continue
		}
		tr := Traffic{
			trafficRecord: &rec,
			trafficType:   t,
		}
		go bf.processTraffic(ctx, tr)
	}
}

func (bf *Factory) processTraffic(ctx context.Context, traffic Traffic) {
	log := bf.logger
	var event bytecountTrafficEventT
	if traffic.trafficRecord != nil {
		if err := binary.Read(bytes.NewBuffer(traffic.trafficRecord.RawSample), bf.nativeEndian, &event); err != nil {
			log.Infof("Failed to decode received data: %+v", err)
			return
		}
		t := traffic.trafficType
		if err := bf.submit(ctx, &event, t); err != nil {
			log.Infof("Failed to submit the traffic report: %+v", err)
			return
		}
	}
}

func (bf *Factory) submit(ctx context.Context, event *bytecountTrafficEventT, t uint32) error {
	var dir store.TrafficDirection
	__srcIP := make([]uint32, 4)
	__dstIP := make([]uint32, 4)
	var srcPort uint32
	var dstPort uint32
	switch t {
	case IPv4Ingress.TypeInt:
		dir = store.TRAFFIC_DIR_V4_INGRESS
	case IPv4Egress.TypeInt:
		dir = store.TRAFFIC_DIR_V4_EGRESS
	default:
		return nil
	}
	__srcIP[0] = event.SrcIp4
	__dstIP[0] = event.DstIp4
	srcPort = event.SrcPort
	dstPort = uint32(event.DstPort)
	srcIP := util.ToIP(__srcIP[0], nil, 4)
	dstIP := util.ToIP(__dstIP[0], nil, 4)
	report := &store.TrafficReport{
		TrafficReportMeta: store.TrafficReportMetaData{
			SrcIP:   srcIP.String(),
			DstIP:   dstIP.String(),
			SrcPort: srcPort,
			DstPort: dstPort,
		},
		Dir:       dir,
		Protocol:  event.Protocol,
		Family:    event.Family,
		DataBytes: event.Len,
		Identity:  identity.NumericIdentity(event.Identity),
		Timestamp: time.Now(),
	}
	bf.logger.Infof("src_ip: %v => dst_ip: %v; %v bytes sent", srcIP.String(), dstIP.String(), event.Len)
	bf.trStore.AddTrafficReport(ctx, report)
	return nil
}
