// Copyright (c) 2018 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Manage dhcpcd for ports including static
// XXX wwan*? Skip for now since wwan container handles configuring IP

package devicenetwork

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/lf-edge/eve/pkg/pillar/agentlog"
	"github.com/lf-edge/eve/pkg/pillar/base"
	"github.com/lf-edge/eve/pkg/pillar/types"
)

// UpdateDhcpClient starts/modifies/deletes dhcpcd per interface
// Assumes that the caller has checked that the interfaces exist
// We therefore skip any interfaces which do not exist
func UpdateDhcpClient(log *base.LogObject, newConfig, oldConfig types.DevicePortConfig) {

	// Look for adds or changes
	log.Functionf("updateDhcpClient: new %v old %v\n",
		newConfig, oldConfig)
	for _, newU := range newConfig.Ports {
		oldU := lookupOnIfname(oldConfig, newU.IfName)
		if oldU == nil || oldU.Dhcp == types.DT_NONE {
			log.Functionf("updateDhcpClient: new %s\n", newU.IfName)
			// Inactivate in case a dhcpcd is running
			doDhcpClientActivate(log, newU)
		} else {
			log.Functionf("updateDhcpClient: found old %v\n",
				oldU)
			if !reflect.DeepEqual(newU.DhcpConfig, oldU.DhcpConfig) {
				log.Functionf("updateDhcpClient: changed %s\n",
					newU.IfName)
				doDhcpClientInactivate(log, *oldU)
				doDhcpClientActivate(log, newU)
			}
		}
	}
	// Look for deletes from oldConfig to newConfig
	for _, oldU := range oldConfig.Ports {
		newU := lookupOnIfname(newConfig, oldU.IfName)
		if newU == nil || newU.Dhcp == types.DT_NONE {
			log.Functionf("updateDhcpClient: deleted %s\n",
				oldU.IfName)
			doDhcpClientInactivate(log, oldU)
		} else {
			log.Functionf("updateDhcpClient: found new %v\n",
				newU)
		}
	}
}

