// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package saver2v2

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/spdx/tools-golang/spdx"
)

func renderRelationships2_2(relationships []*spdx.Relationship2_2, buf *bytes.Buffer) error {

	var rels []interface{}
	for _, v := range relationships {
		rel := make(map[string]interface{})
		rel["spdxElementId"] = spdx.RenderDocElementID(v.RefA)
		rel["relatedSpdxElement"] = spdx.RenderDocElementID(v.RefB)
		rel["relationshipType"] = v.Relationship
		if v.RelationshipComment != "" {
			rel["comment"] = v.RelationshipComment
		}
		rels = append(rels, rel)
	}
	relsjson, _ := json.Marshal(rels)
	fmt.Fprintf(buf, "\"%s\": %s ,", "relationships", relsjson)
	return nil
}