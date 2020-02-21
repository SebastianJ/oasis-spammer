package genesis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/oasislabs/oasis-core/go/genesis/api"
)

// LoadDocument - loads the genesis document from a given path
func LoadDocument(path string) error {
	genesisPath := fmt.Sprintf("%s/genesis.json", path)

	// Set up the genesis system for the signature system's chain context.
	genesis, err := parseGenesisDocument(genesisPath)
	if err != nil {
		return fmt.Errorf("genesisFile.NewFileProvider: %w", err)
	}
	genesisDoc, err := genesis.GetGenesisDocument()
	if err != nil {
		return fmt.Errorf("genesis.GetGenesisDocument: %w", err)
	}
	fmt.Printf("setting chain context - chain_context: %s\n", genesisDoc.ChainContext())
	genesisDoc.SetChainContext()

	return nil
}

// fileProvider provides the static gensis document that network was
// initialized with.
type fileProvider struct {
	document *api.Document
}

func (p *fileProvider) GetGenesisDocument() (*api.Document, error) {
	return p.document, nil
}

func parseGenesisDocument(path string) (api.Provider, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("failed to open genesis document, err: %s\n", err.Error())
		return nil, err
	}

	var doc api.Document
	if err = json.Unmarshal(raw, &doc); err != nil {
		return nil, fmt.Errorf("genesis: malformed genesis file: %w", err)
	}

	return &fileProvider{document: &doc}, nil
}
