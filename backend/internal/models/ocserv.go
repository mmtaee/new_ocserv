package models

import (
	"fmt"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
	"time"
)

const (
	Free            = "Free"
	MonthlyTransmit = "MonthlyTransmit"
	MonthlyReceive  = "MonthlyReceive"
	TotallyTransmit = "TotallyTransmit"
	TotallyReceive  = "TotallyReceive"
)

const (
	Connected    = "Connected"
	Disconnected = "Disconnected"
	Failed       = "Failed"
)

type OcservRuntimeConfig struct {
	MaxClients      *int    `json:"max-clients" desc:"Maximum number of concurrent client connections. Example: 1024"`
	MaxSameClients  *int    `json:"max-same-clients" desc:"Maximum simultaneous connections per user. Example: 2"`
	KeepAlive       *int    `json:"keepalive" desc:"Interval (in seconds) between keepalive pings. Example: 32400"`
	DPD             *int    `json:"dpd" desc:"Dead Peer Detection timeout in seconds. Example: 90"`
	MobileDPD       *int    `json:"mobile-dpd" desc:"DPD timeout for mobile clients. Example: 1800"`
	AuthTimeout     *int    `json:"auth-timeout" desc:"Timeout in seconds to complete authentication. Example: 240"`
	TryMTUDiscovery *bool   `json:"try-mtu-discovery" desc:"Enable MTU discovery to avoid fragmentation. Example: true"`
	PredictableIPs  *bool   `json:"predictable-ips" desc:"Assign predictable IP addresses based on user. Example: true"`
	TunnelAllDNS    *bool   `json:"tunnel-all-dns" desc:"Tunnel all DNS queries through the VPN. Example: true"`
	DNS             *string `json:"dns" desc:"Comma-separated list of DNS servers. Example: '8.8.8.8,1.1.1.1'"`
	PingLeases      *bool   `json:"ping-leases" desc:"Ping clients to test lease validity. Example: false"`
	MTU             *int    `json:"mtu" desc:"Maximum Transmission Unit for VPN tunnel. Example: 1420"`
	DenyRoaming     *bool   `json:"deny-roaming" desc:"Disconnect clients if their IP address changes. Example: false"`
}

type OcservUserOrGroupConfigs struct {
	DNS                  *string   `json:"dns" desc:"Comma-separated list of DNS servers to assign to the client. Example: '8.8.8.8,1.1.1.1'"`
	NBNS                 *string   `json:"nbns" desc:"Comma-separated list of NetBIOS Name Servers (WINS) for Windows clients. Example: '192.168.1.1'"`
	IPv4Netmask          *string   `json:"ipv4-netmask" desc:"Subnet mask for IPv4 network when not using CIDR. Example: '255.255.255.0'"`
	RxDataPerSec         *string   `json:"rx-data-per-sec" desc:"Maximum receive bandwidth in bytes per second. Example: '100000' for 100 KB/s"`
	TxDataPerSec         *string   `json:"tx-data-per-sec" desc:"Maximum transmit bandwidth in bytes per second. Example: '200000' for 200 KB/s"`
	ExplicitIPv4         *string   `json:"explicit-ipv4" desc:"Static IPv4 address to assign to client. Example: '192.168.100.10'"`
	NoRoute              *string   `json:"no-route" desc:"Network(s) to exclude from routing over VPN. Example: '192.168.0.0/16'"`
	CGroup               *string   `json:"cgroup" desc:"Linux control group to assign the VPN worker process to. Format: 'controller,subsystem:name'. Multiple controllers can be comma-separated. Example: 'cpuset,cpu:test' to assign to the 'test' cgroup under 'cpuset' and 'cpu' subsystems."`
	IRoute               *[]string `json:"iroute" desc:"List of internal routes available only via VPN. Example: ['10.0.0.0/8', '172.16.0.0/12']"`
	Route                *[]string `json:"route" desc:"Routes pushed to the client for routing traffic. Example: ['0.0.0.0/0', '10.10.0.0/16']"`
	NetPriority          *int      `json:"net-priority" desc:"Priority for routes; lower is higher priority. Example: 1"`
	DenyRoaming          *bool     `json:"deny-roaming" desc:"Disconnect client if its IP changes (e.g., due to network switch). Example: true"`
	NoUDP                *bool     `json:"no-udp" desc:"Disables UDP, enforcing TCP-only VPN connection. Example: true"`
	KeepAlive            *int      `json:"keepalive" desc:"Interval in seconds to send keep-alive pings. Example: 60"`
	DPD                  *int      `json:"dpd" desc:"Dead Peer Detection timeout in seconds. Example: 90"`
	MobileDPD            *int      `json:"mobile-dpd" desc:"DPD timeout specifically for mobile clients. Example: 300"`
	MaxSameClients       *int      `json:"max-same-clients" desc:"Maximum simultaneous logins per user. Example: 2"`
	TunnelAllDNS         *bool     `json:"tunnel-all-dns" desc:"Force all DNS traffic through the VPN tunnel. Example: true"`
	RestrictUserToRoutes *bool     `json:"restrict-user-to-routes" desc:"Allow client access only to defined routes. Example: true"`
	StatsReportTime      *int      `json:"stats-report-time" desc:"Interval in seconds for stats reporting. Example: 300"`
	MTU                  *int      `json:"mtu" desc:"Tunnel interface MTU to avoid fragmentation. Example: 1400"`
	IdleTimeout          *int      `json:"idle-timeout" desc:"Time in seconds before disconnecting idle clients. Example: 600"`
	MobileIdleTimeout    *int      `json:"mobile-idle-timeout" desc:"Idle timeout for mobile clients. Example: 900"`
	RestrictUserToPorts  *string   `json:"restrict-user-to-ports" desc:"Comma-separated list of allowed (or blocked, if negated) protocols and ports. Supports 'tcp(port)', 'udp(port)', 'icmp()', 'icmpv6()', and negation with '!()'. Example: 'tcp(443), tcp(80), udp(53)', or '!(tcp(22), udp(1194))' to block those ports."`
	SplitDNS             *string   `json:"split-dns" desc:"Comma-separated domains for DNS to go via VPN. Example: 'internal.company.com,corp.local'"`
	SessionTimeout       *int      `json:"session-timeout" desc:"Max session time in seconds before forced disconnect. Example: 3600"`
}

