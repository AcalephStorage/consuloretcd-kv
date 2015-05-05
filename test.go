package main

import (
    "github.com/ashcrow/consuloretcd"
    "net/http"
)

// Example
func main() {
    // You must provide and http.Client
    client := http.Client{}

    // Consul example. Replace "consul" with "etcd" to use etcd.
    consul, _ := consuloretcd.NewClient(
        "consul",
        consuloretcd.Config{
            Endpoint: "http://127.0.0.1",
            Client: client,
            Port: 8500})

//Ceph-common
    consul.PutKey("fsid","4a158d27-f750-41d5-9e7f-26ce4c9d2d45")
    consul.PutKey("cephx","true")
    consul.PutKey("cephx_require_signatures","false")
    consul.PutKey("cephx_cluster_require_signatures","true")
    consul.PutKey("cephx_service_require_signatures","false")
    consul.PutKey("max_open_files","131072")
    consul.PutKey("disable_in_memory_logs","true")

//Monitor
    consul.PutKey("mon_osd_down_out_interval","600")
    consul.PutKey("mon_osd_min_down_reporters","4")
    consul.PutKey("mon_clock_drift_allowed",".15")
    consul.PutKey("mon_clock_drift_warn_backoff","30")
    consul.PutKey("mon_osd_full_ratio",".95")
    consul.PutKey("mon_osd_nearfull_ratio",".85")
    consul.PutKey("mon_osd_report_timeout","300")

//OSD
    consul.PutKey("journal_size","100")
    consul.PutKey("pool_default_pg_num","128")
    consul.PutKey("pool_default_pgp_num","128")
    consul.PutKey("pool_default_size","3")
    consul.PutKey("pool_default_min_size","1")
    consul.PutKey("cluster_network","192.168.42.0/24")
    consul.PutKey("public_network","192.168.42.0/24")
    consul.PutKey("osd_mkfs_type","xfs")
    consul.PutKey("osd_mkfs_options_xfs","-f -i size=2048")
    consul.PutKey("osd_mount_options_xfs","noatime,largeio,inode,swalloc")
    consul.PutKey("osd_mon_heartbeat_interva","30")

//crush
    consul.PutKey("pool_default_crush_rule","0")
    consul.PutKey("osd_crush_update_on_start","true")

//Object backend
    consul.PutKey("osd_objectstore","filestore")

//Performance tuning
    consul.PutKey("filestore_merge_threshold","40")
    consul.PutKey("filestore_split_multiple","8")
    consul.PutKey("osd_op_threads","8")
    consul.PutKey("filestore_op_threads","8")
    consul.PutKey("filestore_max_sync_interval","5")
    consul.PutKey("osd_max_scrubs","1")

//Recovery tuning
    consul.PutKey("osd_recovery_max_active","5")
    consul.PutKey("osd_max_backfills","2")
    consul.PutKey("osd_recovery_op_priority","2")
    consul.PutKey("osd_client_op_priority","63")
    consul.PutKey("osd_recovery_max_chunk","1048576")
    consul.PutKey("osd_recovery_threads","1")

//Ports    
    consul.PutKey("mon_port","6789")
    consul.PutKey("ms_bind_port_min","6800")
    consul.PutKey("ms_bind_port_max","7100")

//rbd
    consul.PutKey("rbd_cache_enabled","true")
    consul.PutKey("rbd_cache_writethrough_until_flush","true")


//
    consul.PutKey("use_inktank_ceph_repo","true")
    consul.PutKey("iscsi_support","false")
    consul.PutKey("ceph_reduced_log_verbosity","false")
    consul.PutKey("radosgw","false")
}