// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// PoolSRVConfigList holds a list of PoolSRV configuration.
type PoolSRVConfigList struct {
	Items    []PoolSRVConfig `json:"items"`
	Kind     string          `json:"kind"`
	SelfLink string          `json:"selflink"`
}

// PoolSRVConfig holds the configuration of a single PoolSRV.
type PoolSRVConfig struct {
}

// PoolSRVEndpoint represents the REST resource for managing PoolSRV.
const PoolSRVEndpoint = "/pool/srv"

// PoolSRVResource provides an API to manage PoolSRV configurations.
type PoolSRVResource struct {
	c *f5.Client
}

// ListAll  lists all the PoolSRV configurations.
func (r *PoolSRVResource) ListAll() (*PoolSRVConfigList, error) {
	var list PoolSRVConfigList
	if err := r.c.ReadQuery(BasePath+PoolSRVEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single PoolSRV configuration identified by id.
func (r *PoolSRVResource) Get(id string) (*PoolSRVConfig, error) {
	var item PoolSRVConfig
	if err := r.c.ReadQuery(BasePath+PoolSRVEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new PoolSRV configuration.
func (r *PoolSRVResource) Create(item PoolSRVConfig) error {
	if err := r.c.ModQuery("POST", BasePath+PoolSRVEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a PoolSRV configuration identified by id.
func (r *PoolSRVResource) Edit(id string, item PoolSRVConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+PoolSRVEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single PoolSRV configuration identified by id.
func (r *PoolSRVResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+PoolSRVEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