func doDhcpClientActivate(log *base.LogObject, nuc types.NetworkPortConfig) {

	log.Functionf("doDhcpClientActivate(%s) dhcp %v addr %s gateway %s\n",
		nuc.IfName, nuc.Dhcp, nuc.AddrSubnet,
		nuc.Gateway.String())
	if strings.HasPrefix(nuc.IfName, "wwan") {
		log.Functionf("doDhcpClientActivate: skipping %s\n",
			nuc.IfName)
		return
	}

	// Check the ifname exists to avoid waiting for a dhcpcd below
	_, err := IfnameToIndex(log, nuc.IfName)
	if err != nil {
		// Caller intends us to proceed without this interface
		log.Warnf("doDhcpClientActivate(%s) failed %s", nuc.IfName, err)
		return
	}
	switch nuc.Dhcp {
	case types.DT_NONE:
		log.Functionf("doDhcpClientActivate(%s) DT_NONE is a no-op\n",
			nuc.IfName)
		return
	case types.DT_CLIENT:
		for dhcpcdExists(log, nuc.IfName) {
			log.Warnf("dhcpcd %s already exists", nuc.IfName)
			time.Sleep(1 * time.Second)
		}
		log.Functionf("dhcpcd %s not running", nuc.IfName)
		extras := []string{"-f", "/dhcpcd.conf", "--noipv4ll", "-b", "-t", "0"}
		switch nuc.Type {
		case types.NtIpv4Only:
			extras = []string{"-f", "/dhcpcd.conf", "--noipv4ll", "--ipv4only", "-b", "-t", "0"}
		case types.NtIpv6Only:
			extras = []string{"-f", "/dhcpcd.conf", "--ipv6only", "-b", "-t", "0"}
		case types.NT_NOOP:
		case types.NT_IPV4:
		case types.NT_IPV6:
		case types.NtDualStack:
		default:
		}
		if nuc.Gateway != nil && nuc.Gateway.String() == "0.0.0.0" {
			extras = append(extras, "--nogateway")
		}
		if !dhcpcdCmd(log, "--request", extras, nuc.IfName, true) {
			log.Errorf("doDhcpClientActivate: request failed for %s\n",
				nuc.IfName)
		}
		// Wait for a bit then give up
		waitCount := 0
		failed := false
		for !dhcpcdExists(log, nuc.IfName) {
			log.Warnf("dhcpcd %s not yet running", nuc.IfName)
			waitCount++
			if waitCount >= 3 {
				log.Errorf("dhcpcd %s not yet running", nuc.IfName)
				failed = true
				break
			}
			time.Sleep(1 * time.Second)
		}
		if !failed {
			log.Functionf("dhcpcd %s is running", nuc.IfName)
		}

	case types.DT_STATIC:
		if nuc.AddrSubnet == "" {
			log.Errorf("doDhcpClientActivate: missing AddrSubnet for %s\n",
				nuc.IfName)
			return
		}
		// Check that we can parse it
		_, _, err := net.ParseCIDR(nuc.AddrSubnet)
		if err != nil {
			log.Errorf("doDhcpClientActivate: failed to parse %s for %s: %s\n",
				nuc.AddrSubnet, nuc.IfName, err)
			return
		}
		for dhcpcdExists(log, nuc.IfName) {
			log.Warnf("dhcpcd %s already exists", nuc.IfName)
			time.Sleep(1 * time.Second)
		}
		log.Functionf("dhcpcd %s not running", nuc.IfName)
		args := []string{fmt.Sprintf("ip_address=%s", nuc.AddrSubnet)}

		extras := []string{"-f", "/dhcpcd.conf", "-b", "-t", "0"}
		if nuc.Gateway == nil || nuc.Gateway.String() == "0.0.0.0" {
			extras = append(extras, "--nogateway")
		} else if nuc.Gateway.String() != "" {
			args = append(args, "--static",
				fmt.Sprintf("routers=%s", nuc.Gateway.String()))
		}
		// XXX do we need to calculate a list for option?
		for _, dns := range nuc.DnsServers {
			args = append(args, "--static",
				fmt.Sprintf("domain_name_servers=%s", dns.String()))
		}
		if nuc.DomainName != "" {
			args = append(args, "--static",
				fmt.Sprintf("domain_name=%s", nuc.DomainName))
		}
		if nuc.NtpServer != nil && !nuc.NtpServer.IsUnspecified() {
			args = append(args, "--static",
				fmt.Sprintf("ntp_servers=%s",
					nuc.NtpServer.String()))
		}

		args = append(args, extras...)
		if !dhcpcdCmd(log, "--static", args, nuc.IfName, true) {
			log.Errorf("doDhcpClientActivate: request failed for %s\n",
				nuc.IfName)
		}
		// Wait for a bit then give up
		waitCount := 0
		failed := false
		for !dhcpcdExists(log, nuc.IfName) {
			log.Warnf("dhcpcd %s not yet running", nuc.IfName)
			waitCount++
			if waitCount >= 3 {
				log.Errorf("dhcpcd %s not yet running", nuc.IfName)
				failed = true
				break
			}
			time.Sleep(1 * time.Second)
		}
		if !failed {
			log.Functionf("dhcpcd %s is running", nuc.IfName)
		}
	default:
		log.Errorf("doDhcpClientActivate: unsupported dhcp %v\n",
			nuc.Dhcp)
	}
}

