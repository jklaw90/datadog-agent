// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build linux

package probes

import manager "github.com/DataDog/ebpf-manager"

// xattrProbes holds the list of probes used to track xattr events
var xattrProbes = []*manager.Probe{
	{
		ProbeIdentificationPair: manager.ProbeIdentificationPair{
			UID:          SecurityAgentUID,
			EBPFFuncName: "kprobe_vfs_setxattr",
		},
	},
	{
		ProbeIdentificationPair: manager.ProbeIdentificationPair{
			UID:          SecurityAgentUID,
			EBPFFuncName: "kprobe_vfs_removexattr",
		},
	},
}

func getXattrProbes(fentry bool) []*manager.Probe {
	xattrProbes = append(xattrProbes, ExpandSyscallProbes(&manager.Probe{
		ProbeIdentificationPair: manager.ProbeIdentificationPair{
			UID: SecurityAgentUID,
		},
		SyscallFuncName: "setxattr",
	}, fentry, EntryAndExit|SupportFentry)...)
	xattrProbes = append(xattrProbes, ExpandSyscallProbes(&manager.Probe{
		ProbeIdentificationPair: manager.ProbeIdentificationPair{
			UID: SecurityAgentUID,
		},
		SyscallFuncName: "fsetxattr",
	}, fentry, EntryAndExit|SupportFentry)...)
	xattrProbes = append(xattrProbes, ExpandSyscallProbes(&manager.Probe{
		ProbeIdentificationPair: manager.ProbeIdentificationPair{
			UID: SecurityAgentUID,
		},
		SyscallFuncName: "lsetxattr",
	}, fentry, EntryAndExit|SupportFentry)...)
	xattrProbes = append(xattrProbes, ExpandSyscallProbes(&manager.Probe{
		ProbeIdentificationPair: manager.ProbeIdentificationPair{
			UID: SecurityAgentUID,
		},
		SyscallFuncName: "removexattr",
	}, fentry, EntryAndExit|SupportFentry)...)
	xattrProbes = append(xattrProbes, ExpandSyscallProbes(&manager.Probe{
		ProbeIdentificationPair: manager.ProbeIdentificationPair{
			UID: SecurityAgentUID,
		},
		SyscallFuncName: "fremovexattr",
	}, fentry, EntryAndExit|SupportFentry)...)
	xattrProbes = append(xattrProbes, ExpandSyscallProbes(&manager.Probe{
		ProbeIdentificationPair: manager.ProbeIdentificationPair{
			UID: SecurityAgentUID,
		},
		SyscallFuncName: "lremovexattr",
	}, fentry, EntryAndExit|SupportFentry)...)
	return xattrProbes
}
