/*
 * Copyright 2019-2020 by Nedim Sabic Sabic
 * https://www.fibratus.io
 * All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ktypes

import (
	"encoding/binary"
	"github.com/rabbitstack/fibratus/pkg/sys/etw"
	"github.com/rabbitstack/fibratus/pkg/util/hashers"
	"golang.org/x/sys/windows"
)

// ProvidersCount designates the number of interesting providers.
// Remember to increment if a new event source is introduced.
const ProvidersCount = 11

// EventSource is the type that designates the provenance of the event
type EventSource uint8

const (
	// SystemLogger event is emitted by the system provider
	SystemLogger EventSource = iota
	// AuditAPICallsLogger event is emitted by Audit API calls provider
	AuditAPICallsLogger
	// DNSLogger event is emitted by DNS provider
	DNSLogger
)

// Ktype identifies an event type. It comprises the event GUID + hook ID to uniquely identify the event
type Ktype [18]byte

var (
	// ProcessEventGUID represents process provider event GUID
	ProcessEventGUID = windows.GUID{Data1: 0x3d6fa8d0, Data2: 0xfe05, Data3: 0x11d0, Data4: [8]byte{0x9d, 0xda, 0x0, 0xc0, 0x4f, 0xd7, 0xba, 0x7c}}
	// ThreadEventGUID represents thread provider event GUID
	ThreadEventGUID = windows.GUID{Data1: 0x3d6fa8d1, Data2: 0xfe05, Data3: 0x11d0, Data4: [8]byte{0x9d, 0xda, 0x0, 0xc0, 0x4f, 0xd7, 0xba, 0x7c}}
	// ImageEventGUID represents image provider event GUID
	ImageEventGUID = windows.GUID{Data1: 0x2cb15d1d, Data2: 0x5fc1, Data3: 0x11d2, Data4: [8]byte{0xab, 0xe1, 0x0, 0xa0, 0xc9, 0x11, 0xf5, 0x18}}
	// FileEventGUID represents file provider event GUID
	FileEventGUID = windows.GUID{Data1: 0x90cbdc39, Data2: 0x4a3e, Data3: 0x11d1, Data4: [8]byte{0x84, 0xf4, 0x0, 0x0, 0xf8, 0x04, 0x64, 0xe3}}
	// RegistryEventGUID represents registry provider event GUID
	RegistryEventGUID = windows.GUID{Data1: 0xae53722e, Data2: 0xc863, Data3: 0x11d2, Data4: [8]byte{0x86, 0x59, 0x0, 0xc0, 0x4f, 0xa3, 0x21, 0xa1}}
	// NetworkTCPEventGUID represents network TCP provider event GUID
	NetworkTCPEventGUID = windows.GUID{Data1: 0x9a280ac0, Data2: 0xc8e0, Data3: 0x11d1, Data4: [8]byte{0x84, 0xe2, 0x0, 0xc0, 0x4f, 0xb9, 0x98, 0xa2}}
	// NetworkUDPEventGUID represents network UDP provider event GUID
	NetworkUDPEventGUID = windows.GUID{Data1: 0xbf3a50c5, Data2: 0xa9c9, Data3: 0x4988, Data4: [8]byte{0xa0, 0x05, 0x2d, 0xf0, 0xb7, 0xc8, 0x0f, 0x80}}
	// HandleEventGUID represents handle provider event GUID
	HandleEventGUID = windows.GUID{Data1: 0x89497f50, Data2: 0xeffe, Data3: 0x4440, Data4: [8]byte{0x8c, 0xf2, 0xce, 0x6b, 0x1c, 0xdc, 0xac, 0xa7}}
	// MemEventGUID represents memory provider event GUID
	MemEventGUID = windows.GUID{Data1: 0x3d6fa8d3, Data2: 0xfe05, Data3: 0x11d0, Data4: [8]byte{0x9d, 0xda, 0x00, 0xc0, 0x4f, 0xd7, 0xba, 0x7c}}
	// AuditAPIEventGUID represents audit API calls event GUID
	AuditAPIEventGUID = windows.GUID{Data1: 0xe02a841c, Data2: 0x75a3, Data3: 0x4fa7, Data4: [8]byte{0xaf, 0xc8, 0xae, 0x09, 0xcf, 0x9b, 0x7f, 0x23}}
	// DNSEventGUID represents DNS provider event GUID
	DNSEventGUID = windows.GUID{Data1: 0x1c95126e, Data2: 0x7eea, Data3: 0x49a9, Data4: [8]byte{0xa3, 0xfe, 0xa3, 0x78, 0xb0, 0x3d, 0xdb, 0x4d}}
)

var (
	// CreateProcess identifies process creation kernel events
	CreateProcess = pack(ProcessEventGUID, 1)
	// TerminateProcess identifies process termination kernel events
	TerminateProcess = pack(ProcessEventGUID, 2)
	// ProcessRundown represents the start data collection process event that enumerates processes that are currently running at the time the kernel session starts
	ProcessRundown = pack(ProcessEventGUID, 3)
	// OpenProcess identifies the kernel events that are triggered when the process handle is acquired
	OpenProcess = pack(AuditAPIEventGUID, 5)

	// CreateThread identifies thread creation kernel events
	CreateThread = pack(ThreadEventGUID, 1)
	// TerminateThread identifies thread termination kernel events
	TerminateThread = pack(ThreadEventGUID, 2)
	// ThreadRundown represents the start data collection thread event that enumerates threads that are currently running at the time the kernel session starts
	ThreadRundown = pack(ThreadEventGUID, 3)
	// OpenThread identifies the kernel events that are triggered when the process acquires a thread handle
	OpenThread = pack(AuditAPIEventGUID, 6)
	// SetThreadContext identifies the kernel event that is fired when the thread context is changed
	SetThreadContext = pack(AuditAPIEventGUID, 4)

	// MapViewFile represents events that map a view of a file mapping into the address space of a calling process
	MapViewFile = pack(FileEventGUID, 37)
	// UnmapViewFile represents events that unmap a view of a file mapping from the address space of a calling process
	UnmapViewFile = pack(FileEventGUID, 38)
	// MapFileRundown represents the event that is emitted at the start of the tracing session to enumerate I/O mapped files
	MapFileRundown = pack(FileEventGUID, 39)

	// FileRundown events are generated by kernel rundown logger to enumerate all open files on the start of the kernel session
	FileRundown = pack(FileEventGUID, 36)
	// CreateFile represents events that create/open a file or I/O device
	CreateFile = pack(FileEventGUID, 64)
	// ReleaseFile represents events that occur when the last file handle is disposed
	ReleaseFile = pack(FileEventGUID, 65)
	// CloseFile represents events that dispose existing kernel file objects
	CloseFile = pack(FileEventGUID, 66)
	// ReadFile represents events that read data from the file or I/O device
	ReadFile = pack(FileEventGUID, 67)
	// WriteFile represents events that write data to the file or I/O device
	WriteFile = pack(FileEventGUID, 68)
	// SetFileInformation represents events that set file information
	SetFileInformation = pack(FileEventGUID, 69)
	// DeleteFile identifies file deletion events
	DeleteFile = pack(FileEventGUID, 70)
	// RenameFile identifies events that are responsible for renaming files
	RenameFile = pack(FileEventGUID, 71)
	// EnumDirectory identifies enumerate directory and directory notification events
	EnumDirectory = pack(FileEventGUID, 72)
	// FileOpEnd signals the finalization of the file operation
	FileOpEnd = pack(FileEventGUID, 76)

	// RegCreateKey represents registry key creation kernel events
	RegCreateKey = pack(RegistryEventGUID, 10)
	// RegOpenKey represents registry open key kernel events
	RegOpenKey = pack(RegistryEventGUID, 11)
	// RegCloseKey represents registry close key kernel event.
	RegCloseKey = pack(RegistryEventGUID, 27)
	// RegDeleteKey represents registry key deletion kernel events
	RegDeleteKey = pack(RegistryEventGUID, 12)
	// RegQueryKey represents registry query key kernel events
	RegQueryKey = pack(RegistryEventGUID, 13)
	// RegSetValue represents registry set value kernel events
	RegSetValue = pack(RegistryEventGUID, 14)
	// RegDeleteValue are kernel events for registry value removals
	RegDeleteValue = pack(RegistryEventGUID, 15)
	// RegQueryValue are kernel events for registry value queries
	RegQueryValue = pack(RegistryEventGUID, 16)
	// RegCreateKCB represents kernel events for KCB (Key Control Block) creation requests
	RegCreateKCB = pack(RegistryEventGUID, 22)
	// RegDeleteKCB represents kernel events for KCB(Key Control Block) closures
	RegDeleteKCB = pack(RegistryEventGUID, 23)
	// RegKCBRundown enumerates the registry keys open at the start of the kernel session.
	RegKCBRundown = pack(RegistryEventGUID, 25)

	// UnloadImage represents unload image kernel events
	UnloadImage = pack(ImageEventGUID, 2)
	// ImageRundown represents kernel events that is triggered to enumerate all loaded images
	ImageRundown = pack(ImageEventGUID, 3)
	// LoadImage represents load image kernel events that are triggered when a DLL or executable file  is loaded
	LoadImage = pack(ImageEventGUID, 10)

	// AcceptTCPv4 represents the TCPv4 kernel events for accepting connection requests from the socket queue.
	AcceptTCPv4 = pack(NetworkTCPEventGUID, 15)
	// AcceptTCPv6 represents the TCPv6 kernel events for accepting connection requests from the socket queue.
	AcceptTCPv6 = pack(NetworkTCPEventGUID, 31)
	// SendTCPv4 represents the TCPv4 kernel events for sending data to the connected socket.
	SendTCPv4 = pack(NetworkTCPEventGUID, 10)
	// SendTCPv6 represents the TCPv6 kernel events for sending data to the connected socket.
	SendTCPv6 = pack(NetworkTCPEventGUID, 26)
	// SendUDPv4 represents the UDPv4 kernel events for sending datagrams to connectionless sockets.
	SendUDPv4 = pack(NetworkUDPEventGUID, 10)
	// SendUDPv6 represents the UDPv6 kernel events for sending datagrams to connectionless sockets.
	SendUDPv6 = pack(NetworkUDPEventGUID, 26)
	// RecvTCPv4 represents the TCP IPv4 network receive event.
	RecvTCPv4 = pack(NetworkTCPEventGUID, 11)
	// RecvTCPv6 represents the TCP IPv6 network receive event.
	RecvTCPv6 = pack(NetworkTCPEventGUID, 27)
	// RecvUDPv4 represents the UDP IPv4 network receive event.
	RecvUDPv4 = pack(NetworkUDPEventGUID, 11)
	// RecvUDPv6 represents the UDP IPv6 network receive event.
	RecvUDPv6 = pack(NetworkUDPEventGUID, 27)
	// ConnectTCPv4 represents the TCP IPv4 network connect event.
	ConnectTCPv4 = pack(NetworkTCPEventGUID, 12)
	// ConnectTCPv6 represents the TCP IPv6 network connect event.
	ConnectTCPv6 = pack(NetworkTCPEventGUID, 28)
	// DisconnectTCPv4 is the TCP IPv4 network disconnect event.
	DisconnectTCPv4 = pack(NetworkTCPEventGUID, 13)
	// DisconnectTCPv6 is the TCP IPv6 network disconnect event.
	DisconnectTCPv6 = pack(NetworkTCPEventGUID, 29)
	// ReconnectTCPv4 is the TCP IPv4 network reconnect event.
	ReconnectTCPv4 = pack(NetworkTCPEventGUID, 16)
	// ReconnectTCPv6 is the TCP IPv6 network reconnect event.
	ReconnectTCPv6 = pack(NetworkTCPEventGUID, 32)
	// RetransmitTCPv4 is the TCP IPv4 network retransmit event.
	RetransmitTCPv4 = pack(NetworkTCPEventGUID, 14)
	// RetransmitTCPv6 is the TCP IPv6 network retransmit event.
	RetransmitTCPv6 = pack(NetworkTCPEventGUID, 30)

	// CreateHandle represents handle creation event
	CreateHandle = pack(HandleEventGUID, 32)
	// CloseHandle represents handle closure event
	CloseHandle = pack(HandleEventGUID, 33)
	// DuplicateHandle represents handle duplication event
	DuplicateHandle = pack(HandleEventGUID, 34)

	// VirtualAlloc represents virtual memory allocation event
	VirtualAlloc = pack(MemEventGUID, 98)
	// VirtualFree represents virtual memory release event
	VirtualFree = pack(MemEventGUID, 99)

	// QueryDNS represents DNS query events
	QueryDNS = pack(DNSEventGUID, 3006)
	// ReplyDNS represents the DNS response events
	ReplyDNS = pack(DNSEventGUID, 3008)

	// StackWalk represents stack walk event with the collection of return addresses
	StackWalk = pack(windows.GUID{Data1: 0xdef2fe46, Data2: 0x7bd6, Data3: 0x4b80, Data4: [8]byte{0xbd, 0x94, 0xf5, 0x7f, 0xe2, 0x0d, 0x0c, 0xe3}}, 32)

	// UnknownKtype designates unknown kernel event type
	UnknownKtype = pack(windows.GUID{}, 0)
)

// NewFromEventRecord creates a new event type from ETW event record.
func NewFromEventRecord(ev *etw.EventRecord) Ktype {
	return pack(ev.Header.ProviderID, ev.HookID())
}

// String returns the string representation of the event type. Returns an empty string
// if the event type is not recognized.
func (k Ktype) String() string {
	switch k {
	case CreateProcess:
		return "CreateProcess"
	case TerminateProcess:
		return "TerminateProcess"
	case ProcessRundown:
		return "ProcessRundown"
	case OpenProcess:
		return "OpenProcess"
	case CreateThread:
		return "CreateThread"
	case TerminateThread:
		return "TerminateThread"
	case ThreadRundown:
		return "ThreadRundown"
	case OpenThread:
		return "OpenThread"
	case SetThreadContext:
		return "SetThreadContext"
	case CreateFile:
		return "CreateFile"
	case CloseFile:
		return "CloseFile"
	case ReleaseFile:
		return "ReleaseFile"
	case ReadFile:
		return "ReadFile"
	case WriteFile:
		return "WriteFile"
	case SetFileInformation:
		return "SetFileInformation"
	case DeleteFile:
		return "DeleteFile"
	case RenameFile:
		return "RenameFile"
	case EnumDirectory:
		return "EnumDirectory"
	case FileOpEnd:
		return "FileOpEnd"
	case FileRundown:
		return "FileRundown"
	case MapViewFile:
		return "MapViewFile"
	case UnmapViewFile:
		return "UnmapViewFile"
	case MapFileRundown:
		return "MapFileRundown"
	case CreateHandle:
		return "CreateHandle"
	case CloseHandle:
		return "CloseHandle"
	case DuplicateHandle:
		return "DuplicateHandle"
	case RegKCBRundown:
		return "RegKCBRundown"
	case RegOpenKey:
		return "RegOpenKey"
	case RegCloseKey:
		return "RegCloseKey"
	case RegCreateKey:
		return "RegCreateKey"
	case RegDeleteKey:
		return "RegDeleteKey"
	case RegDeleteValue:
		return "RegDeleteValue"
	case RegQueryKey:
		return "RegQueryKey"
	case RegQueryValue:
		return "RegQueryValue"
	case RegCreateKCB:
		return "RegCreateKCB"
	case RegSetValue:
		return "RegSetValue"
	case LoadImage:
		return "LoadImage"
	case UnloadImage:
		return "UnloadImage"
	case ImageRundown:
		return "ImageRundown"
	case AcceptTCPv4, AcceptTCPv6:
		return "Accept"
	case SendTCPv4, SendTCPv6, SendUDPv4, SendUDPv6:
		return "Send"
	case RecvTCPv4, RecvTCPv6, RecvUDPv4, RecvUDPv6:
		return "Recv"
	case ConnectTCPv4, ConnectTCPv6:
		return "Connect"
	case ReconnectTCPv4, ReconnectTCPv6:
		return "Reconnect"
	case DisconnectTCPv4, DisconnectTCPv6:
		return "Disconnect"
	case RetransmitTCPv4, RetransmitTCPv6:
		return "Retransmit"
	case VirtualAlloc:
		return "VirtualAlloc"
	case VirtualFree:
		return "VirtualFree"
	case QueryDNS:
		return "QueryDns"
	case ReplyDNS:
		return "ReplyDns"
	case StackWalk:
		return "StackWalk"
	default:
		return ""
	}
}

// Category determines the category to which the event type pertains.
func (k Ktype) Category() Category {
	switch k {
	case CreateProcess, TerminateProcess, OpenProcess, ProcessRundown:
		return Process
	case CreateThread, TerminateThread, OpenThread, SetThreadContext, ThreadRundown, StackWalk:
		return Thread
	case LoadImage, UnloadImage, ImageRundown:
		return Image
	case CreateFile, ReadFile, WriteFile, EnumDirectory, DeleteFile, RenameFile, CloseFile, SetFileInformation,
		FileRundown, FileOpEnd, ReleaseFile, MapViewFile, UnmapViewFile, MapFileRundown:
		return File
	case RegCreateKey, RegDeleteKey, RegOpenKey, RegCloseKey, RegQueryKey, RegQueryValue, RegSetValue, RegDeleteValue,
		RegKCBRundown, RegDeleteKCB, RegCreateKCB:
		return Registry
	case AcceptTCPv4, AcceptTCPv6,
		ConnectTCPv4, ConnectTCPv6,
		ReconnectTCPv4, ReconnectTCPv6,
		RetransmitTCPv4, RetransmitTCPv6,
		DisconnectTCPv4, DisconnectTCPv6,
		SendTCPv4, SendTCPv6, SendUDPv4, SendUDPv6,
		RecvTCPv4, RecvTCPv6, RecvUDPv4, RecvUDPv6,
		QueryDNS, ReplyDNS:
		return Net
	case CreateHandle, CloseHandle, DuplicateHandle:
		return Handle
	case VirtualAlloc, VirtualFree:
		return Mem
	default:
		return Unknown
	}
}

// Subcategory determines the event subcategory, if any.
func (k Ktype) Subcategory() Subcategory {
	switch k {
	case QueryDNS, ReplyDNS:
		return DNS
	default:
		return None
	}
}

// Description returns a brief description of the event type.
func (k Ktype) Description() string {
	switch k {
	case CreateProcess:
		return "Creates a new process and its primary thread"
	case TerminateProcess:
		return "Terminates the process and all of its threads"
	case OpenProcess:
		return "Opens the process handle"
	case CreateThread:
		return "Creates a thread to execute within the virtual address space of the calling process"
	case TerminateThread:
		return "Terminates a thread within the process"
	case OpenThread:
		return "Opens the thread handle"
	case SetThreadContext:
		return "Sets the thread context"
	case ReadFile:
		return "Reads data from the file or I/O device"
	case WriteFile:
		return "Writes data to the file or I/O device"
	case CreateFile:
		return "Creates or opens a file or I/O device"
	case CloseFile:
		return "Closes the file handle"
	case DeleteFile:
		return "Removes the file from the file system"
	case RenameFile:
		return "Changes the file name"
	case SetFileInformation:
		return "Sets the file meta information"
	case EnumDirectory:
		return "Enumerates a directory or dispatches a directory change notification to registered listeners"
	case MapViewFile:
		return "Maps a view of a file mapping into the address space of a calling process"
	case UnmapViewFile:
		return "Unmaps a mapped view of a file from the calling process's address space"
	case RegCreateKey:
		return "Creates a registry key or opens it if the key already exists"
	case RegOpenKey:
		return "Opens the registry key"
	case RegCloseKey:
		return "Closes the registry key"
	case RegSetValue:
		return "Sets the data for the value of a registry key"
	case RegQueryValue:
		return "Reads the data for the value of a registry key"
	case RegQueryKey:
		return "Enumerates subkeys of the parent key"
	case RegDeleteKey:
		return "Removes the registry key"
	case RegDeleteValue:
		return "Removes the registry value"
	case AcceptTCPv4, AcceptTCPv6:
		return "Accepts the connection request from the socket queue"
	case ConnectTCPv4, ConnectTCPv6:
		return "Connects establishes a connection to the socket"
	case DisconnectTCPv4, DisconnectTCPv6:
		return "Terminates data reception on the socket"
	case ReconnectTCPv4, ReconnectTCPv6:
		return "Reconnects to the socket"
	case RetransmitTCPv4, RetransmitTCPv6:
		return "Retransmits unacknowledged TCP segments"
	case SendTCPv4, SendUDPv4, SendTCPv6, SendUDPv6:
		return "Sends data over the wire"
	case RecvTCPv4, RecvUDPv4, RecvTCPv6, RecvUDPv6:
		return "Receives data from the socket"
	case LoadImage:
		return "Loads the module into the address space of the calling process"
	case UnloadImage:
		return "Unloads the module from the address space of the calling process"
	case CreateHandle:
		return "Creates a new handle"
	case CloseHandle:
		return "Closes the handle"
	case DuplicateHandle:
		return "Duplicates the handle"
	case VirtualAlloc:
		return "Reserves, commits, or changes the state of a region of memory within the process virtual address space"
	case VirtualFree:
		return "Releases or decommits a region of memory within the process virtual address space"
	case QueryDNS:
		return "Sends a DNS query to the name server"
	case ReplyDNS:
		return "Receives the response from the DNS server"
	default:
		return ""
	}
}

// Hash calculates the hash number of the event type.
func (k Ktype) Hash() uint32 {
	if k == UnknownKtype {
		return 0
	}
	return hashers.FnvUint32([]byte(k.String()))
}

// Exists determines whether particular event type exists.
func (k Ktype) Exists() bool {
	return k.String() != ""
}

// OnlyState determines whether the event type is solely used for state management.
func (k Ktype) OnlyState() bool {
	switch k {
	case ProcessRundown,
		ThreadRundown,
		ImageRundown,
		FileRundown,
		RegKCBRundown,
		FileOpEnd,
		ReleaseFile,
		MapFileRundown,
		RegCreateKCB,
		RegDeleteKCB:
		return true
	default:
		return false
	}
}

// CanEnrichStack determines if the event can be enriched with a callstack.
func (k Ktype) CanEnrichStack() bool {
	switch k {
	case CreateProcess,
		CreateThread,
		TerminateThread,
		LoadImage,
		RegCreateKey,
		RegDeleteKey,
		RegSetValue,
		RegDeleteValue,
		DeleteFile,
		RenameFile:
		return true
	default:
		return false
	}
}

// UnmarshalYAML converts the ktype name to ktype array type.
func (k *Ktype) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var ktyp string
	err := unmarshal(&ktyp)
	if err != nil {
		return err
	}
	*k = KeventNameToKtype(ktyp)
	return nil
}

// GUID returns the event GUID from the raw ktype.
func (k *Ktype) GUID() windows.GUID {
	return windows.GUID{
		Data1: binary.BigEndian.Uint32(k[0:4]),
		Data2: binary.BigEndian.Uint16(k[4:6]),
		Data3: binary.BigEndian.Uint16(k[6:8]),
		Data4: [8]byte{k[8], k[9], k[10], k[11], k[12], k[13], k[14], k[15]},
	}
}

// HookID returns the event operation code (hook ID) from the raw ktype.
func (k *Ktype) HookID() uint16 {
	return binary.BigEndian.Uint16(k[16:])
}

// Source designates the provenance of this event type.
func (k Ktype) Source() EventSource {
	switch k {
	case OpenProcess, OpenThread, SetThreadContext:
		return AuditAPICallsLogger
	case QueryDNS, ReplyDNS:
		return DNSLogger
	default:
		return SystemLogger
	}
}

// pack merges event provider GUID and the hook ID into `Ktype` array.
// The type provides a convenient way for comparing event types.
func pack(g windows.GUID, id uint16) Ktype {
	return [18]byte{
		byte(g.Data1 >> 24), byte(g.Data1 >> 16), byte(g.Data1 >> 8), byte(g.Data1),
		byte(g.Data2 >> 8), byte(g.Data2),
		byte(g.Data3 >> 8), byte(g.Data3),
		g.Data4[0],
		g.Data4[1],
		g.Data4[2],
		g.Data4[3],
		g.Data4[4],
		g.Data4[5],
		g.Data4[6],
		g.Data4[7],
		byte(id >> 8), byte(id),
	}
}
