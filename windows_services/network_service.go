package windows_services

/*
#include <winsock2.h>
#include <ws2ipdef.h>
#include <iphlpapi.h>
#include <stdio.h>

#pragma comment(lib, "iphlpapi.lib")

#define MALLOC(x) HeapAlloc(GetProcessHeap(), 0, (x))
#define FREE(x) HeapFree(GetProcessHeap(), 0, (x))

void test() {
// Declare and initialize variables
    PIP_INTERFACE_INFO pInfo = NULL;
    ULONG ulOutBufLen = 0;

    DWORD dwRetVal = 0;
    int iReturn = 1;

    int i;

// Make an initial call to GetInterfaceInfo to get
// the necessary size in the ulOutBufLen variable


    dwRetVal = GetInterfaceInfo(NULL, &ulOutBufLen);
    if (dwRetVal == ERROR_INSUFFICIENT_BUFFER) {
        pInfo = (IP_INTERFACE_INFO *) MALLOC(ulOutBufLen);
        if (pInfo == NULL) {
            printf
                ("Unable to allocate memory needed to call GetInterfaceInfo\n");
            return;
        }
    }
// Make a second call to GetInterfaceInfo to get
// the actual data we need
    dwRetVal = GetInterfaceInfo(pInfo, &ulOutBufLen);
    if (dwRetVal == NO_ERROR) {
        printf("Number of Adapters: %ld\n\n", pInfo->NumAdapters);
        for (i = 0; i < pInfo->NumAdapters; i++) {
            printf("Adapter Index[%d]: %ld\n", i,
                   pInfo->Adapter[i].Index);
            printf("Adapter Name[%d]: %ws\n\n", i,
                   pInfo->Adapter[i].Name);
        }
        iReturn = 0;
    } else if (dwRetVal == ERROR_NO_DATA) {
        printf
            ("There are no network adapters with IPv4 enabled on the local system\n");
        iReturn = 0;
    } else {
        printf("GetInterfaceInfo failed with error: %d\n", dwRetVal);
        iReturn = 1;
    }

    FREE(pInfo);
	print(iReturn);
    return;
}
*/
import "C"

import (
	"golang.org/x/sys/windows"
	"os"
	syscall2 "syscall"
	"unsafe"
	"vm-init-utils/modules"
)

func (ws *WindowsService) SetNetWork(conf *modules.Network) error {

	return nil
}

func GetMACAddress() (string, error) {
	C.test()
	//netInterfaces, err := net.Interfaces()
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//mac, macerr := "", errors.New("no valid mac address")
	//for i := 0; i < len(netInterfaces); i++ {
	//	flagStr := netInterfaces[i].Flags.String()
	//	if (netInterfaces[i].Flags & net.FlagLoopback) == 0 {
	//		if strings.Contains(flagStr, "up") {
	//			if strings.Contains(flagStr, "broadcast") {
	//				if strings.Contains(flagStr, "multicast") {
	//					index := netInterfaces[i].Index
	//					if isEthernet(index) {
	//						mac = netInterfaces[i].HardwareAddr.String()
	//						fmt.Println(flagStr)
	//						//return mac, nil
	//					}
	//				}
	//			}
	//		}
	//	}
	//}
	//return mac, macerr
	return "", nil
}

// 根据网卡接口 Index 判断其是否为 Ethernet 网卡
func isEthernet(ifindex int) bool {
	aas, err := adapterAddresses()
	if err != nil {
		return false
	}
	result := false
	for _, aa := range aas {
		index := aa.IfIndex
		if ifindex == int(index) {
			switch aa.IfType {
			case windows.IF_TYPE_ETHERNET_CSMACD:
				result = true
			}

			if result {
				break
			}
		}
	}
	return result
}

// 从 net/interface_windows.go 中复制过来
func adapterAddresses() ([]*windows.IpAdapterAddresses, error) {
	var b []byte
	l := uint32(15000) // recommended initial size
	for {
		b = make([]byte, l)
		err := windows.GetAdaptersAddresses(syscall2.AF_UNSPEC, windows.GAA_FLAG_INCLUDE_PREFIX, 0, (*windows.IpAdapterAddresses)(unsafe.Pointer(&b[0])), &l)
		if err == nil {
			if l == 0 {
				return nil, nil
			}
			break
		}
		if err.(syscall2.Errno) != syscall2.ERROR_BUFFER_OVERFLOW {
			return nil, os.NewSyscallError("getadaptersaddresses", err)
		}
		if l <= uint32(len(b)) {
			return nil, os.NewSyscallError("getadaptersaddresses", err)
		}
	}
	var aas []*windows.IpAdapterAddresses
	for aa := (*windows.IpAdapterAddresses)(unsafe.Pointer(&b[0])); aa != nil; aa = aa.Next {
		aas = append(aas, aa)
	}
	return aas, nil
}

//func (w *modules.Windows) SetNetWork() error {
//	ipArgs := []string{
//		"interface",
//		"ip",
//		"set",
//		"address",
//		w.Network.Name,
//		"static",
//		w.Network.Addr,
//		w.Network.Mask,
//		w.Network.Gateway,
//	}
//	ipCmd := exec.Command("netsh", ipArgs...)
//	ipCmd.Stdin = os.Stdin
//	ipCmd.Stdout = os.Stdout
//	ipCmd.Stderr = os.Stderr
//	err := ipCmd.Run()
//	if err != nil {
//		return err
//	}
//	dnsArgs := []string{
//		"interface",
//		"ip",
//		"set",
//		"dns",
//		w.Network.Name,
//		"static",
//		w.Network.Gateway,
//	}
//	dnsCmd := exec.Command("netsh", dnsArgs...)
//	dnsCmd.Stdin = os.Stdin
//	dnsCmd.Stdout = os.Stdout
//	dnsCmd.Stderr = os.Stderr
//	return dnsCmd.Run()
//}
//
//func (w *Windows) SetPWD() error {
//	args := []string{
//		"user",
//		w.Sys.UserName,
//		w.Sys.PWD,
//	}
//	cmd := exec.Command("net", args...)
//	cmd.Stdin = os.Stdin
//	cmd.Stdout = os.Stdout
//	cmd.Stderr = os.Stderr
//	return cmd.Run()
//}
