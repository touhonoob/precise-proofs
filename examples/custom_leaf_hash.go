/* Example of using a custom hash function for leafs
 */
package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/centrifuge/precise-proofs/examples/documents"
	"github.com/centrifuge/precise-proofs/proofs"
)

func main() {
	// ExampleDocument is a protobuf message
	document := documentspb.ExampleDocument{
		Value1:      1,
		ValueA:      "Foo",
		ValueB:      "Bar",
		ValueBytes1: []byte("foobar"),
	}

	doctree := proofs.NewDocumentTree(proofs.TreeOptions{
		Hash:     sha256.New(),
		LeafHash: md5.New(), // Specify a custom hash function for leaves
	})

	checkErr(doctree.AddLeavesFromDocument(&document))
	checkErr(doctree.Generate())
	fmt.Printf("Generated tree: %s\n", doctree.String())

	// Generate the actual proof for a field. In this case the field called "ValueA".
	proof, err := doctree.CreateProof("valueA")
	checkErr(err)
	proofJson, _ := json.Marshal(proof)
	fmt.Println("Proof:\n", string(proofJson))

	// Validate the proof that was just generated
	valid, err := doctree.ValidateProof(&proof)
	checkErr(err)

	fmt.Printf("Proof validated: %v\n", valid)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}