package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const countersFilePath = "/sys/class/infiniband/xscale_0/counters/counters"

//const countersFilePath = "./counters"

// rdma_tx_pkts                  16
// rdma_tx_bytes                 17568
// rdma_rx_pkts                  2
// rdma_rx_bytes                 124
// np_cnp_sent                   0
// rp_cnp_handled                0
// np_ecn_marked_roce_packets    0
// rp_cnp_ignored                0
// out_of_sequence               0
// packet_seq_err                0
// out_of_buffer                 0
// rnr_nak_retry_err             0
// local_ack_timeout_err         0
// tx_pause                      0
// rx_pause                      0
var (
	txPackets      int
	rxPackets      int
	rdmaTxPktsDesc = prometheus.NewDesc(
		"rdma_tx_pkts",
		"rdma_tx_pkts",
		nil, nil,
	)

	rdmaRxPktsDesc = prometheus.NewDesc(
		"rdma_rx_pkts",
		"rdma_rx_pkts",
		nil, nil,
	)
	npCnpSentDesc = prometheus.NewDesc(
		"np_cnp_sent",
		"np_cnp_sent",
		nil, nil,
	)
	rpCnpHandledDesc = prometheus.NewDesc(
		"rp_cnp_handled",
		"rp_cnp_handled",
		nil, nil,
	)
	npEcnMarkedRocePacketsDesc = prometheus.NewDesc(
		"np_ecn_marked_roce_packets",
		"np_ecn_marked_roce_packets",
		nil, nil,
	)
	rpCnpIgnoredDesc = prometheus.NewDesc(
		"rp_cnp_ignored",
		"rp_cnp_ignored",
		nil, nil,
	)
	outOfSequenceDesc = prometheus.NewDesc(
		"out_of_sequence",
		"out_of_sequence",
		nil, nil,
	)
	packetSeqErrDesc = prometheus.NewDesc(
		"packet_seq_err",
		"packet_seq_err",
		nil, nil,
	)
	outOfBufferDesc = prometheus.NewDesc(
		"out_of_buffer",
		"out_of_buffer",
		nil, nil,
	)
	rnrNakRetryErrDesc = prometheus.NewDesc(
		"rnr_nak_retry_err",
		"rnr_nak_retry_err",
		nil, nil,
	)
	localAckTimeoutErrDesc = prometheus.NewDesc(
		"local_ack_timeout_err",
		"local_ack_timeout_err",
		nil, nil,
	)
	txPauseDesc = prometheus.NewDesc(
		"tx_pause",
		"tx_pause",
		nil, nil,
	)
	rxPauseDesc = prometheus.NewDesc(
		"rx_pause",
		"rx_pause",
		nil, nil,
	)
)

type myCollector struct {
}

func (c *myCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- rdmaTxPktsDesc
	ch <- rdmaRxPktsDesc
}

func (c *myCollector) Collect(ch chan<- prometheus.Metric) {
	// your logic should be placed here

	currentTime := time.Now()
	content, err := ioutil.ReadFile(countersFilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		//fmt.Printf("%s: %s\n", fields[0], fields[1])
		if fields[0] == "rdma_tx_pkts" {
			txPackets, _ = strconv.Atoi(fields[1])
			fmt.Println("txPackets: ", txPackets)

			ch <- prometheus.NewMetricWithTimestamp(currentTime,
				prometheus.MustNewConstMetric(rdmaTxPktsDesc, prometheus.GaugeValue, float64(txPackets)))
		}
		if fields[0] == "rdma_rx_pkts" {
			rxPackets, _ = strconv.Atoi(fields[1])
			ch <- prometheus.NewMetricWithTimestamp(currentTime,
				prometheus.MustNewConstMetric(rdmaRxPktsDesc, prometheus.GaugeValue, float64(rxPackets)))
		}
		if fields[0] == "np_cnp_sent" {
			npCnpSent, _ := strconv.Atoi(fields[1])
			ch <- prometheus.NewMetricWithTimestamp(currentTime,
				prometheus.MustNewConstMetric(npCnpSentDesc, prometheus.GaugeValue, float64(npCnpSent)))
		}
		if fields[0] == "rp_cnp_handled" {
			rpCnpHandled, _ := strconv.Atoi(fields[1])
			ch <- prometheus.NewMetricWithTimestamp(currentTime,
				prometheus.MustNewConstMetric(rpCnpHandledDesc, prometheus.GaugeValue, float64(rpCnpHandled)))
		}
		if fields[0] == "np_ecn_marked_roce_packets" {
			npEcnMarkedRocePackets, _ := strconv.Atoi(fields[1])
			ch <- prometheus.NewMetricWithTimestamp(currentTime,
				prometheus.MustNewConstMetric(npEcnMarkedRocePacketsDesc, prometheus.GaugeValue, float64(npEcnMarkedRocePackets)))
		}
		if fields[0] == "rp_cnp_ignored" {
			rpCnpIgnored, _ := strconv.Atoi(fields[1])
			ch <- prometheus.NewMetricWithTimestamp(currentTime,
				prometheus.MustNewConstMetric(rpCnpIgnoredDesc, prometheus.GaugeValue, float64(rpCnpIgnored)))
		}
		if fields[0] == "out_of_sequence" {
			outOfSequence, _ := strconv.Atoi(fields[1])
			ch <- prometheus.NewMetricWithTimestamp(currentTime,
				prometheus.MustNewConstMetric(outOfSequenceDesc, prometheus.GaugeValue, float64(outOfSequence)))
		}
		if fields[0] == "packet_seq_err" {
			packetSeqErr, _ := strconv.Atoi(fields[1])
			ch <- prometheus.NewMetricWithTimestamp(currentTime,
				prometheus.MustNewConstMetric(packetSeqErrDesc, prometheus.GaugeValue, float64(packetSeqErr)))
		}
		if fields[0] == "out_of_buffer" {
			outOfBuffer, _ := strconv.Atoi(fields[1])
			ch <- prometheus.NewMetricWithTimestamp(currentTime,
				prometheus.MustNewConstMetric(outOfBufferDesc, prometheus.GaugeValue, float64(outOfBuffer)))
		}
		if fields[0] == "rnr_nak_retry_err" {
			rnrNakRetryErr, _ := strconv.Atoi(fields[1])
			ch <- prometheus.NewMetricWithTimestamp(currentTime,
				prometheus.MustNewConstMetric(rnrNakRetryErrDesc, prometheus.GaugeValue, float64(rnrNakRetryErr)))
		}
		if fields[0] == "local_ack_timeout_err" {
			localAckTimeoutErr, _ := strconv.Atoi(fields[1])
			ch <- prometheus.NewMetricWithTimestamp(currentTime,
				prometheus.MustNewConstMetric(localAckTimeoutErrDesc, prometheus.GaugeValue, float64(localAckTimeoutErr)))
		}
		if fields[0] == "tx_pause" {
			txPause, _ := strconv.Atoi(fields[1])
			ch <- prometheus.NewMetricWithTimestamp(currentTime,
				prometheus.MustNewConstMetric(txPauseDesc, prometheus.GaugeValue, float64(txPause)))
		}
		if fields[0] == "rx_pause" {
			rxPause, _ := strconv.Atoi(fields[1])
			ch <- prometheus.NewMetricWithTimestamp(currentTime,
				prometheus.MustNewConstMetric(rxPauseDesc, prometheus.GaugeValue, float64(rxPause)))
		}

	}

	//读取文件内容

}

func main() {

	collector := &myCollector{}
	prometheus.MustRegister(collector)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8000", nil)
}
