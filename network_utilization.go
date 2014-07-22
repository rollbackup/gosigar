package sigar

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type NetworkUtilization map[string]DeviceNetworkUtilization

type DeviceNetworkUtilization struct {
	RxBytes          int64
	RxPackets        int64
	RxErrors         int64
	RxDroppedPackets int64
	TxBytes          int64
	TxPackets        int64
	TxErrors         int64
	TxDroppedPackets int64
}

func (self *NetworkUtilization) Get() error {
	statFile, err := ioutil.ReadFile("/proc/net/dev")
	if err != nil {
		return err
	}

	lines := strings.Split(string(statFile), "\n")
	if len(lines) <= 2 {
		return fmt.Errorf("/proc/net/dev doesn't have the expected format")
	}

	for _, line := range lines[2:] {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) < 16 {
			return fmt.Errorf("/proc/net/dev doesn't have the expected format. Expected 16 fields found %d", len(fields))
		}
		name := strings.Trim(fields[0], ":")
		utilization := DeviceNetworkUtilization{}
		if utilization.RxBytes, err = strconv.ParseInt(fields[1], 10, 64); err != nil {
			return err
		}
		if utilization.RxPackets, err = strconv.ParseInt(fields[2], 10, 64); err != nil {
			return err
		}
		if utilization.RxErrors, err = strconv.ParseInt(fields[3], 10, 64); err != nil {
			return err
		}
		if utilization.RxDroppedPackets, err = strconv.ParseInt(fields[4], 10, 64); err != nil {
			return err
		}
		if utilization.TxBytes, err = strconv.ParseInt(fields[9], 10, 64); err != nil {
			return err
		}
		if utilization.TxPackets, err = strconv.ParseInt(fields[10], 10, 64); err != nil {
			return err
		}
		if utilization.TxErrors, err = strconv.ParseInt(fields[11], 10, 64); err != nil {
			return err
		}
		if utilization.TxDroppedPackets, err = strconv.ParseInt(fields[12], 10, 64); err != nil {
			return err
		}
		(*self)[name] = utilization
	}
	return nil
}
