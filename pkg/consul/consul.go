package consul

import (
	"fmt"
	"strconv"

	consularisv1alpha1 "github.com/fire-ant/consularis/pkg/apis/consularis.io/v1alpha1"
	"github.com/fire-ant/consularis/pkg/config"
	consulapi "github.com/hashicorp/consul/api"
)

// NewClient returns a KV client for Consul
func NewClient(conf *config.Config) (*consulapi.KV, error) {
	dc := consulapi.DefaultConfig()
	dc.Address = conf.Consul
	dc.Datacenter = "vm"
	if conf.Port != "" {
		dc.Address = dc.Address + ":" + conf.Port
	}
	client, err := consulapi.NewClient(dc)
	if err != nil {
		return nil, err
	}
	kv := client.KV()
	return kv, nil
}

// KVUpdate will update the consul KV with the ConsulObject KV
func KVUpdate(client *consulapi.KV, object *consularisv1alpha1.ConsulObject) error {
	// p := &consulapi.KVPair{Key: "testkey1", Value: []byte("testvalue1")}
	// _, err := client.Put(p, nil)
	// if err != nil {
	// 	fmt.Printf("test put err: %s", err.Error())
	// }
	var ctxns consulapi.KVTxnOps
	kvs := *object.Spec.KV
	for _, kv := range kvs {
		var f uint64
		f = 0
		var err error
		if kv.Flag != "" {
			f, err = strconv.ParseUint(kv.Flag, 10, 64)
			if err != nil {
				return fmt.Errorf("failed to convert Flag %v for Key %v to uint64, make sure it's a 64-bit number", kv.Flag, kv.Key)
			}
		}
		ctxns = append(ctxns, &consulapi.KVTxnOp{
			Verb:  consulapi.KVSet,
			Key:   kv.Key,
			Flags: f,
			Value: []byte(kv.Value),
		})
	}
	ok, _, _, err := client.Txn(ctxns, nil)
	if !ok {
		return fmt.Errorf("consul transaction failed, error: %s", err.Error())
	}
	return nil

}

// KVDelete will delete the consul KV if the ConsulObject resource is deleted
func KVDelete(client *consulapi.KV, object *consularisv1alpha1.ConsulObject) error {
	var ctxns consulapi.KVTxnOps
	kvs := *object.Spec.KV
	for _, kv := range kvs {
		ctxns = append(ctxns, &consulapi.KVTxnOp{
			Verb: consulapi.KVDelete,
			Key:  kv.Key,
		})
	}
	ok, _, _, err := client.Txn(ctxns, nil)
	if !ok {
		return fmt.Errorf("consul transaction failed, error: %s", err.Error())
	}
	return nil

}