type OcservUser struct {
	ID            uint       `json:"-" gorm:"primaryKey;autoIncrement"`
	UID           string     `json:"uid" gorm:"type:varchar(26);not null;unique"`
	Group         string     `json:"group" gorm:"type:varchar(16);default:'defaults'"`
	Username      string     `json:"username" gorm:"type:varchar(16);not null;unique"`
	Password      string     `json:"password" gorm:"type:varchar(16);not null"`
	IsLocked      bool       `json:"is_locked" gorm:"default(false)"`
	CreatedAt     time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	ExpireAt      *time.Time `json:"expire_at"`
	DeactivatedAt *time.Time `json:"deactivated_at" gorm:"default:NULL"`
	TrafficType   string     `json:"traffic_type" gorm:"type:varchar(32);not null;default:1" enums:"Free,MonthlyTransmit,MonthlyReceive,TotallyTransmit,TotallyReceive"`
	TrafficSize   int        `json:"traffic_size" gorm:"not null;default:10"` // in GiB  >> x * 1024 ** 3
	Rx            int        `json:"rx" gorm:"not null;default:0"`            // Receive in bytes
	Tx            int        `json:"tx" gorm:"not null;default:0"`            // Transmit in bytes
	Description   string     `json:"description" gorm:"type:text"`
	IsOnline      bool       `json:"is_online" gorm:"-:migration;->"`
}

type OcUserActivity struct {
	ID        uint      `json:"-" gorm:"primaryKey;autoIncrement"`
	OcUserID  uint      `json:"-" gorm:"index;constraint:OnDelete:CASCADE;"`
	Type      string    `json:"type" gorm:"type:varchar(32);not null;default:1" enums:"Connected,Disconnected,Failed"`
	Log       string    `json:"log" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type OcUserTrafficStatistics struct {
	ID        uint      `json:"-" gorm:"primaryKey;autoIncrement"`
	OcUserID  uint      `json:"-" gorm:"index;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	Rx        int       `json:"rx" gorm:"default:0"` // in bytes
	Tx        int       `json:"tx" gorm:"default:0"` // in bytes
}

func (o *OcservUser) BeforeSave(tx *gorm.DB) (err error) {
	if !validateTrafficType(o.TrafficType) {
		return fmt.Errorf("invalid TrafficType: %s", o.TrafficType)
	}
	if o.TrafficType == Free {
		o.TrafficSize = 0
	}
	return nil
}

func (o *OcservUser) BeforeCreate(tx *gorm.DB) (err error) {
	o.UID = ulid.Make().String()
	return
}

func (a *OcUserActivity) BeforeSave(tx *gorm.DB) (err error) {
	if !validateActivityType(a.Type) {
		return fmt.Errorf("invalid Type: %s", a.Type)
	}
	return nil
}

func validateActivityType(t string) bool {
	switch t {
	case Connected, Disconnected, Failed:
		return true
	default:
		return false
	}
}

func validateTrafficType(trafficType string) bool {
	switch trafficType {
	case Free, MonthlyTransmit, MonthlyReceive, TotallyTransmit, TotallyReceive:
		return true
	default:
		return false
	}
}
