//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package bleve

import (
	"github.com/blevesearch/bleve/index/upside_down"
	"github.com/blevesearch/bleve/registry"
)

func openNonFileIndexUsing(path string, mapping *IndexMapping, kvstore string, kvconfig map[string]interface{}) (*indexImpl, error) {
	// first validate the mapping
	err := mapping.validate()
	if err != nil {
		return nil, err
	}

	if path == "" {
		return newMemIndex(mapping)
	}

	if kvconfig == nil {
		kvconfig = map[string]interface{}{}
	}

	rv := indexImpl{
		path:  path,
		m:     mapping,
		meta:  newIndexMeta(kvstore, kvconfig),
		stats: &IndexStat{},
	}
	storeConstructor := registry.KVStoreConstructorByName(rv.meta.Storage)
	if storeConstructor == nil {
		return nil, ErrorUnknownStorageType
	}
	// at this point there is hope that we can be successful, so save index meta
	// do not save meta to file
	// err = rv.meta.Save(path)
	// if err != nil {
	// 	return nil, err
	// }
	kvconfig["create_if_missing"] = true
	kvconfig["error_if_exists"] = true
	kvconfig["path"] = indexStorePath(path)

	// now create the store
	rv.s, err = storeConstructor(kvconfig)
	if err != nil {
		return nil, err
	}

	// open the index
	rv.i = upside_down.NewUpsideDownCouch(rv.s, Config.analysisQueue)
	err = rv.i.Open()
	if err != nil {
		return nil, err
	}
	rv.stats.indexStat = rv.i.Stats()

	// now persist the mapping
	// do not save the mapping info
	// mappingBytes, err := json.Marshal(mapping)
	// if err != nil {
	// 	return nil, err
	// }
	// err = rv.i.SetInternal(mappingInternalKey, mappingBytes)
	// if err != nil {
	// 	return nil, err
	// }

	// mark the index as open
	rv.mutex.Lock()
	defer rv.mutex.Unlock()
	rv.open = true
	return &rv, nil
}

func OpenNonFileIndexUsing(path string, mapping *IndexMapping, kvstore string, kvconfig map[string]interface{}) (Index, error) {
	return openNonFileIndexUsing(path, mapping, kvstore, kvconfig)
}
