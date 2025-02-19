// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build linux

package probes

import manager "github.com/DataDog/ebpf-manager"

// mprotectProbes holds the list of probes used to track mprotect events
var mprotectProbes = []*manager.Probe{
	{
		ProbeIdentificationPair: manager.ProbeIdentificationPair{
			UID:          SecurityAgentUID,
			EBPFFuncName: "hook_security_file_mprotect",
		},
	},
}

func getMProtectProbes(fentry bool) []*manager.Probe {
	mprotectProbes = append(mprotectProbes, ExpandSyscallProbes(&manager.Probe{
		ProbeIdentificationPair: manager.ProbeIdentificationPair{
			UID: SecurityAgentUID,
		},
		SyscallFuncName: "mprotect",
	}, fentry, EntryAndExit|SupportFentry)...)
	return mprotectProbes
}