func doDhcpClientInactivate(log *base.LogObject, nuc types.NetworkPortConfig) {

	log.Functionf("doDhcpClientInactivate(%s) dhcp %v addr %s gateway %s\n",
		nuc.IfName, nuc.Dhcp, nuc.AddrSubnet,
		nuc.Gateway.String())
	if strings.HasPrefix(nuc.IfName, "wwan") {
		log.Functionf("doDhcpClientInactivate: skipping %s\n",
			nuc.IfName)
		return
	}
	switch nuc.Dhcp {
	case types.DT_NONE:
		log.Functionf("doDhcpClientInactivate(%s) DT_NONE is a no-op\n",
			nuc.IfName)
	case types.DT_STATIC, types.DT_CLIENT:
		startTime := time.Now()
		var extras []string
		// run release, wait for a bit, then exit and give up
		waitCount := 0
		failed := false
		for {
			//it waits up to 10 seconds https://github.com/NetworkConfiguration/dhcpcd/blob/dhcpcd-8.1.6/src/dhcpcd.c#L1950-L1957
			if !dhcpcdCmd(log, "--release", extras, nuc.IfName, false) {
				log.Errorf("doDhcpClientInactivate: release failed for %s, elapsed time %v", nuc.IfName, time.Since(startTime))
			}
			if !dhcpcdExists(log, nuc.IfName) {
				break
			}
			waitCount++
			if waitCount >= 3 {
				log.Errorf("doDhcpClientInactivate: dhcpcd %s still running, will exit it, elapsed time %v", nuc.IfName, time.Since(startTime))
				failed = true
				break
			}
			log.Warnf("doDhcpClientInactivate: dhcpcd %s still running, elapsed time %v", nuc.IfName, time.Since(startTime))
			time.Sleep(1 * time.Second)
		}
		if !failed {
			log.Functionf("dhcpcd %s gone, elapsed time %v", nuc.IfName, time.Since(startTime))
			return
		}
		//exit dhcpcd on interface
		//it waits up to 10 seconds https://github.com/NetworkConfiguration/dhcpcd/blob/dhcpcd-8.1.6/src/dhcpcd.c#L1950-L1957
		if !dhcpcdCmd(log, "--exit", extras, nuc.IfName, false) {
			log.Errorf("doDhcpClientInactivate: exit failed for %s, elapsed time %v", nuc.IfName, time.Since(startTime))
		}
		if !dhcpcdExists(log, nuc.IfName) {
			log.Noticef("dhcpcd %s gone after exit, elapsed time %v", nuc.IfName, time.Since(startTime))
			return
		}
		log.Errorf("doDhcpClientInactivate: exiting dhcpcd %s still running, elapsed time %v", nuc.IfName, time.Since(startTime))
	default:
		log.Errorf("doDhcpClientInactivate: unsupported dhcp %v\n",
			nuc.Dhcp)
	}
}

func dhcpcdCmd(log *base.LogObject, op string, extras []string, ifname string, background bool) bool {
	name := "/sbin/dhcpcd"
	args := append([]string{op}, extras...)
	args = append(args, ifname)
	if background {
		cmd := exec.Command(name, args...)
		cmd.Stdout = nil
		cmd.Stderr = nil

		log.Functionf("Background command %s %v\n", name, args)
		log.Functionf("Creating %s at %s", "func", agentlog.GetMyStack())
		go func() {
			if err := cmd.Run(); err != nil {
				log.Errorf("%s %v: failed: %s",
					name, args, err)
			}
		}()
	} else {
		log.Functionf("Calling command %s %v\n", name, args)
		out, err := base.Exec(log, name, args...).CombinedOutput()
		if err != nil {
			errStr := fmt.Sprintf("dhcpcd command %s failed %s output %s",
				args, err, out)
			log.Errorln(errStr)
			return false
		}
	}
	return true
}

func dhcpcdExists(log *base.LogObject, ifname string) bool {

	log.Functionf("dhcpcdExists(%s)", ifname)
	// XXX should we use dhcpcd -P <ifname> to get name of pidfile? Hardcoded path here
	pidfileName := fmt.Sprintf("/run/dhcpcd-%s.pid", ifname)
	val, t := statAndRead(pidfileName)
	if val == "" {
		log.Functionf("dhcpcdExists(%s) not exist", ifname)
		return false
	}
	log.Functionf("dhcpcdExists(%s) found modtime %v", ifname, t)

	pid, err := strconv.Atoi(strings.TrimSpace(val))
	if err != nil {
		log.Errorf("Atoi of %s failed %s; ignored\n", val, err)
		return true // Guess since we dont' know
	}
	// Does the pid exist?
	p, err := os.FindProcess(pid)
	if err != nil {
		log.Functionf("dhcpcdExists(%s) pid %d not found: %s", ifname, pid,
			err)
		return false
	}
	err = p.Signal(syscall.Signal(0))
	if err != nil {
		log.Errorf("dhcpcdExists(%s) Signal failed %s", ifname, err)
		return false
	}
	log.Functionf("dhcpcdExists(%s) Signal 0 OK for %d", ifname, pid)
	return true
}

// Returns content and Modtime
func statAndRead(filename string) (string, time.Time) {
	fi, err := os.Stat(filename)
	if err != nil {
		// File doesn't exist
		return "", time.Time{}
	}
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", fi.ModTime()
	}
	return string(content), fi.ModTime()
}
