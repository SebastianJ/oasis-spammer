package crypto

import (
	"fmt"

	"github.com/oasislabs/oasis-core/go/common/crypto/signature"
	fileSigner "github.com/oasislabs/oasis-core/go/common/crypto/signature/signers/file"
	"github.com/oasislabs/oasis-core/go/common/entity"
)

// LoadSigner - loads the signer from the PEM file
func LoadSigner(path string) (signature.Signer, error) {
	fmt.Printf("Path is: %s\n", path)
	_, signer, err := loadEntity(path)
	if err != nil {
		fmt.Printf("failed to load account entity, err: %s\n", err.Error())
		return nil, err
	}

	return signer, nil
}

func loadEntity(entityDir string) (*entity.Entity, signature.Signer, error) {
	factory, err := fileSigner.NewFactory(entityDir, signature.SignerEntity)
	if err != nil {
		return nil, nil, err
	}
	return entity.Load(entityDir, factory)
}
