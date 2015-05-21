# consul-kv-ceph

A simple CLI program that connects to consul or etcd KV store based on `https://github.com/ashcrow/consuloretcd` and `https://github.com/JeanMertz/consulctl`


# Usage:

# Example<br />
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; `consuloretcd -A 192.168.200.90 -p 8500 get -consul key1`<br />
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; `consuloretcd -A 127.0.0.1 -p 8500 put -etcd key1 value`<br />
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; `consuloretcd -A 10.0.2.12 -p 8500 delete -consul key1` <br />
